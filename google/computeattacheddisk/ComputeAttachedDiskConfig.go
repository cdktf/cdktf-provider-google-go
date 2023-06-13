package computeattacheddisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeAttachedDiskConfig struct {
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
	// name or self_link of the disk that will be attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#disk ComputeAttachedDisk#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
	// name or self_link of the compute instance that the disk will be attached to.
	//
	// If the self_link is provided then zone and project are extracted from the self link. If only the name is used then zone and project must be defined as properties on the resource or provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#instance ComputeAttachedDisk#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance.
	//
	// This name can be used to reference the device for mounting, resizing, and so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#device_name ComputeAttachedDisk#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#id ComputeAttachedDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY.
	//
	// If not specified, the default is to attach the disk in READ_WRITE mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#mode ComputeAttachedDisk#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The project that the referenced compute instance is a part of.
	//
	// If instance is referenced by its self_link the project defined in the link will take precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#project ComputeAttachedDisk#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#timeouts ComputeAttachedDisk#timeouts}
	Timeouts *ComputeAttachedDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone that the referenced compute instance is located within.
	//
	// If instance is referenced by its self_link the zone defined in the link will take precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_attached_disk#zone ComputeAttachedDisk#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

