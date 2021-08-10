// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

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
	registry.AddResourceTypeFactory("aws_nimblestudio_studio_component", studioComponentResourceType)
}

// studioComponentResourceType returns the Terraform aws_nimblestudio_studio_component resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NimbleStudio::StudioComponent resource type.
func studioComponentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "ActiveDirectoryConfiguration": {
			         "additionalProperties": false,
			         "properties": {
			           "ComputerAttributes": {
			             "items": {
			               "additionalProperties": false,
			               "properties": {
			                 "Name": {
			                   "type": "string"
			                 },
			                 "Value": {
			                   "type": "string"
			                 }
			               },
			               "$ref": "#/definitions/ActiveDirectoryComputerAttribute",
			               "type": "object"
			             },
			             "$ref": "#/definitions/ActiveDirectoryComputerAttributeList",
			             "type": "array"
			           },
			           "DirectoryId": {
			             "type": "string"
			           },
			           "OrganizationalUnitDistinguishedName": {
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/ActiveDirectoryConfiguration",
			         "type": "object"
			       },
			       "ComputeFarmConfiguration": {
			         "additionalProperties": false,
			         "properties": {
			           "ActiveDirectoryUser": {
			             "type": "string"
			           },
			           "Endpoint": {
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/ComputeFarmConfiguration",
			         "type": "object"
			       },
			       "LicenseServiceConfiguration": {
			         "additionalProperties": false,
			         "properties": {
			           "Endpoint": {
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/LicenseServiceConfiguration",
			         "type": "object"
			       },
			       "SharedFileSystemConfiguration": {
			         "additionalProperties": false,
			         "properties": {
			           "Endpoint": {
			             "type": "string"
			           },
			           "FileSystemId": {
			             "type": "string"
			           },
			           "LinuxMountPoint": {
			             "type": "string"
			           },
			           "ShareName": {
			             "type": "string"
			           },
			           "WindowsMountDrive": {
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/SharedFileSystemConfiguration",
			         "type": "object"
			       }
			     },
			     "$ref": "#/definitions/StudioComponentConfiguration",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"active_directory_configuration": {
						// Property: ActiveDirectoryConfiguration
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "ComputerAttributes": {
						         "items": {
						           "additionalProperties": false,
						           "properties": {
						             "Name": {
						               "type": "string"
						             },
						             "Value": {
						               "type": "string"
						             }
						           },
						           "$ref": "#/definitions/ActiveDirectoryComputerAttribute",
						           "type": "object"
						         },
						         "$ref": "#/definitions/ActiveDirectoryComputerAttributeList",
						         "type": "array"
						       },
						       "DirectoryId": {
						         "type": "string"
						       },
						       "OrganizationalUnitDistinguishedName": {
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/ActiveDirectoryConfiguration",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"computer_attributes": {
									// Property: ComputerAttributes
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "additionalProperties": false,
									       "properties": {
									         "Name": {
									           "type": "string"
									         },
									         "Value": {
									           "type": "string"
									         }
									       },
									       "$ref": "#/definitions/ActiveDirectoryComputerAttribute",
									       "type": "object"
									     },
									     "$ref": "#/definitions/ActiveDirectoryComputerAttributeList",
									     "type": "array"
									   }
									*/
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"name": {
												// Property: Name
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"value": {
												// Property: Value
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"directory_id": {
									// Property: DirectoryId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"organizational_unit_distinguished_name": {
									// Property: OrganizationalUnitDistinguishedName
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"compute_farm_configuration": {
						// Property: ComputeFarmConfiguration
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "ActiveDirectoryUser": {
						         "type": "string"
						       },
						       "Endpoint": {
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/ComputeFarmConfiguration",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"active_directory_user": {
									// Property: ActiveDirectoryUser
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"endpoint": {
									// Property: Endpoint
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"license_service_configuration": {
						// Property: LicenseServiceConfiguration
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Endpoint": {
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/LicenseServiceConfiguration",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"endpoint": {
									// Property: Endpoint
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"shared_file_system_configuration": {
						// Property: SharedFileSystemConfiguration
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Endpoint": {
						         "type": "string"
						       },
						       "FileSystemId": {
						         "type": "string"
						       },
						       "LinuxMountPoint": {
						         "type": "string"
						       },
						       "ShareName": {
						         "type": "string"
						       },
						       "WindowsMountDrive": {
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/SharedFileSystemConfiguration",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"endpoint": {
									// Property: Endpoint
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"file_system_id": {
									// Property: FileSystemId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"linux_mount_point": {
									// Property: LinuxMountPoint
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"share_name": {
									// Property: ShareName
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"windows_mount_drive": {
									// Property: WindowsMountDrive
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"ec_2_security_group_ids": {
			// Property: Ec2SecurityGroupIds
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "$ref": "#/definitions/SecurityGroupId",
			       "type": "string"
			     },
			     "$ref": "#/definitions/StudioComponentSecurityGroupIdList",
			     "type": "array"
			   }
			*/
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"initialization_scripts": {
			// Property: InitializationScripts
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "LaunchProfileProtocolVersion": {
			           "type": "string"
			         },
			         "Platform": {
			           "type": "string"
			         },
			         "RunContext": {
			           "type": "string"
			         },
			         "Script": {
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/StudioComponentInitializationScript",
			       "type": "object"
			     },
			     "$ref": "#/definitions/StudioComponentInitializationScriptList",
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"launch_profile_protocol_version": {
						// Property: LaunchProfileProtocolVersion
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"platform": {
						// Property: Platform
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"run_context": {
						// Property: RunContext
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"script": {
						// Property: Script
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"script_parameters": {
			// Property: ScriptParameters
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/ScriptParameterKeyValue",
			       "type": "object"
			     },
			     "$ref": "#/definitions/StudioComponentScriptParameterKeyValueList",
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"studio_component_id": {
			// Property: StudioComponentId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// StudioId is a force-new attribute.
		},
		"subtype": {
			// Property: Subtype
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// Subtype is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "patternProperties": {
			       "": {
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::NimbleStudio::StudioComponent.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StudioComponent").WithTerraformTypeName("aws_nimblestudio_studio_component").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_nimblestudio_studio_component", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
