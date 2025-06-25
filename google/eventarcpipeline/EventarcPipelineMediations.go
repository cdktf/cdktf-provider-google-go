// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineMediations struct {
	// transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/eventarc_pipeline#transformation EventarcPipeline#transformation}
	Transformation *EventarcPipelineMediationsTransformation `field:"optional" json:"transformation" yaml:"transformation"`
}

