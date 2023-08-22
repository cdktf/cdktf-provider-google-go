package computeregioncommitment


type ComputeRegionCommitmentLicenseResource struct {
	// Any applicable license URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_commitment#license ComputeRegionCommitment#license}
	License *string `field:"required" json:"license" yaml:"license"`
	// The number of licenses purchased.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_commitment#amount ComputeRegionCommitment#amount}
	Amount *string `field:"optional" json:"amount" yaml:"amount"`
	// Specifies the core range of the instance for which this license applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_commitment#cores_per_license ComputeRegionCommitment#cores_per_license}
	CoresPerLicense *string `field:"optional" json:"coresPerLicense" yaml:"coresPerLicense"`
}

