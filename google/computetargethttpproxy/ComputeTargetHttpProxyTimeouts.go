// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computetargethttpproxy


type ComputeTargetHttpProxyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/compute_target_http_proxy#create ComputeTargetHttpProxy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/compute_target_http_proxy#delete ComputeTargetHttpProxy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/compute_target_http_proxy#update ComputeTargetHttpProxy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

