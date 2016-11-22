package cli_test

import (
	"fmt"
	"os/exec"
	"sync"

	"bytes"
	"context"
	"io/ioutil"
	"math/rand"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

//TestSpec contains all the CommandSpec objects
type TestSpec struct {
	Name     string
	Timeout  time.Duration
	Commands []CommandSpec
}

//CommandSpec defines the commands with arguments and options
type CommandSpec struct {
	Cmd         string   `yaml:"cmd"`
	Args        []string `yaml:"args"`
	Options     []string `yaml:"options"`
	Expectation string   `yaml:"expectation"`
	Retry       int      `yaml:"retry"`
	Timeout     string   `yaml:"timeout"`
	Delay       string   `yaml:"delay"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	testDir      = "./samples"
	lookupDir    = "./lookup"
	regexMap     map[string]string
	suiteTimeout string
	wg sync.WaitGroup
)

//read, parse and execute test commands
func TestCmds(t *testing.T) {
	suiteTimeout = "1m"
	duration, err := time.ParseDuration(suiteTimeout)
	if err != nil {
		t.Errorf("Unable to generate suite timeout, reason: %v", err)
		return
	}
	//test suite context
	ctxSuite, cancelSuite := context.WithTimeout(context.Background(), duration)
	defer cancelSuite()
	go checkTimeout(t, ctxSuite, "Suite")
	err = loadRegexLookup()
	if err != nil {
		t.Errorf("Unable to load lookup specs, reason: %v", err)
		return
	}
	tests, err := loadTestSpecs()
	if err != nil {
		t.Errorf("Unable to load test specs, reason: %v", err)
		return
	}
	wg.Add(len(tests))
	for _, test := range tests {
		t.Logf("-----------------------------------------------------------------------------------------")
		t.Logf("Running spec: %s", test.Name)
		//test spec context
		ctxSpec, cancelSpec := context.WithTimeout(context.Background(), test.Timeout)
		defer cancelSpec()
		go checkTimeout(t, ctxSpec, test.Name)
		go runTestSpec(t, test)
	}
	wg.Wait()
	t.Log("Finished!")
}

//read test_samples directory by parsing its contents
func loadTestSpecs() ([]*TestSpec, error) {
	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		return nil, err
	}
	tests := []*TestSpec{}
	for _, file := range files {
		test, err := loadTestSpec(path.Join(testDir, file.Name()))
		if err != nil {
			return nil, err
		}
		if test != nil {
			tests = append(tests, test)
		}
	}
	return tests, nil
}

//parse test_samples directory and unmarshal its contents
func loadTestSpec(fileName string) (*TestSpec, error) {
	if filepath.Ext(fileName) != ".yml" {
		return nil, nil
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to load test spec: %s. Error: %v", fileName, err)
	}
	duration, duraErr := time.ParseDuration("0ms")
	if duraErr != nil {
		return nil, fmt.Errorf("Unable to create duration for timeout: %s. Error: %v", fileName, err)
	}
	testSpec := &TestSpec{
		Name:    fileName,
		Timeout: duration,
	}

	var commandMap []CommandSpec
	if err := yaml.Unmarshal(content, &commandMap); err != nil {
		return nil, fmt.Errorf("Unable to parse test spec: %s. Error: %v", fileName, err)
	}

	for _, command := range commandMap {
		if command.Timeout == "" {
			command.Timeout = "5000ms"
		}
		testSpecTimeout := "1m"
		duration, duraErr := time.ParseDuration(testSpecTimeout)
		if duraErr != nil {
			return nil, fmt.Errorf("Unable to create duration for timeout: %s. Error: %v", fileName, err)
		}
		testSpec.Timeout = duration
		testSpec.Commands = append(testSpec.Commands, command)
	}
	return testSpec, nil
}

//execute commands and check for timeout, delay and retry
func runTestSpec(t *testing.T, test *TestSpec) {
	defer wg.Done()
	var i int
	var cache = map[string]string{}
	var err error
	//iterate through all the testSpec
	for _, cmdSpec := range test.Commands {
		var cmdTmplString []string
		duration, duraErr := time.ParseDuration(cmdSpec.Timeout)
		if duraErr != nil {
			t.Log("Parsing duration failed: %v", err)
			t.Fail()
		}
		//cmd Spec context
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		go checkTimeout(t, ctx, cmdSpec.Cmd)
		for i = -1; i < cmdSpec.Retry; i++ {
			//err is set to nil a the beginning of the loop to ensure that each time a
			//command is retried or executed atleast once without the error assigned
			//from the previous executions
			err = nil

			//generate command string from cmdSpec
			cmdString := generateCmdString(&cmdSpec)

			//perform templating on cmdString
			cmdTmplOutput, tmplErr := performTemplating(strings.Join(cmdString, " "), cache)
			if tmplErr != nil {
				t.Log("Executing templating failed: %s", tmplErr)
				t.Fail()
			}
			cmdTmplString = strings.Fields(cmdTmplOutput)

			//execute command
			cmdOutput, cmdErr := exec.Command(cmdTmplString[0], cmdTmplString[1:]...).CombinedOutput()
			t.Logf("Running: %s", strings.Join(cmdTmplString, " "))

			//perform templating on RegEx string
			regexTmplOutput, tmplErr := performTemplating(cmdSpec.Expectation, cache)
			if tmplErr != nil {
					t.Log("Executing templating failed: %s", tmplErr)
					t.Fail()
			}

			//check if the command output matches the RegEx
			expectedOutput := regexp.MustCompile(regexTmplOutput)
			if !expectedOutput.MatchString(string(cmdOutput)) {
				err = fmt.Errorf("Mismatched expected output: %s : Error: %v", cmdOutput, cmdErr)
				t.Log(err)
			}

			//if no error after retries, break the loop to continue command execution
			if err == nil {
				break
			}
			//add delay (in Millisecond) to wait for command execution
			if cmdSpec.Delay != "" {
				del, delErr := time.ParseDuration(cmdSpec.Delay)
				if delErr != nil {
					t.Log("Invalid delay specified: %s : Error: %v", cmdSpec.Delay, delErr)
					t.Fail()
				}
				time.Sleep(del)
			}
		}
		if i > 0 {
			t.Log("This command :", cmdTmplString, "has re-run", i, "times.")
		}
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		cancel()
	}
}

//create an array of strings representing the commands by concatenating
//all the fields from the yml files in test_samples directory
func generateCmdString(cmdSpec *CommandSpec) (cmdString []string) {
	cmdSplit := strings.Fields(cmdSpec.Cmd)
	optionsSplit := []string{}
	for _, val := range cmdSpec.Options {
		optionsSplit = append(optionsSplit, strings.Fields(val)...)
	}
	cmdString = append(cmdSplit, cmdSpec.Args...)
	cmdString = append(cmdString, optionsSplit...)
	if regexMap[cmdSpec.Expectation] != "" {
		cmdSpec.Expectation = regexMap[cmdSpec.Expectation]
	}
	return
}

//read lookup directory by parsing its contents
func loadRegexLookup() error {
	files, err := ioutil.ReadDir(lookupDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		err := parseLookup(path.Join(lookupDir, file.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

//parse lookup directory and unmarshal its contents
func parseLookup(file string) error {
	if filepath.Ext(file) != ".yml" {
		return nil
	}
	pairs, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Unable to load regex lookup: %s. Error: %v", file, err)
	}
	if err := yaml.Unmarshal(pairs, &regexMap); err != nil {
		return fmt.Errorf("Unable to parse regex lookup: %s. Error: %v", file, err)
	}
	return nil
}

//create, parse and execute a template to generate unique values
func performTemplating(s string, cache map[string]string) (output string, err error) {
	var t *template.Template
	t, err = template.New("Command").Parse(s)
	if err != nil {
		return
	}
	//custom function to create a unique name with a randomly generated string
	name := func(in string) string {
		if val, ok := cache[in]; ok {
			return val
		}
		out := in + randString(10)
		cache[in] = out
		return out
	}
	//custom function to randomly generate a port number
	port := func(in string, min, max int) string {
		if val, ok := cache[in]; ok {
			return val
		}
		out := strconv.Itoa(rand.Intn(max-min) + min)
		cache[in] = out
		return out
	}
	var doc bytes.Buffer
	//add the custom functions to template for execution
	var fMap = template.FuncMap{
		"uniq": func(in string) string { return name(in) },
		"port": func(in string, min, max int) string { return port(in, min, max) },
	}
	//execute the parsed template
	err = t.Execute(&doc, fMap)
	if err != nil {
		return
	}
	output = doc.String()
	return
}

//generate a random string consisting of uppercase and lowercase characters
func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//repeatedly call the done() method for the context, fail if the deadline exceeds
func checkTimeout(t *testing.T, ctx context.Context, name string){
  for {
    select {
    case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				t.Fatal("Deadline exceeded:", name)
			}
			return
    }
  }
}