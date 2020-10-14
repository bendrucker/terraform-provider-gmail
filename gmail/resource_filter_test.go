package gmail

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGmailFilter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "gmail_filter" "foo" {
	criteria {
		from = "foo@example.com"
	}

	action {
		remove_label_ids = ["INBOX"]
	}
}
								`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("gmail_filter.foo", "criteria.0.from", "foo@example.com"),
				),
			},
		},
	})
}
