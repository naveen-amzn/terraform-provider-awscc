---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_volume Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::Volume
---

# awscc_ec2_volume (Resource)

Resource Type definition for AWS::EC2::Volume



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `availability_zone` (String)

### Optional

- `auto_enable_io` (Boolean)
- `encrypted` (Boolean)
- `iops` (Number)
- `kms_key_id` (String)
- `multi_attach_enabled` (Boolean)
- `outpost_arn` (String)
- `size` (Number)
- `snapshot_id` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `throughput` (Number)
- `volume_type` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `volume_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_volume.example <resource ID>
```