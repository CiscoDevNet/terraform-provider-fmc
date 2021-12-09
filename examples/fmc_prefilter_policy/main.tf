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

resource "fmc_prefilter_policy" "prefilter_policy1" {
  name        = "Prefilter Policy"
  description = "Terraform Prefilter Policy description"
  default_action { 
    log_begin = false
    send_events_to_fmc = false
    action = "BLOCK_TUNNELS"
  }
}

resource "fmc_prefilter_policy" "prefilter_policy2" {
  name        = "Another Prefilter Policy"
  description = "Terraform Prefilter Policy description"
  default_action { 
    log_begin = false
    send_events_to_fmc = false
    action = "ANALYZE_TUNNELS" 
  }
}