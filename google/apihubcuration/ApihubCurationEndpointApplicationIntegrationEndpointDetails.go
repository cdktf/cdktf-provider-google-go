// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubcuration


type ApihubCurationEndpointApplicationIntegrationEndpointDetails struct {
	// The API trigger ID of the Application Integration workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_curation#trigger_id ApihubCuration#trigger_id}
	TriggerId *string `field:"required" json:"triggerId" yaml:"triggerId"`
	// The endpoint URI should be a valid REST URI for triggering an Application Integration. Format: 'https://integrations.googleapis.com/v1/{name=projects/* /locations/* /integrations/*}:execute' or 'https://{location}-integrations.googleapis.com/v1/{name=projects/* /locations/* /integrations/*}:execute'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_curation#uri ApihubCuration#uri}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

