// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudidentitygroupmembership


type CloudIdentityGroupMembershipRoles struct {
	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: ["OWNER", "MANAGER", "MEMBER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/cloud_identity_group_membership#name CloudIdentityGroupMembership#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

