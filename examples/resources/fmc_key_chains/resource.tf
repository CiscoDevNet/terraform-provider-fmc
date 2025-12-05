resource "fmc_key_chains" "example" {
  items = {
    my_key_chains = {
      description = "My Host object"
      keys = [
        {
          id                       = 1
          key                      = "my_secret_key"
          accept_lifetime_start    = "2025-08-25T12:14:23"
          accept_lifetime_end_type = "DATETIME"
          accept_lifetime_end      = "2026-08-25T12:14:23"
          send_lifetime_start      = "2025-08-25T12:14:23"
          send_lifetime_end_type   = "DURATION"
          send_lifetime_end        = "172800"
        }
      ]
    }
  }
}
