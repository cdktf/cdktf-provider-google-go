// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/iamworkloadidentitypoolprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/iam_workload_identity_pool_provider google_iam_workload_identity_pool_provider}.
type IamWorkloadIdentityPoolProvider interface {
	cdktf.TerraformResource
	AttributeCondition() *string
	SetAttributeCondition(val *string)
	AttributeConditionInput() *string
	AttributeMapping() *map[string]*string
	SetAttributeMapping(val *map[string]*string)
	AttributeMappingInput() *map[string]*string
	Aws() IamWorkloadIdentityPoolProviderAwsOutputReference
	AwsInput() *IamWorkloadIdentityPoolProviderAws
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
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
	Oidc() IamWorkloadIdentityPoolProviderOidcOutputReference
	OidcInput() *IamWorkloadIdentityPoolProviderOidc
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
	// Experimental.
	RawOverrides() interface{}
	Saml() IamWorkloadIdentityPoolProviderSamlOutputReference
	SamlInput() *IamWorkloadIdentityPoolProviderSaml
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IamWorkloadIdentityPoolProviderTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkloadIdentityPoolId() *string
	SetWorkloadIdentityPoolId(val *string)
	WorkloadIdentityPoolIdInput() *string
	WorkloadIdentityPoolProviderId() *string
	SetWorkloadIdentityPoolProviderId(val *string)
	WorkloadIdentityPoolProviderIdInput() *string
	X509() IamWorkloadIdentityPoolProviderX509OutputReference
	X509Input() *IamWorkloadIdentityPoolProviderX509
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
	PutAws(value *IamWorkloadIdentityPoolProviderAws)
	PutOidc(value *IamWorkloadIdentityPoolProviderOidc)
	PutSaml(value *IamWorkloadIdentityPoolProviderSaml)
	PutTimeouts(value *IamWorkloadIdentityPoolProviderTimeouts)
	PutX509(value *IamWorkloadIdentityPoolProviderX509)
	ResetAttributeCondition()
	ResetAttributeMapping()
	ResetAws()
	ResetDescription()
	ResetDisabled()
	ResetDisplayName()
	ResetId()
	ResetOidc()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSaml()
	ResetTimeouts()
	ResetX509()
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

// The jsii proxy struct for IamWorkloadIdentityPoolProvider
type jsiiProxy_IamWorkloadIdentityPoolProvider struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) AttributeCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) AttributeConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) AttributeMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) AttributeMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Aws() IamWorkloadIdentityPoolProviderAwsOutputReference {
	var returns IamWorkloadIdentityPoolProviderAwsOutputReference
	_jsii_.Get(
		j,
		"aws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) AwsInput() *IamWorkloadIdentityPoolProviderAws {
	var returns *IamWorkloadIdentityPoolProviderAws
	_jsii_.Get(
		j,
		"awsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Oidc() IamWorkloadIdentityPoolProviderOidcOutputReference {
	var returns IamWorkloadIdentityPoolProviderOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) OidcInput() *IamWorkloadIdentityPoolProviderOidc {
	var returns *IamWorkloadIdentityPoolProviderOidc
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Saml() IamWorkloadIdentityPoolProviderSamlOutputReference {
	var returns IamWorkloadIdentityPoolProviderSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) SamlInput() *IamWorkloadIdentityPoolProviderSaml {
	var returns *IamWorkloadIdentityPoolProviderSaml
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) Timeouts() IamWorkloadIdentityPoolProviderTimeoutsOutputReference {
	var returns IamWorkloadIdentityPoolProviderTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) WorkloadIdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadIdentityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) WorkloadIdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadIdentityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) WorkloadIdentityPoolProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadIdentityPoolProviderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) WorkloadIdentityPoolProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadIdentityPoolProviderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) X509() IamWorkloadIdentityPoolProviderX509OutputReference {
	var returns IamWorkloadIdentityPoolProviderX509OutputReference
	_jsii_.Get(
		j,
		"x509",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider) X509Input() *IamWorkloadIdentityPoolProviderX509 {
	var returns *IamWorkloadIdentityPoolProviderX509
	_jsii_.Get(
		j,
		"x509Input",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/iam_workload_identity_pool_provider google_iam_workload_identity_pool_provider} Resource.
func NewIamWorkloadIdentityPoolProvider(scope constructs.Construct, id *string, config *IamWorkloadIdentityPoolProviderConfig) IamWorkloadIdentityPoolProvider {
	_init_.Initialize()

	if err := validateNewIamWorkloadIdentityPoolProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkloadIdentityPoolProvider{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/iam_workload_identity_pool_provider google_iam_workload_identity_pool_provider} Resource.
func NewIamWorkloadIdentityPoolProvider_Override(i IamWorkloadIdentityPoolProvider, scope constructs.Construct, id *string, config *IamWorkloadIdentityPoolProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetAttributeCondition(val *string) {
	if err := j.validateSetAttributeConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeCondition",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetAttributeMapping(val *map[string]*string) {
	if err := j.validateSetAttributeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeMapping",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetWorkloadIdentityPoolId(val *string) {
	if err := j.validateSetWorkloadIdentityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadIdentityPoolId",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProvider)SetWorkloadIdentityPoolProviderId(val *string) {
	if err := j.validateSetWorkloadIdentityPoolProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadIdentityPoolProviderId",
		val,
	)
}

// Generates CDKTF code for importing a IamWorkloadIdentityPoolProvider resource upon running "cdktf plan <stack-name>".
func IamWorkloadIdentityPoolProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIamWorkloadIdentityPoolProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
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
func IamWorkloadIdentityPoolProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamWorkloadIdentityPoolProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IamWorkloadIdentityPoolProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamWorkloadIdentityPoolProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IamWorkloadIdentityPoolProvider_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamWorkloadIdentityPoolProvider_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IamWorkloadIdentityPoolProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) PutAws(value *IamWorkloadIdentityPoolProviderAws) {
	if err := i.validatePutAwsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAws",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) PutOidc(value *IamWorkloadIdentityPoolProviderOidc) {
	if err := i.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOidc",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) PutSaml(value *IamWorkloadIdentityPoolProviderSaml) {
	if err := i.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSaml",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) PutTimeouts(value *IamWorkloadIdentityPoolProviderTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) PutX509(value *IamWorkloadIdentityPoolProviderX509) {
	if err := i.validatePutX509Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putX509",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetAttributeCondition() {
	_jsii_.InvokeVoid(
		i,
		"resetAttributeCondition",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetAttributeMapping() {
	_jsii_.InvokeVoid(
		i,
		"resetAttributeMapping",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetAws() {
	_jsii_.InvokeVoid(
		i,
		"resetAws",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetDisabled() {
	_jsii_.InvokeVoid(
		i,
		"resetDisabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetDisplayName() {
	_jsii_.InvokeVoid(
		i,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetOidc() {
	_jsii_.InvokeVoid(
		i,
		"resetOidc",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetProject() {
	_jsii_.InvokeVoid(
		i,
		"resetProject",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetSaml() {
	_jsii_.InvokeVoid(
		i,
		"resetSaml",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ResetX509() {
	_jsii_.InvokeVoid(
		i,
		"resetX509",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

