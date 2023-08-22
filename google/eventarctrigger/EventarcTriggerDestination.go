package eventarctrigger


type EventarcTriggerDestination struct {
	// [WARNING] Configuring a Cloud Function in Trigger is not supported as of today.
	//
	// The Cloud Function resource name. Format: projects/{project}/locations/{location}/functions/{function}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/eventarc_trigger#cloud_function EventarcTrigger#cloud_function}
	CloudFunction *string `field:"optional" json:"cloudFunction" yaml:"cloudFunction"`
	// cloud_run_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/eventarc_trigger#cloud_run_service EventarcTrigger#cloud_run_service}
	CloudRunService *EventarcTriggerDestinationCloudRunService `field:"optional" json:"cloudRunService" yaml:"cloudRunService"`
	// gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/eventarc_trigger#gke EventarcTrigger#gke}
	Gke *EventarcTriggerDestinationGke `field:"optional" json:"gke" yaml:"gke"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/eventarc_trigger#workflow EventarcTrigger#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

