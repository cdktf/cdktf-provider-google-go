package dataplextask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexTaskConfig struct {
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
	// execution_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#execution_spec DataplexTask#execution_spec}
	ExecutionSpec *DataplexTaskExecutionSpec `field:"required" json:"executionSpec" yaml:"executionSpec"`
	// trigger_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#trigger_spec DataplexTask#trigger_spec}
	TriggerSpec *DataplexTaskTriggerSpec `field:"required" json:"triggerSpec" yaml:"triggerSpec"`
	// User-provided description of the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#description DataplexTask#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#display_name DataplexTask#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#id DataplexTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#labels DataplexTask#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The lake in which the task will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#lake DataplexTask#lake}
	Lake *string `field:"optional" json:"lake" yaml:"lake"`
	// The location in which the task will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#location DataplexTask#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// notebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#notebook DataplexTask#notebook}
	Notebook *DataplexTaskNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#project DataplexTask#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#spark DataplexTask#spark}
	Spark *DataplexTaskSpark `field:"optional" json:"spark" yaml:"spark"`
	// The task Id of the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#task_id DataplexTask#task_id}
	TaskId *string `field:"optional" json:"taskId" yaml:"taskId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#timeouts DataplexTask#timeouts}
	Timeouts *DataplexTaskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

