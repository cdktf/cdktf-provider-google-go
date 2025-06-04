// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagergcpuseraccessbinding


type AccessContextManagerGcpUserAccessBindingScopedAccessSettings struct {
	// active_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/access_context_manager_gcp_user_access_binding#active_settings AccessContextManagerGcpUserAccessBinding#active_settings}
	ActiveSettings *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings `field:"optional" json:"activeSettings" yaml:"activeSettings"`
	// dry_run_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/access_context_manager_gcp_user_access_binding#dry_run_settings AccessContextManagerGcpUserAccessBinding#dry_run_settings}
	DryRunSettings *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings `field:"optional" json:"dryRunSettings" yaml:"dryRunSettings"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/access_context_manager_gcp_user_access_binding#scope AccessContextManagerGcpUserAccessBinding#scope}
	Scope *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope `field:"optional" json:"scope" yaml:"scope"`
}

