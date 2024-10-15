// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntime


type NotebooksRuntimeVirtualMachine struct {
	// virtual_machine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/notebooks_runtime#virtual_machine_config NotebooksRuntime#virtual_machine_config}
	VirtualMachineConfig *NotebooksRuntimeVirtualMachineVirtualMachineConfig `field:"optional" json:"virtualMachineConfig" yaml:"virtualMachineConfig"`
}

