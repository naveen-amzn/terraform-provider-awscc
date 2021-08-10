// Code generated by generators/resource/main.go; DO NOT EDIT.

package emrcontainers_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSEMRContainersVirtualCluster_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EMRContainers::VirtualCluster", "aws_emrcontainers_virtual_cluster", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
