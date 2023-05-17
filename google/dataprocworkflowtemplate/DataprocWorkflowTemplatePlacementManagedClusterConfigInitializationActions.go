package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions struct {
	// Required. Cloud Storage URI of executable file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#executable_file DataprocWorkflowTemplate#executable_file}
	ExecutableFile *string `field:"optional" json:"executableFile" yaml:"executableFile"`
	// Optional.
	//
	// Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#execution_timeout DataprocWorkflowTemplate#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
}

