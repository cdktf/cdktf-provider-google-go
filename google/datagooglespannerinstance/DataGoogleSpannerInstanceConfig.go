package datagooglespannerinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleSpannerInstanceConfig struct {
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
	// A unique identifier for the instance, which cannot be changed after the instance is created.
	//
	// The name must be between 6 and 30 characters
	// in length.
	//
	//
	// If not provided, a random string starting with 'tf-' will be selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/spanner_instance#name DataGoogleSpannerInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the instance's configuration (similar but not quite the same as a region) which defines the geographic placement and replication of your databases in this instance.
	//
	// It determines where your data
	// is stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/spanner_instance#config DataGoogleSpannerInstance#config}
	Config *string `field:"optional" json:"config" yaml:"config"`
	// The descriptive name for this instance as it appears in UIs.
	//
	// Must be
	// unique per project and between 4 and 30 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/spanner_instance#display_name DataGoogleSpannerInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/spanner_instance#id DataGoogleSpannerInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/spanner_instance#project DataGoogleSpannerInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

