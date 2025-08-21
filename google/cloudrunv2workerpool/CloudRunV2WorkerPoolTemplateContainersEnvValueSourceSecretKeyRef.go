// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateContainersEnvValueSourceSecretKeyRef struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// Format: {secretName} if the secret is in the same project. projects/{project}/secrets/{secretName} if the secret is in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_run_v2_worker_pool#secret CloudRunV2WorkerPool#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_run_v2_worker_pool#version CloudRunV2WorkerPool#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

