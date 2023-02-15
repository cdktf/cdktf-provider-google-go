package cloudbuildtrigger


type CloudbuildTriggerPubsubConfig struct {
	// The name of the topic from which this subscription is receiving messages.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#topic CloudbuildTrigger#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Service account that will make the push request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#service_account_email CloudbuildTrigger#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

