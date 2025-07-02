// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateContainersStartupProbeHttpGetHttpHeaders struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_run_v2_job#name CloudRunV2Job#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_run_v2_job#value CloudRunV2Job#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

