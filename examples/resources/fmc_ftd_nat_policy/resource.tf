resource "fmc_ftd_nat_policy" "example" {
  name        = "nat_policy_1"
  description = "My nat policy"
  manual_nat_rules = [
    {
      description          = "My manual nat rule 1"
      enabled              = true
      section              = "BEFORE_AUTO"
      nat_type             = "STATIC"
      fall_through         = false
      original_source_id   = "76d24097-41c4-4558-a4d0-a8c07ac08471"
      translated_source_id = "76d24097-41c4-4558-a4d0-a8c07ac08476"
    }
  ]
  auto_nat_rules = [
    {
      nat_type              = "STATIC"
      original_network_id   = "76d24097-41c4-4558-a4d0-a8c07ac08481"
      translated_network_id = "76d24097-41c4-4558-a4d0-a8c07ac08483"
    }
  ]
}
