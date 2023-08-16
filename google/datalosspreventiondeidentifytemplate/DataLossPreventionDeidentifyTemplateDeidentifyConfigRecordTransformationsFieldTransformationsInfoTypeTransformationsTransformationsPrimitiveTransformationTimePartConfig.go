package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig struct {
	// The part of the time to keep. Possible values: ["YEAR", "MONTH", "DAY_OF_MONTH", "DAY_OF_WEEK", "WEEK_OF_YEAR", "HOUR_OF_DAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#part_to_extract DataLossPreventionDeidentifyTemplate#part_to_extract}
	PartToExtract *string `field:"required" json:"partToExtract" yaml:"partToExtract"`
}

