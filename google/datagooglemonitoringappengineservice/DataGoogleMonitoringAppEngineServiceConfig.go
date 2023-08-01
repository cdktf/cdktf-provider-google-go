package datagooglemonitoringappengineservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleMonitoringAppEngineServiceConfig struct {
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
	// The ID of the App Engine module underlying this service.
	//
	// Corresponds to the 'moduleId' resource label for a 'gae_app'
	// monitored resource(see https://cloud.google.com/monitoring/api/resources#tag_gae_app)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/data-sources/monitoring_app_engine_service#module_id DataGoogleMonitoringAppEngineService#module_id}
	ModuleId *string `field:"required" json:"moduleId" yaml:"moduleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/data-sources/monitoring_app_engine_service#id DataGoogleMonitoringAppEngineService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/data-sources/monitoring_app_engine_service#project DataGoogleMonitoringAppEngineService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

