// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowEventHandlersTriggerFulfillmentConditionalCases struct {
	// A JSON encoded list of cascading if-else conditions.
	//
	// Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.
	// See [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_flow#cases DialogflowCxFlow#cases}
	Cases *string `field:"optional" json:"cases" yaml:"cases"`
}

