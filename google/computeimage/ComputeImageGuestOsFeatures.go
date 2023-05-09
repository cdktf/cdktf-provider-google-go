package computeimage


type ComputeImageGuestOsFeatures struct {
	// The type of supported feature.
	//
	// Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC", "SEV_LIVE_MIGRATABLE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/compute_image#type ComputeImage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

