// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeRegexesPatternsCloudStorageRegex struct {
	// Regex to test the bucket name against.
	//
	// If empty, all buckets match. Example: "marketing2021" or "(marketing)\d{4}" will both match the bucket gs://marketing2021
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/data_loss_prevention_discovery_config#bucket_name_regex DataLossPreventionDiscoveryConfig#bucket_name_regex}
	BucketNameRegex *string `field:"optional" json:"bucketNameRegex" yaml:"bucketNameRegex"`
	// For organizations, if unset, will match all projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/data_loss_prevention_discovery_config#project_id_regex DataLossPreventionDiscoveryConfig#project_id_regex}
	ProjectIdRegex *string `field:"optional" json:"projectIdRegex" yaml:"projectIdRegex"`
}

