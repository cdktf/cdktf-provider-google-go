// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalAdminClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#location GkeonpremBareMetalAdminCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The bare metal admin cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#name GkeonpremBareMetalAdminCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Annotations on the Bare Metal Admin Cluster.
	//
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#annotations GkeonpremBareMetalAdminCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// A human readable description of this Bare Metal Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#bare_metal_version GkeonpremBareMetalAdminCluster#bare_metal_version}
	BareMetalVersion *string `field:"optional" json:"bareMetalVersion" yaml:"bareMetalVersion"`
	// cluster_operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#cluster_operations GkeonpremBareMetalAdminCluster#cluster_operations}
	ClusterOperations *GkeonpremBareMetalAdminClusterClusterOperations `field:"optional" json:"clusterOperations" yaml:"clusterOperations"`
	// control_plane block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#control_plane GkeonpremBareMetalAdminCluster#control_plane}
	ControlPlane *GkeonpremBareMetalAdminClusterControlPlane `field:"optional" json:"controlPlane" yaml:"controlPlane"`
	// A human readable description of this Bare Metal Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#description GkeonpremBareMetalAdminCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#id GkeonpremBareMetalAdminCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#load_balancer GkeonpremBareMetalAdminCluster#load_balancer}
	LoadBalancer *GkeonpremBareMetalAdminClusterLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// maintenance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#maintenance_config GkeonpremBareMetalAdminCluster#maintenance_config}
	MaintenanceConfig *GkeonpremBareMetalAdminClusterMaintenanceConfig `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#network_config GkeonpremBareMetalAdminCluster#network_config}
	NetworkConfig *GkeonpremBareMetalAdminClusterNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// node_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#node_access_config GkeonpremBareMetalAdminCluster#node_access_config}
	NodeAccessConfig *GkeonpremBareMetalAdminClusterNodeAccessConfig `field:"optional" json:"nodeAccessConfig" yaml:"nodeAccessConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#node_config GkeonpremBareMetalAdminCluster#node_config}
	NodeConfig *GkeonpremBareMetalAdminClusterNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#project GkeonpremBareMetalAdminCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#proxy GkeonpremBareMetalAdminCluster#proxy}
	Proxy *GkeonpremBareMetalAdminClusterProxy `field:"optional" json:"proxy" yaml:"proxy"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#security_config GkeonpremBareMetalAdminCluster#security_config}
	SecurityConfig *GkeonpremBareMetalAdminClusterSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#storage GkeonpremBareMetalAdminCluster#storage}
	Storage *GkeonpremBareMetalAdminClusterStorage `field:"optional" json:"storage" yaml:"storage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/gkeonprem_bare_metal_admin_cluster#timeouts GkeonpremBareMetalAdminCluster#timeouts}
	Timeouts *GkeonpremBareMetalAdminClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

