// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/integrationsclient/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationsClientCloudKmsConfigOutputReference interface {
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
	InternalValue() *IntegrationsClientCloudKmsConfig
	SetInternalValue(val *IntegrationsClientCloudKmsConfig)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	KeyVersion() *string
	SetKeyVersion(val *string)
	KeyVersionInput() *string
	KmsLocation() *string
	SetKmsLocation(val *string)
	KmsLocationInput() *string
	KmsProjectId() *string
	SetKmsProjectId(val *string)
	KmsProjectIdInput() *string
	KmsRing() *string
	SetKmsRing(val *string)
	KmsRingInput() *string
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
	ResetKeyVersion()
	ResetKmsProjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationsClientCloudKmsConfigOutputReference
type jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) InternalValue() *IntegrationsClientCloudKmsConfig {
	var returns *IntegrationsClientCloudKmsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KeyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KeyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KmsLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KmsLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KmsProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KmsProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KmsRing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsRing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) KmsRingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsRingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationsClientCloudKmsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationsClientCloudKmsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationsClientCloudKmsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationsClient.IntegrationsClientCloudKmsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationsClientCloudKmsConfigOutputReference_Override(i IntegrationsClientCloudKmsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationsClient.IntegrationsClientCloudKmsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetInternalValue(val *IntegrationsClientCloudKmsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetKeyVersion(val *string) {
	if err := j.validateSetKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVersion",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetKmsLocation(val *string) {
	if err := j.validateSetKmsLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsLocation",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetKmsProjectId(val *string) {
	if err := j.validateSetKmsProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsProjectId",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetKmsRing(val *string) {
	if err := j.validateSetKmsRingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsRing",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) ResetKeyVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetKeyVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) ResetKmsProjectId() {
	_jsii_.InvokeVoid(
		i,
		"resetKmsProjectId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationsClientCloudKmsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

