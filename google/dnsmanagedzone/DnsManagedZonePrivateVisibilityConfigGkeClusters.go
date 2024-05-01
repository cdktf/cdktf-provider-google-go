// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsmanagedzone


type DnsManagedZonePrivateVisibilityConfigGkeClusters struct {
	// The resource name of the cluster to bind this ManagedZone to.
	//
	// This should be specified in the format like
	// 'projects/* /locations/* /clusters/*'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/dns_managed_zone#gke_cluster_name DnsManagedZone#gke_cluster_name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	GkeClusterName *string `field:"required" json:"gkeClusterName" yaml:"gkeClusterName"`
}

