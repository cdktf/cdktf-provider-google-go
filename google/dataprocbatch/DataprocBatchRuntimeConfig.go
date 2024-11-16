// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchRuntimeConfig struct {
	// Optional custom container image for the job runtime environment. If not specified, a default container image will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/dataproc_batch#container_image DataprocBatch#container_image}
	ContainerImage *string `field:"optional" json:"containerImage" yaml:"containerImage"`
	// A mapping of property names to values, which are used to configure workload execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/dataproc_batch#properties DataprocBatch#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Version of the batch runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/dataproc_batch#version DataprocBatch#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

