terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = "api"
  fmc_password = "CXsecurity!@34"
  fmc_host = "10.106.107.228"
  fmc_insecure_skip_verify = true
}

data "fmc_security_zones" "inside" {
    name = "inside"
}

data "fmc_security_zones" "outside" {
    name = "outside"
}

data "fmc_host_objects" "CUCMPub" {
  name = "CUCM-Pub"
}

data "fmc_network_objects" "private" {
    name = "VLAN825-Private"
}

data "fmc_network_objects" "public" {
    name = "VLAN825-Public"
}

resource "fmc_ftd_nat_policies" "nat_policy" {
    name = "Terraform NAT Policy"
    description = "New NAT policy!"
}

resource "fmc_ftd_manualnat_rules" "new_rule" {
    nat_policy = fmc_ftd_nat_policies.nat_policy.id
    description = "Testing Manual NAT priv-pub"
    nat_type = "static"
    source_interface {
        id = data.fmc_security_zones.inside.id
        type = data.fmc_security_zones.inside.type
    }
    destination_interface {
        id = data.fmc_security_zones.outside.id
        type = data.fmc_security_zones.outside.type
    }
    original_source {
        id = data.fmc_network_objects.public.id
        type = data.fmc_network_objects.public.type
    }
    translated_destination {
        id = data.fmc_network_objects.public.id
        type = data.fmc_network_objects.public.type
    }
    interface_in_original_destination = true
    interface_in_translated_source = true
    ipv6 = true
}

resource "fmc_ftd_manualnat_rules" "new_rule_after" {
    nat_policy = fmc_ftd_nat_policies.nat_policy.id
    description = "Testing Manual NAT priv-pub"
    nat_type = "static"
    section = "after_auto"
    source_interface {
        id = data.fmc_security_zones.inside.id
        type = data.fmc_security_zones.inside.type
    }
    destination_interface {
        id = data.fmc_security_zones.outside.id
        type = data.fmc_security_zones.outside.type
    }
    original_source {
        id = data.fmc_network_objects.public.id
        type = data.fmc_network_objects.public.type
    }
    translated_destination {
        id = data.fmc_network_objects.public.id
        type = data.fmc_network_objects.public.type
    }
    interface_in_original_destination = true
    interface_in_translated_source = true
    ipv6 = true
}

resource "fmc_ftd_manualnat_rules" "new_rule_before_1" {
    nat_policy = fmc_ftd_nat_policies.nat_policy.id
    description = "Testing Manual NAT before priv-pub"
    nat_type = "static"
    section = "before_auto"
    target_index = 1
    source_interface {
        id = data.fmc_security_zones.inside.id
        type = data.fmc_security_zones.inside.type
    }
    destination_interface {
        id = data.fmc_security_zones.outside.id
        type = data.fmc_security_zones.outside.type
    }
    original_source {
        id = data.fmc_network_objects.public.id
        type = data.fmc_network_objects.public.type
    }
    translated_destination {
        id = data.fmc_host_objects.CUCMPub.id
        type = data.fmc_host_objects.CUCMPub.type
    }
    interface_in_original_destination = true
    interface_in_translated_source = true
    ipv6 = true
}

output "new_ftd_nat_policy" {
    value = fmc_ftd_nat_policies.nat_policy
}

output "new_ftd_manualnat_rule" {
    value = fmc_ftd_manualnat_rules.new_rule
}

output "new_ftd_manualnat_rule_after" {
    value = fmc_ftd_manualnat_rules.new_rule_after
}

output "new_ftd_manualnat_rule_before_1" {
    value = fmc_ftd_manualnat_rules.new_rule_before_1
}
