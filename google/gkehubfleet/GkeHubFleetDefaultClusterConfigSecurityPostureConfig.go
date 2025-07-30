// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfleet


type GkeHubFleetDefaultClusterConfigSecurityPostureConfig struct {
	// Sets which mode to use for Security Posture features. Possible values: ["DISABLED", "BASIC", "ENTERPRISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_hub_fleet#mode GkeHubFleet#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Sets which mode to use for vulnerability scanning. Possible values: ["VULNERABILITY_DISABLED", "VULNERABILITY_BASIC", "VULNERABILITY_ENTERPRISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_hub_fleet#vulnerability_mode GkeHubFleet#vulnerability_mode}
	VulnerabilityMode *string `field:"optional" json:"vulnerabilityMode" yaml:"vulnerabilityMode"`
}

