---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_security_key Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::SecurityKey
---

# awscc_connect_security_key (Data Source)

Data Source schema for AWS::Connect::SecurityKey



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `association_id` (String) An associationID is automatically generated when a storage config is associated with an instance
- `instance_id` (String) Amazon Connect instance identifier
- `key` (String) A valid security key in PEM format.


