// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeorganization


type ApigeeOrganizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/apigee_organization#create ApigeeOrganization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/apigee_organization#delete ApigeeOrganization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/apigee_organization#update ApigeeOrganization#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

