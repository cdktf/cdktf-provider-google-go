package datastoreindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastoreIndexConfig struct {
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
	// The entity kind which the index applies to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#kind DatastoreIndex#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Policy for including ancestors in the index. Default value: "NONE" Possible values: ["NONE", "ALL_ANCESTORS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#ancestor DatastoreIndex#ancestor}
	Ancestor *string `field:"optional" json:"ancestor" yaml:"ancestor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#id DatastoreIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#project DatastoreIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#properties DatastoreIndex#properties}
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#timeouts DatastoreIndex#timeouts}
	Timeouts *DatastoreIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

