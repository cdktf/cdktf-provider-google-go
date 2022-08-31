// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId struct {
	// The ID of the project to which the entities belong.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#project_id DataLossPreventionJobTrigger#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// If not empty, the ID of the namespace to which the entities belong.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#namespace_id DataLossPreventionJobTrigger#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
}

