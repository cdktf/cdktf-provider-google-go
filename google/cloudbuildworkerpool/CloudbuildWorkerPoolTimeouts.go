// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildworkerpool


type CloudbuildWorkerPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/cloudbuild_worker_pool#create CloudbuildWorkerPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/cloudbuild_worker_pool#delete CloudbuildWorkerPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/cloudbuild_worker_pool#update CloudbuildWorkerPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

