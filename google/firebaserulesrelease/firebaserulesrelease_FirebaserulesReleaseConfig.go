package firebaserulesrelease

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaserulesReleaseConfig struct {
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
	// Format: `projects/{project_id}/releases/{release_id}`\Firestore Rules Releases will **always** have the name 'cloud.firestore'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firebaserules_release#name FirebaserulesRelease#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firebaserules_release#ruleset_name FirebaserulesRelease#ruleset_name}
	RulesetName *string `field:"required" json:"rulesetName" yaml:"rulesetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firebaserules_release#id FirebaserulesRelease#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firebaserules_release#project FirebaserulesRelease#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firebaserules_release#timeouts FirebaserulesRelease#timeouts}
	Timeouts *FirebaserulesReleaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

