package apigeeaddonsconfig


type ApigeeAddonsConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_addons_config#create ApigeeAddonsConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_addons_config#delete ApigeeAddonsConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_addons_config#update ApigeeAddonsConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
