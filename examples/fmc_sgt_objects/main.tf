terraform {
  required_providers {
    
    fmc = {
      source = "registry.terraform.io/CiscoDevNet/fmc"
    }
  }
}


provider "fmc" {
    fmc_username = "admin"
    fmc_password = "P@ssw123"
    fmc_host = "192.168.25.174"
    fmc_insecure_skip_verify = true
}

# Create a access control policy
resource "fmc_access_policies" "access_policy" {
    name = "terraform-acp"
    default_action = "permit"
}

data "fmc_access_policies" "acp1" {
    name = fmc_access_policies.access_policy.name
}

# Create security zones
resource "fmc_security_zone" "inside" {
  name          = "inside"
  interface_mode = "ROUTED"
}
resource "fmc_security_zone" "outside" {
  name          = "outside"
  interface_mode = "ROUTED"
}

data "fmc_security_zones" "inside" {
    name = fmc_security_zone.inside.name
}
data "fmc_security_zones" "outside" {
    name = fmc_security_zone.outside.name
}

# Get the security group tag
resource "fmc_sgt_objects" "my_sgt1" {
  name = "another_one"
  description = "bites the dust"
  tag = "26"
}

resource "fmc_sgt_objects" "my_sgt2" {
  name = "and_another_one"
  description = "gone and another one bites the dust"
  tag = "27"
}

data "fmc_sgt_objects" "sgt1" {
    name = fmc_sgt_objects.my_sgt1.name
}

data "fmc_sgt_objects" "sgt2" {
    name = fmc_sgt_objects.my_sgt2.name
}
# create a access rule
resource "fmc_access_rules" "access_rule_1" {
    acp = data.fmc_access_policies.acp1.id
    section = "default"
    name = "Test rule 1"
    action = "allow"
    enabled = true
    enable_syslog = true
    syslog_severity = "alert"
    send_events_to_fmc = true
    log_files = false
    log_end = true
    source_zones {
        source_zone {
            id = data.fmc_security_zones.inside.id
            type =  data.fmc_security_zones.inside.type
        }
        source_zone {
            id = data.fmc_security_zones.outside.id
            type =  data.fmc_security_zones.outside.type
        }
    }
    destination_zones {
        destination_zone {
            id = data.fmc_security_zones.outside.id
            type =  data.fmc_security_zones.outside.type
        }
    }

    source_security_group_tags {
        source_security_group_tag {
            id = data.fmc_sgt_objects.sgt2.id
            type = data.fmc_sgt_objects.sgt2.type
        }
    }

    destination_security_group_tags {
        destination_security_group_tag {
            id = data.fmc_sgt_objects.sgt1.id
            type = data.fmc_sgt_objects.sgt1.type
        }
    }

    new_comments = [ "New", "comment" ]
}
