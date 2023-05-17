package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue struct {
	// Day of month.
	//
	// Must be from 1 to 31 and valid for the year and month, or 0 if specifying a
	// year by itself or a year and month where the day is not significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#day DataLossPreventionDeidentifyTemplate#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Month of year.
	//
	// Must be from 1 to 12, or 0 if specifying a year without a month and day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#month DataLossPreventionDeidentifyTemplate#month}
	Month *float64 `field:"optional" json:"month" yaml:"month"`
	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#year DataLossPreventionDeidentifyTemplate#year}
	Year *float64 `field:"optional" json:"year" yaml:"year"`
}

