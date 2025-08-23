// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageFormParametersFillBehaviorRepromptEventHandlers struct {
	// The name of the event to handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_page#event DialogflowCxPage#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The target flow to transition to. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_page#target_flow DialogflowCxPage#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_page#target_page DialogflowCxPage#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_page#trigger_fulfillment DialogflowCxPage#trigger_fulfillment}
	TriggerFulfillment *DialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

