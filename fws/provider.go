package fws

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-fauxwebservices/client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FWS_HOST", nil),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FWS_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fauxwebservices_server":   resourceServer(),
			"fauxwebservices_bucket":   resourceBucket(),
			"fauxwebservices_database": resourceDatabase(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	t := &token

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := client.NewClient(host, t)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create FWS client",
			Detail:   "Unable to create anonymous FWS client",
		})
		return nil, diags
	}

	return c, diags
}
