// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingprojectbucketconfig


type LoggingProjectBucketConfigIndexConfigs struct {
	// The LogEntry field path to index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/logging_project_bucket_config#field_path LoggingProjectBucketConfig#field_path}
	FieldPath *string `field:"required" json:"fieldPath" yaml:"fieldPath"`
	// The type of data in this index Note that some paths are automatically indexed, and other paths are not eligible for indexing.
	//
	// See [indexing documentation]( https://cloud.google.com/logging/docs/view/advanced-queries#indexed-fields) for details.
	// For example: jsonPayload.request.status
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/logging_project_bucket_config#type LoggingProjectBucketConfig#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

