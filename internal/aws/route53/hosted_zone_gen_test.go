// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSRoute53HostedZone_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53::HostedZone", "aws_route53_hosted_zone", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
