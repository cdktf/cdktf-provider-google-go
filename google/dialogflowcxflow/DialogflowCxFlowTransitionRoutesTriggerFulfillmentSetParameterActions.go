// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowTransitionRoutesTriggerFulfillmentSetParameterActions struct {
	// Display name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dialogflow_cx_flow#parameter DialogflowCxFlow#parameter}
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The new JSON-encoded value of the parameter. A null value clears the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dialogflow_cx_flow#value DialogflowCxFlow#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

