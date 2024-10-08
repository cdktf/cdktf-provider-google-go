// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterAuthorization struct {
	// Groups that can perform operations as a cluster admin.
	//
	// A managed
	// ClusterRoleBinding will be created to grant the 'cluster-admin' ClusterRole
	// to the groups. Up to ten admin groups can be provided.
	//
	// For more info on RBAC, see
	// https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/container_attached_cluster#admin_groups ContainerAttachedCluster#admin_groups}
	AdminGroups *[]*string `field:"optional" json:"adminGroups" yaml:"adminGroups"`
	// Users that can perform operations as a cluster admin.
	//
	// A managed
	// ClusterRoleBinding will be created to grant the 'cluster-admin' ClusterRole
	// to the users. Up to ten admin users can be provided.
	//
	// For more info on RBAC, see
	// https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/container_attached_cluster#admin_users ContainerAttachedCluster#admin_users}
	AdminUsers *[]*string `field:"optional" json:"adminUsers" yaml:"adminUsers"`
}

