resource "fmc_intrusion_policy" "example" {
  name            = "my_intrusion_policy"
  description     = "My IPS Policy"
  base_policy_id  = "0050568A-4E02-1ed3-0000-004294969198"
  inspection_mode = "PREVENTION"
}
