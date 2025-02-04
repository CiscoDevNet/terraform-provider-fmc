resource "fmc_device_bfd" "example" {
  device_id              = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  hop_type               = "SINGLE_HOP"
  bfd_template_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  interface_logical_name = "outside"
  interface_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
