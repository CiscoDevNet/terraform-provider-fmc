resource "fmc_intrusion_rule" "example" {
  rule_data = "alert icmp any any -> any any ( sid:10000301; gid:2000; msg:\"CUSTOM RULE1\"; classtype:icmp-event; rev:1; )"
  rule_groups = [
    {
      id   = "12da34567890-1234-5678-90ab-cdef12345678"
      name = "Custom Rule Group 1"
    }
  ]
}
