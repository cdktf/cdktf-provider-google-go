package mlenginemodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MlEngineModelConfig struct {
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
	// The name specified for the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#name MlEngineModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// default_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#default_version MlEngineModel#default_version}
	DefaultVersion *MlEngineModelDefaultVersion `field:"optional" json:"defaultVersion" yaml:"defaultVersion"`
	// The description specified for the model when it was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#description MlEngineModel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#id MlEngineModel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// One or more labels that you can add, to organize your models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#labels MlEngineModel#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#online_prediction_console_logging MlEngineModel#online_prediction_console_logging}
	OnlinePredictionConsoleLogging interface{} `field:"optional" json:"onlinePredictionConsoleLogging" yaml:"onlinePredictionConsoleLogging"`
	// If true, online prediction access logs are sent to StackDriver Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#online_prediction_logging MlEngineModel#online_prediction_logging}
	OnlinePredictionLogging interface{} `field:"optional" json:"onlinePredictionLogging" yaml:"onlinePredictionLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#project MlEngineModel#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The list of regions where the model is going to be deployed. Currently only one region per model is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#regions MlEngineModel#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#timeouts MlEngineModel#timeouts}
	Timeouts *MlEngineModelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

