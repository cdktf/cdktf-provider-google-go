// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntime


type NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImages struct {
	// The path to the container image repository. For example: gcr.io/{project_id}/{imageName}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/notebooks_runtime#repository NotebooksRuntime#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/notebooks_runtime#tag NotebooksRuntime#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

