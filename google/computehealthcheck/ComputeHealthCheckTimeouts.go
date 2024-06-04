// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computehealthcheck


type ComputeHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/compute_health_check#create ComputeHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/compute_health_check#delete ComputeHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/compute_health_check#update ComputeHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

