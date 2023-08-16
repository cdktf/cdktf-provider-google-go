package datagooglecomputeinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_instance#id DataGoogleComputeInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the instance. One of name or self_link must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_instance#name DataGoogleComputeInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the project in which the resource belongs.
	//
	// If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_instance#project DataGoogleComputeInstance#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The URI of the created resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_instance#self_link DataGoogleComputeInstance#self_link}
	SelfLink *string `field:"optional" json:"selfLink" yaml:"selfLink"`
	// The zone of the instance.
	//
	// If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_instance#zone DataGoogleComputeInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

