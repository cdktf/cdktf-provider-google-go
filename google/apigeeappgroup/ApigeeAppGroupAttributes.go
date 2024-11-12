// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeappgroup


type ApigeeAppGroupAttributes struct {
	// Key of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/apigee_app_group#name ApigeeAppGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/apigee_app_group#value ApigeeAppGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

