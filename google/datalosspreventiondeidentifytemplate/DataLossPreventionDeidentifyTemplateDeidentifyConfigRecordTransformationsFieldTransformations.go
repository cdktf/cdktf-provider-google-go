package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#fields DataLossPreventionDeidentifyTemplate#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// primitive_transformation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#primitive_transformation DataLossPreventionDeidentifyTemplate#primitive_transformation}
	PrimitiveTransformation *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation `field:"required" json:"primitiveTransformation" yaml:"primitiveTransformation"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#condition DataLossPreventionDeidentifyTemplate#condition}
	Condition *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition `field:"optional" json:"condition" yaml:"condition"`
}

