// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parametermanagerregionalparameter


type ParameterManagerRegionalParameterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/parameter_manager_regional_parameter#create ParameterManagerRegionalParameter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/parameter_manager_regional_parameter#delete ParameterManagerRegionalParameter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/parameter_manager_regional_parameter#update ParameterManagerRegionalParameter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

