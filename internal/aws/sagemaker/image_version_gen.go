// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

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
	registry.AddResourceTypeFactory("aws_sagemaker_image_version", imageVersionResourceType)
}

// imageVersionResourceType returns the Terraform aws_sagemaker_image_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::ImageVersion resource type.
func imageVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"base_image": {
			// Property: BaseImage
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The registry path of the container image on which this image version is based.",
			     "maxLength": 255,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/BaseImage",
			     "type": "string"
			   }
			*/
			Description: "The registry path of the container image on which this image version is based.",
			Type:        types.StringType,
			Required:    true,
			// BaseImage is a force-new attribute.
		},
		"container_image": {
			// Property: ContainerImage
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The registry path of the container image that contains this image version.",
			     "maxLength": 255,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ContainerImage",
			     "type": "string"
			   }
			*/
			Description: "The registry path of the container image that contains this image version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_arn": {
			// Property: ImageArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the parent image.",
			     "maxLength": 256,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ImageArn",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the parent image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_name": {
			// Property: ImageName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the image this version belongs to.",
			     "maxLength": 63,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ImageName",
			     "type": "string"
			   }
			*/
			Description: "The name of the image this version belongs to.",
			Type:        types.StringType,
			Required:    true,
			// ImageName is a force-new attribute.
		},
		"image_version_arn": {
			// Property: ImageVersionArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the image version.",
			     "maxLength": 256,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/ImageVersionArn",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the image version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The version number of the image version.",
			     "$ref": "#/definitions/Version",
			     "type": "integer"
			   }
			*/
			Description: "The version number of the image version.",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SageMaker::ImageVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::ImageVersion").WithTerraformTypeName("aws_sagemaker_image_version").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_sagemaker_image_version", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
