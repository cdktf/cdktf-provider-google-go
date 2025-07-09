// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowKnowledgeConnectorSettings struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_flow#data_store_connections DialogflowCxFlow#data_store_connections}
	DataStoreConnections interface{} `field:"optional" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// Whether Knowledge Connector is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_flow#enabled DialogflowCxFlow#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The target flow to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>.
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_flow#target_flow DialogflowCxFlow#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>.
	// The page must be in the same host flow (the flow that owns this 'KnowledgeConnectorSettings').
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_flow#target_page DialogflowCxFlow#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_flow#trigger_fulfillment DialogflowCxFlow#trigger_fulfillment}
	TriggerFulfillment *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

