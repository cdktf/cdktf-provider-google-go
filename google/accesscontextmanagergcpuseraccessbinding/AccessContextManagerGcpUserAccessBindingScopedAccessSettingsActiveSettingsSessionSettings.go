// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagergcpuseraccessbinding


type AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsSessionSettings struct {
	// Optional.
	//
	// How long a user is allowed to take between actions before a new access token must be issued. Only set for Google Cloud apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/access_context_manager_gcp_user_access_binding#max_inactivity AccessContextManagerGcpUserAccessBinding#max_inactivity}
	MaxInactivity *string `field:"optional" json:"maxInactivity" yaml:"maxInactivity"`
	// Optional.
	//
	// The session length. Setting this field to zero is equal to disabling session. Also can set infinite session by flipping the enabled bit to false below. If useOidcMaxAge is true, for OIDC apps, the session length will be the minimum of this field and OIDC max_age param.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/access_context_manager_gcp_user_access_binding#session_length AccessContextManagerGcpUserAccessBinding#session_length}
	SessionLength *string `field:"optional" json:"sessionLength" yaml:"sessionLength"`
	// Optional.
	//
	// This field enables or disables Google Cloud session length. When false, all fields set above will be disregarded and the session length is basically infinite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/access_context_manager_gcp_user_access_binding#session_length_enabled AccessContextManagerGcpUserAccessBinding#session_length_enabled}
	SessionLengthEnabled interface{} `field:"optional" json:"sessionLengthEnabled" yaml:"sessionLengthEnabled"`
	// Optional.
	//
	// The session challenges proposed to users when the Google Cloud session length is up. Possible values: ["LOGIN", "SECURITY_KEY", "PASSWORD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/access_context_manager_gcp_user_access_binding#session_reauth_method AccessContextManagerGcpUserAccessBinding#session_reauth_method}
	SessionReauthMethod *string `field:"optional" json:"sessionReauthMethod" yaml:"sessionReauthMethod"`
	// Optional.
	//
	// Only useful for OIDC apps. When false, the OIDC max_age param, if passed in the authentication request will be ignored. When true, the re-auth period will be the minimum of the sessionLength field and the max_age OIDC param.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/access_context_manager_gcp_user_access_binding#use_oidc_max_age AccessContextManagerGcpUserAccessBinding#use_oidc_max_age}
	UseOidcMaxAge interface{} `field:"optional" json:"useOidcMaxAge" yaml:"useOidcMaxAge"`
}

