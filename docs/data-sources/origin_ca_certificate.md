---
page_title: "cloudflare_origin_ca_certificate Data Source - Cloudflare"
subcategory: ""
description: |-
  Use this data source to retrieve an existing origin ca certificate.
---

# cloudflare_origin_ca_certificate (Data Source)

Use this data source to retrieve an existing origin ca certificate.

## Example Usage

```terraform
data "cloudflare_origin_ca_certificate" "example" {
  id = "REPLACE_ME"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Origin CA Certificate unique identifier.

### Read-Only

- `certificate` (String) The Origin CA certificate.
- `expires_on` (String) The timestamp when the certificate will expire.
- `hostnames` (List of String) A list of hostnames or wildcard names bound to the certificate.
- `request_type` (String) The signature type desired on the certificate. Available values: `origin-rsa`, `origin-ecc`, `keyless-certificate`
- `revoked_at` (String) The timestamp when the certificate was revoked.


