// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerrepository


type SecureSourceManagerRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/secure_source_manager_repository#create SecureSourceManagerRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/secure_source_manager_repository#delete SecureSourceManagerRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

