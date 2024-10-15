// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowfulfillment


type DialogflowFulfillmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/dialogflow_fulfillment#create DialogflowFulfillment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/dialogflow_fulfillment#delete DialogflowFulfillment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/dialogflow_fulfillment#update DialogflowFulfillment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

