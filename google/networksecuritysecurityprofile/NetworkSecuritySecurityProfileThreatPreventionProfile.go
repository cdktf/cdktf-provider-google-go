// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileThreatPreventionProfile struct {
	// antivirus_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_security_profile#antivirus_overrides NetworkSecuritySecurityProfile#antivirus_overrides}
	AntivirusOverrides interface{} `field:"optional" json:"antivirusOverrides" yaml:"antivirusOverrides"`
	// severity_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_security_profile#severity_overrides NetworkSecuritySecurityProfile#severity_overrides}
	SeverityOverrides interface{} `field:"optional" json:"severityOverrides" yaml:"severityOverrides"`
	// threat_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_security_profile#threat_overrides NetworkSecuritySecurityProfile#threat_overrides}
	ThreatOverrides interface{} `field:"optional" json:"threatOverrides" yaml:"threatOverrides"`
}

