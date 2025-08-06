// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityprofilev2


type ApigeeSecurityProfileV2ProfileAssessmentConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/apigee_security_profile_v2#assessment ApigeeSecurityProfileV2#assessment}.
	Assessment *string `field:"required" json:"assessment" yaml:"assessment"`
	// The weight of the assessment. Possible values: ["MINOR", "MODERATE", "MAJOR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/apigee_security_profile_v2#weight ApigeeSecurityProfileV2#weight}
	Weight *string `field:"required" json:"weight" yaml:"weight"`
}

