package ultradns

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
		"ultradns": testAccProvider,
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
	if v := os.Getenv("ULTRADNS_USERNAME"); v == "" {
		t.Fatal("ULTRADNS_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("ULTRADNS_PASSWORD"); v == "" {
		t.Fatal("ULTRADNS_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("ULTRADNS_DOMAIN"); v == "" {
		t.Fatal("ULTRADNS_DOMAIN must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}
