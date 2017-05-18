{{ source "default.ikt" }}
{{ source "file:///infrakit/env.ikt" }}
{{ $workerSize := var "/swarm/size/worker" }}
[
  {
    "Plugin": "group",
    "Properties": {
      "ID": "amp-worker-{{ var "/aws/vpcid" }}",
      "Properties": {
        "Allocation": {
          "Size": 1
        },
        "Instance": {
          "Plugin": "instance-aws/autoscalinggroup-autoscalinggroup",
          "Properties": {
            "CreateAutoScalingGroupInput": {
              "DesiredCapacity": {{ $workerSize }},
              "HealthCheckGracePeriod": "200",
              "HealthCheckType": "EC2",
              "LaunchConfigurationName": "{{ var "/aws/stackname" }}-LaunchConfiguration",
              "MaxSize": "5",
              "MinSize": "0",
              "Tags": {
                "Name": "{{ var "/aws/stackname" }}-worker",
                "Deployment": "Infrakit",
                "Role" : "worker"
              },
              "VPCZoneIdentifier": [ "{{ var "/aws/subnetid1" }}", "{{ var "/aws/subnetid1" }}", "{{ var "/aws/subnetid1" }}" ]
            }
          }
        },
        "Instance": {
          "Plugin": "instance-aws/autoscalinggroup-launchconfiguration",
          "Properties": {
            "CreateLaunchConfigurationInput": {
              "AssociatePublicIpAddress": "true",
              "IamInstanceProfile": "{{ var "/aws/instanceprofile" }}",
              "ImageId": "{{ var "/aws/amiid" }}",
              "InstanceType": "{{ var "/aws/instancetype" }}",
              "KeyName": "{{ var "/aws/keyname" }}",
              "SecurityGroups": [ "{{ var "/aws/securitygroupid" }}" ]
            }
          }
        },
        "Flavor": {
          "Plugin": "flavor-combo",
          "Properties": {
            "Flavors": [
              {
                "Plugin": "flavor-vanilla",
                "Properties": {
                  "Init": [
                    "#!/bin/bash",
                    "apt-get install -y awscli jq"
                  ]
                }
              }, {
                "Plugin": "flavor-swarm/worker",
                "Properties": {
                  "InitScriptTemplateURL": "{{ var "/script/baseurl" }}/worker-init.tpl",
                  "SwarmJoinIP": "{{ var "/bootstrap/ip" }}",
                  "Docker" : {
                    "Host" : "unix:///var/run/docker.sock"
                  }
                }
              }
            ]
          }
        }
      }
    }
  }
]
