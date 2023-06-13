package bigqueryroutine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryRoutineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the dataset containing this routine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#dataset_id BigqueryRoutine#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The body of the routine.
	//
	// For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#definition_body BigqueryRoutine#definition_body}
	DefinitionBody *string `field:"required" json:"definitionBody" yaml:"definitionBody"`
	// The ID of the the routine.
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#routine_id BigqueryRoutine#routine_id}
	RoutineId *string `field:"required" json:"routineId" yaml:"routineId"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#arguments BigqueryRoutine#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// The description of the routine if defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#description BigqueryRoutine#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#determinism_level BigqueryRoutine#determinism_level}
	DeterminismLevel *string `field:"optional" json:"determinismLevel" yaml:"determinismLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#id BigqueryRoutine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#imported_libraries BigqueryRoutine#imported_libraries}
	ImportedLibraries *[]*string `field:"optional" json:"importedLibraries" yaml:"importedLibraries"`
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#language BigqueryRoutine#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#project BigqueryRoutine#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	//
	// If absent, the return table type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the columns in the evaluated table result will
	// be cast to match the column types specificed in return table type, at query time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#return_table_type BigqueryRoutine#return_table_type}
	ReturnTableType *string `field:"optional" json:"returnTableType" yaml:"returnTableType"`
	// A JSON schema for the return type.
	//
	// Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#return_type BigqueryRoutine#return_type}
	ReturnType *string `field:"optional" json:"returnType" yaml:"returnType"`
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#routine_type BigqueryRoutine#routine_type}
	RoutineType *string `field:"optional" json:"routineType" yaml:"routineType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_routine#timeouts BigqueryRoutine#timeouts}
	Timeouts *BigqueryRoutineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

