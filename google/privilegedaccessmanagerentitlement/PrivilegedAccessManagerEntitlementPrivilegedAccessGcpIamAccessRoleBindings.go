// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementPrivilegedAccessGcpIamAccessRoleBindings struct {
	// IAM role to be granted. https://cloud.google.com/iam/docs/roles-overview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/privileged_access_manager_entitlement#role PrivilegedAccessManagerEntitlement#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The expression field of the IAM condition to be associated with the role.
	//
	// If specified, a user with an active grant for this entitlement would be able to access the resource only if this condition evaluates to true for their request.
	// https://cloud.google.com/iam/docs/conditions-overview#attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/privileged_access_manager_entitlement#condition_expression PrivilegedAccessManagerEntitlement#condition_expression}
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
}

