resource "fmc_vpn_ra_secure_client_customization" "example" {
  vpn_ra_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  gui_and_text_messages = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
  icons_and_images = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
  scripts = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
  binaries = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
  custom_installer_transforms = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
  localized_installer_transforms = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
}
