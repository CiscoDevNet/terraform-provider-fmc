terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/cisco/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_security_zones" "inside" {
    name = "inside"
}

data "fmc_security_zones" "outside" {
    name = "outside"
}

data "fmc_network_objects" "private" {
    name = "VLAN825-Private"
}

data "fmc_network_objects" "public" {
    name = "VLAN825-Public"
}

data "fmc_host_objects" "CUCMPub" {
  name = "CUCM-Pub"
}

data "fmc_host_objects" "CUCMPubDR" {
  name = "CUCM-Pub-DR"
}

resource "fmc_ftd_nat_policies" "nat_policy" {
    name = "Terraform NAT Policy"
    description = "New NAT policy!"
}

resource "fmc_ftd_autonat_rules" "new_rule" {
    nat_policy = fmc_ftd_nat_policies.nat_policy.id
    description = "Testing Auto NAT priv-pub"
    nat_type = "static"
    source_interface {
        id = data.fmc_security_zones.inside.id
        type = data.fmc_security_zones.inside.type
    }
    destination_interface {
        id = data.fmc_security_zones.outside.id
        type = data.fmc_security_zones.outside.type
    }
    original_network {
        id = data.fmc_network_objects.private.id
        type = data.fmc_network_objects.private.type
    }
    translated_network {
        id = data.fmc_network_objects.public.id
        type = data.fmc_network_objects.public.type
    }
    translated_network_is_destination_interface = false
    original_port {
        port = 53
        protocol = "udp"
    }
    translated_port = 5353
    ipv6 = true
}

resource "fmc_ftd_autonat_rules" "new_rule_2" {
    nat_policy = fmc_ftd_nat_policies.nat_policy.id
    description = "Testing Auto NAT pub-priv"
    nat_type = "dynamic"
    source_interface {
        id = data.fmc_security_zones.inside.id
        type = data.fmc_security_zones.inside.type
    }
    destination_interface {
        id = data.fmc_security_zones.outside.id
        type = data.fmc_security_zones.outside.type
    }
    original_network {
        id = data.fmc_host_objects.CUCMPub.id
        type = data.fmc_host_objects.CUCMPub.type
    }
    translated_network_is_destination_interface = false
    pat_options {
        pat_pool_address {
            id = data.fmc_network_objects.private.id
            type = data.fmc_network_objects.private.type
        }
        extended_pat_table = true
        round_robin = true
    }
    ipv6 = true
}

resource "fmc_ftd_autonat_rules" "new_rule_3" {
    nat_policy = fmc_ftd_nat_policies.nat_policy.id
    description = "Testing Auto NAT pub-priv without PAT pool"
    nat_type = "dynamic"
    source_interface {
        id = data.fmc_security_zones.outside.id
        # type = data.fmc_security_zones.outside.type
    }
    destination_interface {
        id = data.fmc_security_zones.outside.id
        # type = data.fmc_security_zones.outside.type
    }
    original_network {
        id = data.fmc_host_objects.CUCMPubDR.id
        # type = data.fmc_host_objects.CUCMPubDR.type
    }
    pat_options {
        interface_pat = true
        round_robin = true
    }
    ipv6 = true
}

output "new_ftd_nat_policy" {
    value = fmc_ftd_nat_policies.nat_policy
}

output "new_ftd_autonat_rule" {
    value = fmc_ftd_autonat_rules.new_rule
}

# output "new_ftd_autonat_rule_2" {
#     value = fmc_ftd_autonat_rules.new_rule_2
# }
