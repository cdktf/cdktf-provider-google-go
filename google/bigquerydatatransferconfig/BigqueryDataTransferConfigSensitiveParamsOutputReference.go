// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/bigquerydatatransferconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryDataTransferConfigSensitiveParamsOutputReference interface {
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
	InternalValue() *BigqueryDataTransferConfigSensitiveParams
	SetInternalValue(val *BigqueryDataTransferConfigSensitiveParams)
	SecretAccessKey() *string
	SetSecretAccessKey(val *string)
	SecretAccessKeyInput() *string
	SecretAccessKeyWo() *string
	SetSecretAccessKeyWo(val *string)
	SecretAccessKeyWoInput() *string
	SecretAccessKeyWoVersion() *string
	SetSecretAccessKeyWoVersion(val *string)
	SecretAccessKeyWoVersionInput() *string
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
	ResetSecretAccessKey()
	ResetSecretAccessKeyWo()
	ResetSecretAccessKeyWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryDataTransferConfigSensitiveParamsOutputReference
type jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) InternalValue() *BigqueryDataTransferConfigSensitiveParams {
	var returns *BigqueryDataTransferConfigSensitiveParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) SecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) SecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) SecretAccessKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) SecretAccessKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) SecretAccessKeyWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) SecretAccessKeyWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigqueryDataTransferConfigSensitiveParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryDataTransferConfigSensitiveParamsOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryDataTransferConfigSensitiveParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfigSensitiveParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryDataTransferConfigSensitiveParamsOutputReference_Override(b BigqueryDataTransferConfigSensitiveParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfigSensitiveParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetInternalValue(val *BigqueryDataTransferConfigSensitiveParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetSecretAccessKey(val *string) {
	if err := j.validateSetSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKey",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetSecretAccessKeyWo(val *string) {
	if err := j.validateSetSecretAccessKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKeyWo",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetSecretAccessKeyWoVersion(val *string) {
	if err := j.validateSetSecretAccessKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ResetSecretAccessKey() {
	_jsii_.InvokeVoid(
		b,
		"resetSecretAccessKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ResetSecretAccessKeyWo() {
	_jsii_.InvokeVoid(
		b,
		"resetSecretAccessKeyWo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ResetSecretAccessKeyWoVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetSecretAccessKeyWoVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfigSensitiveParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

