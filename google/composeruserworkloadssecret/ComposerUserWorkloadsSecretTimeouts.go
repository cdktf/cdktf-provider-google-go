// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composeruserworkloadssecret


type ComposerUserWorkloadsSecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/composer_user_workloads_secret#create ComposerUserWorkloadsSecret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/composer_user_workloads_secret#delete ComposerUserWorkloadsSecret#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/composer_user_workloads_secret#update ComposerUserWorkloadsSecret#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

