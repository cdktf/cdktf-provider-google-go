package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations struct {
	// primitive_transformation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#primitive_transformation DataLossPreventionDeidentifyTemplate#primitive_transformation}
	PrimitiveTransformation *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation `field:"required" json:"primitiveTransformation" yaml:"primitiveTransformation"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#info_types DataLossPreventionDeidentifyTemplate#info_types}
	InfoTypes interface{} `field:"optional" json:"infoTypes" yaml:"infoTypes"`
}

