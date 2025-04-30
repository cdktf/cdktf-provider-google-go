// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagergcpuseraccessbinding


type AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope struct {
	// client_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/access_context_manager_gcp_user_access_binding#client_scope AccessContextManagerGcpUserAccessBinding#client_scope}
	ClientScope *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeClientScope `field:"optional" json:"clientScope" yaml:"clientScope"`
}

