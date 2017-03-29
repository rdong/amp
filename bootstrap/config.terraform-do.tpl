{{ source "default.ikt" }}
{{ source "file:///infrakit/env.ikt" }}
{{ $workerSize := ref "/swarm/size/worker" }}
[
  {
    "Plugin": "group",
    "Properties": {
      "ID": "amp-worker-{{ ref "/do/stackname" }}",
      "Properties": {
        "Allocation": {
          "Size": {{ $workerSize }}
        },
        "Instance": {
          "Plugin": "instance-terraform",
          "Properties": {
            "type": "digitalocean_droplet",
            "value": {
              "image": "${var.cluster_image}"
              "name": "${var.do_name}"
              "region": "${var.cluster_region}"
              "size": "${var.cluster_size}"
              "ssh_keys" "${var.cluster_ssh_key}"
              "tags": {
                SwarmRole = "worker"
                Project = "${var.do_name}"
              }
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
                  ]
                }
              }, {
                "Plugin": "flavor-swarm/worker",
                "Properties": {
                  "InitScriptTemplateURL": "{{ ref "/script/baseurl" }}/worker-init.tpl",
                  "SwarmJoinIP": "{{ ref "/bootstrap/ip" }}",
                  "Docker" : {
                    {{ if ref "/certificate/ca/service" }}"Host" : "unix:///var/run/docker.sock",
                    "TLS" : {
                      "CAFile": "{{ ref "/docker/remoteapi/cafile" }}",
                      "CertFile": "{{ ref "/docker/remoteapi/certfile" }}",
                      "KeyFile": "{{ ref "/docker/remoteapi/keyfile" }}",
                      "InsecureSkipVerify": false
                    }
                    {{ else }}"Host" : "unix:///var/run/docker.sock"
                    {{ end }}
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
