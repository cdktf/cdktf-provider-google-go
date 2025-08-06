// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstancedesiredusercreatedendpoints


type MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection struct {
	// The URI of the consumer side forwarding rule. Format: projects/{project}/regions/{region}/forwardingRules/{forwarding_rule}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance_desired_user_created_endpoints#forwarding_rule MemorystoreInstanceDesiredUserCreatedEndpoints#forwarding_rule}
	ForwardingRule *string `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// The IP allocated on the consumer network for the PSC forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance_desired_user_created_endpoints#ip_address MemorystoreInstanceDesiredUserCreatedEndpoints#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The consumer network where the IP address resides, in the form of projects/{project_id}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance_desired_user_created_endpoints#network MemorystoreInstanceDesiredUserCreatedEndpoints#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The PSC connection id of the forwarding rule connected to the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance_desired_user_created_endpoints#psc_connection_id MemorystoreInstanceDesiredUserCreatedEndpoints#psc_connection_id}
	PscConnectionId *string `field:"required" json:"pscConnectionId" yaml:"pscConnectionId"`
	// The service attachment which is the target of the PSC connection, in the form of projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance_desired_user_created_endpoints#service_attachment MemorystoreInstanceDesiredUserCreatedEndpoints#service_attachment}
	ServiceAttachment *string `field:"required" json:"serviceAttachment" yaml:"serviceAttachment"`
	// The consumer project_id where the forwarding rule is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance_desired_user_created_endpoints#project_id MemorystoreInstanceDesiredUserCreatedEndpoints#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

