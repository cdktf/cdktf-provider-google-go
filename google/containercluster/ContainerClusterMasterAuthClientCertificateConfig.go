// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMasterAuthClientCertificateConfig struct {
	// Whether client certificate authorization is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/container_cluster#issue_client_certificate ContainerCluster#issue_client_certificate}
	IssueClientCertificate interface{} `field:"required" json:"issueClientCertificate" yaml:"issueClientCertificate"`
}

