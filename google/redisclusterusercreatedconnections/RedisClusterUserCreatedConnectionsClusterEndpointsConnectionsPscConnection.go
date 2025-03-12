// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redisclusterusercreatedconnections


type RedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnection struct {
	// The IP allocated on the consumer network for the PSC forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/redis_cluster_user_created_connections#address RedisClusterUserCreatedConnections#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// The URI of the consumer side forwarding rule. Format: projects/{project}/regions/{region}/forwardingRules/{forwarding_rule}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/redis_cluster_user_created_connections#forwarding_rule RedisClusterUserCreatedConnections#forwarding_rule}
	ForwardingRule *string `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// The consumer network where the IP address resides, in the form of projects/{project_id}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/redis_cluster_user_created_connections#network RedisClusterUserCreatedConnections#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The PSC connection id of the forwarding rule connected to the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/redis_cluster_user_created_connections#psc_connection_id RedisClusterUserCreatedConnections#psc_connection_id}
	PscConnectionId *string `field:"required" json:"pscConnectionId" yaml:"pscConnectionId"`
	// The service attachment which is the target of the PSC connection, in the form of projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/redis_cluster_user_created_connections#service_attachment RedisClusterUserCreatedConnections#service_attachment}
	ServiceAttachment *string `field:"required" json:"serviceAttachment" yaml:"serviceAttachment"`
	// The consumer project_id where the forwarding rule is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/redis_cluster_user_created_connections#project_id RedisClusterUserCreatedConnections#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

