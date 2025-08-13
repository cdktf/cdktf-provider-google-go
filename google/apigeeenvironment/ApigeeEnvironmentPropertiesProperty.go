// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeenvironment


type ApigeeEnvironmentPropertiesProperty struct {
	// The property key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_environment#name ApigeeEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The property value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_environment#value ApigeeEnvironment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

