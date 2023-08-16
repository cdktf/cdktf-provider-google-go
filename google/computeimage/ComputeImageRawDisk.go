package computeimage


type ComputeImageRawDisk struct {
	// The full Google Cloud Storage URL where disk storage is stored You must provide either this property or the sourceDisk property but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#source ComputeImage#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The format used to encode and transmit the block device, which should be TAR.
	//
	// This is just a container and transmission format
	// and not a runtime format. Provided by the client when the disk
	// image is created. Default value: "TAR" Possible values: ["TAR"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#container_type ComputeImage#container_type}
	ContainerType *string `field:"optional" json:"containerType" yaml:"containerType"`
	// An optional SHA1 checksum of the disk image before unpackaging.
	//
	// This is provided by the client when the disk image is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_image#sha1 ComputeImage#sha1}
	Sha1 *string `field:"optional" json:"sha1" yaml:"sha1"`
}

