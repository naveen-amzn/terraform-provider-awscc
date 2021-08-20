// Code generated by generators/resource/main.go; DO NOT EDIT.

package wafv2

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_wafv2_regex_pattern_set", regexPatternSetResourceType)
}

// regexPatternSetResourceType returns the Terraform awscc_wafv2_regex_pattern_set resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::WAFv2::RegexPatternSet resource type.
func regexPatternSetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the WAF entity.",
			//   "type": "string"
			// }
			Description: "ARN of the WAF entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the entity.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Description of the entity.",
			Type:        types.StringType,
			Optional:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the RegexPatternSet",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Id of the RegexPatternSet",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the RegexPatternSet.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the RegexPatternSet.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"regular_expression_list": {
			// Property: RegularExpressionList
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			// {
			//   "description": "Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.",
			//   "enum": [
			//     "CLOUDFRONT",
			//     "REGIONAL"
			//   ],
			//   "type": "string"
			// }
			Description: "Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.",
			Type:        types.StringType,
			Required:    true,
			// Scope is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Contains a list of Regular expressions based on the provided inputs. RegexPatternSet can be used with other WAF entities with RegexPatternSetReferenceStatement to perform other actions .",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::RegexPatternSet").WithTerraformTypeName("awscc_wafv2_regex_pattern_set").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_wafv2_regex_pattern_set", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}