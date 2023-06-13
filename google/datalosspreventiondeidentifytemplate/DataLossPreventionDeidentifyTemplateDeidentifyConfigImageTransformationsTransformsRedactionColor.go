package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor struct {
	// The amount of blue in the color as a value in the interval [0, 1].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#blue DataLossPreventionDeidentifyTemplate#blue}
	Blue *float64 `field:"optional" json:"blue" yaml:"blue"`
	// The amount of green in the color as a value in the interval [0, 1].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#green DataLossPreventionDeidentifyTemplate#green}
	Green *float64 `field:"optional" json:"green" yaml:"green"`
	// The amount of red in the color as a value in the interval [0, 1].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#red DataLossPreventionDeidentifyTemplate#red}
	Red *float64 `field:"optional" json:"red" yaml:"red"`
}

