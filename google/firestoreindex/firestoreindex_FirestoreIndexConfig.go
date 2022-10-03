package firestoreindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirestoreIndexConfig struct {
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
	// The collection being indexed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#collection FirestoreIndex#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#fields FirestoreIndex#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The Firestore database id. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#database FirestoreIndex#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#id FirestoreIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#project FirestoreIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#query_scope FirestoreIndex#query_scope}
	QueryScope *string `field:"optional" json:"queryScope" yaml:"queryScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#timeouts FirestoreIndex#timeouts}
	Timeouts *FirestoreIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

