// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplate


type ColabRuntimeTemplateSoftwareConfigEnv struct {
	// Name of the environment variable. Must be a valid C identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_runtime_template#name ColabRuntimeTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Variables that reference a $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables.
	//
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_runtime_template#value ColabRuntimeTemplate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

