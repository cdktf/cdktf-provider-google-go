// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides struct {
	// Threat action override. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/network_security_security_profile#action NetworkSecuritySecurityProfile#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Severity level to match. Possible values: ["CRITICAL", "HIGH", "INFORMATIONAL", "LOW", "MEDIUM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/network_security_security_profile#severity NetworkSecuritySecurityProfile#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
}

