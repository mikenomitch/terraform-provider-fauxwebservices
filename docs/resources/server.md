---
page_title: "order Resource - terraform-provider-hashicups"
subcategory: ""
description: |-
  The order resource allows you to configure a HashiCups order.
---

# Resource `hashicups_order`

-> Visit the [Perform CRUD operations with Providers](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers&utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) Learn tutorial for an interactive getting started experience.

The order resource allows you to configure a HashiCups order.

## Example Usage

```terraform
resource "fws_server" "demoserver" {
  name = "demo-server"
}
```

## Argument Reference

- `name` - (Required) The name of the server
