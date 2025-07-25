resource "fmc_ftd_manual_nat_rule" "example" {
  ftd_nat_policy_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  description          = "My manual nat rule 1"
  enabled              = true
  section              = "BEFORE_AUTO"
  nat_type             = "STATIC"
  fall_through         = false
  original_source_id   = "76d24097-41c4-4558-a4d0-a8c07ac08471"
  translated_source_id = "76d24097-41c4-4558-a4d0-a8c07ac08476"
}
