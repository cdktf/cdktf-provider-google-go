package datagoogleiampolicy


type DataGoogleIamPolicyAuditConfig struct {
	// audit_log_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#audit_log_configs DataGoogleIamPolicy#audit_log_configs}
	AuditLogConfigs interface{} `field:"required" json:"auditLogConfigs" yaml:"auditLogConfigs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#service DataGoogleIamPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

