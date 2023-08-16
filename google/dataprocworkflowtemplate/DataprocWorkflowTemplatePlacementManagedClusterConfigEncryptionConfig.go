package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig struct {
	// Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#gce_pd_kms_key_name DataprocWorkflowTemplate#gce_pd_kms_key_name}
	GcePdKmsKeyName *string `field:"optional" json:"gcePdKmsKeyName" yaml:"gcePdKmsKeyName"`
}

