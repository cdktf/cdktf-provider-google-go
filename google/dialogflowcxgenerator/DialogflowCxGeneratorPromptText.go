// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerator


type DialogflowCxGeneratorPromptText struct {
	// Text input which can be used for prompt or banned phrases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_generator#text DialogflowCxGenerator#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

