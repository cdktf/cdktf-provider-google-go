package appengineapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineApplicationConfig struct {
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
	// The location to serve the app from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#location_id AppEngineApplication#location_id}
	LocationId *string `field:"required" json:"locationId" yaml:"locationId"`
	// The domain to authenticate users with when using App Engine's User API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#auth_domain AppEngineApplication#auth_domain}
	AuthDomain *string `field:"optional" json:"authDomain" yaml:"authDomain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#database_type AppEngineApplication#database_type}.
	DatabaseType *string `field:"optional" json:"databaseType" yaml:"databaseType"`
	// feature_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#feature_settings AppEngineApplication#feature_settings}
	FeatureSettings *AppEngineApplicationFeatureSettings `field:"optional" json:"featureSettings" yaml:"featureSettings"`
	// iap block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#iap AppEngineApplication#iap}
	Iap *AppEngineApplicationIap `field:"optional" json:"iap" yaml:"iap"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#id AppEngineApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project ID to create the application under.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#project AppEngineApplication#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The serving status of the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#serving_status AppEngineApplication#serving_status}
	ServingStatus *string `field:"optional" json:"servingStatus" yaml:"servingStatus"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#timeouts AppEngineApplication#timeouts}
	Timeouts *AppEngineApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

