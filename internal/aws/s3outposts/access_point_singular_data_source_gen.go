// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3outposts_access_point", accessPointDataSource)
}

// accessPointDataSource returns the Terraform awscc_s3outposts_access_point data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3Outposts::AccessPoint resource.
func accessPointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified AccessPoint.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/accesspoint\\/[^:]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/bucket\\/[^:]+$",
		//	  "type": "string"
		//	}
		"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the AccessPoint.",
		//	  "maxLength": 50,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9]([a-z0-9\\\\-]*[a-z0-9])?$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the AccessPoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The access point policy associated with this access point.",
		//	  "type": "object"
		//	}
		"policy": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The access point policy associated with this access point.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.",
		//	  "properties": {
		//	    "VpcId": {
		//	      "description": "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vpc_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3Outposts::AccessPoint",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::AccessPoint").WithTerraformTypeName("awscc_s3outposts_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"bucket":            "Bucket",
		"name":              "Name",
		"policy":            "Policy",
		"vpc_configuration": "VpcConfiguration",
		"vpc_id":            "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
