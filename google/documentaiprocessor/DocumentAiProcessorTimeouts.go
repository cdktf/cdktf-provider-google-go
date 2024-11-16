// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package documentaiprocessor


type DocumentAiProcessorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/document_ai_processor#create DocumentAiProcessor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/document_ai_processor#delete DocumentAiProcessor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

