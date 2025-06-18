// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigA struct {
	// ad_breaks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#ad_breaks TranscoderJobTemplate#ad_breaks}
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// edit_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#edit_list TranscoderJobTemplate#edit_list}
	EditList interface{} `field:"optional" json:"editList" yaml:"editList"`
	// elementary_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#elementary_streams TranscoderJobTemplate#elementary_streams}
	ElementaryStreams interface{} `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// encryptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#encryptions TranscoderJobTemplate#encryptions}
	Encryptions interface{} `field:"optional" json:"encryptions" yaml:"encryptions"`
	// inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#inputs TranscoderJobTemplate#inputs}
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
	// manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#manifests TranscoderJobTemplate#manifests}
	Manifests interface{} `field:"optional" json:"manifests" yaml:"manifests"`
	// mux_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#mux_streams TranscoderJobTemplate#mux_streams}
	MuxStreams interface{} `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#output TranscoderJobTemplate#output}
	Output *TranscoderJobTemplateConfigOutput `field:"optional" json:"output" yaml:"output"`
	// overlays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#overlays TranscoderJobTemplate#overlays}
	Overlays interface{} `field:"optional" json:"overlays" yaml:"overlays"`
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/transcoder_job_template#pubsub_destination TranscoderJobTemplate#pubsub_destination}
	PubsubDestination *TranscoderJobTemplateConfigPubsubDestination `field:"optional" json:"pubsubDestination" yaml:"pubsubDestination"`
}

