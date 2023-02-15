package osconfigpatchdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentConfig struct {
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
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#instance_filter OsConfigPatchDeployment#instance_filter}
	InstanceFilter *OsConfigPatchDeploymentInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// A name for the patch deployment in the project.
	//
	// When creating a name the following rules apply:
	// Must contain only lowercase letters, numbers, and hyphens.
	// Must start with a letter.
	// Must be between 1-63 characters.
	// Must end with a number or a letter.
	// Must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#patch_deployment_id OsConfigPatchDeployment#patch_deployment_id}
	PatchDeploymentId *string `field:"required" json:"patchDeploymentId" yaml:"patchDeploymentId"`
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#description OsConfigPatchDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Duration of the patch.
	//
	// After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#duration OsConfigPatchDeployment#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#id OsConfigPatchDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// one_time_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#one_time_schedule OsConfigPatchDeployment#one_time_schedule}
	OneTimeSchedule *OsConfigPatchDeploymentOneTimeSchedule `field:"optional" json:"oneTimeSchedule" yaml:"oneTimeSchedule"`
	// patch_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#patch_config OsConfigPatchDeployment#patch_config}
	PatchConfig *OsConfigPatchDeploymentPatchConfig `field:"optional" json:"patchConfig" yaml:"patchConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#project OsConfigPatchDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// recurring_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#recurring_schedule OsConfigPatchDeployment#recurring_schedule}
	RecurringSchedule *OsConfigPatchDeploymentRecurringSchedule `field:"optional" json:"recurringSchedule" yaml:"recurringSchedule"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#rollout OsConfigPatchDeployment#rollout}
	Rollout *OsConfigPatchDeploymentRollout `field:"optional" json:"rollout" yaml:"rollout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#timeouts OsConfigPatchDeployment#timeouts}
	Timeouts *OsConfigPatchDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

