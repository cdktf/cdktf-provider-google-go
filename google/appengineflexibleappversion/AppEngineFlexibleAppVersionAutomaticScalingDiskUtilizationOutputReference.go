// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference interface {
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
	InternalValue() *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	SetInternalValue(val *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization)
	TargetReadBytesPerSecond() *float64
	SetTargetReadBytesPerSecond(val *float64)
	TargetReadBytesPerSecondInput() *float64
	TargetReadOpsPerSecond() *float64
	SetTargetReadOpsPerSecond(val *float64)
	TargetReadOpsPerSecondInput() *float64
	TargetWriteBytesPerSecond() *float64
	SetTargetWriteBytesPerSecond(val *float64)
	TargetWriteBytesPerSecondInput() *float64
	TargetWriteOpsPerSecond() *float64
	SetTargetWriteOpsPerSecond(val *float64)
	TargetWriteOpsPerSecondInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetTargetReadBytesPerSecond()
	ResetTargetReadOpsPerSecond()
	ResetTargetWriteBytesPerSecond()
	ResetTargetWriteOpsPerSecond()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) InternalValue() *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization {
	var returns *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadOpsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadOpsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadOpsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadOpsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteOpsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteOpsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteOpsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteOpsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference_Override(a AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetReadBytesPerSecond(val *float64) {
	if err := j.validateSetTargetReadBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReadBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetReadOpsPerSecond(val *float64) {
	if err := j.validateSetTargetReadOpsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReadOpsPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetWriteBytesPerSecond(val *float64) {
	if err := j.validateSetTargetWriteBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetWriteBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetWriteOpsPerSecond(val *float64) {
	if err := j.validateSetTargetWriteOpsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetWriteOpsPerSecond",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetReadBytesPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetReadBytesPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetReadOpsPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetReadOpsPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetWriteBytesPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetWriteBytesPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetWriteOpsPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetWriteOpsPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

