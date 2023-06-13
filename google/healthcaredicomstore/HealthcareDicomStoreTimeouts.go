package healthcaredicomstore


type HealthcareDicomStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/healthcare_dicom_store#create HealthcareDicomStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/healthcare_dicom_store#delete HealthcareDicomStore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/healthcare_dicom_store#update HealthcareDicomStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

