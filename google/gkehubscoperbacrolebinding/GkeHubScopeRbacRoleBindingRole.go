// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubscoperbacrolebinding


type GkeHubScopeRbacRoleBindingRole struct {
	// CustomRole is the custom Kubernetes ClusterRole to be used.
	//
	// The custom role format must be allowlisted in the rbacrolebindingactuation feature and RFC 1123 compliant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gke_hub_scope_rbac_role_binding#custom_role GkeHubScopeRbacRoleBinding#custom_role}
	CustomRole *string `field:"optional" json:"customRole" yaml:"customRole"`
	// PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: ["UNKNOWN", "ADMIN", "EDIT", "VIEW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gke_hub_scope_rbac_role_binding#predefined_role GkeHubScopeRbacRoleBinding#predefined_role}
	PredefinedRole *string `field:"optional" json:"predefinedRole" yaml:"predefinedRole"`
}

