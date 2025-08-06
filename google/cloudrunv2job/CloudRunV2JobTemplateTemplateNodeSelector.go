// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateNodeSelector struct {
	// The GPU to attach to an instance. See https://cloud.google.com/run/docs/configuring/jobs/gpu for configuring GPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_job#accelerator CloudRunV2Job#accelerator}
	Accelerator *string `field:"required" json:"accelerator" yaml:"accelerator"`
}

