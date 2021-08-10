// Code generated by generators/resource/main.go; DO NOT EDIT.

package athena_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSAthenaDataCatalog_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Athena::DataCatalog", "aws_athena_data_catalog", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
