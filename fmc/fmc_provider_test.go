package fmc

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"fmc": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FMC_HOST"); v == "" {
		t.Fatal("FMC_HOST must be set for acceptance tests")
	}
	if v := os.Getenv("FMC_USERNAME"); v == "" {
		t.Fatal("FMC_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("FMC_PASSWORD"); v == "" {
		t.Fatal("FMC_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("FMC_INSECURE_SKIP_VERIFY"); v == "" {
		t.Fatal("FMC_INSECURE_SKIP_VERIFY must be set for acceptance tests")
	}
	if v := os.Getenv("FTD_HOST"); v == "" {
		t.Fatal("FTD_HOST must be set for acceptance tests")
	}
}
