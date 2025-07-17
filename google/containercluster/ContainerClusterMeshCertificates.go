// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMeshCertificates struct {
	// When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/container_cluster#enable_certificates ContainerCluster#enable_certificates}
	EnableCertificates interface{} `field:"required" json:"enableCertificates" yaml:"enableCertificates"`
}

