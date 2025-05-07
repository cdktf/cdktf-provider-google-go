// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigMuxStreamsSegmentSettings struct {
	// Duration of the segments in seconds. The default is '6.0s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/transcoder_job#segment_duration TranscoderJob#segment_duration}
	SegmentDuration *string `field:"optional" json:"segmentDuration" yaml:"segmentDuration"`
}

