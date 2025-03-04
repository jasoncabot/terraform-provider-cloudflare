---
page_title: "cloudflare_tunnel_virtual_network Data Source - Cloudflare"
subcategory: ""
description: |-
  Use this datasource to lookup a tunnel virtual network in an account.
---

# cloudflare_tunnel_virtual_network (Data Source)

Use this datasource to lookup a tunnel virtual network in an account.

## Example Usage

```terraform
data "cloudflare_tunnel_virtual_network" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `name` (String) The Virtual Network Name.

### Read-Only

- `comment` (String) The Virtual Network Comment.
- `id` (String) The ID of this resource.
- `is_default` (Boolean) If true, only include deleted virtual networks. If false, exclude deleted virtual networks. If empty, all virtual networks will be included.


