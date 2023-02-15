package notebooksruntime


type NotebooksRuntimeVirtualMachine struct {
	// virtual_machine_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime#virtual_machine_config NotebooksRuntime#virtual_machine_config}
	VirtualMachineConfig *NotebooksRuntimeVirtualMachineVirtualMachineConfig `field:"optional" json:"virtualMachineConfig" yaml:"virtualMachineConfig"`
}

