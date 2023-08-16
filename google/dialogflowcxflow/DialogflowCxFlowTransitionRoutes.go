package dialogflowcxflow


type DialogflowCxFlowTransitionRoutes struct {
	// The condition to evaluate against form parameters or session parameters.
	//
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#condition DialogflowCxFlow#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The unique identifier of an Intent.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#intent DialogflowCxFlow#intent}
	Intent *string `field:"optional" json:"intent" yaml:"intent"`
	// The target flow to transition to. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#target_flow DialogflowCxFlow#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#target_page DialogflowCxFlow#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#trigger_fulfillment DialogflowCxFlow#trigger_fulfillment}
	TriggerFulfillment *DialogflowCxFlowTransitionRoutesTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

