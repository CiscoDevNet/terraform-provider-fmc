terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = "jay"
  fmc_password = "CXsecurity!@34"
  fmc_host = "10.106.36.90"
  fmc_insecure_skip_verify = true
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
