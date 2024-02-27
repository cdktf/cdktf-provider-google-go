// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterOsEnvironmentConfig struct {
	// Whether the package repo should not be included when initializing bare metal machines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/gkeonprem_bare_metal_cluster#package_repo_excluded GkeonpremBareMetalCluster#package_repo_excluded}
	PackageRepoExcluded interface{} `field:"required" json:"packageRepoExcluded" yaml:"packageRepoExcluded"`
}

