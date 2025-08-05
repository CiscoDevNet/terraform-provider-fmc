resource "fmc_policy_assignment" "example" {
  policy_id               = "0050568A-2561-0ed3-0000-004294972836"
  policy_type             = "FTDNatPolicy"
  after_destroy_policy_id = "0050568A-2561-0ed3-0000-004294972836"
  targets = [
    {
      id   = "9862719c-8d5f-11ef-99a6-aef0794da1c1"
      type = "Device"
      name = "my_ftd_device"
    }
  ]
}
