// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformtenant


type IdentityPlatformTenantClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/identity_platform_tenant#permissions IdentityPlatformTenant#permissions}
	Permissions *IdentityPlatformTenantClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

