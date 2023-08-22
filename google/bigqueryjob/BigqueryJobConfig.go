package bigqueryjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryJobConfig struct {
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
	// The ID of the job.
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#job_id BigqueryJob#job_id}
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// copy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#copy BigqueryJob#copy}
	Copy *BigqueryJobCopy `field:"optional" json:"copy" yaml:"copy"`
	// extract block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#extract BigqueryJob#extract}
	Extract *BigqueryJobExtract `field:"optional" json:"extract" yaml:"extract"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#id BigqueryJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#job_timeout_ms BigqueryJob#job_timeout_ms}
	JobTimeoutMs *string `field:"optional" json:"jobTimeoutMs" yaml:"jobTimeoutMs"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#labels BigqueryJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// load block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#load BigqueryJob#load}
	Load *BigqueryJobLoad `field:"optional" json:"load" yaml:"load"`
	// The geographic location of the job. The default value is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#location BigqueryJob#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#project BigqueryJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#query BigqueryJob#query}
	Query *BigqueryJobQuery `field:"optional" json:"query" yaml:"query"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#timeouts BigqueryJob#timeouts}
	Timeouts *BigqueryJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

