// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetConditionsCloudStorageConditions struct {
	// Only objects with the specified attributes will be scanned. Defaults to [ALL_SUPPORTED_BUCKETS] if unset. Possible values: ["ALL_SUPPORTED_BUCKETS", "AUTOCLASS_DISABLED", "AUTOCLASS_ENABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/data_loss_prevention_discovery_config#included_bucket_attributes DataLossPreventionDiscoveryConfig#included_bucket_attributes}
	IncludedBucketAttributes *[]*string `field:"optional" json:"includedBucketAttributes" yaml:"includedBucketAttributes"`
	// Only objects with the specified attributes will be scanned.
	//
	// If an object has one of the specified attributes but is inside an excluded bucket, it will not be scanned. Defaults to [ALL_SUPPORTED_OBJECTS]. A profile will be created even if no objects match the included_object_attributes. Possible values: ["ALL_SUPPORTED_OBJECTS", "STANDARD", "NEARLINE", "COLDLINE", "ARCHIVE", "REGIONAL", "MULTI_REGIONAL", "DURABLE_REDUCED_AVAILABILITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/data_loss_prevention_discovery_config#included_object_attributes DataLossPreventionDiscoveryConfig#included_object_attributes}
	IncludedObjectAttributes *[]*string `field:"optional" json:"includedObjectAttributes" yaml:"includedObjectAttributes"`
}

