// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateVpcAccessNetworkInterfaces struct {
	// The VPC network that the Cloud Run resource will be able to send traffic to.
	//
	// At least one of network or subnetwork must be specified. If both
	// network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If network is not specified, it will be
	// looked up from the subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#network CloudRunV2WorkerPool#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The VPC subnetwork that the Cloud Run resource will get IPs from.
	//
	// At least one of network or subnetwork must be specified. If both
	// network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the
	// subnetwork with the same name with the network will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#subnetwork CloudRunV2WorkerPool#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// Network tags applied to this Cloud Run WorkerPool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#tags CloudRunV2WorkerPool#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

