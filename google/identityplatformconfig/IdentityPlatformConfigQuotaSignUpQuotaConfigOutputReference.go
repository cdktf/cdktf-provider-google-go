// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/identityplatformconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference interface {
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
	InternalValue() *IdentityPlatformConfigQuotaSignUpQuotaConfig
	SetInternalValue(val *IdentityPlatformConfigQuotaSignUpQuotaConfig)
	Quota() *float64
	SetQuota(val *float64)
	QuotaDuration() *string
	SetQuotaDuration(val *string)
	QuotaDurationInput() *string
	QuotaInput() *float64
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	ResetQuota()
	ResetQuotaDuration()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference
type jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) InternalValue() *IdentityPlatformConfigQuotaSignUpQuotaConfig {
	var returns *IdentityPlatformConfigQuotaSignUpQuotaConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) Quota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) QuotaDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) QuotaDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) QuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIdentityPlatformConfigQuotaSignUpQuotaConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference_Override(i IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetInternalValue(val *IdentityPlatformConfigQuotaSignUpQuotaConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetQuota(val *float64) {
	if err := j.validateSetQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quota",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetQuotaDuration(val *string) {
	if err := j.validateSetQuotaDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaDuration",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ResetQuota() {
	_jsii_.InvokeVoid(
		i,
		"resetQuota",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ResetQuotaDuration() {
	_jsii_.InvokeVoid(
		i,
		"resetQuotaDuration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		i,
		"resetStartTime",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfigQuotaSignUpQuotaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

