// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSRDSDBProxyTargetGroup_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RDS::DBProxyTargetGroup", "aws_rds_db_proxy_target_group", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
