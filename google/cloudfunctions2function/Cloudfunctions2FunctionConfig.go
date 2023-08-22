package cloudfunctions2function

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudfunctions2FunctionConfig struct {
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
	// A user-defined name of the function. Function names must be unique globally and match pattern 'projects/*\/locations/*\/functions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#name Cloudfunctions2Function#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// build_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#build_config Cloudfunctions2Function#build_config}
	BuildConfig *Cloudfunctions2FunctionBuildConfig `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// User-provided description of a function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#description Cloudfunctions2Function#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// event_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#event_trigger Cloudfunctions2Function#event_trigger}
	EventTrigger *Cloudfunctions2FunctionEventTrigger `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#id Cloudfunctions2Function#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
	//
	// It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#kms_key_name Cloudfunctions2Function#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// A set of key/value label pairs associated with this Cloud Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#labels Cloudfunctions2Function#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of this cloud function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#location Cloudfunctions2Function#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#project Cloudfunctions2Function#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#service_config Cloudfunctions2Function#service_config}
	ServiceConfig *Cloudfunctions2FunctionServiceConfig `field:"optional" json:"serviceConfig" yaml:"serviceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions2_function#timeouts Cloudfunctions2Function#timeouts}
	Timeouts *Cloudfunctions2FunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

