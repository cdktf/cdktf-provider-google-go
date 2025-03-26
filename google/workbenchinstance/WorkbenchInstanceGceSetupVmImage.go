// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetupVmImage struct {
	// Optional. Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/workbench_instance#family WorkbenchInstance#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Optional. Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/workbench_instance#name WorkbenchInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the Google Cloud project that this VM image belongs to. Format: {project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/workbench_instance#project WorkbenchInstance#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

