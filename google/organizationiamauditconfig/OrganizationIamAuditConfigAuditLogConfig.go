package organizationiamauditconfig


type OrganizationIamAuditConfigAuditLogConfig struct {
	// Permission type for which logging is to be configured. Must be one of DATA_READ, DATA_WRITE, or ADMIN_READ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/organization_iam_audit_config#log_type OrganizationIamAuditConfig#log_type}
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Identities that do not cause logging for this type of permission.
	//
	// Each entry can have one of the following values:user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com. serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/organization_iam_audit_config#exempted_members OrganizationIamAuditConfig#exempted_members}
	ExemptedMembers *[]*string `field:"optional" json:"exemptedMembers" yaml:"exemptedMembers"`
}

