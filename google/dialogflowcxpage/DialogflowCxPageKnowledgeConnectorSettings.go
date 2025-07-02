// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageKnowledgeConnectorSettings struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_page#data_store_connections DialogflowCxPage#data_store_connections}
	DataStoreConnections interface{} `field:"optional" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// Whether Knowledge Connector is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_page#enabled DialogflowCxPage#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The target flow to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>.
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_page#target_flow DialogflowCxPage#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>.
	// The page must be in the same host flow (the flow that owns this 'KnowledgeConnectorSettings').
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_page#target_page DialogflowCxPage#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_page#trigger_fulfillment DialogflowCxPage#trigger_fulfillment}
	TriggerFulfillment *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

