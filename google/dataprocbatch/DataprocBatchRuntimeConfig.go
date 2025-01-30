// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchRuntimeConfig struct {
	// autotuning_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/dataproc_batch#autotuning_config DataprocBatch#autotuning_config}
	AutotuningConfig *DataprocBatchRuntimeConfigAutotuningConfig `field:"optional" json:"autotuningConfig" yaml:"autotuningConfig"`
	// Optional. Cohort identifier. Identifies families of the workloads having the same shape, e.g. daily ETL jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/dataproc_batch#cohort DataprocBatch#cohort}
	Cohort *string `field:"optional" json:"cohort" yaml:"cohort"`
	// Optional custom container image for the job runtime environment. If not specified, a default container image will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/dataproc_batch#container_image DataprocBatch#container_image}
	ContainerImage *string `field:"optional" json:"containerImage" yaml:"containerImage"`
	// A mapping of property names to values, which are used to configure workload execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/dataproc_batch#properties DataprocBatch#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Version of the batch runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/dataproc_batch#version DataprocBatch#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

