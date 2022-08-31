// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type HealthcareDatasetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_dataset#create HealthcareDataset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_dataset#delete HealthcareDataset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_dataset#update HealthcareDataset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

