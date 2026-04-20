resource "fmc_intrusion_policy_group_override" "example" {
  intrusion_policy_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  intrusion_rule_group_id = "76d24097-41c4-4558-a4d0-a8c07ac08471"
  override_security_level = "LEVEL_2"
}
