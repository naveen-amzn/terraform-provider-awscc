// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_macie_findings_filter", findingsFilterDataSourceType)
}

// findingsFilterDataSourceType returns the Terraform awscc_macie_findings_filter data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Macie::FindingsFilter resource type.
func findingsFilterDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter action.",
			//   "enum": [
			//     "ARCHIVE",
			//     "NOOP"
			//   ],
			//   "type": "string"
			// }
			Description: "Findings filter action.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter ARN.",
			//   "type": "string"
			// }
			Description: "Findings filter ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter description",
			//   "type": "string"
			// }
			Description: "Findings filter description",
			Type:        types.StringType,
			Computed:    true,
		},
		"finding_criteria": {
			// Property: FindingCriteria
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter criteria.",
			//   "properties": {
			//     "Criterion": {
			//       "description": "Map of filter criteria.",
			//       "patternProperties": {
			//         "": {
			//           "properties": {
			//             "eq": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "gt": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "gte": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "lt": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "lte": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "neq": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Findings filter criteria.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"criterion": {
						// Property: Criterion
						Description: "Map of filter criteria.",
						// Pattern: ""
						Attributes: tfsdk.MapNestedAttributes(
							map[string]tfsdk.Attribute{
								"eq": {
									// Property: eq
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"gt": {
									// Property: gt
									Type:     types.NumberType,
									Computed: true,
								},
								"gte": {
									// Property: gte
									Type:     types.NumberType,
									Computed: true,
								},
								"lt": {
									// Property: lt
									Type:     types.NumberType,
									Computed: true,
								},
								"lte": {
									// Property: lte
									Type:     types.NumberType,
									Computed: true,
								},
								"neq": {
									// Property: neq
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
							tfsdk.MapNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"findings_filter_list_items": {
			// Property: FindingsFilterListItems
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filters list.",
			//   "items": {
			//     "description": "Returned by ListHandler representing filter name and ID.",
			//     "properties": {
			//       "Id": {
			//         "type": "string"
			//       },
			//       "Name": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Findings filters list.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"id": {
						// Property: Id
						Type:     types.StringType,
						Computed: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter ID.",
			//   "type": "string"
			// }
			Description: "Findings filter ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter name",
			//   "type": "string"
			// }
			Description: "Findings filter name",
			Type:        types.StringType,
			Computed:    true,
		},
		"position": {
			// Property: Position
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter position.",
			//   "type": "integer"
			// }
			Description: "Findings filter position.",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Macie::FindingsFilter",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::FindingsFilter").WithTerraformTypeName("awscc_macie_findings_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                     "Action",
		"arn":                        "Arn",
		"criterion":                  "Criterion",
		"description":                "Description",
		"eq":                         "eq",
		"finding_criteria":           "FindingCriteria",
		"findings_filter_list_items": "FindingsFilterListItems",
		"gt":                         "gt",
		"gte":                        "gte",
		"id":                         "Id",
		"lt":                         "lt",
		"lte":                        "lte",
		"name":                       "Name",
		"neq":                        "neq",
		"position":                   "Position",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
