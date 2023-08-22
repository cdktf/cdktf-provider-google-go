package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimeters struct {
	// Resource name for the ServicePerimeter. The short_name component must begin with a letter and only include alphanumeric and '_'. Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#name AccessContextManagerServicePerimeters#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Human readable title. Must be unique within the Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#title AccessContextManagerServicePerimeters#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Description of the ServicePerimeter and its use. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#description AccessContextManagerServicePerimeters#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#perimeter_type AccessContextManagerServicePerimeters#perimeter_type}
	PerimeterType *string `field:"optional" json:"perimeterType" yaml:"perimeterType"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#spec AccessContextManagerServicePerimeters#spec}
	Spec *AccessContextManagerServicePerimetersServicePerimetersSpec `field:"optional" json:"spec" yaml:"spec"`
	// status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#status AccessContextManagerServicePerimeters#status}
	Status *AccessContextManagerServicePerimetersServicePerimetersStatus `field:"optional" json:"status" yaml:"status"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeters#use_explicit_dry_run_spec AccessContextManagerServicePerimeters#use_explicit_dry_run_spec}
	UseExplicitDryRunSpec interface{} `field:"optional" json:"useExplicitDryRunSpec" yaml:"useExplicitDryRunSpec"`
}

