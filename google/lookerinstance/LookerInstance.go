package lookerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v8/lookerinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/looker_instance google_looker_instance}.
type LookerInstance interface {
	cdktf.TerraformResource
	AdminSettings() LookerInstanceAdminSettingsOutputReference
	AdminSettingsInput() *LookerInstanceAdminSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumerNetwork() *string
	SetConsumerNetwork(val *string)
	ConsumerNetworkInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DenyMaintenancePeriod() LookerInstanceDenyMaintenancePeriodOutputReference
	DenyMaintenancePeriodInput() *LookerInstanceDenyMaintenancePeriod
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EgressPublicIp() *string
	EncryptionConfig() LookerInstanceEncryptionConfigOutputReference
	EncryptionConfigInput() *LookerInstanceEncryptionConfig
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
	IngressPrivateIp() *string
	IngressPublicIp() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LookerUri() *string
	LookerVersion() *string
	MaintenanceWindow() LookerInstanceMaintenanceWindowOutputReference
	MaintenanceWindowInput() *LookerInstanceMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OauthConfig() LookerInstanceOauthConfigOutputReference
	OauthConfigInput() *LookerInstanceOauthConfig
	PlatformEdition() *string
	SetPlatformEdition(val *string)
	PlatformEditionInput() *string
	PrivateIpEnabled() interface{}
	SetPrivateIpEnabled(val interface{})
	PrivateIpEnabledInput() interface{}
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpEnabled() interface{}
	SetPublicIpEnabled(val interface{})
	PublicIpEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReservedRange() *string
	SetReservedRange(val *string)
	ReservedRangeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LookerInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	UserMetadata() LookerInstanceUserMetadataOutputReference
	UserMetadataInput() *LookerInstanceUserMetadata
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
	PutAdminSettings(value *LookerInstanceAdminSettings)
	PutDenyMaintenancePeriod(value *LookerInstanceDenyMaintenancePeriod)
	PutEncryptionConfig(value *LookerInstanceEncryptionConfig)
	PutMaintenanceWindow(value *LookerInstanceMaintenanceWindow)
	PutOauthConfig(value *LookerInstanceOauthConfig)
	PutTimeouts(value *LookerInstanceTimeouts)
	PutUserMetadata(value *LookerInstanceUserMetadata)
	ResetAdminSettings()
	ResetConsumerNetwork()
	ResetDenyMaintenancePeriod()
	ResetEncryptionConfig()
	ResetId()
	ResetMaintenanceWindow()
	ResetOauthConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformEdition()
	ResetPrivateIpEnabled()
	ResetProject()
	ResetPublicIpEnabled()
	ResetRegion()
	ResetReservedRange()
	ResetTimeouts()
	ResetUserMetadata()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LookerInstance
type jsiiProxy_LookerInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LookerInstance) AdminSettings() LookerInstanceAdminSettingsOutputReference {
	var returns LookerInstanceAdminSettingsOutputReference
	_jsii_.Get(
		j,
		"adminSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) AdminSettingsInput() *LookerInstanceAdminSettings {
	var returns *LookerInstanceAdminSettings
	_jsii_.Get(
		j,
		"adminSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ConsumerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ConsumerNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) DenyMaintenancePeriod() LookerInstanceDenyMaintenancePeriodOutputReference {
	var returns LookerInstanceDenyMaintenancePeriodOutputReference
	_jsii_.Get(
		j,
		"denyMaintenancePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) DenyMaintenancePeriodInput() *LookerInstanceDenyMaintenancePeriod {
	var returns *LookerInstanceDenyMaintenancePeriod
	_jsii_.Get(
		j,
		"denyMaintenancePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) EgressPublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) EncryptionConfig() LookerInstanceEncryptionConfigOutputReference {
	var returns LookerInstanceEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) EncryptionConfigInput() *LookerInstanceEncryptionConfig {
	var returns *LookerInstanceEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) IngressPrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) IngressPublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) LookerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) LookerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) MaintenanceWindow() LookerInstanceMaintenanceWindowOutputReference {
	var returns LookerInstanceMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) MaintenanceWindowInput() *LookerInstanceMaintenanceWindow {
	var returns *LookerInstanceMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) OauthConfig() LookerInstanceOauthConfigOutputReference {
	var returns LookerInstanceOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) OauthConfigInput() *LookerInstanceOauthConfig {
	var returns *LookerInstanceOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) PlatformEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) PlatformEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) PrivateIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) PrivateIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) PublicIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) PublicIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ReservedRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) ReservedRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) Timeouts() LookerInstanceTimeoutsOutputReference {
	var returns LookerInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) UserMetadata() LookerInstanceUserMetadataOutputReference {
	var returns LookerInstanceUserMetadataOutputReference
	_jsii_.Get(
		j,
		"userMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstance) UserMetadataInput() *LookerInstanceUserMetadata {
	var returns *LookerInstanceUserMetadata
	_jsii_.Get(
		j,
		"userMetadataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/looker_instance google_looker_instance} Resource.
func NewLookerInstance(scope constructs.Construct, id *string, config *LookerInstanceConfig) LookerInstance {
	_init_.Initialize()

	if err := validateNewLookerInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookerInstance{}

	_jsii_.Create(
		"@cdktf/provider-google.lookerInstance.LookerInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/looker_instance google_looker_instance} Resource.
func NewLookerInstance_Override(l LookerInstance, scope constructs.Construct, id *string, config *LookerInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lookerInstance.LookerInstance",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LookerInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetConsumerNetwork(val *string) {
	if err := j.validateSetConsumerNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerNetwork",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetPlatformEdition(val *string) {
	if err := j.validateSetPlatformEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformEdition",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetPrivateIpEnabled(val interface{}) {
	if err := j.validateSetPrivateIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpEnabled",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetPublicIpEnabled(val interface{}) {
	if err := j.validateSetPublicIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpEnabled",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_LookerInstance)SetReservedRange(val *string) {
	if err := j.validateSetReservedRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedRange",
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
func LookerInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLookerInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.lookerInstance.LookerInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LookerInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLookerInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.lookerInstance.LookerInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LookerInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLookerInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.lookerInstance.LookerInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LookerInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.lookerInstance.LookerInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LookerInstance) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LookerInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LookerInstance) PutAdminSettings(value *LookerInstanceAdminSettings) {
	if err := l.validatePutAdminSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdminSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) PutDenyMaintenancePeriod(value *LookerInstanceDenyMaintenancePeriod) {
	if err := l.validatePutDenyMaintenancePeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDenyMaintenancePeriod",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) PutEncryptionConfig(value *LookerInstanceEncryptionConfig) {
	if err := l.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) PutMaintenanceWindow(value *LookerInstanceMaintenanceWindow) {
	if err := l.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) PutOauthConfig(value *LookerInstanceOauthConfig) {
	if err := l.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) PutTimeouts(value *LookerInstanceTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) PutUserMetadata(value *LookerInstanceUserMetadata) {
	if err := l.validatePutUserMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUserMetadata",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstance) ResetAdminSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetConsumerNetwork() {
	_jsii_.InvokeVoid(
		l,
		"resetConsumerNetwork",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetDenyMaintenancePeriod() {
	_jsii_.InvokeVoid(
		l,
		"resetDenyMaintenancePeriod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetPlatformEdition() {
	_jsii_.InvokeVoid(
		l,
		"resetPlatformEdition",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetPrivateIpEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateIpEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetProject() {
	_jsii_.InvokeVoid(
		l,
		"resetProject",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetPublicIpEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicIpEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		l,
		"resetRegion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetReservedRange() {
	_jsii_.InvokeVoid(
		l,
		"resetReservedRange",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) ResetUserMetadata() {
	_jsii_.InvokeVoid(
		l,
		"resetUserMetadata",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookerInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

