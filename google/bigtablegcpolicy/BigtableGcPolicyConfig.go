package bigtablegcpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableGcPolicyConfig struct {
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
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#column_family BigtableGcPolicy#column_family}
	ColumnFamily *string `field:"required" json:"columnFamily" yaml:"columnFamily"`
	// The name of the Bigtable instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#instance_name BigtableGcPolicy#instance_name}
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#table BigtableGcPolicy#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// The deletion policy for the GC policy.
	//
	// Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted
	// in a replicated instance. Possible values are: "ABANDON".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#deletion_policy BigtableGcPolicy#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Serialized JSON string for garbage collection policy. Conflicts with "mode", "max_age" and "max_version".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#gc_rules BigtableGcPolicy#gc_rules}
	GcRules *string `field:"optional" json:"gcRules" yaml:"gcRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#id BigtableGcPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// max_age block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#max_age BigtableGcPolicy#max_age}
	MaxAge *BigtableGcPolicyMaxAge `field:"optional" json:"maxAge" yaml:"maxAge"`
	// max_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#max_version BigtableGcPolicy#max_version}
	MaxVersion interface{} `field:"optional" json:"maxVersion" yaml:"maxVersion"`
	// NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources.
	//
	// This field may be deprecated in the future. If multiple policies are set, you should choose between UNION OR INTERSECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#mode BigtableGcPolicy#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_gc_policy#project BigtableGcPolicy#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

