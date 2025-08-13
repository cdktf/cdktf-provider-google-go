// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineDestinationsNetworkConfig struct {
	// Name of the NetworkAttachment that allows access to the consumer VPC.
	//
	// Format:
	// 'projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}'
	//
	// Required for HTTP endpoint destinations. Must not be specified for
	// Workflows, MessageBus, or Topic destinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/eventarc_pipeline#network_attachment EventarcPipeline#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
}

