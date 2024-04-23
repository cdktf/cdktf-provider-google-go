// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterControlPlaneLocal struct {
	// Only machines matching this filter will be allowed to host control plane nodes.
	//
	// The filtering language accepts strings like "name=<name>",
	// and is documented here: [AIP-160](https://google.aip.dev/160).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/edgecontainer_cluster#machine_filter EdgecontainerCluster#machine_filter}
	MachineFilter *string `field:"optional" json:"machineFilter" yaml:"machineFilter"`
	// The number of nodes to serve as replicas of the Control Plane. Only 1 and 3 are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/edgecontainer_cluster#node_count EdgecontainerCluster#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Name of the Google Distributed Cloud Edge zones where this node pool will be created. For example: 'us-central1-edge-customer-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/edgecontainer_cluster#node_location EdgecontainerCluster#node_location}
	NodeLocation *string `field:"optional" json:"nodeLocation" yaml:"nodeLocation"`
	// Policy configuration about how user applications are deployed. Possible values: ["SHARED_DEPLOYMENT_POLICY_UNSPECIFIED", "ALLOWED", "DISALLOWED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/edgecontainer_cluster#shared_deployment_policy EdgecontainerCluster#shared_deployment_policy}
	SharedDeploymentPolicy *string `field:"optional" json:"sharedDeploymentPolicy" yaml:"sharedDeploymentPolicy"`
}

