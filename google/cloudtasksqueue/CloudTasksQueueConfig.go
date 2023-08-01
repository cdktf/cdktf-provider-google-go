package cloudtasksqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudTasksQueueConfig struct {
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
	// The location of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#location CloudTasksQueue#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// app_engine_routing_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#app_engine_routing_override CloudTasksQueue#app_engine_routing_override}
	AppEngineRoutingOverride *CloudTasksQueueAppEngineRoutingOverride `field:"optional" json:"appEngineRoutingOverride" yaml:"appEngineRoutingOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#id CloudTasksQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The queue name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#name CloudTasksQueue#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#project CloudTasksQueue#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rate_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#rate_limits CloudTasksQueue#rate_limits}
	RateLimits *CloudTasksQueueRateLimits `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// retry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#retry_config CloudTasksQueue#retry_config}
	RetryConfig *CloudTasksQueueRetryConfig `field:"optional" json:"retryConfig" yaml:"retryConfig"`
	// stackdriver_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#stackdriver_logging_config CloudTasksQueue#stackdriver_logging_config}
	StackdriverLoggingConfig *CloudTasksQueueStackdriverLoggingConfig `field:"optional" json:"stackdriverLoggingConfig" yaml:"stackdriverLoggingConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_tasks_queue#timeouts CloudTasksQueue#timeouts}
	Timeouts *CloudTasksQueueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

