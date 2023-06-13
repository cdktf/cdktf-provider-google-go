package cloudfunctionsfunction


type CloudfunctionsFunctionEventTrigger struct {
	// The type of event to observe.
	//
	// For example: "google.storage.object.finalize". See the documentation on calling Cloud Functions for a full reference of accepted triggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloudfunctions_function#event_type CloudfunctionsFunction#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The name or partial URI of the resource from which to observe events. For example, "myBucket" or "projects/my-project/topics/my-topic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloudfunctions_function#resource CloudfunctionsFunction#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// failure_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloudfunctions_function#failure_policy CloudfunctionsFunction#failure_policy}
	FailurePolicy *CloudfunctionsFunctionEventTriggerFailurePolicy `field:"optional" json:"failurePolicy" yaml:"failurePolicy"`
}

