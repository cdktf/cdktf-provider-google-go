// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityclienttlspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/networksecurityclienttlspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityClientTlsPolicyServerValidationCaOutputReference interface {
	cdktf.ComplexObject
	CertificateProviderInstance() NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstanceOutputReference
	CertificateProviderInstanceInput() *NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance
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
	GrpcEndpoint() NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointOutputReference
	GrpcEndpointInput() *NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutCertificateProviderInstance(value *NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance)
	PutGrpcEndpoint(value *NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint)
	ResetCertificateProviderInstance()
	ResetGrpcEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityClientTlsPolicyServerValidationCaOutputReference
type jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) CertificateProviderInstance() NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstanceOutputReference {
	var returns NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstanceOutputReference
	_jsii_.Get(
		j,
		"certificateProviderInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) CertificateProviderInstanceInput() *NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance {
	var returns *NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance
	_jsii_.Get(
		j,
		"certificateProviderInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GrpcEndpoint() NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointOutputReference {
	var returns NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointOutputReference
	_jsii_.Get(
		j,
		"grpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GrpcEndpointInput() *NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint {
	var returns *NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint
	_jsii_.Get(
		j,
		"grpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkSecurityClientTlsPolicyServerValidationCaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkSecurityClientTlsPolicyServerValidationCaOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityClientTlsPolicyServerValidationCaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityClientTlsPolicy.NetworkSecurityClientTlsPolicyServerValidationCaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkSecurityClientTlsPolicyServerValidationCaOutputReference_Override(n NetworkSecurityClientTlsPolicyServerValidationCaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityClientTlsPolicy.NetworkSecurityClientTlsPolicyServerValidationCaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) PutCertificateProviderInstance(value *NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance) {
	if err := n.validatePutCertificateProviderInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCertificateProviderInstance",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) PutGrpcEndpoint(value *NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint) {
	if err := n.validatePutGrpcEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putGrpcEndpoint",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ResetCertificateProviderInstance() {
	_jsii_.InvokeVoid(
		n,
		"resetCertificateProviderInstance",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ResetGrpcEndpoint() {
	_jsii_.InvokeVoid(
		n,
		"resetGrpcEndpoint",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

