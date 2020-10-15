package gmail

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"google.golang.org/api/gmail/v1"
)

func TestAccGmailFilter_basic(t *testing.T) {
	var filter gmail.Filter

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGmailFilterDestroy,
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
					testAccCheckGmailFilterExists("gmail_filter.foo", &filter),
					func(s *terraform.State) error {
						if got := filter.Criteria.From; got != "foo@example.com" {
							return fmt.Errorf("expected filter from foo@example.com, got %s", got)
						}

						if got := filter.Action.RemoveLabelIds; !reflect.DeepEqual([]string{"INBOX"}, got) {
							return fmt.Errorf("expected INBOX in remove labels action, got %v", got)
						}

						return nil
					},
					resource.TestCheckResourceAttr("gmail_filter.foo", "criteria.0.from", "foo@example.com"),
				),
			},
		},
	})
}

func testAccCheckGmailFilterExists(path string, filter *gmail.Filter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[path]
		if !ok {
			return fmt.Errorf("not found: %s", path)
		}

		client, err := testAccProvider.Meta().(*Config).NewService(context.Background())
		if err != nil {
			return err
		}

		out, err := client.Users.Settings.Filters.Get("me", rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		*filter = *out
		return nil
	}
}

func testAccCheckGmailFilterDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(*Config).NewService(context.Background())
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gmail_filter" {
			continue
		}

		if _, err := client.Users.Settings.Filters.Get("me", rs.Primary.ID).Do(); err == nil {
			return fmt.Errorf("filter still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
