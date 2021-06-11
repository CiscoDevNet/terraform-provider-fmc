package fmc

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testAccProviders map[string]*schema.Provider
	testAccProvider  *schema.Provider
)

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func TestAccPreCheck(t *testing.T) {
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
}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"fmc": testAccProvider,
	}
}

func TestAccProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestAccProviderConfigure(t *testing.T) {
}
