package healthcaredataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareDatasetConfig struct {
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
	// The location for the Dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dataset#location HealthcareDataset#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name for the Dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dataset#name HealthcareDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dataset#id HealthcareDataset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dataset#project HealthcareDataset#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dataset#timeouts HealthcareDataset#timeouts}
	Timeouts *HealthcareDatasetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The default timezone used by this dataset.
	//
	// Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dataset#time_zone HealthcareDataset#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

