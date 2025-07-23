// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parametermanagerregionalparameterversion


type ParameterManagerRegionalParameterVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_regional_parameter_version#create ParameterManagerRegionalParameterVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_regional_parameter_version#delete ParameterManagerRegionalParameterVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_regional_parameter_version#update ParameterManagerRegionalParameterVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

