package accesscontextmanagerserviceperimeter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerServicePerimeterConfig struct {
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
	// Resource name for the ServicePerimeter. The short_name component must begin with a letter and only include alphanumeric and '_'. Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#name AccessContextManagerServicePerimeter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AccessPolicy this ServicePerimeter lives in. Format: accessPolicies/{policy_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#parent AccessContextManagerServicePerimeter#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Human readable title. Must be unique within the Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#title AccessContextManagerServicePerimeter#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Description of the ServicePerimeter and its use. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#description AccessContextManagerServicePerimeter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#id AccessContextManagerServicePerimeter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the type of the Perimeter.
	//
	// There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	//
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	//
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves. Default value: "PERIMETER_TYPE_REGULAR" Possible values: ["PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#perimeter_type AccessContextManagerServicePerimeter#perimeter_type}
	PerimeterType *string `field:"optional" json:"perimeterType" yaml:"perimeterType"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#spec AccessContextManagerServicePerimeter#spec}
	Spec *AccessContextManagerServicePerimeterSpec `field:"optional" json:"spec" yaml:"spec"`
	// status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#status AccessContextManagerServicePerimeter#status}
	Status *AccessContextManagerServicePerimeterStatus `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#timeouts AccessContextManagerServicePerimeter#timeouts}
	Timeouts *AccessContextManagerServicePerimeterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Use explicit dry run spec flag.
	//
	// Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#use_explicit_dry_run_spec AccessContextManagerServicePerimeter#use_explicit_dry_run_spec}
	UseExplicitDryRunSpec interface{} `field:"optional" json:"useExplicitDryRunSpec" yaml:"useExplicitDryRunSpec"`
}

