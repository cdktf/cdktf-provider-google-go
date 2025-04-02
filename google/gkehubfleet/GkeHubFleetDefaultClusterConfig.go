// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfleet


type GkeHubFleetDefaultClusterConfig struct {
	// binary_authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gke_hub_fleet#binary_authorization_config GkeHubFleet#binary_authorization_config}
	BinaryAuthorizationConfig *GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfig `field:"optional" json:"binaryAuthorizationConfig" yaml:"binaryAuthorizationConfig"`
	// security_posture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gke_hub_fleet#security_posture_config GkeHubFleet#security_posture_config}
	SecurityPostureConfig *GkeHubFleetDefaultClusterConfigSecurityPostureConfig `field:"optional" json:"securityPostureConfig" yaml:"securityPostureConfig"`
}

