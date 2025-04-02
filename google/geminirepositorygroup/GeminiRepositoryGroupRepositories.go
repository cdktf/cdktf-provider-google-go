// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminirepositorygroup


type GeminiRepositoryGroupRepositories struct {
	// Required. The Git branch pattern used for indexing in RE2 syntax. See https://github.com/google/re2/wiki/syntax for syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gemini_repository_group#branch_pattern GeminiRepositoryGroup#branch_pattern}
	BranchPattern *string `field:"required" json:"branchPattern" yaml:"branchPattern"`
	// Required. The DeveloperConnect repository full resource name, relative resource name or resource URL to be indexed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gemini_repository_group#resource GeminiRepositoryGroup#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
}

