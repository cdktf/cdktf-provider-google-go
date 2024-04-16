// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterClusterOperations struct {
	// Whether collection of application logs/metrics should be enabled (in addition to system logs/metrics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/gkeonprem_bare_metal_admin_cluster#enable_application_logs GkeonpremBareMetalAdminCluster#enable_application_logs}
	EnableApplicationLogs interface{} `field:"optional" json:"enableApplicationLogs" yaml:"enableApplicationLogs"`
}

