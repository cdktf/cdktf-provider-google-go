// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolprojectintelligenceconfig


type StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets struct {
	// List of bucket id regexes to exclude in the storage intelligence plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/storage_control_project_intelligence_config#bucket_id_regexes StorageControlProjectIntelligenceConfig#bucket_id_regexes}
	BucketIdRegexes *[]*string `field:"required" json:"bucketIdRegexes" yaml:"bucketIdRegexes"`
}

