terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}


resource "fmc_icmpv4_objects" "shbharti-icmpv4-5" {
  name        = "shbharti-icmpv4-5"
  icmp_type = "3"
  code  = 2
}

resource "fmc_icmpv4_objects" "shbharti-icmpv4-7" {
  name        = "shbharti-icmpv4-7"
  icmp_type = "3"
}

output "new_fmc_network_object_3" {
  value = fmc_icmpv4_objects.shbharti-icmpv4-5.id
}
