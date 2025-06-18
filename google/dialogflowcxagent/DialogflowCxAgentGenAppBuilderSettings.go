// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxagent


type DialogflowCxAgentGenAppBuilderSettings struct {
	// The full name of the Gen App Builder engine related to this agent if there is one.
	//
	// Format: projects/{Project ID}/locations/{Location ID}/collections/{Collection ID}/engines/{Engine ID}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/dialogflow_cx_agent#engine DialogflowCxAgent#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
}

