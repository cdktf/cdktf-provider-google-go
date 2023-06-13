package computeregioncommitment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionCommitmentConfig struct {
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
	// Name of the resource.
	//
	// The name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#name ComputeRegionCommitment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The plan for this commitment, which determines duration and discount rate.
	//
	// The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years). Possible values: ["TWELVE_MONTH", "THIRTY_SIX_MONTH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#plan ComputeRegionCommitment#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Specifies whether to enable automatic renewal for the commitment.
	//
	// The default value is false if not specified.
	// If the field is set to true, the commitment will be automatically renewed for either
	// one or three years according to the terms of the existing commitment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#auto_renew ComputeRegionCommitment#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// The category of the commitment.
	//
	// Category MACHINE specifies commitments composed of
	// machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE
	// specifies commitments composed of software licenses, listed in licenseResources.
	// Note that only MACHINE commitments should have a Type specified. Possible values: ["LICENSE", "MACHINE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#category ComputeRegionCommitment#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#description ComputeRegionCommitment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#id ComputeRegionCommitment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// license_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#license_resource ComputeRegionCommitment#license_resource}
	LicenseResource *ComputeRegionCommitmentLicenseResource `field:"optional" json:"licenseResource" yaml:"licenseResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#project ComputeRegionCommitment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// URL of the region where this commitment may be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#region ComputeRegionCommitment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#resources ComputeRegionCommitment#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#timeouts ComputeRegionCommitment#timeouts}
	Timeouts *ComputeRegionCommitmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of commitment, which affects the discount rate and the eligible resources.
	//
	// The type could be one of the following value: 'MEMORY_OPTIMIZED', 'ACCELERATOR_OPTIMIZED',
	// 'GENERAL_PURPOSE_N1', 'GENERAL_PURPOSE_N2', 'GENERAL_PURPOSE_N2D', 'GENERAL_PURPOSE_E2',
	// 'GENERAL_PURPOSE_T2D', 'GENERAL_PURPOSE_C3', 'COMPUTE_OPTIMIZED_C2', 'COMPUTE_OPTIMIZED_C2D' and
	// 'GRAPHICS_OPTIMIZED_G2'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_commitment#type ComputeRegionCommitment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

