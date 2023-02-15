package bigtabletableiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableTableIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#instance BigtableTableIamMember#instance}.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#member BigtableTableIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#role BigtableTableIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#table BigtableTableIamMember#table}.
	Table *string `field:"required" json:"table" yaml:"table"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#condition BigtableTableIamMember#condition}
	Condition *BigtableTableIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#id BigtableTableIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table_iam_member#project BigtableTableIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

