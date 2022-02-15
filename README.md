# terradep
Dependency management tool for HashiCorp Terraform

## Command Structure
This project will be having following commands;
1. config - For configuration file management i.e. generating, modifying or resetting(~/.terradep/terradep.yaml and dcf.yaml)
1. dependencies - For managing dependencies i.e. installing, updating, removing, listing

## config
This command will be used for managing configuration file. It involves generating, modifying or resetting configuration files. It will be supporting dcf.yaml and terradep.yaml files. It has following subcommands:

| command | description | 
| ------ | ------ |
| init | Initialize default configuration files - terradep.yaml and dcf.yaml |
| repo | Add/update/delete repository configuration |

## dependencies
This command will be used for managing dependencies. It involves installing, updating, removing, listing dependencies. It has following subcommands:

| command | description |
| ------ | ------ |
| install | Download and install dependencies from dcf.yaml |
| update | Update dependencies from dcf.yaml |
| remove | Delete downloaded dependencies |
| build | Build dependencies and upload to a repository |
