package gmail

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/gmail/v1"
)

func resourceFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilterCreate,
		Read:   resourceFilterRead,
		Schema: map[string]*schema.Schema{
			"criteria": {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: `Matching criteria for the filter`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_chats": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Whether the response should exclude chats",
						},
						"from": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The sender's display name or email address",
						},
						"has_attachment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Whether the message has any attachment",
						},
						"query": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Only return messages matching the specified query. Supports the same query format as the Gmail search box.",
						},
						"negated_query": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Only return messages not matching the specified query. Supports the same query format as the Gmail search box.",
						},
						"size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The size of the entire RFC822 message in bytes, including all headers and attachments",
						},
						"size_comparison": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "How the message size in bytes should be in relation to the size field",
							ValidateFunc: validation.StringInSlice([]string{
								"larger",
								"smaller",
								"unspecified",
							}, false),
						},
						"subject": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed.",
						},
						"to": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: `The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You can use simply the local part of the email address. For example, "example" and "example@" both match "example@gmail.com". This field is case-insensitive.`,
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: `Matching criteria for the filter`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_label_ids": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "List of labels to add to the message",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"remove_label_ids": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "List of labels to remove from the message",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"forward": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Email address that the message should be forwarded to",
						},
					},
				},
			},
		},
	}
}

func resourceFilterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gmail.Service)

	filter := &gmail.Filter{
		Criteria: &gmail.FilterCriteria{},
		Action:   &gmail.FilterAction{},
	}

	if v, ok := d.GetOk("criteria.0.exclude_chats"); ok {
		filter.Criteria.ExcludeChats = v.(bool)
	}
	filter.Criteria.From = d.Get("criteria.0.from").(string)
	if v, ok := d.GetOk("criteria.0.has_attachment"); ok {
		filter.Criteria.HasAttachment = v.(bool)
	}
	filter.Criteria.Query = d.Get("criteria.0.query").(string)
	filter.Criteria.NegatedQuery = d.Get("criteria.0.negated_query").(string)
	filter.Criteria.Size = d.Get("criteria.0.size").(int64)
	filter.Criteria.SizeComparison = d.Get("criteria.0.size_comparison").(string)
	filter.Criteria.Subject = d.Get("criteria.0.subject").(string)
	filter.Criteria.To = d.Get("criteria.0.to").(string)

	filter, err := client.Users.Settings.Filters.Create("me", filter).Do()
	if err != nil {
		return fmt.Errorf("error creating filter: %v", err)
	}

	d.SetId(filter.Id)

	return resourceFilterRead(d, m)
}

func resourceFilterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gmail.Service)

	filter, err := client.Users.Settings.Filters.Get("me", d.Id()).Do()
	if err != nil {
		return fmt.Errorf("error reading filter: %v", err)
	}

	d.Set("criteria.0.exclude_chats", filter.Criteria.ExcludeChats)
	d.Set("criteria.0.from", filter.Criteria.From)
	d.Set("criteria.0.has_attachment", filter.Criteria.HasAttachment)
	d.Set("criteria.0.query", filter.Criteria.Query)
	d.Set("criteria.0.negated_query", filter.Criteria.NegatedQuery)
	d.Set("criteria.0.size", filter.Criteria.Size)
	d.Set("criteria.0.size_comparison", filter.Criteria.SizeComparison)
	d.Set("criteria.0.subject", filter.Criteria.Subject)
	d.Set("criteria.0.to", filter.Criteria.To)

	d.Set("action.0.add_label_ids", filter.Action.AddLabelIds)
	d.Set("action.0.remove_label_ids", filter.Action.RemoveLabelIds)
	d.Set("forward", filter.Action.Forward)

	return nil
}
