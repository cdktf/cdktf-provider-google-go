// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigBlockingFunctionsTriggers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/identity_platform_config#event_type IdentityPlatformConfig#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// HTTP URI trigger for the Cloud Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/identity_platform_config#function_uri IdentityPlatformConfig#function_uri}
	FunctionUri *string `field:"required" json:"functionUri" yaml:"functionUri"`
}

