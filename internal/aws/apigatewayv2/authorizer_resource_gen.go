// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigatewayv2_authorizer", authorizerResource)
}

// authorizerResource returns the Terraform awscc_apigatewayv2_authorizer resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGatewayV2::Authorizer resource.
func authorizerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_id": {
			// Property: ApiId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"authorizer_credentials_arn": {
			// Property: AuthorizerCredentialsArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"authorizer_id": {
			// Property: AuthorizerId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"authorizer_payload_format_version": {
			// Property: AuthorizerPayloadFormatVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"authorizer_result_ttl_in_seconds": {
			// Property: AuthorizerResultTtlInSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"authorizer_type": {
			// Property: AuthorizerType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
		},
		"authorizer_uri": {
			// Property: AuthorizerUri
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enable_simple_responses": {
			// Property: EnableSimpleResponses
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"identity_source": {
			// Property: IdentitySource
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"identity_validation_expression": {
			// Property: IdentityValidationExpression
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"jwt_configuration": {
			// Property: JwtConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Audience": {
			//	      "items": {
			//	        "type": "string"
			//	      },
			//	      "type": "array",
			//	      "uniqueItems": false
			//	    },
			//	    "Issuer": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"audience": {
						// Property: Audience
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"issuer": {
						// Property: Issuer
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ApiGatewayV2::Authorizer",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::Authorizer").WithTerraformTypeName("awscc_apigatewayv2_authorizer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":                            "ApiId",
		"audience":                          "Audience",
		"authorizer_credentials_arn":        "AuthorizerCredentialsArn",
		"authorizer_id":                     "AuthorizerId",
		"authorizer_payload_format_version": "AuthorizerPayloadFormatVersion",
		"authorizer_result_ttl_in_seconds":  "AuthorizerResultTtlInSeconds",
		"authorizer_type":                   "AuthorizerType",
		"authorizer_uri":                    "AuthorizerUri",
		"enable_simple_responses":           "EnableSimpleResponses",
		"identity_source":                   "IdentitySource",
		"identity_validation_expression":    "IdentityValidationExpression",
		"issuer":                            "Issuer",
		"jwt_configuration":                 "JwtConfiguration",
		"name":                              "Name",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}