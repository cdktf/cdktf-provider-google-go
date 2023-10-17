// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionReadinessCheckOutputReference interface {
	cdktf.ComplexObject
	AppStartTimeout() *string
	SetAppStartTimeout(val *string)
	AppStartTimeoutInput() *string
	CheckInterval() *string
	SetCheckInterval(val *string)
	CheckIntervalInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *AppEngineFlexibleAppVersionReadinessCheck
	SetInternalValue(val *AppEngineFlexibleAppVersionReadinessCheck)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	SuccessThreshold() *float64
	SetSuccessThreshold(val *float64)
	SuccessThresholdInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAppStartTimeout()
	ResetCheckInterval()
	ResetFailureThreshold()
	ResetHost()
	ResetSuccessThreshold()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionReadinessCheckOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) AppStartTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appStartTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) AppStartTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appStartTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) CheckInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) CheckIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) InternalValue() *AppEngineFlexibleAppVersionReadinessCheck {
	var returns *AppEngineFlexibleAppVersionReadinessCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) SuccessThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) SuccessThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionReadinessCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionReadinessCheckOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionReadinessCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionReadinessCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionReadinessCheckOutputReference_Override(a AppEngineFlexibleAppVersionReadinessCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionReadinessCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetAppStartTimeout(val *string) {
	if err := j.validateSetAppStartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appStartTimeout",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetCheckInterval(val *string) {
	if err := j.validateSetCheckIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkInterval",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionReadinessCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetSuccessThreshold(val *float64) {
	if err := j.validateSetSuccessThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successThreshold",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ResetAppStartTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetAppStartTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ResetCheckInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		a,
		"resetHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ResetSuccessThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionReadinessCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

