// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeedeveloper


type ApigeeDeveloperAttributes struct {
	// Key of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_developer#name ApigeeDeveloper#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_developer#value ApigeeDeveloper#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

