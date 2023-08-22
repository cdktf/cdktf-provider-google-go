package notebooksruntime


type NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig struct {
	// Count of cores of this accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/notebooks_runtime#core_count NotebooksRuntime#core_count}
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Accelerator model. For valid values, see 'https://cloud.google.com/vertex-ai/docs/workbench/reference/ rest/v1/projects.locations.runtimes#AcceleratorType'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/notebooks_runtime#type NotebooksRuntime#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

