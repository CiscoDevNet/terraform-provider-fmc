resource "fmc_intrusion_policy_rule_override" "example" {
  intrusion_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  intrusion_rule_id   = "76d24097-41c4-4558-a4d0-a8c07ac08471"
  override_state      = "BLOCK"
  rule_data           = "alert icmp any any -> any any ( sid:10000301; gid:2000; msg:\"CUSTOM RULE1\"; classtype:icmp-event; rev:1; )"
}
