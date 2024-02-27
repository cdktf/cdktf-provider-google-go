// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarctrigger


type EventarcTriggerDestination struct {
	// cloud_run_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/eventarc_trigger#cloud_run_service EventarcTrigger#cloud_run_service}
	CloudRunService *EventarcTriggerDestinationCloudRunService `field:"optional" json:"cloudRunService" yaml:"cloudRunService"`
	// gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/eventarc_trigger#gke EventarcTrigger#gke}
	Gke *EventarcTriggerDestinationGke `field:"optional" json:"gke" yaml:"gke"`
	// http_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/eventarc_trigger#http_endpoint EventarcTrigger#http_endpoint}
	HttpEndpoint *EventarcTriggerDestinationHttpEndpoint `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/eventarc_trigger#network_config EventarcTrigger#network_config}
	NetworkConfig *EventarcTriggerDestinationNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/eventarc_trigger#workflow EventarcTrigger#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

