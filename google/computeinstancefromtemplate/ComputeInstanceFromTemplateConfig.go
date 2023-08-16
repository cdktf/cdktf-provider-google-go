package computeinstancefromtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceFromTemplateConfig struct {
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
	// The name of the instance. One of name or self_link must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#name ComputeInstanceFromTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name or self link of an instance template to create the instance based on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#source_instance_template ComputeInstanceFromTemplate#source_instance_template}
	SourceInstanceTemplate *string `field:"required" json:"sourceInstanceTemplate" yaml:"sourceInstanceTemplate"`
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#advanced_machine_features ComputeInstanceFromTemplate#advanced_machine_features}
	AdvancedMachineFeatures *ComputeInstanceFromTemplateAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// If true, allows Terraform to stop the instance to update its properties.
	//
	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#allow_stopping_for_update ComputeInstanceFromTemplate#allow_stopping_for_update}
	AllowStoppingForUpdate interface{} `field:"optional" json:"allowStoppingForUpdate" yaml:"allowStoppingForUpdate"`
	// List of disks attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#attached_disk ComputeInstanceFromTemplate#attached_disk}
	AttachedDisk interface{} `field:"optional" json:"attachedDisk" yaml:"attachedDisk"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#boot_disk ComputeInstanceFromTemplate#boot_disk}
	BootDisk *ComputeInstanceFromTemplateBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#can_ip_forward ComputeInstanceFromTemplate#can_ip_forward}
	CanIpForward interface{} `field:"optional" json:"canIpForward" yaml:"canIpForward"`
	// confidential_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#confidential_instance_config ComputeInstanceFromTemplate#confidential_instance_config}
	ConfidentialInstanceConfig *ComputeInstanceFromTemplateConfidentialInstanceConfig `field:"optional" json:"confidentialInstanceConfig" yaml:"confidentialInstanceConfig"`
	// Whether deletion protection is enabled on this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#deletion_protection ComputeInstanceFromTemplate#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A brief description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#description ComputeInstanceFromTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#desired_status ComputeInstanceFromTemplate#desired_status}
	DesiredStatus *string `field:"optional" json:"desiredStatus" yaml:"desiredStatus"`
	// Whether the instance has virtual displays enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#enable_display ComputeInstanceFromTemplate#enable_display}
	EnableDisplay interface{} `field:"optional" json:"enableDisplay" yaml:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#guest_accelerator ComputeInstanceFromTemplate#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// A custom hostname for the instance.
	//
	// Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#hostname ComputeInstanceFromTemplate#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#id ComputeInstanceFromTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs assigned to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#labels ComputeInstanceFromTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The machine type to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#machine_type ComputeInstanceFromTemplate#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Metadata key/value pairs made available within the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#metadata ComputeInstanceFromTemplate#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Metadata startup scripts made available within the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#metadata_startup_script ComputeInstanceFromTemplate#metadata_startup_script}
	MetadataStartupScript *string `field:"optional" json:"metadataStartupScript" yaml:"metadataStartupScript"`
	// The minimum CPU platform specified for the VM instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#min_cpu_platform ComputeInstanceFromTemplate#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#network_interface ComputeInstanceFromTemplate#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// network_performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#network_performance_config ComputeInstanceFromTemplate#network_performance_config}
	NetworkPerformanceConfig *ComputeInstanceFromTemplateNetworkPerformanceConfig `field:"optional" json:"networkPerformanceConfig" yaml:"networkPerformanceConfig"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#params ComputeInstanceFromTemplate#params}
	Params *ComputeInstanceFromTemplateParams `field:"optional" json:"params" yaml:"params"`
	// The ID of the project in which the resource belongs.
	//
	// If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#project ComputeInstanceFromTemplate#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#reservation_affinity ComputeInstanceFromTemplate#reservation_affinity}
	ReservationAffinity *ComputeInstanceFromTemplateReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// A list of self_links of resource policies to attach to the instance.
	//
	// Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#resource_policies ComputeInstanceFromTemplate#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#scheduling ComputeInstanceFromTemplate#scheduling}
	Scheduling *ComputeInstanceFromTemplateScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// The scratch disks attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#scratch_disk ComputeInstanceFromTemplate#scratch_disk}
	ScratchDisk interface{} `field:"optional" json:"scratchDisk" yaml:"scratchDisk"`
	// The service account to attach to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#service_account ComputeInstanceFromTemplate#service_account}
	ServiceAccount interface{} `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#shielded_instance_config ComputeInstanceFromTemplate#shielded_instance_config}
	ShieldedInstanceConfig *ComputeInstanceFromTemplateShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// The list of tags attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#tags ComputeInstanceFromTemplate#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#timeouts ComputeInstanceFromTemplate#timeouts}
	Timeouts *ComputeInstanceFromTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone of the instance.
	//
	// If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#zone ComputeInstanceFromTemplate#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

