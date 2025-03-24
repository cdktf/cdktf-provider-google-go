// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob


type HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource struct {
	// Directory path where all the Whistle files are located. Example: gs://{bucket-id}/{path/to/import-root/dir}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/healthcare_pipeline_job#import_uri_prefix HealthcarePipelineJob#import_uri_prefix}
	ImportUriPrefix *string `field:"required" json:"importUriPrefix" yaml:"importUriPrefix"`
	// Main configuration file which has the entrypoint or the root function. Example: gs://{bucket-id}/{path/to/import-root/dir}/entrypoint-file-name.wstl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/healthcare_pipeline_job#uri HealthcarePipelineJob#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

