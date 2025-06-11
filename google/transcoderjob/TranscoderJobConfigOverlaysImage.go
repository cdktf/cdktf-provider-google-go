// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigOverlaysImage struct {
	// URI of the image in Cloud Storage. For example, gs://bucket/inputs/image.png.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/transcoder_job#uri TranscoderJob#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

