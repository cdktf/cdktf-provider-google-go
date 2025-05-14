// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityprofilev2


type ApigeeSecurityProfileV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/apigee_security_profile_v2#create ApigeeSecurityProfileV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/apigee_security_profile_v2#delete ApigeeSecurityProfileV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/apigee_security_profile_v2#update ApigeeSecurityProfileV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

