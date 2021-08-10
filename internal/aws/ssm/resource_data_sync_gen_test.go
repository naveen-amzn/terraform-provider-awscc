// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSSSMResourceDataSync_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSM::ResourceDataSync", "aws_ssm_resource_data_sync", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
