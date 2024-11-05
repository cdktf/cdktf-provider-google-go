// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment


type ComposerEnvironmentConfigWebServerNetworkAccessControl struct {
	// allowed_ip_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/composer_environment#allowed_ip_range ComposerEnvironment#allowed_ip_range}
	AllowedIpRange interface{} `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
}

