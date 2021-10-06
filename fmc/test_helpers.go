package fmc

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"math/rand"
	"time"
)

// generates random string with fixed size
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}


//checks if resource with given path (combination of resource type and name) exists, and if properties are set
func testAccCheckResource(n string, properties map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		if properties != nil {
			for key, value := range properties {
				if rs.Primary.Attributes[key] != value {
					return fmt.Errorf("attribute mismatch for key: %s. Expected: %s, got: %s", key, value, rs.Primary.Attributes[key])
				}
			}
		}

		//
		return nil
	}
}