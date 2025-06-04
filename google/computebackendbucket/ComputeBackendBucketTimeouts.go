// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendbucket


type ComputeBackendBucketTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_backend_bucket#create ComputeBackendBucket#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_backend_bucket#delete ComputeBackendBucket#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_backend_bucket#update ComputeBackendBucket#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

