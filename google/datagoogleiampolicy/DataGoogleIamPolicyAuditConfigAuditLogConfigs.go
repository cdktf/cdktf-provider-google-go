package datagoogleiampolicy


type DataGoogleIamPolicyAuditConfigAuditLogConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/iam_policy#log_type DataGoogleIamPolicy#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/iam_policy#exempted_members DataGoogleIamPolicy#exempted_members}.
	ExemptedMembers *[]*string `field:"optional" json:"exemptedMembers" yaml:"exemptedMembers"`
}

