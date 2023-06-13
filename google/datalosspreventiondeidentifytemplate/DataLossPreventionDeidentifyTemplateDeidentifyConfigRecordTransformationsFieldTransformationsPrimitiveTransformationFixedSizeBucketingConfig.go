package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig struct {
	// Size of each bucket (except for minimum and maximum buckets).
	//
	// So if lower_bound = 10, upper_bound = 89, and bucketSize = 10, then the following buckets would be used: -10, 10-20, 20-30, 30-40, 40-50, 50-60, 60-70, 70-80, 80-89, 89+.
	// Precision up to 2 decimals works.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#bucket_size DataLossPreventionDeidentifyTemplate#bucket_size}
	BucketSize *float64 `field:"required" json:"bucketSize" yaml:"bucketSize"`
	// lower_bound block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#lower_bound DataLossPreventionDeidentifyTemplate#lower_bound}
	LowerBound *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound `field:"required" json:"lowerBound" yaml:"lowerBound"`
	// upper_bound block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#upper_bound DataLossPreventionDeidentifyTemplate#upper_bound}
	UpperBound *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound `field:"required" json:"upperBound" yaml:"upperBound"`
}

