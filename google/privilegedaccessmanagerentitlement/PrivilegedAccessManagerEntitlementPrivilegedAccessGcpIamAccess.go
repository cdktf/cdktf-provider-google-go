// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementPrivilegedAccessGcpIamAccess struct {
	// Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/privileged_access_manager_entitlement#resource PrivilegedAccessManagerEntitlement#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The type of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/privileged_access_manager_entitlement#resource_type PrivilegedAccessManagerEntitlement#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// role_bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/privileged_access_manager_entitlement#role_bindings PrivilegedAccessManagerEntitlement#role_bindings}
	RoleBindings interface{} `field:"required" json:"roleBindings" yaml:"roleBindings"`
}

