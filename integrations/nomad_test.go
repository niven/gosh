package integrations_test

import (
	"encoding/json"
	"testing"

	"github.com/niven/gosh/integrations"
	"github.com/stretchr/testify/assert"
)

func TestNomadStructs(t *testing.T) {

	input := `
	{
		"Region": "global",
		"Namespace": "apps",
		"ID": "example",
		"ParentID": "",
		"Name": "example",
		"Type": "batch",
		"Priority": 50,
		"AllAtOnce": false,
		"Datacenters": ["dc1"],
		"NodePool": "prod",
		"Constraints": [
		  {
			"LTarget": "${attr.kernel.name}",
			"RTarget": "linux",
			"Operand": "="
		  }
		],
		"TaskGroups": [
		  {
			"Name": "cache",
			"Count": 1,
			"Constraints": [
			  {
				"LTarget": "${attr.os.signals}",
				"RTarget": "SIGUSR1",
				"Operand": "set_contains"
			  }
			],
			"Affinities": [
			  {
				"LTarget": "${meta.datacenter}",
				"RTarget": "dc1",
				"Operand": "=",
				"Weight": 50
			  }
			],
			"RestartPolicy": {
			  "Attempts": 10,
			  "Interval": 300000000000,
			  "Delay": 25000000000,
			  "Mode": "delay"
			},
			"Tasks": [
			  {
				"Config": {
				  "command": "env",
				  "image": "alpine"
				},
				"Driver": "docker",
				"Lifecycle": {
				  "Hook": "prestart",
				  "Sidecar": false
				},
				"Name": "init",
				"Resources": {
				  "CPU": 100,
				  "MemoryMB": 300
				}
			  },
			  {
				"Name": "redis",
				"Driver": "docker",
				"User": "foo-user",
				"Config": {
				  "image": "redis:latest",
				  "port_map": [
					{
					  "db": 6379
					}
				  ]
				},
				"Env": {
				  "foo": "bar",
				  "baz": "pipe"
				},
				"Services": [
				  {
					"Name": "cache-redis",
					"PortLabel": "db",
					"Tags": ["global", "cache"],
					"Checks": [
					  {
						"Name": "alive",
						"Type": "tcp",
						"Command": "",
						"Args": null,
						"Path": "",
						"Protocol": "",
						"PortLabel": "",
						"Interval": 10000000000,
						"Timeout": 2000000000,
						"InitialStatus": ""
					  }
					]
				  }
				],
				"Vault": null,
				"Templates": [
				  {
					"SourcePath": "local/config.conf.tpl",
					"DestPath": "local/config.conf",
					"EmbeddedTmpl": "",
					"ChangeMode": "signal",
					"ChangeSignal": "SIGUSR1",
					"Splay": 5000000000,
					"Perms": ""
				  }
				],
				"Constraints": null,
				"Affinities": null,
				"Resources": {
				  "CPU": 500,
				  "MemoryMB": 256,
				  "DiskMB": 0,
				  "Networks": [
					{
					  "Device": "",
					  "CIDR": "",
					  "IP": "",
					  "MBits": 10,
					  "ReservedPorts": [
						{
						  "Label": "rpc",
						  "Value": 25566
						}
					  ],
					  "DynamicPorts": [
						{
						  "Label": "db",
						  "Value": 0
						}
					  ]
					}
				  ]
				},
				"DispatchPayload": {
				  "File": "config.json"
				},
				"Meta": {
				  "foo": "bar",
				  "baz": "pipe"
				},
				"KillTimeout": 5000000000,
				"LogConfig": {
				  "Disabled": false,
				  "MaxFiles": 10,
				  "MaxFileSizeMB": 10
				},
				"Artifacts": [
				  {
					"GetterSource": "http://foo.com/artifact.tar.gz",
					"GetterOptions": {
					  "checksum": "md5:c4aa853ad2215426eb7d70a21922e794"
					},
					"RelativeDest": "local/"
				  }
				],
				"Leader": false
			  }
			],
			"EphemeralDisk": {
			  "Sticky": false,
			  "SizeMB": 300,
			  "Migrate": false
			},
			"Meta": {
			  "foo": "bar",
			  "baz": "pipe"
			}
		  }
		],
		"Update": {
		  "Stagger": 10000000000,
		  "MaxParallel": 1
		},
		"Periodic": {
		  "Enabled": true,
		  "Spec": "* * * * *",
		  "SpecType": "cron",
		  "ProhibitOverlap": true
		},
		"ParameterizedJob": {
		  "Payload": "required",
		  "MetaRequired": ["foo"],
		  "MetaOptional": ["bar"]
		},
		"Payload": null,
		"Meta": {
		  "foo": "bar",
		  "baz": "pipe"
		},
		"VaultToken": "",
		"Status": "running",
		"StatusDescription": "",
		"CreateIndex": 7,
		"ModifyIndex": 7,
		"JobModifyIndex": 7
	  } 
	`
	result := integrations.NomadJobInfo{}
	err := json.Unmarshal([]byte(input), &result)

	assert.Nil(t, err)
	assert.Equal(t, "example", result.Name)
	assert.Equal(t, "running", result.Status)
}
