// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memcacheinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemcacheInstanceConfig struct {
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
	// The resource name of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#name MemcacheInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#node_config MemcacheInstance#node_config}
	NodeConfig *MemcacheInstanceNodeConfig `field:"required" json:"nodeConfig" yaml:"nodeConfig"`
	// Number of nodes in the memcache instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#node_count MemcacheInstance#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// The full name of the GCE network to connect the instance to.  If not provided, 'default' will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#authorized_network MemcacheInstance#authorized_network}
	AuthorizedNetwork *string `field:"optional" json:"authorizedNetwork" yaml:"authorizedNetwork"`
	// A user-visible name for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#display_name MemcacheInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#id MemcacheInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#labels MemcacheInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#maintenance_policy MemcacheInstance#maintenance_policy}
	MaintenancePolicy *MemcacheInstanceMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// memcache_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#memcache_parameters MemcacheInstance#memcache_parameters}
	MemcacheParameters *MemcacheInstanceMemcacheParameters `field:"optional" json:"memcacheParameters" yaml:"memcacheParameters"`
	// The major version of Memcached software.
	//
	// If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5", "MEMCACHE_1_6_15"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#memcache_version MemcacheInstance#memcache_version}
	MemcacheVersion *string `field:"optional" json:"memcacheVersion" yaml:"memcacheVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#project MemcacheInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#region MemcacheInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Contains the name of allocated IP address ranges associated with the private service access connection for example, "test-default" associated with IP range 10.0.0.0/29.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#reserved_ip_range_id MemcacheInstance#reserved_ip_range_id}
	ReservedIpRangeId *[]*string `field:"optional" json:"reservedIpRangeId" yaml:"reservedIpRangeId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#timeouts MemcacheInstance#timeouts}
	Timeouts *MemcacheInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zones where memcache nodes should be provisioned.  If not provided, all zones will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/memcache_instance#zones MemcacheInstance#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

