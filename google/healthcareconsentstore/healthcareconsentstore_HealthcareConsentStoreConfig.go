package healthcareconsentstore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareConsentStoreConfig struct {
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
	// Identifies the dataset addressed by this request. Must be in the format 'projects/{project}/locations/{location}/datasets/{dataset}'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#dataset HealthcareConsentStore#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The name of this ConsentStore, for example: "consent1".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#name HealthcareConsentStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Default time to live for consents in this store.
	//
	// Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#default_consent_ttl HealthcareConsentStore#default_consent_ttl}
	DefaultConsentTtl *string `field:"optional" json:"defaultConsentTtl" yaml:"defaultConsentTtl"`
	// If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#enable_consent_create_on_update HealthcareConsentStore#enable_consent_create_on_update}
	EnableConsentCreateOnUpdate interface{} `field:"optional" json:"enableConsentCreateOnUpdate" yaml:"enableConsentCreateOnUpdate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#id HealthcareConsentStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-supplied key-value pairs used to organize Consent stores.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: '[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}'
	//
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: '[\p{Ll}\p{Lo}\p{N}_-]{0,63}'
	//
	// No more than 64 labels can be associated with a given store.
	//
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#labels HealthcareConsentStore#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_consent_store#timeouts HealthcareConsentStore#timeouts}
	Timeouts *HealthcareConsentStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

