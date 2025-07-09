// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappstoragepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappStoragePoolConfig struct {
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
	// Capacity of the storage pool (in GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#capacity_gib NetappStoragePool#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// Name of the location. For zonal Flex pools specify a zone name, in all other cases a region name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#location NetappStoragePool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the storage pool. Needs to be unique per location/region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#name NetappStoragePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// VPC network name with format: 'projects/{{project}}/global/networks/{{network}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#network NetappStoragePool#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Service level of the storage pool. Possible values: ["PREMIUM", "EXTREME", "STANDARD", "FLEX"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#service_level NetappStoragePool#service_level}
	ServiceLevel *string `field:"required" json:"serviceLevel" yaml:"serviceLevel"`
	// Specifies the Active Directory policy to be used.
	//
	// Format: 'projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}'.
	// The policy needs to be in the same location as the storage pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#active_directory NetappStoragePool#active_directory}
	ActiveDirectory *string `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Optional.
	//
	// True if the storage pool supports Auto Tiering enabled volumes. Default is false.
	// Auto-tiering can be enabled after storage pool creation but it can't be disabled once enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#allow_auto_tiering NetappStoragePool#allow_auto_tiering}
	AllowAutoTiering interface{} `field:"optional" json:"allowAutoTiering" yaml:"allowAutoTiering"`
	// Optional. True if using Independent Scaling of capacity and performance (Hyperdisk). Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#custom_performance_enabled NetappStoragePool#custom_performance_enabled}
	CustomPerformanceEnabled interface{} `field:"optional" json:"customPerformanceEnabled" yaml:"customPerformanceEnabled"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#description NetappStoragePool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#id NetappStoragePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the CMEK policy to be used for volume encryption.
	//
	// Format: 'projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}'.
	// The policy needs to be in the same location as the storage pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#kms_config NetappStoragePool#kms_config}
	KmsConfig *string `field:"optional" json:"kmsConfig" yaml:"kmsConfig"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#labels NetappStoragePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// When enabled, the volumes uses Active Directory as LDAP name service for UID/GID lookups.
	//
	// Required to enable extended group support for NFSv3,
	// using security identifiers for NFSv4.1 or principal names for kerberized NFSv4.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#ldap_enabled NetappStoragePool#ldap_enabled}
	LdapEnabled interface{} `field:"optional" json:"ldapEnabled" yaml:"ldapEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#project NetappStoragePool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Specifies the replica zone for regional Flex pools. 'zone' and 'replica_zone' values can be swapped to initiate a [zone switch](https://cloud.google.com/netapp/volumes/docs/configure-and-use/storage-pools/edit-or-delete-storage-pool#switch_active_and_replica_zones).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#replica_zone NetappStoragePool#replica_zone}
	ReplicaZone *string `field:"optional" json:"replicaZone" yaml:"replicaZone"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#timeouts NetappStoragePool#timeouts}
	Timeouts *NetappStoragePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional. Custom Performance Total IOPS of the pool If not provided, it will be calculated based on the totalThroughputMibps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#total_iops NetappStoragePool#total_iops}
	TotalIops *string `field:"optional" json:"totalIops" yaml:"totalIops"`
	// Optional. Custom Performance Total Throughput of the pool (in MiB/s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#total_throughput_mibps NetappStoragePool#total_throughput_mibps}
	TotalThroughputMibps *string `field:"optional" json:"totalThroughputMibps" yaml:"totalThroughputMibps"`
	// Specifies the active zone for regional Flex pools.
	//
	// 'zone' and 'replica_zone' values can be swapped to initiate a
	// [zone switch](https://cloud.google.com/netapp/volumes/docs/configure-and-use/storage-pools/edit-or-delete-storage-pool#switch_active_and_replica_zones).
	// If you want to create a zonal Flex pool, specify a zone name for 'location' and omit 'zone'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_storage_pool#zone NetappStoragePool#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

