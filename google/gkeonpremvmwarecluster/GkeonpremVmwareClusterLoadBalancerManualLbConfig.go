// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterLoadBalancerManualLbConfig struct {
	// NodePort for control plane service.
	//
	// The Kubernetes API server in the admin
	// cluster is implemented as a Service of type NodePort (ex. 30968).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gkeonprem_vmware_cluster#control_plane_node_port GkeonpremVmwareCluster#control_plane_node_port}
	ControlPlaneNodePort *float64 `field:"optional" json:"controlPlaneNodePort" yaml:"controlPlaneNodePort"`
	// NodePort for ingress service's http.
	//
	// The ingress service in the admin
	// cluster is implemented as a Service of type NodePort (ex. 32527).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gkeonprem_vmware_cluster#ingress_http_node_port GkeonpremVmwareCluster#ingress_http_node_port}
	IngressHttpNodePort *float64 `field:"optional" json:"ingressHttpNodePort" yaml:"ingressHttpNodePort"`
	// NodePort for ingress service's https.
	//
	// The ingress service in the admin
	// cluster is implemented as a Service of type NodePort (ex. 30139).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gkeonprem_vmware_cluster#ingress_https_node_port GkeonpremVmwareCluster#ingress_https_node_port}
	IngressHttpsNodePort *float64 `field:"optional" json:"ingressHttpsNodePort" yaml:"ingressHttpsNodePort"`
	// NodePort for konnectivity server service running as a sidecar in each kube-apiserver pod (ex. 30564).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gkeonprem_vmware_cluster#konnectivity_server_node_port GkeonpremVmwareCluster#konnectivity_server_node_port}
	KonnectivityServerNodePort *float64 `field:"optional" json:"konnectivityServerNodePort" yaml:"konnectivityServerNodePort"`
}

