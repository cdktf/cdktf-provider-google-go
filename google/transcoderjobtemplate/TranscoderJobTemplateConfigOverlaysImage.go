// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigOverlaysImage struct {
	// URI of the image in Cloud Storage. For example, gs://bucket/inputs/image.png.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/transcoder_job_template#uri TranscoderJobTemplate#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

