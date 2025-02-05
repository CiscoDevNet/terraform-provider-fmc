resource "fmc_network_analysis_policy" "example" {
  name            = "my_network_analysis_policy"
  description     = "My network analysis policy"
  base_policy_id  = ""
  inspection_mode = "PREVENTION"
}
