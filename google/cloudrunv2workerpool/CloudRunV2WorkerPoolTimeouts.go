// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_run_v2_worker_pool#create CloudRunV2WorkerPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_run_v2_worker_pool#delete CloudRunV2WorkerPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_run_v2_worker_pool#update CloudRunV2WorkerPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

