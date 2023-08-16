package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActions struct {
	// deidentify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#deidentify DataLossPreventionJobTrigger#deidentify}
	Deidentify *DataLossPreventionJobTriggerInspectJobActionsDeidentify `field:"optional" json:"deidentify" yaml:"deidentify"`
	// job_notification_emails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#job_notification_emails DataLossPreventionJobTrigger#job_notification_emails}
	JobNotificationEmails *DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails `field:"optional" json:"jobNotificationEmails" yaml:"jobNotificationEmails"`
	// publish_findings_to_cloud_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#publish_findings_to_cloud_data_catalog DataLossPreventionJobTrigger#publish_findings_to_cloud_data_catalog}
	PublishFindingsToCloudDataCatalog *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog `field:"optional" json:"publishFindingsToCloudDataCatalog" yaml:"publishFindingsToCloudDataCatalog"`
	// publish_summary_to_cscc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#publish_summary_to_cscc DataLossPreventionJobTrigger#publish_summary_to_cscc}
	PublishSummaryToCscc *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc `field:"optional" json:"publishSummaryToCscc" yaml:"publishSummaryToCscc"`
	// publish_to_stackdriver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#publish_to_stackdriver DataLossPreventionJobTrigger#publish_to_stackdriver}
	PublishToStackdriver *DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver `field:"optional" json:"publishToStackdriver" yaml:"publishToStackdriver"`
	// pub_sub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#pub_sub DataLossPreventionJobTrigger#pub_sub}
	PubSub *DataLossPreventionJobTriggerInspectJobActionsPubSub `field:"optional" json:"pubSub" yaml:"pubSub"`
	// save_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#save_findings DataLossPreventionJobTrigger#save_findings}
	SaveFindings *DataLossPreventionJobTriggerInspectJobActionsSaveFindings `field:"optional" json:"saveFindings" yaml:"saveFindings"`
}

