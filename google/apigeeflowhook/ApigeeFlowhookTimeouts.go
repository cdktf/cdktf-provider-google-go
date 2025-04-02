// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeflowhook


type ApigeeFlowhookTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_flowhook#create ApigeeFlowhook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_flowhook#delete ApigeeFlowhook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

