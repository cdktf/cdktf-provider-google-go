package pubsubschema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PubsubSchemaConfig struct {
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
	// The ID to use for the schema, which will become the final component of the schema's resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#name PubsubSchema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The definition of the schema.
	//
	// This should contain a string representing the full definition of the schema
	// that is a valid schema definition of the type specified in type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#definition PubsubSchema#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#id PubsubSchema#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#project PubsubSchema#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#timeouts PubsubSchema#timeouts}
	Timeouts *PubsubSchemaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of the schema definition Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#type PubsubSchema#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

