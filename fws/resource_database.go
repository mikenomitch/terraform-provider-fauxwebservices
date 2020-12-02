package fws

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-fauxwebservices/client"
)

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDatabaseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	database := new(client.Database)
	database.Name = name

	o, err := c.CreateDatabase(database)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.ID)

	resourceDatabaseRead(ctx, d, m)

	return diags
}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	databaseID := d.Id()

	database, err := c.GetDatabase(databaseID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", database.Name); err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	databaseID := d.Id()

	if d.HasChange("name") {
		name := d.Get("name").(string)

		database := new(client.Database)
		database.ID = databaseID
		database.Name = name

		_, err := c.UpdateDatabase(database)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceDatabaseRead(ctx, d, m)
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	databaseID := d.Id()

	err := c.DeleteDatabase(databaseID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
