# terraform files

This folder contains the `*.tf` files read by Terraform.
Place your `terraform.tfvars` file with your custom values, don't update the `variables.tf` file.
Example content:

```
do_token = "77e027c7447f4680...."
do_name = "tgm-ikt"
do_ssh_key = "tgm-sfo-1"
infrakit_config_base_url = "https://raw.githubusercontent.com/appcelerator/amp/branchname/bootstrap"
```
