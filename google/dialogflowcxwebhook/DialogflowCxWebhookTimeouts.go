// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxwebhook


type DialogflowCxWebhookTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/dialogflow_cx_webhook#create DialogflowCxWebhook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/dialogflow_cx_webhook#delete DialogflowCxWebhook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/dialogflow_cx_webhook#update DialogflowCxWebhook#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

