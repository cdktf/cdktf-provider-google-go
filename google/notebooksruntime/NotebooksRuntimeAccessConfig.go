package notebooksruntime


type NotebooksRuntimeAccessConfig struct {
	// The type of access mode this instance. For valid values, see 'https://cloud.google.com/vertex-ai/docs/workbench/reference/ rest/v1/projects.locations.runtimes#RuntimeAccessType'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_runtime#access_type NotebooksRuntime#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// The owner of this runtime after creation. Format: 'alias@example.com'. Currently supports one owner only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_runtime#runtime_owner NotebooksRuntime#runtime_owner}
	RuntimeOwner *string `field:"optional" json:"runtimeOwner" yaml:"runtimeOwner"`
}

