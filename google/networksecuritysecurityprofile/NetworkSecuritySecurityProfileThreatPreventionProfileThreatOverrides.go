// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides struct {
	// Threat action. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/network_security_security_profile#action NetworkSecuritySecurityProfile#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Vendor-specific ID of a threat to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/network_security_security_profile#threat_id NetworkSecuritySecurityProfile#threat_id}
	ThreatId *string `field:"required" json:"threatId" yaml:"threatId"`
}

