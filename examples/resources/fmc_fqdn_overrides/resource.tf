resource "fmc_fqdn_overrides" "example" {
  parent_name = "my_fqdn"
  parent_id   = "12345678-90ab-cdef-1234-567890abcdef"
  overrides = [
    {
      target_id   = "12345678-90ab-cdef-1234-567890abcdef"
      target_type = "Device"
      description = "My FQDN object"
      fqdn        = "sub.example.com"
    }
  ]
}
