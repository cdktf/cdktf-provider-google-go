// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateContainersEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER, and may not exceed 32768 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/cloud_run_v2_worker_pool#name CloudRunV2WorkerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Literal value of the environment variable.
	//
	// Defaults to "" and the maximum allowed length is 32768 characters. Variable references are not supported in Cloud Run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/cloud_run_v2_worker_pool#value CloudRunV2WorkerPool#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/cloud_run_v2_worker_pool#value_source CloudRunV2WorkerPool#value_source}
	ValueSource *CloudRunV2WorkerPoolTemplateContainersEnvValueSource `field:"optional" json:"valueSource" yaml:"valueSource"`
}

