package cloudbuildtrigger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerConfig struct {
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
	// approval_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#approval_config CloudbuildTrigger#approval_config}
	ApprovalConfig *CloudbuildTriggerApprovalConfig `field:"optional" json:"approvalConfig" yaml:"approvalConfig"`
	// bitbucket_server_trigger_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#bitbucket_server_trigger_config CloudbuildTrigger#bitbucket_server_trigger_config}
	BitbucketServerTriggerConfig *CloudbuildTriggerBitbucketServerTriggerConfig `field:"optional" json:"bitbucketServerTriggerConfig" yaml:"bitbucketServerTriggerConfig"`
	// build block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#build CloudbuildTrigger#build}
	BuildAttribute *CloudbuildTriggerBuild `field:"optional" json:"buildAttribute" yaml:"buildAttribute"`
	// Human-readable description of the trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#description CloudbuildTrigger#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#disabled CloudbuildTrigger#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template.
	//
	// Either a filename or build template must be provided. Set this only when using trigger_template or github.
	// When using Pub/Sub, Webhook or Manual set the file name using git_file_source instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#filename CloudbuildTrigger#filename}
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// A Common Expression Language string. Used only with Pub/Sub and Webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#filter CloudbuildTrigger#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// git_file_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#git_file_source CloudbuildTrigger#git_file_source}
	GitFileSource *CloudbuildTriggerGitFileSource `field:"optional" json:"gitFileSource" yaml:"gitFileSource"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#github CloudbuildTrigger#github}
	Github *CloudbuildTriggerGithub `field:"optional" json:"github" yaml:"github"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#id CloudbuildTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for '**'.
	//
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	//
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignored_file globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#ignored_files CloudbuildTrigger#ignored_files}
	IgnoredFiles *[]*string `field:"optional" json:"ignoredFiles" yaml:"ignoredFiles"`
	// Build logs will be sent back to GitHub as part of the checkrun result.
	//
	// Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or
	// INCLUDE_BUILD_LOGS_WITH_STATUS Possible values: ["INCLUDE_BUILD_LOGS_UNSPECIFIED", "INCLUDE_BUILD_LOGS_WITH_STATUS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#include_build_logs CloudbuildTrigger#include_build_logs}
	IncludeBuildLogs *string `field:"optional" json:"includeBuildLogs" yaml:"includeBuildLogs"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for '**'.
	//
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	//
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#included_files CloudbuildTrigger#included_files}
	IncludedFiles *[]*string `field:"optional" json:"includedFiles" yaml:"includedFiles"`
	// The [Cloud Build location](https://cloud.google.com/build/docs/locations) for the trigger. If not specified, "global" is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#location CloudbuildTrigger#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Name of the trigger. Must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#name CloudbuildTrigger#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#project CloudbuildTrigger#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// pubsub_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#pubsub_config CloudbuildTrigger#pubsub_config}
	PubsubConfig *CloudbuildTriggerPubsubConfig `field:"optional" json:"pubsubConfig" yaml:"pubsubConfig"`
	// repository_event_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#repository_event_config CloudbuildTrigger#repository_event_config}
	RepositoryEventConfig *CloudbuildTriggerRepositoryEventConfig `field:"optional" json:"repositoryEventConfig" yaml:"repositoryEventConfig"`
	// The service account used for all user-controlled operations including triggers.patch, triggers.run, builds.create, and builds.cancel.
	//
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	//
	// Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#service_account CloudbuildTrigger#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// source_to_build block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#source_to_build CloudbuildTrigger#source_to_build}
	SourceToBuild *CloudbuildTriggerSourceToBuild `field:"optional" json:"sourceToBuild" yaml:"sourceToBuild"`
	// Substitutions data for Build resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#substitutions CloudbuildTrigger#substitutions}
	Substitutions *map[string]*string `field:"optional" json:"substitutions" yaml:"substitutions"`
	// Tags for annotation of a BuildTrigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#tags CloudbuildTrigger#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#timeouts CloudbuildTrigger#timeouts}
	Timeouts *CloudbuildTriggerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trigger_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#trigger_template CloudbuildTrigger#trigger_template}
	TriggerTemplate *CloudbuildTriggerTriggerTemplate `field:"optional" json:"triggerTemplate" yaml:"triggerTemplate"`
	// webhook_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#webhook_config CloudbuildTrigger#webhook_config}
	WebhookConfig *CloudbuildTriggerWebhookConfig `field:"optional" json:"webhookConfig" yaml:"webhookConfig"`
}

