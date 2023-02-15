package healthcarefhirstoreiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareFhirStoreIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#fhir_store_id HealthcareFhirStoreIamBinding#fhir_store_id}.
	FhirStoreId *string `field:"required" json:"fhirStoreId" yaml:"fhirStoreId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#members HealthcareFhirStoreIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#role HealthcareFhirStoreIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#condition HealthcareFhirStoreIamBinding#condition}
	Condition *HealthcareFhirStoreIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#id HealthcareFhirStoreIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

