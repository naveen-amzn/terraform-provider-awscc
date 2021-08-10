// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

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
	registry.AddResourceTypeFactory("aws_ecs_cluster_capacity_provider_associations", clusterCapacityProviderAssociationsResourceType)
}

// clusterCapacityProviderAssociationsResourceType returns the Terraform aws_ecs_cluster_capacity_provider_associations resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECS::ClusterCapacityProviderAssociations resource type.
func clusterCapacityProviderAssociationsResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"capacity_providers": {
			// Property: CapacityProviders
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "List of capacity providers to associate with the cluster",
			     "items": {
			       "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
			       "$ref": "#/definitions/CapacityProvider",
			       "type": "string"
			     },
			     "$ref": "#/definitions/CapacityProviders",
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "List of capacity providers to associate with the cluster",
			// Ordered set.
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
		},
		"cluster": {
			// Property: Cluster
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the cluster",
			     "maxLength": 2048,
			     "minLength": 1,
			     "$ref": "#/definitions/Cluster",
			     "type": "string"
			   }
			*/
			Description: "The name of the cluster",
			Type:        types.StringType,
			Required:    true,
			// Cluster is a force-new attribute.
		},
		"default_capacity_provider_strategy": {
			// Property: DefaultCapacityProviderStrategy
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "List of capacity providers to associate with the cluster",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Base": {
			           "type": "integer"
			         },
			         "CapacityProvider": {
			           "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
			           "$ref": "#/definitions/CapacityProvider",
			           "type": "string"
			         },
			         "Weight": {
			           "type": "integer"
			         }
			       },
			       "$ref": "#/definitions/CapacityProviderStrategy",
			       "required": [
			         "CapacityProvider"
			       ],
			       "type": "object"
			     },
			     "$ref": "#/definitions/DefaultCapacityProviderStrategy",
			     "type": "array"
			   }
			*/
			Description: "List of capacity providers to associate with the cluster",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"base": {
						// Property: Base
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"capacity_provider": {
						// Property: CapacityProvider
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
						     "$ref": "#/definitions/CapacityProvider",
						     "type": "string"
						   }
						*/
						Description: "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
						Type:        types.StringType,
						Required:    true,
					},
					"weight": {
						// Property: Weight
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
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
		Description: "Associate a set of ECS Capacity Providers with a specified ECS Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::ClusterCapacityProviderAssociations").WithTerraformTypeName("aws_ecs_cluster_capacity_provider_associations").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ecs_cluster_capacity_provider_associations", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
