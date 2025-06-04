// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityservertlspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/networksecurityservertlspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference interface {
	cdktf.ComplexObject
	ClientValidationCa() NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaList
	ClientValidationCaInput() interface{}
	ClientValidationMode() *string
	SetClientValidationMode(val *string)
	ClientValidationModeInput() *string
	ClientValidationTrustConfig() *string
	SetClientValidationTrustConfig(val *string)
	ClientValidationTrustConfigInput() *string
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
	InternalValue() *NetworkSecurityServerTlsPolicyMtlsPolicy
	SetInternalValue(val *NetworkSecurityServerTlsPolicyMtlsPolicy)
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
	PutClientValidationCa(value interface{})
	ResetClientValidationCa()
	ResetClientValidationMode()
	ResetClientValidationTrustConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference
type jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ClientValidationCa() NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaList {
	var returns NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaList
	_jsii_.Get(
		j,
		"clientValidationCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ClientValidationCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientValidationCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ClientValidationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientValidationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ClientValidationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientValidationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ClientValidationTrustConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientValidationTrustConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ClientValidationTrustConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientValidationTrustConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) InternalValue() *NetworkSecurityServerTlsPolicyMtlsPolicy {
	var returns *NetworkSecurityServerTlsPolicyMtlsPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkSecurityServerTlsPolicyMtlsPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityServerTlsPolicyMtlsPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityServerTlsPolicy.NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkSecurityServerTlsPolicyMtlsPolicyOutputReference_Override(n NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityServerTlsPolicy.NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetClientValidationMode(val *string) {
	if err := j.validateSetClientValidationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientValidationMode",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetClientValidationTrustConfig(val *string) {
	if err := j.validateSetClientValidationTrustConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientValidationTrustConfig",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetInternalValue(val *NetworkSecurityServerTlsPolicyMtlsPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) PutClientValidationCa(value interface{}) {
	if err := n.validatePutClientValidationCaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putClientValidationCa",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ResetClientValidationCa() {
	_jsii_.InvokeVoid(
		n,
		"resetClientValidationCa",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ResetClientValidationMode() {
	_jsii_.InvokeVoid(
		n,
		"resetClientValidationMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ResetClientValidationTrustConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetClientValidationTrustConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyMtlsPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

