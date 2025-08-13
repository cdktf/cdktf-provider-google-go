// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityaction


type ApigeeSecurityActionFlagHeaders struct {
	// The header name to be sent to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#name ApigeeSecurityAction#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header value to be sent to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#value ApigeeSecurityAction#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

