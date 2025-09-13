// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagementconnectivitytest


type NetworkManagementConnectivityTestSource struct {
	// app_engine_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#app_engine_version NetworkManagementConnectivityTest#app_engine_version}
	AppEngineVersion *NetworkManagementConnectivityTestSourceAppEngineVersion `field:"optional" json:"appEngineVersion" yaml:"appEngineVersion"`
	// cloud_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#cloud_function NetworkManagementConnectivityTest#cloud_function}
	CloudFunction *NetworkManagementConnectivityTestSourceCloudFunction `field:"optional" json:"cloudFunction" yaml:"cloudFunction"`
	// cloud_run_revision block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#cloud_run_revision NetworkManagementConnectivityTest#cloud_run_revision}
	CloudRunRevision *NetworkManagementConnectivityTestSourceCloudRunRevision `field:"optional" json:"cloudRunRevision" yaml:"cloudRunRevision"`
	// A Cloud SQL instance URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#cloud_sql_instance NetworkManagementConnectivityTest#cloud_sql_instance}
	CloudSqlInstance *string `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// A cluster URI for Google Kubernetes Engine cluster control plane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#gke_master_cluster NetworkManagementConnectivityTest#gke_master_cluster}
	GkeMasterCluster *string `field:"optional" json:"gkeMasterCluster" yaml:"gkeMasterCluster"`
	// A Compute Engine instance URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#instance NetworkManagementConnectivityTest#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// The IP address of the endpoint, which can be an external or internal IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#ip_address NetworkManagementConnectivityTest#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// A VPC network URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#network NetworkManagementConnectivityTest#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Type of the network where the endpoint is located. Possible values: ["GCP_NETWORK", "NON_GCP_NETWORK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#network_type NetworkManagementConnectivityTest#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The IP protocol port of the endpoint. Only applicable when protocol is TCP or UDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#port NetworkManagementConnectivityTest#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Project ID where the endpoint is located.
	//
	// The project ID can be derived from the URI if you provide a endpoint or
	// network URI.
	// The following are two cases where you may need to provide the project ID:
	// 1. Only the IP address is specified, and the IP address is within a Google
	// Cloud project.
	// 2. When you are using Shared VPC and the IP address that you provide is
	// from the service project. In this case, the network that the IP address
	// resides in is defined in the host project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_management_connectivity_test#project_id NetworkManagementConnectivityTest#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

