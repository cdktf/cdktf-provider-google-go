package appenginedomainmapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineDomainMappingConfig struct {
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
	// Relative name of the domain serving the application. Example: example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_domain_mapping#domain_name AppEngineDomainMapping#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_domain_mapping#id AppEngineDomainMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the domain creation should override any existing mappings for this domain.
	//
	// By default, overrides are rejected. Default value: "STRICT" Possible values: ["STRICT", "OVERRIDE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_domain_mapping#override_strategy AppEngineDomainMapping#override_strategy}
	OverrideStrategy *string `field:"optional" json:"overrideStrategy" yaml:"overrideStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_domain_mapping#project AppEngineDomainMapping#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// ssl_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_domain_mapping#ssl_settings AppEngineDomainMapping#ssl_settings}
	SslSettings *AppEngineDomainMappingSslSettings `field:"optional" json:"sslSettings" yaml:"sslSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_domain_mapping#timeouts AppEngineDomainMapping#timeouts}
	Timeouts *AppEngineDomainMappingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

