package datagoogleiampolicy


type DataGoogleIamPolicyAuditConfig struct {
	// audit_log_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/iam_policy#audit_log_configs DataGoogleIamPolicy#audit_log_configs}
	AuditLogConfigs interface{} `field:"required" json:"auditLogConfigs" yaml:"auditLogConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/iam_policy#service DataGoogleIamPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

