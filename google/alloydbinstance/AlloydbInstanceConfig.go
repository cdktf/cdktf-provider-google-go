// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlloydbInstanceConfig struct {
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
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#cluster AlloydbInstance#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The ID of the alloydb instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#instance_id AlloydbInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The type of the instance.
	//
	// If the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the 'depends_on' meta-data attribute.
	// If the instance type is SECONDARY, point to the cluster_type of the associated secondary cluster instead of mentioning SECONDARY.
	// Example: {instance_type = google_alloydb_cluster.<secondary_cluster_name>.cluster_type} instead of {instance_type = SECONDARY}
	// If the instance type is SECONDARY, the terraform delete instance operation does not delete the secondary instance but abandons it instead.
	// Use deletion_policy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.
	// Users can undo the delete secondary instance action by importing the deleted secondary instance by calling terraform import. Possible values: ["PRIMARY", "READ_POOL", "SECONDARY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#instance_type AlloydbInstance#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// 'Specifies whether an instance needs to spin up.
	//
	// Once the instance is
	// active, the activation policy can be updated to the 'NEVER' to stop the
	// instance. Likewise, the activation policy can be updated to 'ALWAYS' to
	// start the instance.
	// There are restrictions around when an instance can/cannot be activated (for
	// example, a read pool instance should be stopped before stopping primary
	// etc.). Please refer to the API documentation for more details.
	// Possible values are: 'ACTIVATION_POLICY_UNSPECIFIED', 'ALWAYS', 'NEVER'.' Possible values: ["ACTIVATION_POLICY_UNSPECIFIED", "ALWAYS", "NEVER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#activation_policy AlloydbInstance#activation_policy}
	ActivationPolicy *string `field:"optional" json:"activationPolicy" yaml:"activationPolicy"`
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#annotations AlloydbInstance#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// 'Availability type of an Instance.
	//
	// Defaults to REGIONAL for both primary and read instances.
	// Note that primary and read instances can have different availability types.
	// Primary instances can be either ZONAL or REGIONAL. Read Pool instances can also be either ZONAL or REGIONAL.
	// Read pools of size 1 can only have zonal availability. Read pools with a node count of 2 or more
	// can have regional availability (nodes are present in 2 or more zones in a region).
	// Possible values are: 'AVAILABILITY_TYPE_UNSPECIFIED', 'ZONAL', 'REGIONAL'.' Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#availability_type AlloydbInstance#availability_type}
	AvailabilityType *string `field:"optional" json:"availabilityType" yaml:"availabilityType"`
	// client_connection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#client_connection_config AlloydbInstance#client_connection_config}
	ClientConnectionConfig *AlloydbInstanceClientConnectionConfig `field:"optional" json:"clientConnectionConfig" yaml:"clientConnectionConfig"`
	// Database flags.
	//
	// Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#database_flags AlloydbInstance#database_flags}
	DatabaseFlags *map[string]*string `field:"optional" json:"databaseFlags" yaml:"databaseFlags"`
	// User-settable and human-readable display name for the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#display_name AlloydbInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#gce_zone AlloydbInstance#gce_zone}
	GceZone *string `field:"optional" json:"gceZone" yaml:"gceZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#id AlloydbInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the alloydb instance.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#labels AlloydbInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// machine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#machine_config AlloydbInstance#machine_config}
	MachineConfig *AlloydbInstanceMachineConfig `field:"optional" json:"machineConfig" yaml:"machineConfig"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#network_config AlloydbInstance#network_config}
	NetworkConfig *AlloydbInstanceNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// psc_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#psc_instance_config AlloydbInstance#psc_instance_config}
	PscInstanceConfig *AlloydbInstancePscInstanceConfig `field:"optional" json:"pscInstanceConfig" yaml:"pscInstanceConfig"`
	// query_insights_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#query_insights_config AlloydbInstance#query_insights_config}
	QueryInsightsConfig *AlloydbInstanceQueryInsightsConfig `field:"optional" json:"queryInsightsConfig" yaml:"queryInsightsConfig"`
	// read_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#read_pool_config AlloydbInstance#read_pool_config}
	ReadPoolConfig *AlloydbInstanceReadPoolConfig `field:"optional" json:"readPoolConfig" yaml:"readPoolConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/alloydb_instance#timeouts AlloydbInstance#timeouts}
	Timeouts *AlloydbInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

