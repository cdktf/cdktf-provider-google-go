package cloudbuildworkerpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildWorkerPoolConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#location CloudbuildWorkerPool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User-defined name of the `WorkerPool`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#name CloudbuildWorkerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#annotations CloudbuildWorkerPool#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#display_name CloudbuildWorkerPool#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#id CloudbuildWorkerPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#network_config CloudbuildWorkerPool#network_config}
	NetworkConfig *CloudbuildWorkerPoolNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#project CloudbuildWorkerPool#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#timeouts CloudbuildWorkerPool#timeouts}
	Timeouts *CloudbuildWorkerPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_worker_pool#worker_config CloudbuildWorkerPool#worker_config}
	WorkerConfig *CloudbuildWorkerPoolWorkerConfig `field:"optional" json:"workerConfig" yaml:"workerConfig"`
}

