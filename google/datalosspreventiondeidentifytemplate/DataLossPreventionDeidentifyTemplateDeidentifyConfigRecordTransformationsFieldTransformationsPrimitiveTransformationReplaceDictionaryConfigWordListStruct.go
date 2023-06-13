package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceDictionaryConfigWordListStruct struct {
	// Words or phrases defining the dictionary.
	//
	// The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#words DataLossPreventionDeidentifyTemplate#words}
	Words *[]*string `field:"required" json:"words" yaml:"words"`
}

