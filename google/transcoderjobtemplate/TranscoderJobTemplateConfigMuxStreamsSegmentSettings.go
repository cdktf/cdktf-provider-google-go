// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigMuxStreamsSegmentSettings struct {
	// Duration of the segments in seconds. The default is '6.0s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#segment_duration TranscoderJobTemplate#segment_duration}
	SegmentDuration *string `field:"optional" json:"segmentDuration" yaml:"segmentDuration"`
}

