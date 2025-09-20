// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterRbacBindingConfig struct {
	// Setting this to true will allow any ClusterRoleBinding and RoleBinding with subjects system:authenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#enable_insecure_binding_system_authenticated ContainerCluster#enable_insecure_binding_system_authenticated}
	EnableInsecureBindingSystemAuthenticated interface{} `field:"optional" json:"enableInsecureBindingSystemAuthenticated" yaml:"enableInsecureBindingSystemAuthenticated"`
	// Setting this to true will allow any ClusterRoleBinding and RoleBinding with subjects system:anonymous or system:unauthenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#enable_insecure_binding_system_unauthenticated ContainerCluster#enable_insecure_binding_system_unauthenticated}
	EnableInsecureBindingSystemUnauthenticated interface{} `field:"optional" json:"enableInsecureBindingSystemUnauthenticated" yaml:"enableInsecureBindingSystemUnauthenticated"`
}

