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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#cluster AlloydbInstance#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The ID of the alloydb instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#instance_id AlloydbInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The type of the instance.
	//
	// If the instance type is READ_POOL, provide the associated PRIMARY instance in the 'depends_on' meta-data attribute. Possible values: ["PRIMARY", "READ_POOL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#instance_type AlloydbInstance#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#annotations AlloydbInstance#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Availability type of an Instance.
	//
	// Defaults to REGIONAL for both primary and read instances. Note that primary and read instances can have different availability types. Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#availability_type AlloydbInstance#availability_type}
	AvailabilityType *string `field:"optional" json:"availabilityType" yaml:"availabilityType"`
	// Database flags.
	//
	// Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#database_flags AlloydbInstance#database_flags}
	DatabaseFlags *map[string]*string `field:"optional" json:"databaseFlags" yaml:"databaseFlags"`
	// User-settable and human-readable display name for the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#display_name AlloydbInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#gce_zone AlloydbInstance#gce_zone}
	GceZone *string `field:"optional" json:"gceZone" yaml:"gceZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#id AlloydbInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the alloydb instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#labels AlloydbInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// machine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#machine_config AlloydbInstance#machine_config}
	MachineConfig *AlloydbInstanceMachineConfig `field:"optional" json:"machineConfig" yaml:"machineConfig"`
	// read_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#read_pool_config AlloydbInstance#read_pool_config}
	ReadPoolConfig *AlloydbInstanceReadPoolConfig `field:"optional" json:"readPoolConfig" yaml:"readPoolConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/alloydb_instance#timeouts AlloydbInstance#timeouts}
	Timeouts *AlloydbInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

