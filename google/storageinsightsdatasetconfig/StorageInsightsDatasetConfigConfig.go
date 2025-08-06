// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageInsightsDatasetConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The user-defined ID of the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#dataset_config_id StorageInsightsDatasetConfig#dataset_config_id}
	DatasetConfigId *string `field:"required" json:"datasetConfigId" yaml:"datasetConfigId"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#identity StorageInsightsDatasetConfig#identity}
	Identity *StorageInsightsDatasetConfigIdentity `field:"required" json:"identity" yaml:"identity"`
	// The location of the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#location StorageInsightsDatasetConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Number of days of history that must be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#retention_period_days StorageInsightsDatasetConfig#retention_period_days}
	RetentionPeriodDays *float64 `field:"required" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
	// An optional user-provided description for the dataset configuration with a maximum length of 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#description StorageInsightsDatasetConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// exclude_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#exclude_cloud_storage_buckets StorageInsightsDatasetConfig#exclude_cloud_storage_buckets}
	ExcludeCloudStorageBuckets *StorageInsightsDatasetConfigExcludeCloudStorageBuckets `field:"optional" json:"excludeCloudStorageBuckets" yaml:"excludeCloudStorageBuckets"`
	// exclude_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#exclude_cloud_storage_locations StorageInsightsDatasetConfig#exclude_cloud_storage_locations}
	ExcludeCloudStorageLocations *StorageInsightsDatasetConfigExcludeCloudStorageLocations `field:"optional" json:"excludeCloudStorageLocations" yaml:"excludeCloudStorageLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#id StorageInsightsDatasetConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// include_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#include_cloud_storage_buckets StorageInsightsDatasetConfig#include_cloud_storage_buckets}
	IncludeCloudStorageBuckets *StorageInsightsDatasetConfigIncludeCloudStorageBuckets `field:"optional" json:"includeCloudStorageBuckets" yaml:"includeCloudStorageBuckets"`
	// include_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#include_cloud_storage_locations StorageInsightsDatasetConfig#include_cloud_storage_locations}
	IncludeCloudStorageLocations *StorageInsightsDatasetConfigIncludeCloudStorageLocations `field:"optional" json:"includeCloudStorageLocations" yaml:"includeCloudStorageLocations"`
	// If set to true, the request includes all the newly created buckets in the dataset that meet the inclusion and exclusion rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#include_newly_created_buckets StorageInsightsDatasetConfig#include_newly_created_buckets}
	IncludeNewlyCreatedBuckets interface{} `field:"optional" json:"includeNewlyCreatedBuckets" yaml:"includeNewlyCreatedBuckets"`
	// A boolean terraform only flag to link/unlink dataset.
	//
	// Setting this field to true while creation will automatically link the created dataset as an additional functionality.
	// -> **Note** A dataset config resource can only be destroyed once it is unlinked,
	// so users must set this field to false to unlink the dataset and destroy the dataset config resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#link_dataset StorageInsightsDatasetConfig#link_dataset}
	LinkDataset interface{} `field:"optional" json:"linkDataset" yaml:"linkDataset"`
	// Organization resource ID that the source projects should belong to.
	//
	// Projects that do not belong to the provided organization are not considered when creating the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#organization_number StorageInsightsDatasetConfig#organization_number}
	OrganizationNumber *string `field:"optional" json:"organizationNumber" yaml:"organizationNumber"`
	// Defines the options for providing a source organization for the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#organization_scope StorageInsightsDatasetConfig#organization_scope}
	OrganizationScope interface{} `field:"optional" json:"organizationScope" yaml:"organizationScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#project StorageInsightsDatasetConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// source_folders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#source_folders StorageInsightsDatasetConfig#source_folders}
	SourceFolders *StorageInsightsDatasetConfigSourceFolders `field:"optional" json:"sourceFolders" yaml:"sourceFolders"`
	// source_projects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#source_projects StorageInsightsDatasetConfig#source_projects}
	SourceProjects *StorageInsightsDatasetConfigSourceProjects `field:"optional" json:"sourceProjects" yaml:"sourceProjects"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#timeouts StorageInsightsDatasetConfig#timeouts}
	Timeouts *StorageInsightsDatasetConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

