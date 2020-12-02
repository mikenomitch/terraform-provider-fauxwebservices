package fws

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-fauxwebservices/client"
)

func resourceBucket() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBucketCreate,
		ReadContext:   resourceBucketRead,
		UpdateContext: resourceBucketUpdate,
		DeleteContext: resourceBucketDelete,
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

func resourceBucketCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	bucket := new(client.Bucket)
	bucket.Name = name

	o, err := c.CreateBucket(bucket)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.ID)

	resourceBucketRead(ctx, d, m)

	return diags
}

func resourceBucketRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucketID := d.Id()

	bucket, err := c.GetBucket(bucketID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", bucket.Name); err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceBucketUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	bucketID := d.Id()

	if d.HasChange("name") {
		name := d.Get("name").(string)

		bucket := new(client.Bucket)
		bucket.ID = bucketID
		bucket.Name = name

		_, err := c.UpdateBucket(bucket)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceBucketRead(ctx, d, m)
}

func resourceBucketDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucketID := d.Id()

	err := c.DeleteBucket(bucketID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
