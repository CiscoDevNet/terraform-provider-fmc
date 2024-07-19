resource "fmc_network_analysis_policy" "example" {
  name            = "net_analysis_policy_1"
  description     = "My network analysis policy"
  base_policy_id  = ""
  inspection_mode = "PREVENTION"
}
