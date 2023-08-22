package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#kerberos_config DataprocWorkflowTemplate#kerberos_config}
	KerberosConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

