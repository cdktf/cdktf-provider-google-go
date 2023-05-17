package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfig struct {
	// image_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#image_transformations DataLossPreventionDeidentifyTemplate#image_transformations}
	ImageTransformations *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations `field:"optional" json:"imageTransformations" yaml:"imageTransformations"`
	// info_type_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#info_type_transformations DataLossPreventionDeidentifyTemplate#info_type_transformations}
	InfoTypeTransformations *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations `field:"optional" json:"infoTypeTransformations" yaml:"infoTypeTransformations"`
	// record_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#record_transformations DataLossPreventionDeidentifyTemplate#record_transformations}
	RecordTransformations *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations `field:"optional" json:"recordTransformations" yaml:"recordTransformations"`
}

