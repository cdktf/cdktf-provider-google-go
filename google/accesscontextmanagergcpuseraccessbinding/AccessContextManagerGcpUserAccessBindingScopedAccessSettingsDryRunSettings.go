// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagergcpuseraccessbinding


type AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings struct {
	// Optional.
	//
	// Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/access_context_manager_gcp_user_access_binding#access_levels AccessContextManagerGcpUserAccessBinding#access_levels}
	AccessLevels *[]*string `field:"optional" json:"accessLevels" yaml:"accessLevels"`
}

