// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/healthcarefhirstore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/healthcare_fhir_store google_healthcare_fhir_store}.
type HealthcareFhirStore interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComplexDataTypeReferenceParsing() *string
	SetComplexDataTypeReferenceParsing(val *string)
	ComplexDataTypeReferenceParsingInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Dataset() *string
	SetDataset(val *string)
	DatasetInput() *string
	DefaultSearchHandlingStrict() interface{}
	SetDefaultSearchHandlingStrict(val interface{})
	DefaultSearchHandlingStrictInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableReferentialIntegrity() interface{}
	SetDisableReferentialIntegrity(val interface{})
	DisableReferentialIntegrityInput() interface{}
	DisableResourceVersioning() interface{}
	SetDisableResourceVersioning(val interface{})
	DisableResourceVersioningInput() interface{}
	EffectiveLabels() cdktf.StringMap
	EnableHistoryImport() interface{}
	SetEnableHistoryImport(val interface{})
	EnableHistoryImportInput() interface{}
	EnableUpdateCreate() interface{}
	SetEnableUpdateCreate(val interface{})
	EnableUpdateCreateInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationConfig() HealthcareFhirStoreNotificationConfigOutputReference
	NotificationConfigInput() *HealthcareFhirStoreNotificationConfig
	NotificationConfigs() HealthcareFhirStoreNotificationConfigsList
	NotificationConfigsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	StreamConfigs() HealthcareFhirStoreStreamConfigsList
	StreamConfigsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() HealthcareFhirStoreTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutNotificationConfig(value *HealthcareFhirStoreNotificationConfig)
	PutNotificationConfigs(value interface{})
	PutStreamConfigs(value interface{})
	PutTimeouts(value *HealthcareFhirStoreTimeouts)
	ResetComplexDataTypeReferenceParsing()
	ResetDefaultSearchHandlingStrict()
	ResetDisableReferentialIntegrity()
	ResetDisableResourceVersioning()
	ResetEnableHistoryImport()
	ResetEnableUpdateCreate()
	ResetId()
	ResetLabels()
	ResetNotificationConfig()
	ResetNotificationConfigs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamConfigs()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HealthcareFhirStore
type jsiiProxy_HealthcareFhirStore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HealthcareFhirStore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ComplexDataTypeReferenceParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexDataTypeReferenceParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ComplexDataTypeReferenceParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexDataTypeReferenceParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DefaultSearchHandlingStrict() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSearchHandlingStrict",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DefaultSearchHandlingStrictInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSearchHandlingStrictInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableReferentialIntegrity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferentialIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableReferentialIntegrityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferentialIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableResourceVersioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableResourceVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableResourceVersioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableResourceVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableHistoryImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHistoryImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableHistoryImportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHistoryImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableUpdateCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUpdateCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableUpdateCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUpdateCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfig() HealthcareFhirStoreNotificationConfigOutputReference {
	var returns HealthcareFhirStoreNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfigInput() *HealthcareFhirStoreNotificationConfig {
	var returns *HealthcareFhirStoreNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfigs() HealthcareFhirStoreNotificationConfigsList {
	var returns HealthcareFhirStoreNotificationConfigsList
	_jsii_.Get(
		j,
		"notificationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) StreamConfigs() HealthcareFhirStoreStreamConfigsList {
	var returns HealthcareFhirStoreStreamConfigsList
	_jsii_.Get(
		j,
		"streamConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) StreamConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Timeouts() HealthcareFhirStoreTimeoutsOutputReference {
	var returns HealthcareFhirStoreTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/healthcare_fhir_store google_healthcare_fhir_store} Resource.
func NewHealthcareFhirStore(scope constructs.Construct, id *string, config *HealthcareFhirStoreConfig) HealthcareFhirStore {
	_init_.Initialize()

	if err := validateNewHealthcareFhirStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareFhirStore{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/healthcare_fhir_store google_healthcare_fhir_store} Resource.
func NewHealthcareFhirStore_Override(h HealthcareFhirStore, scope constructs.Construct, id *string, config *HealthcareFhirStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetComplexDataTypeReferenceParsing(val *string) {
	if err := j.validateSetComplexDataTypeReferenceParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexDataTypeReferenceParsing",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetDataset(val *string) {
	if err := j.validateSetDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataset",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetDefaultSearchHandlingStrict(val interface{}) {
	if err := j.validateSetDefaultSearchHandlingStrictParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSearchHandlingStrict",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetDisableReferentialIntegrity(val interface{}) {
	if err := j.validateSetDisableReferentialIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableReferentialIntegrity",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetDisableResourceVersioning(val interface{}) {
	if err := j.validateSetDisableResourceVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableResourceVersioning",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetEnableHistoryImport(val interface{}) {
	if err := j.validateSetEnableHistoryImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHistoryImport",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetEnableUpdateCreate(val interface{}) {
	if err := j.validateSetEnableUpdateCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUpdateCreate",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a HealthcareFhirStore resource upon running "cdktf plan <stack-name>".
func HealthcareFhirStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func HealthcareFhirStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareFhirStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareFhirStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HealthcareFhirStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutNotificationConfig(value *HealthcareFhirStoreNotificationConfig) {
	if err := h.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutNotificationConfigs(value interface{}) {
	if err := h.validatePutNotificationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNotificationConfigs",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutStreamConfigs(value interface{}) {
	if err := h.validatePutStreamConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStreamConfigs",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutTimeouts(value *HealthcareFhirStoreTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetComplexDataTypeReferenceParsing() {
	_jsii_.InvokeVoid(
		h,
		"resetComplexDataTypeReferenceParsing",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetDefaultSearchHandlingStrict() {
	_jsii_.InvokeVoid(
		h,
		"resetDefaultSearchHandlingStrict",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetDisableReferentialIntegrity() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableReferentialIntegrity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetDisableResourceVersioning() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableResourceVersioning",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetEnableHistoryImport() {
	_jsii_.InvokeVoid(
		h,
		"resetEnableHistoryImport",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetEnableUpdateCreate() {
	_jsii_.InvokeVoid(
		h,
		"resetEnableUpdateCreate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetLabels() {
	_jsii_.InvokeVoid(
		h,
		"resetLabels",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		h,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetNotificationConfigs() {
	_jsii_.InvokeVoid(
		h,
		"resetNotificationConfigs",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetStreamConfigs() {
	_jsii_.InvokeVoid(
		h,
		"resetStreamConfigs",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

