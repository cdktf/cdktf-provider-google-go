// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigInputs struct {
	// A unique key for this input. Must be specified when using advanced mapping and edit lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/transcoder_job_template#key TranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// URI of the media.
	//
	// Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, gs://bucket/inputs/file.mp4).
	// If empty, the value is populated from Job.input_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/transcoder_job_template#uri TranscoderJobTemplate#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

