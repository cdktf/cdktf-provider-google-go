package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_deidentify_template#conditions DataLossPreventionDeidentifyTemplate#conditions}
	Conditions *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// The operator to apply to the result of conditions.
	//
	// Default and currently only supported value is AND. Default value: "AND" Possible values: ["AND"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_deidentify_template#logical_operator DataLossPreventionDeidentifyTemplate#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}

