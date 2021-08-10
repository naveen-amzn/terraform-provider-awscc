// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

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
	registry.AddResourceTypeFactory("aws_datasync_location_object_storage", locationObjectStorageResourceType)
}

// locationObjectStorageResourceType returns the Terraform aws_datasync_location_object_storage resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationObjectStorage resource type.
func locationObjectStorageResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"access_key": {
			// Property: AccessKey
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			     "maxLength": 200,
			     "minLength": 8,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			Type:        types.StringType,
			Optional:    true,
		},
		"agent_arns": {
			// Property: AgentArns
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			     "insertionOrder": false,
			     "items": {
			       "maxLength": 128,
			       "pattern": "",
			       "type": "string"
			     },
			     "maxItems": 4,
			     "minItems": 1,
			     "type": "array"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			// Multiset.
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
		},
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the bucket on the self-managed object storage server.",
			     "maxLength": 63,
			     "minLength": 3,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the bucket on the self-managed object storage server.",
			Type:        types.StringType,
			Required:    true,
			// BucketName is a force-new attribute.
			// BucketName is a write-only attribute.
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the location that is created.",
			     "maxLength": 128,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The URL of the object storage location that was described.",
			     "maxLength": 4356,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The URL of the object storage location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"secret_key": {
			// Property: SecretKey
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			     "maxLength": 200,
			     "minLength": 8,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			Type:        types.StringType,
			Optional:    true,
			// SecretKey is a write-only attribute.
		},
		"server_hostname": {
			// Property: ServerHostname
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			     "maxLength": 255,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			Type:        types.StringType,
			Required:    true,
			// ServerHostname is a force-new attribute.
			// ServerHostname is a write-only attribute.
		},
		"server_port": {
			// Property: ServerPort
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The port that your self-managed server accepts inbound network traffic on.",
			     "type": "integer"
			   }
			*/
			Description: "The port that your self-managed server accepts inbound network traffic on.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"server_protocol": {
			// Property: ServerProtocol
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The protocol that the object storage server uses to communicate.",
			     "enum": [
			       "HTTPS",
			       "HTTP"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The protocol that the object storage server uses to communicate.",
			Type:        types.StringType,
			Optional:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The subdirectory in the self-managed object storage server that is used to read data from.",
			     "maxLength": 4096,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The subdirectory in the self-managed object storage server that is used to read data from.",
			Type:        types.StringType,
			Optional:    true,
			// Subdirectory is a write-only attribute.
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
			           "description": "The key for an AWS resource tag.",
			           "maxLength": 256,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for an AWS resource tag.",
			           "maxLength": 256,
			           "minLength": 1,
			           "pattern": "",
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
			     "maxItems": 50,
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
						     "description": "The key for an AWS resource tag.",
						     "maxLength": 256,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for an AWS resource tag.",
						     "maxLength": 256,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
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
		Description: "Resource schema for AWS::DataSync::LocationObjectStorage.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationObjectStorage").WithTerraformTypeName("aws_datasync_location_object_storage").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SecretKey",
		"/properties/BucketName",
		"/properties/ServerHostname",
		"/properties/Subdirectory",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_datasync_location_object_storage", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
