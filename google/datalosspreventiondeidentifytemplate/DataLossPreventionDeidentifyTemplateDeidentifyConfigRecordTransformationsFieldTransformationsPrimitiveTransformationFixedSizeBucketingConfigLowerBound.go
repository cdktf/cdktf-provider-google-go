package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound struct {
	// A boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#boolean_value DataLossPreventionDeidentifyTemplate#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// date_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#date_value DataLossPreventionDeidentifyTemplate#date_value}
	DateValue *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue `field:"optional" json:"dateValue" yaml:"dateValue"`
	// Represents a day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#day_of_week_value DataLossPreventionDeidentifyTemplate#day_of_week_value}
	DayOfWeekValue *string `field:"optional" json:"dayOfWeekValue" yaml:"dayOfWeekValue"`
	// A float value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#float_value DataLossPreventionDeidentifyTemplate#float_value}
	FloatValue *float64 `field:"optional" json:"floatValue" yaml:"floatValue"`
	// An integer value (int64 format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#integer_value DataLossPreventionDeidentifyTemplate#integer_value}
	IntegerValue *string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// A string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#string_value DataLossPreventionDeidentifyTemplate#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	//
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#timestamp_value DataLossPreventionDeidentifyTemplate#timestamp_value}
	TimestampValue *string `field:"optional" json:"timestampValue" yaml:"timestampValue"`
	// time_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#time_value DataLossPreventionDeidentifyTemplate#time_value}
	TimeValue *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue `field:"optional" json:"timeValue" yaml:"timeValue"`
}

