terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      version = "0.2"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

resource "fmc_access_policies" "access_policy" {
    name = "Terraform Access Policy"
    default_action = "block" 
}

resource "fmc_access_policies_category" "access_policy_test_category" {
    name             = "CategoryTest"
    access_policy_id = fmc_access_policies.access_policy.id
}

resource "fmc_access_rules" "access_rule_1" {
    acp = fmc_access_policies.access_policy.id
    category = fmc_access_policies_category.access_policy_test_category.name
    name = "Test rule 1"
    enabled = true
    action = "allow"
}
