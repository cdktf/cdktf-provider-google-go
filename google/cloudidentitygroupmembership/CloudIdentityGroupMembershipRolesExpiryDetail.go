// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudidentitygroupmembership


type CloudIdentityGroupMembershipRolesExpiryDetail struct {
	// The time at which the MembershipRole will expire.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	//
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/cloud_identity_group_membership#expire_time CloudIdentityGroupMembership#expire_time}
	ExpireTime *string `field:"required" json:"expireTime" yaml:"expireTime"`
}

