resource "fmc_network_analysis_policy" "example" {
  name            = "my_network_analysis_policy"
  description     = "My network analysis policy"
  base_policy_id  = "0050568A-4E02-1ed3-0000-004294969198"
  inspection_mode = "PREVENTION"
}
