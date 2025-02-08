resource "fmc_smart_license" "example" {
  registration_type   = "REGISTER"
  token               = "X2M3YmJlY..."
  force               = false
  retain_registration = false
}
