// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computehealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v9/computehealthcheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/compute_health_check google_compute_health_check}.
type ComputeHealthCheck interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckIntervalSec() *float64
	SetCheckIntervalSec(val *float64)
	CheckIntervalSecInput() *float64
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrpcHealthCheck() ComputeHealthCheckGrpcHealthCheckOutputReference
	GrpcHealthCheckInput() *ComputeHealthCheckGrpcHealthCheck
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	Http2HealthCheck() ComputeHealthCheckHttp2HealthCheckOutputReference
	Http2HealthCheckInput() *ComputeHealthCheckHttp2HealthCheck
	HttpHealthCheck() ComputeHealthCheckHttpHealthCheckOutputReference
	HttpHealthCheckInput() *ComputeHealthCheckHttpHealthCheck
	HttpsHealthCheck() ComputeHealthCheckHttpsHealthCheckOutputReference
	HttpsHealthCheckInput() *ComputeHealthCheckHttpsHealthCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() ComputeHealthCheckLogConfigOutputReference
	LogConfigInput() *ComputeHealthCheckLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	SelfLink() *string
	SslHealthCheck() ComputeHealthCheckSslHealthCheckOutputReference
	SslHealthCheckInput() *ComputeHealthCheckSslHealthCheck
	TcpHealthCheck() ComputeHealthCheckTcpHealthCheckOutputReference
	TcpHealthCheckInput() *ComputeHealthCheckTcpHealthCheck
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeHealthCheckTimeoutsOutputReference
	TimeoutSec() *float64
	SetTimeoutSec(val *float64)
	TimeoutSecInput() *float64
	TimeoutsInput() interface{}
	Type() *string
	UnhealthyThreshold() *float64
	SetUnhealthyThreshold(val *float64)
	UnhealthyThresholdInput() *float64
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
	PutGrpcHealthCheck(value *ComputeHealthCheckGrpcHealthCheck)
	PutHttp2HealthCheck(value *ComputeHealthCheckHttp2HealthCheck)
	PutHttpHealthCheck(value *ComputeHealthCheckHttpHealthCheck)
	PutHttpsHealthCheck(value *ComputeHealthCheckHttpsHealthCheck)
	PutLogConfig(value *ComputeHealthCheckLogConfig)
	PutSslHealthCheck(value *ComputeHealthCheckSslHealthCheck)
	PutTcpHealthCheck(value *ComputeHealthCheckTcpHealthCheck)
	PutTimeouts(value *ComputeHealthCheckTimeouts)
	ResetCheckIntervalSec()
	ResetDescription()
	ResetGrpcHealthCheck()
	ResetHealthyThreshold()
	ResetHttp2HealthCheck()
	ResetHttpHealthCheck()
	ResetHttpsHealthCheck()
	ResetId()
	ResetLogConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSslHealthCheck()
	ResetTcpHealthCheck()
	ResetTimeouts()
	ResetTimeoutSec()
	ResetUnhealthyThreshold()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeHealthCheck
type jsiiProxy_ComputeHealthCheck struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeHealthCheck) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) CheckIntervalSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) CheckIntervalSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) GrpcHealthCheck() ComputeHealthCheckGrpcHealthCheckOutputReference {
	var returns ComputeHealthCheckGrpcHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) GrpcHealthCheckInput() *ComputeHealthCheckGrpcHealthCheck {
	var returns *ComputeHealthCheckGrpcHealthCheck
	_jsii_.Get(
		j,
		"grpcHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Http2HealthCheck() ComputeHealthCheckHttp2HealthCheckOutputReference {
	var returns ComputeHealthCheckHttp2HealthCheckOutputReference
	_jsii_.Get(
		j,
		"http2HealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Http2HealthCheckInput() *ComputeHealthCheckHttp2HealthCheck {
	var returns *ComputeHealthCheckHttp2HealthCheck
	_jsii_.Get(
		j,
		"http2HealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpHealthCheck() ComputeHealthCheckHttpHealthCheckOutputReference {
	var returns ComputeHealthCheckHttpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpHealthCheckInput() *ComputeHealthCheckHttpHealthCheck {
	var returns *ComputeHealthCheckHttpHealthCheck
	_jsii_.Get(
		j,
		"httpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpsHealthCheck() ComputeHealthCheckHttpsHealthCheckOutputReference {
	var returns ComputeHealthCheckHttpsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpsHealthCheckInput() *ComputeHealthCheckHttpsHealthCheck {
	var returns *ComputeHealthCheckHttpsHealthCheck
	_jsii_.Get(
		j,
		"httpsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) LogConfig() ComputeHealthCheckLogConfigOutputReference {
	var returns ComputeHealthCheckLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) LogConfigInput() *ComputeHealthCheckLogConfig {
	var returns *ComputeHealthCheckLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SslHealthCheck() ComputeHealthCheckSslHealthCheckOutputReference {
	var returns ComputeHealthCheckSslHealthCheckOutputReference
	_jsii_.Get(
		j,
		"sslHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SslHealthCheckInput() *ComputeHealthCheckSslHealthCheck {
	var returns *ComputeHealthCheckSslHealthCheck
	_jsii_.Get(
		j,
		"sslHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TcpHealthCheck() ComputeHealthCheckTcpHealthCheckOutputReference {
	var returns ComputeHealthCheckTcpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"tcpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TcpHealthCheckInput() *ComputeHealthCheckTcpHealthCheck {
	var returns *ComputeHealthCheckTcpHealthCheck
	_jsii_.Get(
		j,
		"tcpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Timeouts() ComputeHealthCheckTimeoutsOutputReference {
	var returns ComputeHealthCheckTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/compute_health_check google_compute_health_check} Resource.
func NewComputeHealthCheck(scope constructs.Construct, id *string, config *ComputeHealthCheckConfig) ComputeHealthCheck {
	_init_.Initialize()

	if err := validateNewComputeHealthCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeHealthCheck{}

	_jsii_.Create(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheck",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/compute_health_check google_compute_health_check} Resource.
func NewComputeHealthCheck_Override(c ComputeHealthCheck, scope constructs.Construct, id *string, config *ComputeHealthCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheck",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetCheckIntervalSec(val *float64) {
	if err := j.validateSetCheckIntervalSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkIntervalSec",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetHealthyThreshold(val *float64) {
	if err := j.validateSetHealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetUnhealthyThreshold(val *float64) {
	if err := j.validateSetUnhealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThreshold",
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
func ComputeHealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeHealthCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeHealthCheck_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeHealthCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheck",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeHealthCheck_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeHealthCheck_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheck",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeHealthCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutGrpcHealthCheck(value *ComputeHealthCheckGrpcHealthCheck) {
	if err := c.validatePutGrpcHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpcHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutHttp2HealthCheck(value *ComputeHealthCheckHttp2HealthCheck) {
	if err := c.validatePutHttp2HealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttp2HealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutHttpHealthCheck(value *ComputeHealthCheckHttpHealthCheck) {
	if err := c.validatePutHttpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutHttpsHealthCheck(value *ComputeHealthCheckHttpsHealthCheck) {
	if err := c.validatePutHttpsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpsHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutLogConfig(value *ComputeHealthCheckLogConfig) {
	if err := c.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutSslHealthCheck(value *ComputeHealthCheckSslHealthCheck) {
	if err := c.validatePutSslHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSslHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutTcpHealthCheck(value *ComputeHealthCheckTcpHealthCheck) {
	if err := c.validatePutTcpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTcpHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutTimeouts(value *ComputeHealthCheckTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetCheckIntervalSec() {
	_jsii_.InvokeVoid(
		c,
		"resetCheckIntervalSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetGrpcHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpcHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHttp2HealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttp2HealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHttpHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHttpsHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetLogConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetSslHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetSslHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetTcpHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

