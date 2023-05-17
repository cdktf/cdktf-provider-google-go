package healthcarefhirstore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareFhirStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#dataset HealthcareFhirStore#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The resource name for the FhirStore.
	//
	// * Changing this property may recreate the FHIR store (removing all data) **
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#name HealthcareFhirStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The FHIR specification version. Possible values: ["DSTU2", "STU3", "R4"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#version HealthcareFhirStore#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Whether to disable referential integrity in this FHIR store.
	//
	// This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	//
	// * Changing this property may recreate the FHIR store (removing all data) **
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#disable_referential_integrity HealthcareFhirStore#disable_referential_integrity}
	DisableReferentialIntegrity interface{} `field:"optional" json:"disableReferentialIntegrity" yaml:"disableReferentialIntegrity"`
	// Whether to disable resource versioning for this FHIR store.
	//
	// This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	//
	// * Changing this property may recreate the FHIR store (removing all data) **
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#disable_resource_versioning HealthcareFhirStore#disable_resource_versioning}
	DisableResourceVersioning interface{} `field:"optional" json:"disableResourceVersioning" yaml:"disableResourceVersioning"`
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into the FHIR store.
	//
	// Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	//
	// * Changing this property may recreate the FHIR store (removing all data) **
	//
	// * This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#enable_history_import HealthcareFhirStore#enable_history_import}
	EnableHistoryImport interface{} `field:"optional" json:"enableHistoryImport" yaml:"enableHistoryImport"`
	// Whether this FHIR store has the updateCreate capability.
	//
	// This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#enable_update_create HealthcareFhirStore#enable_update_create}
	EnableUpdateCreate interface{} `field:"optional" json:"enableUpdateCreate" yaml:"enableUpdateCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#id HealthcareFhirStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-supplied key-value pairs used to organize FHIR stores.
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#labels HealthcareFhirStore#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#notification_config HealthcareFhirStore#notification_config}
	NotificationConfig *HealthcareFhirStoreNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// stream_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#stream_configs HealthcareFhirStore#stream_configs}
	StreamConfigs interface{} `field:"optional" json:"streamConfigs" yaml:"streamConfigs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/healthcare_fhir_store#timeouts HealthcareFhirStore#timeouts}
	Timeouts *HealthcareFhirStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

