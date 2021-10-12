// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_lambda_function", functionResourceType)
}

// functionResourceType returns the Terraform awscc_lambda_function resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lambda::Function resource type.
func functionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"architectures": {
			// Property: Architectures
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "enum": [
			//       "x86_64",
			//       "arm64"
			//     ],
			//     "type": "string"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
				validate.UniqueItems(),
				validate.ArrayForEach(validate.StringInSlice([]string{
					"x86_64",
					"arm64",
				})),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier for function resources",
			//   "type": "string"
			// }
			Description: "Unique identifier for function resources",
			Type:        types.StringType,
			Computed:    true,
		},
		"code": {
			// Property: Code
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The code for the function.",
			//   "properties": {
			//     "ImageUri": {
			//       "description": "ImageUri.",
			//       "type": "string"
			//     },
			//     "S3Bucket": {
			//       "description": "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
			//       "maxLength": 63,
			//       "minLength": 3,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "S3Key": {
			//       "description": "The Amazon S3 key of the deployment package.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "S3ObjectVersion": {
			//       "description": "For versioned objects, the version of the deployment package object to use.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "ZipFile": {
			//       "description": "The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The code for the function.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"image_uri": {
						// Property: ImageUri
						Description: "ImageUri.",
						Type:        types.StringType,
						Optional:    true,
					},
					"s3_bucket": {
						// Property: S3Bucket
						Description: "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(3, 63),
						},
					},
					"s3_key": {
						// Property: S3Key
						Description: "The Amazon S3 key of the deployment package.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1024),
						},
					},
					"s3_object_version": {
						// Property: S3ObjectVersion
						Description: "For versioned objects, the version of the deployment package object to use.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1024),
						},
					},
					"zip_file": {
						// Property: ZipFile
						Description: "The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
			// Code is a write-only property.
		},
		"code_signing_config_arn": {
			// Property: CodeSigningConfigArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique Arn for CodeSigningConfig resource",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique Arn for CodeSigningConfig resource",
			Type:        types.StringType,
			Optional:    true,
		},
		"dead_letter_config": {
			// Property: DeadLetterConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.",
			//   "properties": {
			//     "TargetArn": {
			//       "description": "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"target_arn": {
						// Property: TargetArn
						Description: "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the function.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "A description of the function.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
			},
		},
		"environment": {
			// Property: Environment
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Environment variables that are accessible from function code during execution.",
			//   "properties": {
			//     "Variables": {
			//       "additionalProperties": false,
			//       "description": "Environment variable key-value pairs.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Environment variables that are accessible from function code during execution.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"variables": {
						// Property: Variables
						Description: "Environment variable key-value pairs.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"file_system_configs": {
			// Property: FileSystemConfigs
			// CloudFormation resource type schema:
			// {
			//   "description": "Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Arn": {
			//         "description": "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
			//         "maxLength": 200,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "LocalMountPath": {
			//         "description": "The path where the function can access the file system, starting with /mnt/.",
			//         "maxLength": 160,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Arn",
			//       "LocalMountPath"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "type": "array"
			// }
			Description: "Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Description: "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(200),
						},
					},
					"local_mount_path": {
						// Property: LocalMountPath
						Description: "The path where the function can access the file system, starting with /mnt/.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(160),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(1),
			},
		},
		"function_name": {
			// Property: FunctionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"handler": {
			// Property: Handler
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(128),
			},
		},
		"image_config": {
			// Property: ImageConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "ImageConfig",
			//   "properties": {
			//     "Command": {
			//       "description": "Command.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 1500,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "EntryPoint": {
			//       "description": "EntryPoint.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 1500,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "WorkingDirectory": {
			//       "description": "WorkingDirectory.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "ImageConfig",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"command": {
						// Property: Command
						Description: "Command.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(1500),
							validate.UniqueItems(),
						},
					},
					"entry_point": {
						// Property: EntryPoint
						Description: "EntryPoint.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(1500),
							validate.UniqueItems(),
						},
					},
					"working_directory": {
						// Property: WorkingDirectory
						Description: "WorkingDirectory.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"kms_key_arn": {
			// Property: KmsKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
			Type:        types.StringType,
			Optional:    true,
		},
		"layers": {
			// Property: Layers
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"memory_size": {
			// Property: MemorySize
			// CloudFormation resource type schema:
			// {
			//   "description": "The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.",
			//   "type": "integer"
			// }
			Description: "The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"package_type": {
			// Property: PackageType
			// CloudFormation resource type schema:
			// {
			//   "description": "PackageType.",
			//   "enum": [
			//     "Image",
			//     "Zip"
			//   ],
			//   "type": "string"
			// }
			Description: "PackageType.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"Image",
					"Zip",
				}),
			},
		},
		"reserved_concurrent_executions": {
			// Property: ReservedConcurrentExecutions
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of simultaneous executions to reserve for the function.",
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "The number of simultaneous executions to reserve for the function.",
			Type:        types.NumberType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntAtLeast(0),
			},
		},
		"role": {
			// Property: Role
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the function's execution role.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the function's execution role.",
			Type:        types.StringType,
			Required:    true,
		},
		"runtime": {
			// Property: Runtime
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the function's runtime.",
			//   "type": "string"
			// }
			Description: "The identifier of the function's runtime.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the function.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of tags to apply to the function.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"timeout": {
			// Property: Timeout
			// CloudFormation resource type schema:
			// {
			//   "description": "The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.",
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Description: "The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.",
			Type:        types.NumberType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntAtLeast(1),
			},
		},
		"tracing_config": {
			// Property: TracingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.",
			//   "properties": {
			//     "Mode": {
			//       "description": "The tracing mode.",
			//       "enum": [
			//         "Active",
			//         "PassThrough"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"mode": {
						// Property: Mode
						Description: "The tracing mode.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"Active",
								"PassThrough",
							}),
						},
					},
				},
			),
			Optional: true,
		},
		"vpc_config": {
			// Property: VpcConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC.",
			//   "properties": {
			//     "SecurityGroupIds": {
			//       "description": "A list of VPC security groups IDs.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 5,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SubnetIds": {
			//       "description": "A list of VPC subnet IDs.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 16,
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"security_group_ids": {
						// Property: SecurityGroupIds
						Description: "A list of VPC security groups IDs.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(5),
						},
					},
					"subnet_ids": {
						// Property: SubnetIds
						Description: "A list of VPC subnet IDs.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(16),
						},
					},
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Lambda::Function",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::Function").WithTerraformTypeName("awscc_lambda_function")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architectures":                  "Architectures",
		"arn":                            "Arn",
		"code":                           "Code",
		"code_signing_config_arn":        "CodeSigningConfigArn",
		"command":                        "Command",
		"dead_letter_config":             "DeadLetterConfig",
		"description":                    "Description",
		"entry_point":                    "EntryPoint",
		"environment":                    "Environment",
		"file_system_configs":            "FileSystemConfigs",
		"function_name":                  "FunctionName",
		"handler":                        "Handler",
		"image_config":                   "ImageConfig",
		"image_uri":                      "ImageUri",
		"key":                            "Key",
		"kms_key_arn":                    "KmsKeyArn",
		"layers":                         "Layers",
		"local_mount_path":               "LocalMountPath",
		"memory_size":                    "MemorySize",
		"mode":                           "Mode",
		"package_type":                   "PackageType",
		"reserved_concurrent_executions": "ReservedConcurrentExecutions",
		"role":                           "Role",
		"runtime":                        "Runtime",
		"s3_bucket":                      "S3Bucket",
		"s3_key":                         "S3Key",
		"s3_object_version":              "S3ObjectVersion",
		"security_group_ids":             "SecurityGroupIds",
		"subnet_ids":                     "SubnetIds",
		"tags":                           "Tags",
		"target_arn":                     "TargetArn",
		"timeout":                        "Timeout",
		"tracing_config":                 "TracingConfig",
		"value":                          "Value",
		"variables":                      "Variables",
		"vpc_config":                     "VpcConfig",
		"working_directory":              "WorkingDirectory",
		"zip_file":                       "ZipFile",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Code",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
