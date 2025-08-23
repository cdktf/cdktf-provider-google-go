// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineDestinations struct {
	// authentication_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#authentication_config EventarcPipeline#authentication_config}
	AuthenticationConfig *EventarcPipelineDestinationsAuthenticationConfig `field:"optional" json:"authenticationConfig" yaml:"authenticationConfig"`
	// http_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#http_endpoint EventarcPipeline#http_endpoint}
	HttpEndpoint *EventarcPipelineDestinationsHttpEndpoint `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// The resource name of the Message Bus to which events should be published.
	//
	// The Message Bus resource should exist in the same project as
	// the Pipeline. Format:
	// 'projects/{project}/locations/{location}/messageBuses/{message_bus}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#message_bus EventarcPipeline#message_bus}
	MessageBus *string `field:"optional" json:"messageBus" yaml:"messageBus"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#network_config EventarcPipeline#network_config}
	NetworkConfig *EventarcPipelineDestinationsNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// output_payload_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#output_payload_format EventarcPipeline#output_payload_format}
	OutputPayloadFormat *EventarcPipelineDestinationsOutputPayloadFormat `field:"optional" json:"outputPayloadFormat" yaml:"outputPayloadFormat"`
	// The resource name of the Pub/Sub topic to which events should be published. Format: 'projects/{project}/locations/{location}/topics/{topic}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#topic EventarcPipeline#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same
	// project as the Pipeline. Format:
	// 'projects/{project}/locations/{location}/workflows/{workflow}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/eventarc_pipeline#workflow EventarcPipeline#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

