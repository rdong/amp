# Default variables used by the *.tf files
# This file shouldn't be modified, override the variables in a terraform.tfvars file instead
# Any empty variables MUST be overwritten.

variable "do_token" {
  default = ""
}

variable "do_image" {
  default = {
    ubuntu = "ubuntu-16-04-x64"
  }
}

variable "do_name" {
  default = "ikt"
}

variable "do_region" {
  default = "sfo1"
}

variable "do_size" {
  default = "1GB"
}

variable "do_ssh_key" {
  default = ""
}

variable "infrakit_config_base_url" {
  type = "string"
  description = "Base URL for InfraKit configuration. there should be a bootstrap.sh, a variables.ikt and a config.tpl file"
  default = "https://raw.githubusercontent.com/appcelerator/amp/master/bootstrap"
}
