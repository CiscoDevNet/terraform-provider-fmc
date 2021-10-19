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


resource "fmc_time_range_object" "test" {
  name        		     = "test-time-range"
  description 		     = "Test time range"
  effective_start_date = "2019-09-19T15:53"
  effective_end_date   = "2019-09-21T17:53"
}
