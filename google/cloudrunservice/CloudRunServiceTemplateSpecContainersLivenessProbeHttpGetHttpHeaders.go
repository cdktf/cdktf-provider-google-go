// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplateSpecContainersLivenessProbeHttpGetHttpHeaders struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/cloud_run_service#value CloudRunService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

