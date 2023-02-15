package cloudidentitygroupmembership


type CloudIdentityGroupMembershipRoles struct {
	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: ["OWNER", "MANAGER", "MEMBER"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group_membership#name CloudIdentityGroupMembership#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

