// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computetargetsslproxy


type ComputeTargetSslProxyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/compute_target_ssl_proxy#create ComputeTargetSslProxy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/compute_target_ssl_proxy#delete ComputeTargetSslProxy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/compute_target_ssl_proxy#update ComputeTargetSslProxy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

