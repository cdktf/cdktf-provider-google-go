// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectusageexportbucket


type ProjectUsageExportBucketTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/project_usage_export_bucket#create ProjectUsageExportBucket#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/project_usage_export_bucket#delete ProjectUsageExportBucket#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

