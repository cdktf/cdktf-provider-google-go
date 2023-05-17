package bigqueryjob


type BigqueryJobQuery struct {
	// SQL query text to execute.
	//
	// The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	// NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	// ('DELETE', 'UPDATE', 'MERGE', 'INSERT') must specify 'create_disposition = ""' and 'write_disposition = ""'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#query BigqueryJob#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// If true and query uses legacy SQL dialect, allows the query to produce arbitrarily large result tables at a slight cost in performance.
	//
	// Requires destinationTable to be set. For standard SQL queries, this flag is ignored and large results are always allowed.
	// However, you must still set destinationTable when result size exceeds the allowed maximum response size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#allow_large_results BigqueryJob#allow_large_results}
	AllowLargeResults interface{} `field:"optional" json:"allowLargeResults" yaml:"allowLargeResults"`
	// Specifies whether the job is allowed to create new tables.
	//
	// The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
	// CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
	// Creation, truncation and append actions occur as one atomic update upon job completion Default value: "CREATE_IF_NEEDED" Possible values: ["CREATE_IF_NEEDED", "CREATE_NEVER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#create_disposition BigqueryJob#create_disposition}
	CreateDisposition *string `field:"optional" json:"createDisposition" yaml:"createDisposition"`
	// default_dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#default_dataset BigqueryJob#default_dataset}
	DefaultDataset *BigqueryJobQueryDefaultDataset `field:"optional" json:"defaultDataset" yaml:"defaultDataset"`
	// destination_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#destination_encryption_configuration BigqueryJob#destination_encryption_configuration}
	DestinationEncryptionConfiguration *BigqueryJobQueryDestinationEncryptionConfiguration `field:"optional" json:"destinationEncryptionConfiguration" yaml:"destinationEncryptionConfiguration"`
	// destination_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#destination_table BigqueryJob#destination_table}
	DestinationTable *BigqueryJobQueryDestinationTable `field:"optional" json:"destinationTable" yaml:"destinationTable"`
	// If true and query uses legacy SQL dialect, flattens all nested and repeated fields in the query results.
	//
	// allowLargeResults must be true if this is set to false. For standard SQL queries, this flag is ignored and results are never flattened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#flatten_results BigqueryJob#flatten_results}
	FlattenResults interface{} `field:"optional" json:"flattenResults" yaml:"flattenResults"`
	// Limits the billing tier for this job.
	//
	// Queries that have resource usage beyond this tier will fail (without incurring a charge).
	// If unspecified, this will be set to your project default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#maximum_billing_tier BigqueryJob#maximum_billing_tier}
	MaximumBillingTier *float64 `field:"optional" json:"maximumBillingTier" yaml:"maximumBillingTier"`
	// Limits the bytes billed for this job.
	//
	// Queries that will have bytes billed beyond this limit will fail (without incurring a charge).
	// If unspecified, this will be set to your project default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#maximum_bytes_billed BigqueryJob#maximum_bytes_billed}
	MaximumBytesBilled *string `field:"optional" json:"maximumBytesBilled" yaml:"maximumBytesBilled"`
	// Standard SQL only.
	//
	// Set to POSITIONAL to use positional (?) query parameters or to NAMED to use named (@myparam) query parameters in this query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#parameter_mode BigqueryJob#parameter_mode}
	ParameterMode *string `field:"optional" json:"parameterMode" yaml:"parameterMode"`
	// Specifies a priority for the query. Default value: "INTERACTIVE" Possible values: ["INTERACTIVE", "BATCH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#priority BigqueryJob#priority}
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Allows the schema of the destination table to be updated as a side effect of the query job.
	//
	// Schema update options are supported in two cases: when writeDisposition is WRITE_APPEND;
	// when writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table,
	// specified by partition decorators. For normal tables, WRITE_TRUNCATE will always overwrite the schema.
	// One or more of the following values are specified:
	// ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
	// ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#schema_update_options BigqueryJob#schema_update_options}
	SchemaUpdateOptions *[]*string `field:"optional" json:"schemaUpdateOptions" yaml:"schemaUpdateOptions"`
	// script_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#script_options BigqueryJob#script_options}
	ScriptOptions *BigqueryJobQueryScriptOptions `field:"optional" json:"scriptOptions" yaml:"scriptOptions"`
	// Specifies whether to use BigQuery's legacy SQL dialect for this query.
	//
	// The default value is true.
	// If set to false, the query will use BigQuery's standard SQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#use_legacy_sql BigqueryJob#use_legacy_sql}
	UseLegacySql interface{} `field:"optional" json:"useLegacySql" yaml:"useLegacySql"`
	// Whether to look for the result in the query cache.
	//
	// The query cache is a best-effort cache that will be flushed whenever
	// tables in the query are modified. Moreover, the query cache is only available when a query does not have a destination table specified.
	// The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#use_query_cache BigqueryJob#use_query_cache}
	UseQueryCache interface{} `field:"optional" json:"useQueryCache" yaml:"useQueryCache"`
	// user_defined_function_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#user_defined_function_resources BigqueryJob#user_defined_function_resources}
	UserDefinedFunctionResources interface{} `field:"optional" json:"userDefinedFunctionResources" yaml:"userDefinedFunctionResources"`
	// Specifies the action that occurs if the destination table already exists.
	//
	// The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.
	// WRITE_APPEND: If the table already exists, BigQuery appends the data to the table.
	// WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.
	// Each action is atomic and only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update upon job completion. Default value: "WRITE_EMPTY" Possible values: ["WRITE_TRUNCATE", "WRITE_APPEND", "WRITE_EMPTY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_job#write_disposition BigqueryJob#write_disposition}
	WriteDisposition *string `field:"optional" json:"writeDisposition" yaml:"writeDisposition"`
}

