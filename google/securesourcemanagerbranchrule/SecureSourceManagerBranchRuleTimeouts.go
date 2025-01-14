// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerbranchrule


type SecureSourceManagerBranchRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/secure_source_manager_branch_rule#create SecureSourceManagerBranchRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/secure_source_manager_branch_rule#delete SecureSourceManagerBranchRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/secure_source_manager_branch_rule#update SecureSourceManagerBranchRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

