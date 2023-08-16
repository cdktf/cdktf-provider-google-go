package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActionsPubSub struct {
	// Cloud Pub/Sub topic to send notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#topic DataLossPreventionJobTrigger#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

