resource "fmc_intrusion_policy" "example" {
  name            = "ips_policy_1"
  description     = "My IPS Policy"
  base_policy_id  = ""
  inspection_mode = "PREVENTION"
}
