// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2service


type CloudRunV2ServiceScaling struct {
	// Total instance count for the service in manual scaling mode.
	//
	// This number of instances is divided among all revisions with specified traffic based on the percent of traffic they are receiving.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/cloud_run_v2_service#manual_instance_count CloudRunV2Service#manual_instance_count}
	ManualInstanceCount *float64 `field:"optional" json:"manualInstanceCount" yaml:"manualInstanceCount"`
	// Minimum number of instances for the service, to be divided among all revisions receiving traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/cloud_run_v2_service#min_instance_count CloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// The [scaling mode](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#scalingmode) for the service. Possible values: ["AUTOMATIC", "MANUAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/cloud_run_v2_service#scaling_mode CloudRunV2Service#scaling_mode}
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

