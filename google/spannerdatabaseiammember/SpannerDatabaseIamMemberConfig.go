package spannerdatabaseiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpannerDatabaseIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#database SpannerDatabaseIamMember#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#instance SpannerDatabaseIamMember#instance}.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#member SpannerDatabaseIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#role SpannerDatabaseIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#condition SpannerDatabaseIamMember#condition}
	Condition *SpannerDatabaseIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#id SpannerDatabaseIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/spanner_database_iam_member#project SpannerDatabaseIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

