---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_robomaker_robot Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS::RoboMaker::Robot resource creates an AWS RoboMaker fleet.
---

# awscc_robomaker_robot (Resource)

AWS::RoboMaker::Robot resource creates an AWS RoboMaker fleet.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **architecture** (String) The target architecture of the robot.
- **greengrass_group_id** (String) The Greengrass group id.

### Optional

- **fleet** (String) The Amazon Resource Name (ARN) of the fleet.
- **name** (String) The name for the robot.
- **tags** (Map of String) A key-value pair to associate with a resource.

### Read-Only

- **arn** (String)
- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_robomaker_robot.example <resource ID>
```