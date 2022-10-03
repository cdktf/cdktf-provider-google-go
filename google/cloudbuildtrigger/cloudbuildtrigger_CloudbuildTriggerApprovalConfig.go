package cloudbuildtrigger


type CloudbuildTriggerApprovalConfig struct {
	// Whether or not approval is needed.
	//
	// If this is set on a build, it will become pending when run,
	// and will need to be explicitly approved to start.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#approval_required CloudbuildTrigger#approval_required}
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
}

