// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computehttphealthcheck


type ComputeHttpHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/compute_http_health_check#create ComputeHttpHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/compute_http_health_check#delete ComputeHttpHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/compute_http_health_check#update ComputeHttpHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

