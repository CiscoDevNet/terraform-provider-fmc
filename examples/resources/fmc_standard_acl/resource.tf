resource "fmc_standard_acl" "example" {
  name        = "stdacl1"
  description = "My standard ACL"
  entries = [
    {
      action = "DENY"
      objects = [
        {
        }
      ]
      literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
    }
  ]
}
