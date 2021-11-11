# terraform config 
terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}


# Imput variables

variable "aws_access_key" { 
    type = string
}

variable "aws_secret_key" { 
    type = string
}

variable "region" { 
    type = string
}

variable "fmc_password" { 
    type = string
}

variable "fmc_vm_key_name" { 
    type = string
}


# providers configs 

provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.region
}


# resources 

module "fmc" { 
    source = "./fmc/" 

    password = var.fmc_password
    key_name = var.fmc_vm_key_name
}


# outputs 

output "fmc_ip" {
  value = module.fmc.FMCv_EIPs[0]
}

output "fmc_username" {
    // default value, can't be changed
    value = "admin"
}

output "fmc_password" { 
  value = var.fmc_password
}
