// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composeruserworkloadsconfigmap


type ComposerUserWorkloadsConfigMapTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/composer_user_workloads_config_map#create ComposerUserWorkloadsConfigMap#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/composer_user_workloads_config_map#delete ComposerUserWorkloadsConfigMap#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/composer_user_workloads_config_map#update ComposerUserWorkloadsConfigMap#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

