package healthcaredicomstore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareDicomStoreConfig struct {
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
	// Identifies the dataset addressed by this request. Must be in the format 'projects/{project}/locations/{location}/datasets/{dataset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dicom_store#dataset HealthcareDicomStore#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The resource name for the DicomStore.
	//
	// * Changing this property may recreate the Dicom store (removing all data) **
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dicom_store#name HealthcareDicomStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dicom_store#id HealthcareDicomStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-supplied key-value pairs used to organize DICOM stores.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	// No more than 64 labels can be associated with a given store.
	//
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dicom_store#labels HealthcareDicomStore#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dicom_store#notification_config HealthcareDicomStore#notification_config}
	NotificationConfig *HealthcareDicomStoreNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_dicom_store#timeouts HealthcareDicomStore#timeouts}
	Timeouts *HealthcareDicomStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

