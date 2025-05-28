// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecomputestoragepooltypes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeStoragePoolTypesConfig struct {
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
	// Name of the storage pool type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/compute_storage_pool_types#storage_pool_type DataGoogleComputeStoragePoolTypes#storage_pool_type}
	StoragePoolType *string `field:"required" json:"storagePoolType" yaml:"storagePoolType"`
	// The name of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/compute_storage_pool_types#zone DataGoogleComputeStoragePoolTypes#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/compute_storage_pool_types#project DataGoogleComputeStoragePoolTypes#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

