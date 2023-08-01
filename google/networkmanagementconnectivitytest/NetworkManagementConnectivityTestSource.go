package networkmanagementconnectivitytest


type NetworkManagementConnectivityTestSource struct {
	// A Compute Engine instance URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_management_connectivity_test#instance NetworkManagementConnectivityTest#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// The IP address of the endpoint, which can be an external or internal IP.
	//
	// An IPv6 address is only allowed when the test's
	// destination is a global load balancer VIP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_management_connectivity_test#ip_address NetworkManagementConnectivityTest#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// A Compute Engine network URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_management_connectivity_test#network NetworkManagementConnectivityTest#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Type of the network where the endpoint is located. Possible values: ["GCP_NETWORK", "NON_GCP_NETWORK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_management_connectivity_test#network_type NetworkManagementConnectivityTest#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The IP protocol port of the endpoint. Only applicable when protocol is TCP or UDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_management_connectivity_test#port NetworkManagementConnectivityTest#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Project ID where the endpoint is located.
	//
	// The Project ID can be
	// derived from the URI if you provide a VM instance or network URI.
	// The following are two cases where you must provide the project ID:
	//
	// 1. Only the IP address is specified, and the IP address is
	// within a GCP project.
	// 2. When you are using Shared VPC and the IP address
	// that you provide is from the service project. In this case,
	// the network that the IP address resides in is defined in the
	// host project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_management_connectivity_test#project_id NetworkManagementConnectivityTest#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

