// Code generated by generators/resource/main.go; DO NOT EDIT.

package applicationinsights_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSApplicationInsightsApplication_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApplicationInsights::Application", "aws_applicationinsights_application", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
