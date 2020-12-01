---
page_title: "Provider: Faux Web Services"
subcategory: ""
description: |-
  Terraform provider for interacting with Faux Web Services API.
---

# Faux Web Services Provider

TODO: Add description.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "fauxwebservices" {
  host = "app.terraform.io"
  token = "SOMETOKEN"
}
```

## Schema

### Optional

- **host** (String, Optional) FWS API address (defaults to `app.terraform.io`)
- **token** (String) Token used to authenticate into FWS API (defaults to `FWS_TOKEN` env var)
