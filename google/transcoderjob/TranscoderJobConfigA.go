// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigA struct {
	// ad_breaks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#ad_breaks TranscoderJob#ad_breaks}
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// edit_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#edit_list TranscoderJob#edit_list}
	EditList interface{} `field:"optional" json:"editList" yaml:"editList"`
	// elementary_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#elementary_streams TranscoderJob#elementary_streams}
	ElementaryStreams interface{} `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// encryptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#encryptions TranscoderJob#encryptions}
	Encryptions interface{} `field:"optional" json:"encryptions" yaml:"encryptions"`
	// inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#inputs TranscoderJob#inputs}
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
	// manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#manifests TranscoderJob#manifests}
	Manifests interface{} `field:"optional" json:"manifests" yaml:"manifests"`
	// mux_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#mux_streams TranscoderJob#mux_streams}
	MuxStreams interface{} `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#output TranscoderJob#output}
	Output *TranscoderJobConfigOutput `field:"optional" json:"output" yaml:"output"`
	// overlays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#overlays TranscoderJob#overlays}
	Overlays interface{} `field:"optional" json:"overlays" yaml:"overlays"`
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/transcoder_job#pubsub_destination TranscoderJob#pubsub_destination}
	PubsubDestination *TranscoderJobConfigPubsubDestination `field:"optional" json:"pubsubDestination" yaml:"pubsubDestination"`
}

