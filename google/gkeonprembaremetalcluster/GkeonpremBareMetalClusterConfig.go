// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalClusterConfig struct {
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
	// The Admin Cluster this Bare Metal User Cluster belongs to.
	//
	// This is the full resource name of the Admin Cluster's hub membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#admin_cluster_membership GkeonpremBareMetalCluster#admin_cluster_membership}
	AdminClusterMembership *string `field:"required" json:"adminClusterMembership" yaml:"adminClusterMembership"`
	// A human readable description of this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#bare_metal_version GkeonpremBareMetalCluster#bare_metal_version}
	BareMetalVersion *string `field:"required" json:"bareMetalVersion" yaml:"bareMetalVersion"`
	// control_plane block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#control_plane GkeonpremBareMetalCluster#control_plane}
	ControlPlane *GkeonpremBareMetalClusterControlPlane `field:"required" json:"controlPlane" yaml:"controlPlane"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#load_balancer GkeonpremBareMetalCluster#load_balancer}
	LoadBalancer *GkeonpremBareMetalClusterLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#location GkeonpremBareMetalCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The bare metal cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#name GkeonpremBareMetalCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#network_config GkeonpremBareMetalCluster#network_config}
	NetworkConfig *GkeonpremBareMetalClusterNetworkConfig `field:"required" json:"networkConfig" yaml:"networkConfig"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#storage GkeonpremBareMetalCluster#storage}
	Storage *GkeonpremBareMetalClusterStorage `field:"required" json:"storage" yaml:"storage"`
	// Annotations on the Bare Metal User Cluster.
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#annotations GkeonpremBareMetalCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#binary_authorization GkeonpremBareMetalCluster#binary_authorization}
	BinaryAuthorization *GkeonpremBareMetalClusterBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// cluster_operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#cluster_operations GkeonpremBareMetalCluster#cluster_operations}
	ClusterOperations *GkeonpremBareMetalClusterClusterOperations `field:"optional" json:"clusterOperations" yaml:"clusterOperations"`
	// A human readable description of this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#description GkeonpremBareMetalCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#id GkeonpremBareMetalCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#maintenance_config GkeonpremBareMetalCluster#maintenance_config}
	MaintenanceConfig *GkeonpremBareMetalClusterMaintenanceConfig `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// node_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#node_access_config GkeonpremBareMetalCluster#node_access_config}
	NodeAccessConfig *GkeonpremBareMetalClusterNodeAccessConfig `field:"optional" json:"nodeAccessConfig" yaml:"nodeAccessConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#node_config GkeonpremBareMetalCluster#node_config}
	NodeConfig *GkeonpremBareMetalClusterNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// os_environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#os_environment_config GkeonpremBareMetalCluster#os_environment_config}
	OsEnvironmentConfig *GkeonpremBareMetalClusterOsEnvironmentConfig `field:"optional" json:"osEnvironmentConfig" yaml:"osEnvironmentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#project GkeonpremBareMetalCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#proxy GkeonpremBareMetalCluster#proxy}
	Proxy *GkeonpremBareMetalClusterProxy `field:"optional" json:"proxy" yaml:"proxy"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#security_config GkeonpremBareMetalCluster#security_config}
	SecurityConfig *GkeonpremBareMetalClusterSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#timeouts GkeonpremBareMetalCluster#timeouts}
	Timeouts *GkeonpremBareMetalClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#upgrade_policy GkeonpremBareMetalCluster#upgrade_policy}
	UpgradePolicy *GkeonpremBareMetalClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
}

