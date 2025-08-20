// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterPlatformConfig struct {
	// The required platform version e.g. 1.13.1. If the current platform version is lower than the target version, the platform version will be updated to the target version. If the target version is not installed in the platform (bundle versions), download the target version bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gkeonprem_vmware_admin_cluster#required_platform_version GkeonpremVmwareAdminCluster#required_platform_version}
	RequiredPlatformVersion *string `field:"optional" json:"requiredPlatformVersion" yaml:"requiredPlatformVersion"`
}

