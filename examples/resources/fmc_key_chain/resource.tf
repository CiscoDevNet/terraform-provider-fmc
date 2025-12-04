resource "fmc_key_chain" "example" {
  name        = "my_key_chain"
  description = "My Host object"
  keys = [
    {
      id                       = 1
      key                      = "my_secret_key"
      accept_lifetime_start    = "2025-08-25T12:14:23"
      accept_lifetime_end_type = "DATETIME"
      accept_lifetime_end      = "2026-08-25T12:14:23"
      send_lifetime_start      = "2025-08-25T12:14:23"
      send_lifetime_end_type   = "DATETIME"
      send_lifetime_end        = "2026-08-25T12:14:23"
    }
  ]
}
