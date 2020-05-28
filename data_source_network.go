package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePreprodNetworkRead,

		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"az": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePreprodNetworkRead(d *schema.ResourceData, meta interface{}) error {
	subnets := make([]interface{}, 0)

	namespace := d.Get("namespace").(string)

	switch namespace {
	case "production":
		subnets = append(subnets, map[string]interface{}{
			"id": "ARN::subnet:production",
			"az": "c",
		})

	default:
		subnets = append(subnets, map[string]interface{}{
			"id": "ARN::subnet:default",
			"az": "c",
		})
	}

	if err := d.Set("private_subnets", subnets); err != nil {
		return fmt.Errorf("Error setting subnets: %s", err)
	}

	if err := d.Set("vpc", "ARN:VPC"+namespace); err != nil {
		return fmt.Errorf("Error setting vpc: %s", err)
	}

	d.SetId(time.Now().UTC().String())
	return nil
}
