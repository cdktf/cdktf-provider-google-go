// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigManifests struct {
	// The name of the generated file. The default is 'manifest'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/transcoder_job_template#file_name TranscoderJobTemplate#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// List of user supplied MuxStream.key values that should appear in this manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/transcoder_job_template#mux_streams TranscoderJobTemplate#mux_streams}
	MuxStreams *[]*string `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// Type of the manifest. Possible values: ["MANIFEST_TYPE_UNSPECIFIED", "HLS", "DASH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/transcoder_job_template#type TranscoderJobTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

