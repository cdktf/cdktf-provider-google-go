package appengineflexibleappversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionConfig struct {
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
	// liveness_check block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#liveness_check AppEngineFlexibleAppVersion#liveness_check}
	LivenessCheck *AppEngineFlexibleAppVersionLivenessCheck `field:"required" json:"livenessCheck" yaml:"livenessCheck"`
	// readiness_check block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#readiness_check AppEngineFlexibleAppVersion#readiness_check}
	ReadinessCheck *AppEngineFlexibleAppVersionReadinessCheck `field:"required" json:"readinessCheck" yaml:"readinessCheck"`
	// Desired runtime. Example python27.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#runtime AppEngineFlexibleAppVersion#runtime}
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// AppEngine service resource. Can contain numbers, letters, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#service AppEngineFlexibleAppVersion#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// api_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#api_config AppEngineFlexibleAppVersion#api_config}
	ApiConfig *AppEngineFlexibleAppVersionApiConfig `field:"optional" json:"apiConfig" yaml:"apiConfig"`
	// automatic_scaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#automatic_scaling AppEngineFlexibleAppVersion#automatic_scaling}
	AutomaticScaling *AppEngineFlexibleAppVersionAutomaticScaling `field:"optional" json:"automaticScaling" yaml:"automaticScaling"`
	// Metadata settings that are supplied to this version to enable beta runtime features.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#beta_settings AppEngineFlexibleAppVersion#beta_settings}
	BetaSettings *map[string]*string `field:"optional" json:"betaSettings" yaml:"betaSettings"`
	// Duration that static files should be cached by web proxies and browsers.
	//
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#default_expiration AppEngineFlexibleAppVersion#default_expiration}
	DefaultExpiration *string `field:"optional" json:"defaultExpiration" yaml:"defaultExpiration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#delete_service_on_destroy AppEngineFlexibleAppVersion#delete_service_on_destroy}.
	DeleteServiceOnDestroy interface{} `field:"optional" json:"deleteServiceOnDestroy" yaml:"deleteServiceOnDestroy"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#deployment AppEngineFlexibleAppVersion#deployment}
	Deployment *AppEngineFlexibleAppVersionDeployment `field:"optional" json:"deployment" yaml:"deployment"`
	// endpoints_api_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#endpoints_api_service AppEngineFlexibleAppVersion#endpoints_api_service}
	EndpointsApiService *AppEngineFlexibleAppVersionEndpointsApiService `field:"optional" json:"endpointsApiService" yaml:"endpointsApiService"`
	// entrypoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#entrypoint AppEngineFlexibleAppVersion#entrypoint}
	Entrypoint *AppEngineFlexibleAppVersionEntrypoint `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Environment variables available to the application.
	//
	// As these are not returned in the API request, Terraform will not detect any changes made outside of the Terraform config.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#env_variables AppEngineFlexibleAppVersion#env_variables}
	EnvVariables *map[string]*string `field:"optional" json:"envVariables" yaml:"envVariables"`
	// handlers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#handlers AppEngineFlexibleAppVersion#handlers}
	Handlers interface{} `field:"optional" json:"handlers" yaml:"handlers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#id AppEngineFlexibleAppVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of the types of messages that this application is able to receive.
	//
	// Possible values: ["INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE", "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE", "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#inbound_services AppEngineFlexibleAppVersion#inbound_services}
	InboundServices *[]*string `field:"optional" json:"inboundServices" yaml:"inboundServices"`
	// Instance class that is used to run this version.
	//
	// Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#instance_class AppEngineFlexibleAppVersion#instance_class}
	InstanceClass *string `field:"optional" json:"instanceClass" yaml:"instanceClass"`
	// manual_scaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#manual_scaling AppEngineFlexibleAppVersion#manual_scaling}
	ManualScaling *AppEngineFlexibleAppVersionManualScaling `field:"optional" json:"manualScaling" yaml:"manualScaling"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#network AppEngineFlexibleAppVersion#network}
	Network *AppEngineFlexibleAppVersionNetwork `field:"optional" json:"network" yaml:"network"`
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#nobuild_files_regex AppEngineFlexibleAppVersion#nobuild_files_regex}
	NobuildFilesRegex *string `field:"optional" json:"nobuildFilesRegex" yaml:"nobuildFilesRegex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#noop_on_destroy AppEngineFlexibleAppVersion#noop_on_destroy}.
	NoopOnDestroy interface{} `field:"optional" json:"noopOnDestroy" yaml:"noopOnDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#project AppEngineFlexibleAppVersion#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#resources AppEngineFlexibleAppVersion#resources}
	Resources *AppEngineFlexibleAppVersionResources `field:"optional" json:"resources" yaml:"resources"`
	// The version of the API in the given runtime environment.
	//
	// Please see the app.yaml reference for valid values at 'https://cloud.google.com/appengine/docs/standard/<language>/config/appref'\
	// Substitute '<language>' with 'python', 'java', 'php', 'ruby', 'go' or 'nodejs'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#runtime_api_version AppEngineFlexibleAppVersion#runtime_api_version}
	RuntimeApiVersion *string `field:"optional" json:"runtimeApiVersion" yaml:"runtimeApiVersion"`
	// The channel of the runtime to use. Only available for some runtimes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#runtime_channel AppEngineFlexibleAppVersion#runtime_channel}
	RuntimeChannel *string `field:"optional" json:"runtimeChannel" yaml:"runtimeChannel"`
	// The path or name of the app's main executable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#runtime_main_executable_path AppEngineFlexibleAppVersion#runtime_main_executable_path}
	RuntimeMainExecutablePath *string `field:"optional" json:"runtimeMainExecutablePath" yaml:"runtimeMainExecutablePath"`
	// The identity that the deployed version will run as.
	//
	// Admin API will use the App Engine Appspot service account as
	// default if this field is neither provided in app.yaml file nor through CLI flag.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#service_account AppEngineFlexibleAppVersion#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Current serving status of this version.
	//
	// Only the versions with a SERVING status create instances and can be billed. Default value: "SERVING" Possible values: ["SERVING", "STOPPED"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#serving_status AppEngineFlexibleAppVersion#serving_status}
	ServingStatus *string `field:"optional" json:"servingStatus" yaml:"servingStatus"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#timeouts AppEngineFlexibleAppVersion#timeouts}
	Timeouts *AppEngineFlexibleAppVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Relative name of the version within the service.
	//
	// For example, 'v1'. Version names can contain only lowercase letters, numbers, or hyphens.
	// Reserved names,"default", "latest", and any name with the prefix "ah-".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#version_id AppEngineFlexibleAppVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// vpc_access_connector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#vpc_access_connector AppEngineFlexibleAppVersion#vpc_access_connector}
	VpcAccessConnector *AppEngineFlexibleAppVersionVpcAccessConnector `field:"optional" json:"vpcAccessConnector" yaml:"vpcAccessConnector"`
}

