// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateVolumesSecretItems struct {
	// The relative path of the secret in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/cloud_run_v2_worker_pool#path CloudRunV2WorkerPool#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal).
	//
	// If 0 or not set, the Volume's default mode will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/cloud_run_v2_worker_pool#mode CloudRunV2WorkerPool#mode}
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/cloud_run_v2_worker_pool#version CloudRunV2WorkerPool#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

