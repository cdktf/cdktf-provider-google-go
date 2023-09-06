// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/cloud_run_service#metadata CloudRunService#metadata}
	Metadata *CloudRunServiceTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/cloud_run_service#spec CloudRunService#spec}
	Spec *CloudRunServiceTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

