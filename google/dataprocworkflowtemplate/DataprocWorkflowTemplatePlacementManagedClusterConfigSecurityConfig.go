package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#kerberos_config DataprocWorkflowTemplate#kerberos_config}
	KerberosConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

