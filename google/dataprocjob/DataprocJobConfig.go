package dataprocjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocJobConfig struct {
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
	// placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#placement DataprocJob#placement}
	Placement *DataprocJobPlacement `field:"required" json:"placement" yaml:"placement"`
	// By default, you can only delete inactive jobs within Dataproc.
	//
	// Setting this to true, and calling destroy, will ensure that the job is first cancelled before issuing the delete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#force_delete DataprocJob#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// hadoop_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#hadoop_config DataprocJob#hadoop_config}
	HadoopConfig *DataprocJobHadoopConfig `field:"optional" json:"hadoopConfig" yaml:"hadoopConfig"`
	// hive_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#hive_config DataprocJob#hive_config}
	HiveConfig *DataprocJobHiveConfig `field:"optional" json:"hiveConfig" yaml:"hiveConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#id DataprocJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. The labels to associate with this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#labels DataprocJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// pig_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#pig_config DataprocJob#pig_config}
	PigConfig *DataprocJobPigConfig `field:"optional" json:"pigConfig" yaml:"pigConfig"`
	// presto_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#presto_config DataprocJob#presto_config}
	PrestoConfig *DataprocJobPrestoConfig `field:"optional" json:"prestoConfig" yaml:"prestoConfig"`
	// The project in which the cluster can be found and jobs subsequently run against.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#project DataprocJob#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// pyspark_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#pyspark_config DataprocJob#pyspark_config}
	PysparkConfig *DataprocJobPysparkConfig `field:"optional" json:"pysparkConfig" yaml:"pysparkConfig"`
	// reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#reference DataprocJob#reference}
	Reference *DataprocJobReference `field:"optional" json:"reference" yaml:"reference"`
	// The Cloud Dataproc region.
	//
	// This essentially determines which clusters are available for this job to be submitted to. If not specified, defaults to global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#region DataprocJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#scheduling DataprocJob#scheduling}
	Scheduling *DataprocJobScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// spark_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#spark_config DataprocJob#spark_config}
	SparkConfig *DataprocJobSparkConfig `field:"optional" json:"sparkConfig" yaml:"sparkConfig"`
	// sparksql_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#sparksql_config DataprocJob#sparksql_config}
	SparksqlConfig *DataprocJobSparksqlConfig `field:"optional" json:"sparksqlConfig" yaml:"sparksqlConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_job#timeouts DataprocJob#timeouts}
	Timeouts *DataprocJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

