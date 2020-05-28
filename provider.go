package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NAMESPACE", nil),
				Description: "a namespace for config?",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sergiobuj_network": dataSourceNetwork(),
		},
	}
}
