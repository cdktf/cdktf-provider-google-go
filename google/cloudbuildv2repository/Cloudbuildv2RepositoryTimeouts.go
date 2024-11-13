// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildv2repository


type Cloudbuildv2RepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloudbuildv2_repository#create Cloudbuildv2Repository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloudbuildv2_repository#delete Cloudbuildv2Repository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

