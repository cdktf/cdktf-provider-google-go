// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeserviceattachment


type ComputeServiceAttachmentConsumerAcceptLists struct {
	// The number of consumer forwarding rules the consumer project can create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_service_attachment#connection_limit ComputeServiceAttachment#connection_limit}
	ConnectionLimit *float64 `field:"required" json:"connectionLimit" yaml:"connectionLimit"`
	// The network that is allowed to connect to this service attachment. Only one of project_id_or_num and network_url may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_service_attachment#network_url ComputeServiceAttachment#network_url}
	NetworkUrl *string `field:"optional" json:"networkUrl" yaml:"networkUrl"`
	// A project that is allowed to connect to this service attachment. Only one of project_id_or_num and network_url may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_service_attachment#project_id_or_num ComputeServiceAttachment#project_id_or_num}
	ProjectIdOrNum *string `field:"optional" json:"projectIdOrNum" yaml:"projectIdOrNum"`
}

