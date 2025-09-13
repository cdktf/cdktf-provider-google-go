// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityaction


type ApigeeSecurityActionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_security_action#create ApigeeSecurityAction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_security_action#delete ApigeeSecurityAction#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

