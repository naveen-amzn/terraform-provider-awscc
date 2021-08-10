// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_cloudwatch_composite_alarm", compositeAlarmResourceType)
}

// compositeAlarmResourceType returns the Terraform aws_cloudwatch_composite_alarm resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudWatch::CompositeAlarm resource type.
func compositeAlarmResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"actions_enabled": {
			// Property: ActionsEnabled
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			     "type": "boolean"
			   }
			*/
			Description: "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"alarm_actions": {
			// Property: AlarmActions
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			     "items": {
			       "description": "Amazon Resource Name (ARN) of the action",
			       "maxLength": 1024,
			       "minLength": 1,
			       "type": "string"
			     },
			     "maxItems": 5,
			     "type": "array"
			   }
			*/
			Description: "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"alarm_description": {
			// Property: AlarmDescription
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The description of the alarm",
			     "maxLength": 1024,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "The description of the alarm",
			Type:        types.StringType,
			Optional:    true,
		},
		"alarm_name": {
			// Property: AlarmName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the Composite Alarm",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The name of the Composite Alarm",
			Type:        types.StringType,
			Required:    true,
			// AlarmName is a force-new attribute.
		},
		"alarm_rule": {
			// Property: AlarmRule
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			     "maxLength": 10240,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			Type:        types.StringType,
			Required:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Amazon Resource Name (ARN) of the alarm",
			     "maxLength": 1600,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Amazon Resource Name (ARN) of the alarm",
			Type:        types.StringType,
			Computed:    true,
		},
		"insufficient_data_actions": {
			// Property: InsufficientDataActions
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			     "items": {
			       "description": "Amazon Resource Name (ARN) of the action",
			       "maxLength": 1024,
			       "minLength": 1,
			       "type": "string"
			     },
			     "maxItems": 5,
			     "type": "array"
			   }
			*/
			Description: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"ok_actions": {
			// Property: OKActions
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			     "items": {
			       "description": "Amazon Resource Name (ARN) of the action",
			       "maxLength": 1024,
			       "minLength": 1,
			       "type": "string"
			     },
			     "maxItems": 5,
			     "type": "array"
			   }
			*/
			Description: "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::CompositeAlarm").WithTerraformTypeName("aws_cloudwatch_composite_alarm").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_cloudwatch_composite_alarm", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
