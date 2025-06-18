// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sourcereporepository


type SourcerepoRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/sourcerepo_repository#create SourcerepoRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/sourcerepo_repository#delete SourcerepoRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/sourcerepo_repository#update SourcerepoRepository#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

