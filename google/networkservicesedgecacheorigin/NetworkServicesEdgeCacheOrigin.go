// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesedgecacheorigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v12/networkservicesedgecacheorigin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/network_services_edge_cache_origin google_network_services_edge_cache_origin}.
type NetworkServicesEdgeCacheOrigin interface {
	cdktf.TerraformResource
	AwsV4Authentication() NetworkServicesEdgeCacheOriginAwsV4AuthenticationOutputReference
	AwsV4AuthenticationInput() *NetworkServicesEdgeCacheOriginAwsV4Authentication
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
	EffectiveLabels() cdktf.StringMap
	FailoverOrigin() *string
	SetFailoverOrigin(val *string)
	FailoverOriginInput() *string
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
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OriginAddress() *string
	SetOriginAddress(val *string)
	OriginAddressInput() *string
	OriginOverrideAction() NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference
	OriginOverrideActionInput() *NetworkServicesEdgeCacheOriginOriginOverrideAction
	OriginRedirect() NetworkServicesEdgeCacheOriginOriginRedirectOutputReference
	OriginRedirectInput() *NetworkServicesEdgeCacheOriginOriginRedirect
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	RetryConditions() *[]*string
	SetRetryConditions(val *[]*string)
	RetryConditionsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() NetworkServicesEdgeCacheOriginTimeoutOutputReference
	TimeoutInput() *NetworkServicesEdgeCacheOriginTimeout
	Timeouts() NetworkServicesEdgeCacheOriginTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAwsV4Authentication(value *NetworkServicesEdgeCacheOriginAwsV4Authentication)
	PutOriginOverrideAction(value *NetworkServicesEdgeCacheOriginOriginOverrideAction)
	PutOriginRedirect(value *NetworkServicesEdgeCacheOriginOriginRedirect)
	PutTimeout(value *NetworkServicesEdgeCacheOriginTimeout)
	PutTimeouts(value *NetworkServicesEdgeCacheOriginTimeouts)
	ResetAwsV4Authentication()
	ResetDescription()
	ResetFailoverOrigin()
	ResetId()
	ResetLabels()
	ResetMaxAttempts()
	ResetOriginOverrideAction()
	ResetOriginRedirect()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProject()
	ResetProtocol()
	ResetRetryConditions()
	ResetTimeout()
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

// The jsii proxy struct for NetworkServicesEdgeCacheOrigin
type jsiiProxy_NetworkServicesEdgeCacheOrigin struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) AwsV4Authentication() NetworkServicesEdgeCacheOriginAwsV4AuthenticationOutputReference {
	var returns NetworkServicesEdgeCacheOriginAwsV4AuthenticationOutputReference
	_jsii_.Get(
		j,
		"awsV4Authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) AwsV4AuthenticationInput() *NetworkServicesEdgeCacheOriginAwsV4Authentication {
	var returns *NetworkServicesEdgeCacheOriginAwsV4Authentication
	_jsii_.Get(
		j,
		"awsV4AuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FailoverOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FailoverOriginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginOverrideAction() NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference {
	var returns NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference
	_jsii_.Get(
		j,
		"originOverrideAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginOverrideActionInput() *NetworkServicesEdgeCacheOriginOriginOverrideAction {
	var returns *NetworkServicesEdgeCacheOriginOriginOverrideAction
	_jsii_.Get(
		j,
		"originOverrideActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginRedirect() NetworkServicesEdgeCacheOriginOriginRedirectOutputReference {
	var returns NetworkServicesEdgeCacheOriginOriginRedirectOutputReference
	_jsii_.Get(
		j,
		"originRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginRedirectInput() *NetworkServicesEdgeCacheOriginOriginRedirect {
	var returns *NetworkServicesEdgeCacheOriginOriginRedirect
	_jsii_.Get(
		j,
		"originRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) RetryConditions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) RetryConditionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Timeout() NetworkServicesEdgeCacheOriginTimeoutOutputReference {
	var returns NetworkServicesEdgeCacheOriginTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TimeoutInput() *NetworkServicesEdgeCacheOriginTimeout {
	var returns *NetworkServicesEdgeCacheOriginTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Timeouts() NetworkServicesEdgeCacheOriginTimeoutsOutputReference {
	var returns NetworkServicesEdgeCacheOriginTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/network_services_edge_cache_origin google_network_services_edge_cache_origin} Resource.
func NewNetworkServicesEdgeCacheOrigin(scope constructs.Construct, id *string, config *NetworkServicesEdgeCacheOriginConfig) NetworkServicesEdgeCacheOrigin {
	_init_.Initialize()

	if err := validateNewNetworkServicesEdgeCacheOriginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesEdgeCacheOrigin{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/network_services_edge_cache_origin google_network_services_edge_cache_origin} Resource.
func NewNetworkServicesEdgeCacheOrigin_Override(n NetworkServicesEdgeCacheOrigin, scope constructs.Construct, id *string, config *NetworkServicesEdgeCacheOriginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetFailoverOrigin(val *string) {
	if err := j.validateSetFailoverOriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverOrigin",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetMaxAttempts(val *float64) {
	if err := j.validateSetMaxAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetOriginAddress(val *string) {
	if err := j.validateSetOriginAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAddress",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetRetryConditions(val *[]*string) {
	if err := j.validateSetRetryConditionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryConditions",
		val,
	)
}

// Generates CDKTF code for importing a NetworkServicesEdgeCacheOrigin resource upon running "cdktf plan <stack-name>".
func NetworkServicesEdgeCacheOrigin_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
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
func NetworkServicesEdgeCacheOrigin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkServicesEdgeCacheOrigin_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkServicesEdgeCacheOrigin_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkServicesEdgeCacheOrigin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutAwsV4Authentication(value *NetworkServicesEdgeCacheOriginAwsV4Authentication) {
	if err := n.validatePutAwsV4AuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAwsV4Authentication",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutOriginOverrideAction(value *NetworkServicesEdgeCacheOriginOriginOverrideAction) {
	if err := n.validatePutOriginOverrideActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putOriginOverrideAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutOriginRedirect(value *NetworkServicesEdgeCacheOriginOriginRedirect) {
	if err := n.validatePutOriginRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putOriginRedirect",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutTimeout(value *NetworkServicesEdgeCacheOriginTimeout) {
	if err := n.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeout",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutTimeouts(value *NetworkServicesEdgeCacheOriginTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetAwsV4Authentication() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsV4Authentication",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetFailoverOrigin() {
	_jsii_.InvokeVoid(
		n,
		"resetFailoverOrigin",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetMaxAttempts() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxAttempts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetOriginOverrideAction() {
	_jsii_.InvokeVoid(
		n,
		"resetOriginOverrideAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetOriginRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetOriginRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetPort() {
	_jsii_.InvokeVoid(
		n,
		"resetPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetRetryConditions() {
	_jsii_.InvokeVoid(
		n,
		"resetRetryConditions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

