// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctions2function


type Cloudfunctions2FunctionBuildConfigSource struct {
	// repo_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/cloudfunctions2_function#repo_source Cloudfunctions2Function#repo_source}
	RepoSource *Cloudfunctions2FunctionBuildConfigSourceRepoSource `field:"optional" json:"repoSource" yaml:"repoSource"`
	// storage_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/cloudfunctions2_function#storage_source Cloudfunctions2Function#storage_source}
	StorageSource *Cloudfunctions2FunctionBuildConfigSourceStorageSource `field:"optional" json:"storageSource" yaml:"storageSource"`
}

