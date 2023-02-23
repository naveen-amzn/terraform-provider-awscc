// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iam_role", roleDataSource)
}

// roleDataSource returns the Terraform awscc_iam_role data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::Role resource.
func roleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the role.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssumeRolePolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The trust policy that is associated with this role.",
		//	  "type": "string"
		//	}
		"assume_role_policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The trust policy that is associated with this role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the role that you provide.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the role that you provide.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ManagedPolicyArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. ",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"managed_policy_arns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxSessionDuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. ",
		//	  "type": "integer"
		//	}
		"max_session_duration": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The path to the role.",
		//	  "type": "string"
		//	}
		"path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path to the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PermissionsBoundary
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the policy used to set the permissions boundary for the role.",
		//	  "type": "string"
		//	}
		"permissions_boundary": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the policy used to set the permissions boundary for the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Policies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Adds or updates an inline policy document that is embedded in the specified IAM role. ",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The inline policy document that is embedded in the specified IAM role.",
		//	    "properties": {
		//	      "PolicyDocument": {
		//	        "description": "The policy document.",
		//	        "type": "string"
		//	      },
		//	      "PolicyName": {
		//	        "description": "The friendly name (not ARN) identifying the policy.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "PolicyName",
		//	      "PolicyDocument"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"policies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PolicyDocument
					"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The policy document.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: PolicyName
					"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The friendly name (not ARN) identifying the policy.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Adds or updates an inline policy document that is embedded in the specified IAM role. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stable and unique string identifying the role.",
		//	  "type": "string"
		//	}
		"role_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The stable and unique string identifying the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the IAM role, up to 64 characters in length.",
		//	  "type": "string"
		//	}
		"role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the IAM role, up to 64 characters in length.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags that are attached to the role.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags that are attached to the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::Role",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::Role").WithTerraformTypeName("awscc_iam_role")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"assume_role_policy_document": "AssumeRolePolicyDocument",
		"description":                 "Description",
		"key":                         "Key",
		"managed_policy_arns":         "ManagedPolicyArns",
		"max_session_duration":        "MaxSessionDuration",
		"path":                        "Path",
		"permissions_boundary":        "PermissionsBoundary",
		"policies":                    "Policies",
		"policy_document":             "PolicyDocument",
		"policy_name":                 "PolicyName",
		"role_id":                     "RoleId",
		"role_name":                   "RoleName",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
