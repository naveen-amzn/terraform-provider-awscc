---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_comprehend_flywheel Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::Comprehend::Flywheel resource creates an Amazon Comprehend Flywheel that enables customer to train their model.
---

# awscc_comprehend_flywheel (Resource)

The AWS::Comprehend::Flywheel resource creates an Amazon Comprehend Flywheel that enables customer to train their model.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_access_role_arn` (String)
- `data_lake_s3_uri` (String)
- `flywheel_name` (String)

### Optional

- `active_model_arn` (String)
- `data_security_config` (Attributes) (see [below for nested schema](#nestedatt--data_security_config))
- `model_type` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `task_config` (Attributes) (see [below for nested schema](#nestedatt--task_config))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--data_security_config"></a>
### Nested Schema for `data_security_config`

Optional:

- `data_lake_kms_key_id` (String)
- `model_kms_key_id` (String)
- `volume_kms_key_id` (String)
- `vpc_config` (Attributes) (see [below for nested schema](#nestedatt--data_security_config--vpc_config))

<a id="nestedatt--data_security_config--vpc_config"></a>
### Nested Schema for `data_security_config.vpc_config`

Required:

- `security_group_ids` (Set of String)
- `subnets` (Set of String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--task_config"></a>
### Nested Schema for `task_config`

Required:

- `language_code` (String)

Optional:

- `document_classification_config` (Attributes) (see [below for nested schema](#nestedatt--task_config--document_classification_config))
- `entity_recognition_config` (Attributes) (see [below for nested schema](#nestedatt--task_config--entity_recognition_config))

<a id="nestedatt--task_config--document_classification_config"></a>
### Nested Schema for `task_config.document_classification_config`

Required:

- `mode` (String)

Optional:

- `labels` (Set of String)


<a id="nestedatt--task_config--entity_recognition_config"></a>
### Nested Schema for `task_config.entity_recognition_config`

Optional:

- `entity_types` (Attributes Set) (see [below for nested schema](#nestedatt--task_config--entity_recognition_config--entity_types))

<a id="nestedatt--task_config--entity_recognition_config--entity_types"></a>
### Nested Schema for `task_config.entity_recognition_config.entity_types`

Required:

- `type` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_comprehend_flywheel.example <resource ID>
```