// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotcoredeviceadvisor

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_iotcoredeviceadvisor_suite_definition", suiteDefinitionResourceType)
}

// suiteDefinitionResourceType returns the Terraform aws_iotcoredeviceadvisor_suite_definition resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTCoreDeviceAdvisor::SuiteDefinition resource type.
func suiteDefinitionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"suite_definition_arn": {
			// Property: SuiteDefinitionArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource name for the suite definition.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource name for the suite definition.",
			Type:        types.StringType,
			Computed:    true,
		},
		"suite_definition_configuration": {
			// Property: SuiteDefinitionConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "DevicePermissionRoleArn": {
			         "description": "The device permission role arn of the test suite.",
			         "maxLength": 2048,
			         "minLength": 20,
			         "$ref": "#/definitions/DevicePermissionRoleArn",
			         "type": "string"
			       },
			       "Devices": {
			         "description": "The devices being tested in the test suite",
			         "items": {
			           "additionalProperties": false,
			           "properties": {
			             "CertificateArn": {
			               "maxLength": 2048,
			               "minLength": 20,
			               "type": "string"
			             },
			             "ThingArn": {
			               "maxLength": 2048,
			               "minLength": 20,
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/DeviceUnderTest",
			           "type": "object"
			         },
			         "maxItems": 2,
			         "minItems": 0,
			         "$ref": "#/definitions/Devices",
			         "type": "array"
			       },
			       "IntendedForQualification": {
			         "description": "Whether the tests are intended for qualification in a suite.",
			         "$ref": "#/definitions/IntendedForQualification",
			         "type": "boolean"
			       },
			       "RootGroup": {
			         "description": "The root group of the test suite.",
			         "maxLength": 2048,
			         "minLength": 1,
			         "$ref": "#/definitions/RootGroup",
			         "type": "string"
			       },
			       "SuiteDefinitionName": {
			         "description": "The Name of the suite definition.",
			         "maxLength": 256,
			         "minLength": 1,
			         "$ref": "#/definitions/SuiteDefinitionName",
			         "type": "string"
			       }
			     },
			     "required": [
			       "DevicePermissionRoleArn",
			       "RootGroup"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"device_permission_role_arn": {
						// Property: DevicePermissionRoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The device permission role arn of the test suite.",
						     "maxLength": 2048,
						     "minLength": 20,
						     "$ref": "#/definitions/DevicePermissionRoleArn",
						     "type": "string"
						   }
						*/
						Description: "The device permission role arn of the test suite.",
						Type:        types.StringType,
						Required:    true,
					},
					"devices": {
						// Property: Devices
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The devices being tested in the test suite",
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "CertificateArn": {
						           "maxLength": 2048,
						           "minLength": 20,
						           "type": "string"
						         },
						         "ThingArn": {
						           "maxLength": 2048,
						           "minLength": 20,
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/DeviceUnderTest",
						       "type": "object"
						     },
						     "maxItems": 2,
						     "minItems": 0,
						     "$ref": "#/definitions/Devices",
						     "type": "array"
						   }
						*/
						Description: "The devices being tested in the test suite",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"certificate_arn": {
									// Property: CertificateArn
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 2048,
									     "minLength": 20,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"thing_arn": {
									// Property: ThingArn
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 2048,
									     "minLength": 20,
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 0,
								MaxItems: 2,
							},
						),
						Optional: true,
					},
					"intended_for_qualification": {
						// Property: IntendedForQualification
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Whether the tests are intended for qualification in a suite.",
						     "$ref": "#/definitions/IntendedForQualification",
						     "type": "boolean"
						   }
						*/
						Description: "Whether the tests are intended for qualification in a suite.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"root_group": {
						// Property: RootGroup
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The root group of the test suite.",
						     "maxLength": 2048,
						     "minLength": 1,
						     "$ref": "#/definitions/RootGroup",
						     "type": "string"
						   }
						*/
						Description: "The root group of the test suite.",
						Type:        types.StringType,
						Required:    true,
					},
					"suite_definition_name": {
						// Property: SuiteDefinitionName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The Name of the suite definition.",
						     "maxLength": 256,
						     "minLength": 1,
						     "$ref": "#/definitions/SuiteDefinitionName",
						     "type": "string"
						   }
						*/
						Description: "The Name of the suite definition.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
		},
		"suite_definition_id": {
			// Property: SuiteDefinitionId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The unique identifier for the suite definition.",
			     "maxLength": 36,
			     "minLength": 12,
			     "type": "string"
			   }
			*/
			Description: "The unique identifier for the suite definition.",
			Type:        types.StringType,
			Computed:    true,
		},
		"suite_definition_version": {
			// Property: SuiteDefinitionVersion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The suite definition version of a test suite.",
			     "maxLength": 255,
			     "minLength": 2,
			     "type": "string"
			   }
			*/
			Description: "The suite definition version of a test suite.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An array of key-value pairs to apply to this resource.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			           "maxLength": 256,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTCoreDeviceAdvisor::SuiteDefinition").WithTerraformTypeName("aws_iotcoredeviceadvisor_suite_definition").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iotcoredeviceadvisor_suite_definition", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
