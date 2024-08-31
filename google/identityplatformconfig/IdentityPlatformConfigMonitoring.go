// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigMonitoring struct {
	// request_logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/identity_platform_config#request_logging IdentityPlatformConfig#request_logging}
	RequestLogging *IdentityPlatformConfigMonitoringRequestLogging `field:"optional" json:"requestLogging" yaml:"requestLogging"`
}

