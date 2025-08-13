// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containeranalysisnote


type ContainerAnalysisNoteRelatedUrl struct {
	// Specific URL associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/container_analysis_note#url ContainerAnalysisNote#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Label to describe usage of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/container_analysis_note#label ContainerAnalysisNote#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

