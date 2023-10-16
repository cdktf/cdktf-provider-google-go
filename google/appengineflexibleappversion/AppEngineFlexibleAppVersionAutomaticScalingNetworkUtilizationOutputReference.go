// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	SetInternalValue(val *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization)
	TargetReceivedBytesPerSecond() *float64
	SetTargetReceivedBytesPerSecond(val *float64)
	TargetReceivedBytesPerSecondInput() *float64
	TargetReceivedPacketsPerSecond() *float64
	SetTargetReceivedPacketsPerSecond(val *float64)
	TargetReceivedPacketsPerSecondInput() *float64
	TargetSentBytesPerSecond() *float64
	SetTargetSentBytesPerSecond(val *float64)
	TargetSentBytesPerSecondInput() *float64
	TargetSentPacketsPerSecond() *float64
	SetTargetSentPacketsPerSecond(val *float64)
	TargetSentPacketsPerSecondInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetTargetReceivedBytesPerSecond()
	ResetTargetReceivedPacketsPerSecond()
	ResetTargetSentBytesPerSecond()
	ResetTargetSentPacketsPerSecond()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) InternalValue() *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization {
	var returns *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedPacketsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedPacketsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedPacketsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedPacketsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentPacketsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentPacketsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentPacketsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentPacketsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference_Override(a AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetReceivedBytesPerSecond(val *float64) {
	if err := j.validateSetTargetReceivedBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReceivedBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetReceivedPacketsPerSecond(val *float64) {
	if err := j.validateSetTargetReceivedPacketsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReceivedPacketsPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetSentBytesPerSecond(val *float64) {
	if err := j.validateSetTargetSentBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSentBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetSentPacketsPerSecond(val *float64) {
	if err := j.validateSetTargetSentPacketsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSentPacketsPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetReceivedBytesPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetReceivedBytesPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetReceivedPacketsPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetReceivedPacketsPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetSentBytesPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetSentBytesPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetSentPacketsPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetSentPacketsPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

