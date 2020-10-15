package gmail

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]*schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"gmail": testAccProvider,
	}

}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("GOOGLE_CLIENT_CREDENTIALS") == "" {
		t.Fatal("GOOGLE_CLIENT_CREDENTIALS must be set for testing")
	}

	if os.Getenv("GOOGLE_TOKEN") == "" {
		t.Fatal("GOOGLE_TOKEN must be set for testing")
	}
}
