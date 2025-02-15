---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_logs_destination Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.
---

# awscc_logs_destination (Resource)

The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_name` (String) The name of the destination resource
- `role_arn` (String) The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
- `target_arn` (String) The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)

### Optional

- `destination_policy` (String) An IAM policy document that governs which AWS accounts can create subscription filters against this destination.

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_logs_destination.example <resource ID>
```
