package computeimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeImageConfig struct {
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
	// Name of the resource;
	//
	// provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#name ComputeImage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#description ComputeImage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#disk_size_gb ComputeImage#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The name of the image family to which this image belongs.
	//
	// You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#family ComputeImage#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// guest_os_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#guest_os_features ComputeImage#guest_os_features}
	GuestOsFeatures interface{} `field:"optional" json:"guestOsFeatures" yaml:"guestOsFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#id ComputeImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#image_encryption_key ComputeImage#image_encryption_key}
	ImageEncryptionKey *ComputeImageImageEncryptionKey `field:"optional" json:"imageEncryptionKey" yaml:"imageEncryptionKey"`
	// Labels to apply to this Image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#labels ComputeImage#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Any applicable license URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#licenses ComputeImage#licenses}
	Licenses *[]*string `field:"optional" json:"licenses" yaml:"licenses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#project ComputeImage#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// raw_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#raw_disk ComputeImage#raw_disk}
	RawDisk *ComputeImageRawDisk `field:"optional" json:"rawDisk" yaml:"rawDisk"`
	// The source disk to create this image based on.
	//
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#source_disk ComputeImage#source_disk}
	SourceDisk *string `field:"optional" json:"sourceDisk" yaml:"sourceDisk"`
	// URL of the source image used to create this image.
	//
	// In order to create an image, you must provide the full or partial
	// URL of one of the following:
	//
	// The selfLink URL
	// This property
	// The rawDisk.source URL
	// The sourceDisk URL
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#source_image ComputeImage#source_image}
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	//
	// The selfLink URL
	// This property
	// The sourceImage URL
	// The rawDisk.source URL
	// The sourceDisk URL
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#source_snapshot ComputeImage#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#storage_locations ComputeImage#storage_locations}
	StorageLocations *[]*string `field:"optional" json:"storageLocations" yaml:"storageLocations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#timeouts ComputeImage#timeouts}
	Timeouts *ComputeImageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

