resource "fmc_health_policy" "example" {
  name              = "my_health_policy"
  description       = "My health policy"
  policy_type       = "DevicePolicy"
  is_default_policy = false
  health_modules = [
    {
      name    = "hm_asp_drop"
      enabled = true
      type    = "FTD"
      alert_configs = [
        {
          name    = "nat-xlate-failed"
          enabled = true
        }
      ]
    }
  ]
  health_module_run_time_interval = 5
  metric_collection_interval      = 1
}
