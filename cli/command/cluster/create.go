package cluster

import (
	"context"
	"errors"
	"strings"

	"github.com/appcelerator/amp/api/registration"
	"github.com/appcelerator/amp/api/rpc/cluster"
	"github.com/appcelerator/amp/cli"
	"github.com/appcelerator/amp/data/clusters"
	"github.com/spf13/cobra"
)

// NewCreateCommand returns a new instance of the create command for bootstrapping an cluster.
func NewCreateCommand(c cli.Interface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [OPTIONS]",
		Short:   "Create an amp cluster",
		PreRunE: cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return create(c, cmd)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.workers, "workers", "w", 2, "Initial number of worker nodes")
	flags.IntVarP(&opts.managers, "managers", "m", 3, "Intial number of manager nodes")
	flags.StringVar(&opts.provider, "provider", "local", "Cluster provider")
	flags.StringVar(&opts.name, "name", "", "Cluster Label")
	flags.StringVarP(&opts.tag, "tag", "t", "latest", "Specify tag for cluster images (default is 'latest', use 'local' for development)")
	flags.StringVarP(&opts.registration, "registration", "r", registration.Default, "Specify the registration policy (default is 'email', possible values are 'none' or 'email')")
	flags.StringVarP(&opts.organization, "organization", "o", "", "Organization for which the cluster will be created")
	return cmd
}

// Map cli cluster flags to target bootstrap cluster command flags and update the cluster
func create(c cli.Interface, cmd *cobra.Command) error {
	// This is a map from cli cluster flag name to bootstrap script flag name
	m := map[string]string{
		"workers":      "-w",
		"managers":     "-m",
		"provider":     "-t",
		"name":         "-l",
		"tag":          "-T",
		"registration": "-r",
	}

	if opts.provider == "local" {
		// local deployment is not a remote call to amplifier, it is done by the CLI
		// the following ensures that flags are added before the final command arg
		// TODO: refactor reflag to handle this
		args := []string{"bin/deploy"}
		args = reflag(cmd, m, args)
		args = append(args, DefaultLocalClusterID)
		env := map[string]string{"TAG": opts.tag, "REGISTRATION": opts.registration}
		return queryCluster(c, args, env)
	} else {
		request := cluster.CreateRequest{}
		if val, ok := clusters.CloudProvider_value[strings.ToUpper(opts.provider)]; ok {
			request.Provider = clusters.CloudProvider(val)
		} else {
			return errors.New("unknown provider")
		}
		request.Name = opts.name
		request.ManagerCount = int64(opts.managers)
		conn := c.ClientConn()
		cc := cluster.NewClusterClient(conn)
		ctx := context.Background()
		reply, err := cc.Create(ctx, &request)
		if err != nil {
			return err
		}
		c.Console().Printf("%+v\n", reply)
		return nil
	}
}
