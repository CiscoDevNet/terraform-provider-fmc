resource "fmc_ftd_platform_settings_time_synchronization" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ntp_mode                 = "SYNC_VIA_MGMT_CENTER_NTP"
}
