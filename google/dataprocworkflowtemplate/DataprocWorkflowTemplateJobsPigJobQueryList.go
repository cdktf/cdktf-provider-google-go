package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsPigJobQueryList struct {
	// Required.
	//
	// The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: "hiveJob": { "queryList": { "queries": [ "query1", "query2", "query3;query4", ] } }
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#queries DataprocWorkflowTemplate#queries}
	Queries *[]*string `field:"required" json:"queries" yaml:"queries"`
}

