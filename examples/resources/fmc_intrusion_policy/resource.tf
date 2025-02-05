resource "fmc_intrusion_policy" "example" {
  name            = "fmc_intrusion_policy"
  description     = "My IPS Policy"
  base_policy_id  = ""
  inspection_mode = "PREVENTION"
}
