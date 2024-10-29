// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigElementaryStreamsVideoStream struct {
	// h264 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/transcoder_job_template#h264 TranscoderJobTemplate#h264}
	H264 *TranscoderJobTemplateConfigElementaryStreamsVideoStreamH264 `field:"optional" json:"h264" yaml:"h264"`
}

