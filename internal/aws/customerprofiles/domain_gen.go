// Code generated by generators/resource/main.go; DO NOT EDIT.

package customerprofiles

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
	registry.AddResourceTypeFactory("aws_customerprofiles_domain", domainResourceType)
}

// domainResourceType returns the Terraform aws_customerprofiles_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CustomerProfiles::Domain resource type.
func domainResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The time of this integration got created",
			     "type": "string"
			   }
			*/
			Description: "The time of this integration got created",
			Type:        types.StringType,
			Computed:    true,
		},
		"dead_letter_queue_url": {
			// Property: DeadLetterQueueUrl
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The URL of the SQS dead letter queue",
			     "maxLength": 255,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "The URL of the SQS dead letter queue",
			Type:        types.StringType,
			Optional:    true,
		},
		"default_encryption_key": {
			// Property: DefaultEncryptionKey
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The default encryption key",
			     "maxLength": 255,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "The default encryption key",
			Type:        types.StringType,
			Optional:    true,
		},
		"default_expiration_days": {
			// Property: DefaultExpirationDays
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The default number of days until the data within the domain expires.",
			     "type": "integer"
			   }
			*/
			Description: "The default number of days until the data within the domain expires.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The unique name of the domain.",
			     "maxLength": 64,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The unique name of the domain.",
			Type:        types.StringType,
			Required:    true,
			// DomainName is a force-new attribute.
		},
		"last_updated_at": {
			// Property: LastUpdatedAt
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The time of this integration got last updated at",
			     "type": "string"
			   }
			*/
			Description: "The time of this integration got last updated at",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags (keys and values) associated with the domain",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
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
			     "maxItems": 50,
			     "minItems": 0,
			     "type": "array"
			   }
			*/
			Description: "The tags (keys and values) associated with the domain",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 0,
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
		Description: "A domain defined for 3rd party data source in Profile Service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::Domain").WithTerraformTypeName("aws_customerprofiles_domain").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_customerprofiles_domain", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
