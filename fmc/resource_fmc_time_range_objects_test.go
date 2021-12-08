package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcTimeRangeObjectBasic(t *testing.T) {
	name := "test_range_obj"
	description := "Range test"
	startDate := "2019-09-19T15:53"
	endDate := "2019-09-21T17:53"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcTimeRangeObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcTimeRangeObjectConfigBasic(name, startDate, endDate, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcTimeRangeObjectExists("fmc_time_range_object.test"),
				),
			},
		},
	})
}

func testAccCheckFmcTimeRangeObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_time_range_object" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcTimeRangeObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcTimeRangeObjectConfigBasic(name, startDate, endDate, description string) string {
	return fmt.Sprintf(`
		resource "fmc_time_range_object" "test" {
		  name        		     = "%s"
		  description 		     = "%s"
		  effective_start_date = "%s"
		  effective_end_date   = "%s"
		  
		  recurrence {
			recurrence_type  = "DAILY_INTERVAL"
			days             = ["MON", "TUE"]
			daily_start_time = "09:00"
			daily_end_time   = "11:00"
		  }
		
		  recurrence {
			recurrence_type = "RANGE"
		
			start_time = "09:00"
			start_day  = "MON" 
			end_time   = "11:00"
			end_day    = "FRI"
		  }
		}  
    `, name, description, startDate, endDate)
}

func testAccCheckFmcTimeRangeObjectExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		return nil
	}
}
