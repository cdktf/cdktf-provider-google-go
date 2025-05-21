// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminicoderepositoryindex


type GeminiCodeRepositoryIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gemini_code_repository_index#create GeminiCodeRepositoryIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gemini_code_repository_index#delete GeminiCodeRepositoryIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gemini_code_repository_index#update GeminiCodeRepositoryIndex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

