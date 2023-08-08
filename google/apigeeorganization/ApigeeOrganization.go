package apigeeorganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v8/apigeeorganization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.77.0/docs/resources/apigee_organization google_apigee_organization}.
type ApigeeOrganization interface {
	cdktf.TerraformResource
	AnalyticsRegion() *string
	SetAnalyticsRegion(val *string)
	AnalyticsRegionInput() *string
	ApigeeProjectId() *string
	AuthorizedNetwork() *string
	SetAuthorizedNetwork(val *string)
	AuthorizedNetworkInput() *string
	BillingType() *string
	SetBillingType(val *string)
	BillingTypeInput() *string
	CaCertificate() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableVpcPeering() interface{}
	SetDisableVpcPeering(val interface{})
	DisableVpcPeeringInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	Properties() ApigeeOrganizationPropertiesOutputReference
	PropertiesInput() *ApigeeOrganizationProperties
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
	Retention() *string
	SetRetention(val *string)
	RetentionInput() *string
	RuntimeDatabaseEncryptionKeyName() *string
	SetRuntimeDatabaseEncryptionKeyName(val *string)
	RuntimeDatabaseEncryptionKeyNameInput() *string
	RuntimeType() *string
	SetRuntimeType(val *string)
	RuntimeTypeInput() *string
	SubscriptionType() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApigeeOrganizationTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutProperties(value *ApigeeOrganizationProperties)
	PutTimeouts(value *ApigeeOrganizationTimeouts)
	ResetAnalyticsRegion()
	ResetAuthorizedNetwork()
	ResetBillingType()
	ResetDescription()
	ResetDisableVpcPeering()
	ResetDisplayName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProperties()
	ResetRetention()
	ResetRuntimeDatabaseEncryptionKeyName()
	ResetRuntimeType()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApigeeOrganization
type jsiiProxy_ApigeeOrganization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigeeOrganization) AnalyticsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) AnalyticsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) ApigeeProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) AuthorizedNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) AuthorizedNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) BillingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) BillingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) DisableVpcPeering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableVpcPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) DisableVpcPeeringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableVpcPeeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Properties() ApigeeOrganizationPropertiesOutputReference {
	var returns ApigeeOrganizationPropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) PropertiesInput() *ApigeeOrganizationProperties {
	var returns *ApigeeOrganizationProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Retention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) RetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) RuntimeDatabaseEncryptionKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeDatabaseEncryptionKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) RuntimeDatabaseEncryptionKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeDatabaseEncryptionKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) RuntimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) RuntimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) Timeouts() ApigeeOrganizationTimeoutsOutputReference {
	var returns ApigeeOrganizationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeOrganization) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.77.0/docs/resources/apigee_organization google_apigee_organization} Resource.
func NewApigeeOrganization(scope constructs.Construct, id *string, config *ApigeeOrganizationConfig) ApigeeOrganization {
	_init_.Initialize()

	if err := validateNewApigeeOrganizationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeOrganization{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeOrganization.ApigeeOrganization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.77.0/docs/resources/apigee_organization google_apigee_organization} Resource.
func NewApigeeOrganization_Override(a ApigeeOrganization, scope constructs.Construct, id *string, config *ApigeeOrganizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeOrganization.ApigeeOrganization",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetAnalyticsRegion(val *string) {
	if err := j.validateSetAnalyticsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticsRegion",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetAuthorizedNetwork(val *string) {
	if err := j.validateSetAuthorizedNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedNetwork",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetBillingType(val *string) {
	if err := j.validateSetBillingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingType",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetDisableVpcPeering(val interface{}) {
	if err := j.validateSetDisableVpcPeeringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableVpcPeering",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetRetention(val *string) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetRuntimeDatabaseEncryptionKeyName(val *string) {
	if err := j.validateSetRuntimeDatabaseEncryptionKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeDatabaseEncryptionKeyName",
		val,
	)
}

func (j *jsiiProxy_ApigeeOrganization)SetRuntimeType(val *string) {
	if err := j.validateSetRuntimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeType",
		val,
	)
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
func ApigeeOrganization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeOrganization_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeOrganization.ApigeeOrganization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeOrganization_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeOrganization_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeOrganization.ApigeeOrganization",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeOrganization_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeOrganization_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeOrganization.ApigeeOrganization",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigeeOrganization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.apigeeOrganization.ApigeeOrganization",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigeeOrganization) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigeeOrganization) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigeeOrganization) PutProperties(value *ApigeeOrganizationProperties) {
	if err := a.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeOrganization) PutTimeouts(value *ApigeeOrganizationTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetAnalyticsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalyticsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetAuthorizedNetwork() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizedNetwork",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetBillingType() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetDisableVpcPeering() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableVpcPeering",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetRuntimeDatabaseEncryptionKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeDatabaseEncryptionKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetRuntimeType() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeOrganization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeOrganization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

