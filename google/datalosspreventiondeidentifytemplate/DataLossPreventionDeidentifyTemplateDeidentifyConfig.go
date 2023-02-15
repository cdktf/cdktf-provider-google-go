package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfig struct {
	// info_type_transformations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#info_type_transformations DataLossPreventionDeidentifyTemplate#info_type_transformations}
	InfoTypeTransformations *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations `field:"optional" json:"infoTypeTransformations" yaml:"infoTypeTransformations"`
	// record_transformations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#record_transformations DataLossPreventionDeidentifyTemplate#record_transformations}
	RecordTransformations *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations `field:"optional" json:"recordTransformations" yaml:"recordTransformations"`
}

