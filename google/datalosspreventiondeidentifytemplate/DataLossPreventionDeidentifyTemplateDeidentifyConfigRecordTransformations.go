package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations struct {
	// field_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_deidentify_template#field_transformations DataLossPreventionDeidentifyTemplate#field_transformations}
	FieldTransformations interface{} `field:"optional" json:"fieldTransformations" yaml:"fieldTransformations"`
	// record_suppressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_deidentify_template#record_suppressions DataLossPreventionDeidentifyTemplate#record_suppressions}
	RecordSuppressions interface{} `field:"optional" json:"recordSuppressions" yaml:"recordSuppressions"`
}

