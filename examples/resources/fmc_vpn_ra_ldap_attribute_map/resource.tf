resource "fmc_vpn_ra_ldap_attribute_map" "example" {
  vpn_ra_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  realms = [
    {
      realm_ad_ldap_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      attribute_maps = [
        {
          ldap_attribute_name  = "access_time"
          cisco_attribute_name = "Access-Hours"
          value_maps = [
            {
              ldap_attribute_value  = "ldap_value_example"
              cisco_attribute_value = "cisco_value_example"
            }
          ]
        }
      ]
    }
  ]
}
