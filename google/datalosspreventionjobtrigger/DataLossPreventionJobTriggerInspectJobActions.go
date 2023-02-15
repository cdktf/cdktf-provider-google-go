package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActions struct {
	// publish_findings_to_cloud_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#publish_findings_to_cloud_data_catalog DataLossPreventionJobTrigger#publish_findings_to_cloud_data_catalog}
	PublishFindingsToCloudDataCatalog *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog `field:"optional" json:"publishFindingsToCloudDataCatalog" yaml:"publishFindingsToCloudDataCatalog"`
	// publish_summary_to_cscc block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#publish_summary_to_cscc DataLossPreventionJobTrigger#publish_summary_to_cscc}
	PublishSummaryToCscc *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc `field:"optional" json:"publishSummaryToCscc" yaml:"publishSummaryToCscc"`
	// pub_sub block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#pub_sub DataLossPreventionJobTrigger#pub_sub}
	PubSub *DataLossPreventionJobTriggerInspectJobActionsPubSub `field:"optional" json:"pubSub" yaml:"pubSub"`
	// save_findings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#save_findings DataLossPreventionJobTrigger#save_findings}
	SaveFindings *DataLossPreventionJobTriggerInspectJobActionsSaveFindings `field:"optional" json:"saveFindings" yaml:"saveFindings"`
}

