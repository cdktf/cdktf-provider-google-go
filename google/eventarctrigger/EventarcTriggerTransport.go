// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarctrigger


type EventarcTriggerTransport struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/eventarc_trigger#pubsub EventarcTrigger#pubsub}
	Pubsub *EventarcTriggerTransportPubsub `field:"optional" json:"pubsub" yaml:"pubsub"`
}

