// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsPigJob struct {
	// Optional.
	//
	// Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#continue_on_failure DataprocWorkflowTemplate#continue_on_failure}
	ContinueOnFailure interface{} `field:"optional" json:"continueOnFailure" yaml:"continueOnFailure"`
	// Optional.
	//
	// HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#jar_file_uris DataprocWorkflowTemplate#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#logging_config DataprocWorkflowTemplate#logging_config}
	LoggingConfig *DataprocWorkflowTemplateJobsPigJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Optional.
	//
	// A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#properties DataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains the Pig queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#query_file_uri DataprocWorkflowTemplate#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// query_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#query_list DataprocWorkflowTemplate#query_list}
	QueryList *DataprocWorkflowTemplateJobsPigJobQueryListStruct `field:"optional" json:"queryList" yaml:"queryList"`
	// Optional. Mapping of query variable names to values (equivalent to the Pig command: `name=[value]`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_workflow_template#script_variables DataprocWorkflowTemplate#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

