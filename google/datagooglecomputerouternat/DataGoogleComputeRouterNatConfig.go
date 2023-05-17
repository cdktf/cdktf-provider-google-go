package datagooglecomputerouternat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeRouterNatConfig struct {
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
	// Name of the NAT service. The name must be 1-63 characters long and comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/compute_router_nat#name DataGoogleComputeRouterNat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Cloud Router in which this NAT will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/compute_router_nat#router DataGoogleComputeRouterNat#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/compute_router_nat#id DataGoogleComputeRouterNat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/compute_router_nat#project DataGoogleComputeRouterNat#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the router and NAT reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/compute_router_nat#region DataGoogleComputeRouterNat#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

