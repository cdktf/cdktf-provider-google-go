// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomaintrust

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/activedirectorydomaintrust/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/active_directory_domain_trust google_active_directory_domain_trust}.
type ActiveDirectoryDomainTrust interface {
	cdktf.TerraformResource
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
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
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
	// The tree node.
	Node() constructs.Node
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
	SelectiveAuthentication() interface{}
	SetSelectiveAuthentication(val interface{})
	SelectiveAuthenticationInput() interface{}
	TargetDnsIpAddresses() *[]*string
	SetTargetDnsIpAddresses(val *[]*string)
	TargetDnsIpAddressesInput() *[]*string
	TargetDomainName() *string
	SetTargetDomainName(val *string)
	TargetDomainNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ActiveDirectoryDomainTrustTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustDirection() *string
	SetTrustDirection(val *string)
	TrustDirectionInput() *string
	TrustHandshakeSecret() *string
	SetTrustHandshakeSecret(val *string)
	TrustHandshakeSecretInput() *string
	TrustType() *string
	SetTrustType(val *string)
	TrustTypeInput() *string
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
	PutTimeouts(value *ActiveDirectoryDomainTrustTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSelectiveAuthentication()
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

// The jsii proxy struct for ActiveDirectoryDomainTrust
type jsiiProxy_ActiveDirectoryDomainTrust struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) SelectiveAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectiveAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) SelectiveAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectiveAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TargetDnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetDnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TargetDnsIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetDnsIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TargetDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TargetDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) Timeouts() ActiveDirectoryDomainTrustTimeoutsOutputReference {
	var returns ActiveDirectoryDomainTrustTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TrustDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TrustDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TrustHandshakeSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustHandshakeSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TrustHandshakeSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustHandshakeSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TrustType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust) TrustTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/active_directory_domain_trust google_active_directory_domain_trust} Resource.
func NewActiveDirectoryDomainTrust(scope constructs.Construct, id *string, config *ActiveDirectoryDomainTrustConfig) ActiveDirectoryDomainTrust {
	_init_.Initialize()

	if err := validateNewActiveDirectoryDomainTrustParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveDirectoryDomainTrust{}

	_jsii_.Create(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/active_directory_domain_trust google_active_directory_domain_trust} Resource.
func NewActiveDirectoryDomainTrust_Override(a ActiveDirectoryDomainTrust, scope constructs.Construct, id *string, config *ActiveDirectoryDomainTrustConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetSelectiveAuthentication(val interface{}) {
	if err := j.validateSetSelectiveAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectiveAuthentication",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetTargetDnsIpAddresses(val *[]*string) {
	if err := j.validateSetTargetDnsIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDnsIpAddresses",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetTargetDomainName(val *string) {
	if err := j.validateSetTargetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDomainName",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetTrustDirection(val *string) {
	if err := j.validateSetTrustDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDirection",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetTrustHandshakeSecret(val *string) {
	if err := j.validateSetTrustHandshakeSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustHandshakeSecret",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainTrust)SetTrustType(val *string) {
	if err := j.validateSetTrustTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustType",
		val,
	)
}

// Generates CDKTF code for importing a ActiveDirectoryDomainTrust resource upon running "cdktf plan <stack-name>".
func ActiveDirectoryDomainTrust_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainTrust_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
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
func ActiveDirectoryDomainTrust_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainTrust_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActiveDirectoryDomainTrust_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainTrust_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActiveDirectoryDomainTrust_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainTrust_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ActiveDirectoryDomainTrust_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.activeDirectoryDomainTrust.ActiveDirectoryDomainTrust",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActiveDirectoryDomainTrust) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) PutTimeouts(value *ActiveDirectoryDomainTrustTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ResetProject() {
	_jsii_.InvokeVoid(
		a,
		"resetProject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ResetSelectiveAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetSelectiveAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainTrust) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

