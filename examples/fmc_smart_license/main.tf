terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.1.1"
    }
  }
}

provider "fmc" {
  fmc_username             = var.fmc_username
  fmc_password             = var.fmc_password
  fmc_host                 = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

resource "fmc_smart_license" "registration" {
  registration_type = "EVALUATION"
}

# Runs successfully as long as the registration status is REGISTERED
# resource "fmc_smart_license" "registration" {
#   registration_type = "REGISTER"
#   token             = "X2M3YmJlY..."
# }

# Force to de-register and register with the token provided
# resource "fmc_smart_license" "registration" {
#   registration_type = "REGISTER"
#   token             = "X2M3YmJlY..."
#   force             = true
# }

# De-register on terraform destroy
# resource "fmc_smart_license" "registration" {
#   registration_type = "REGISTER"
#   token             = "X2M3YmJlY..."
#   retain            = false
# }

data "fmc_smart_license" "license" {}

output "smart_license_resource" {
  value     = fmc_smart_license.registration
  sensitive = true
}

output "smart_license_data" {
  value = data.fmc_smart_license.license
}
