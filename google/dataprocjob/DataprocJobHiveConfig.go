package dataprocjob


type DataprocJobHiveConfig struct {
	// Whether to continue executing queries if a query fails.
	//
	// The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#continue_on_failure DataprocJob#continue_on_failure}
	ContinueOnFailure interface{} `field:"optional" json:"continueOnFailure" yaml:"continueOnFailure"`
	// HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks.
	//
	// Can contain Hive SerDes and UDFs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#jar_file_uris DataprocJob#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// A mapping of property names and values, used to configure Hive.
	//
	// Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#properties DataprocJob#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// HCFS URI of file containing Hive script to execute as the job. Conflicts with query_list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#query_file_uri DataprocJob#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// The list of Hive queries or statements to execute as part of the job. Conflicts with query_file_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#query_list DataprocJob#query_list}
	QueryList *[]*string `field:"optional" json:"queryList" yaml:"queryList"`
	// Mapping of query variable names to values (equivalent to the Hive command: SET name="value";).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#script_variables DataprocJob#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

