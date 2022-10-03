package bigtabletable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the Bigtable instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#instance_name BigtableTable#instance_name}
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#name BigtableTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// column_family block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#column_family BigtableTable#column_family}
	ColumnFamily interface{} `field:"optional" json:"columnFamily" yaml:"columnFamily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#id BigtableTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#project BigtableTable#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A list of predefined keys to split the table on.
	//
	// !> Warning: Modifying the split_keys of an existing table will cause Terraform to delete/recreate the entire google_bigtable_table resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#split_keys BigtableTable#split_keys}
	SplitKeys *[]*string `field:"optional" json:"splitKeys" yaml:"splitKeys"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#timeouts BigtableTable#timeouts}
	Timeouts *BigtableTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

