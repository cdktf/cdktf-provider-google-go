package bigqueryroutine


type BigqueryRoutineArguments struct {
	// Defaults to FIXED_TYPE. Default value: "FIXED_TYPE" Possible values: ["FIXED_TYPE", "ANY_TYPE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#argument_kind BigqueryRoutine#argument_kind}
	ArgumentKind *string `field:"optional" json:"argumentKind" yaml:"argumentKind"`
	// A JSON schema for the data type.
	//
	// Required unless argumentKind = ANY_TYPE.
	// ~>**NOTE**: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#data_type BigqueryRoutine#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Specifies whether the argument is input or output. Can be set for procedures only. Possible values: ["IN", "OUT", "INOUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#mode BigqueryRoutine#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of this argument. Can be absent for function return argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#name BigqueryRoutine#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

