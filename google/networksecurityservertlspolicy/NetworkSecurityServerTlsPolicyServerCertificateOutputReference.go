// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityservertlspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/networksecurityservertlspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityServerTlsPolicyServerCertificateOutputReference interface {
	cdktf.ComplexObject
	CertificateProviderInstance() NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstanceOutputReference
	CertificateProviderInstanceInput() *NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance
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
	GrpcEndpoint() NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointOutputReference
	GrpcEndpointInput() *NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint
	InternalValue() *NetworkSecurityServerTlsPolicyServerCertificate
	SetInternalValue(val *NetworkSecurityServerTlsPolicyServerCertificate)
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
	PutCertificateProviderInstance(value *NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance)
	PutGrpcEndpoint(value *NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint)
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

// The jsii proxy struct for NetworkSecurityServerTlsPolicyServerCertificateOutputReference
type jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) CertificateProviderInstance() NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstanceOutputReference {
	var returns NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstanceOutputReference
	_jsii_.Get(
		j,
		"certificateProviderInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) CertificateProviderInstanceInput() *NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance {
	var returns *NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance
	_jsii_.Get(
		j,
		"certificateProviderInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GrpcEndpoint() NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointOutputReference {
	var returns NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointOutputReference
	_jsii_.Get(
		j,
		"grpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GrpcEndpointInput() *NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint {
	var returns *NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint
	_jsii_.Get(
		j,
		"grpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) InternalValue() *NetworkSecurityServerTlsPolicyServerCertificate {
	var returns *NetworkSecurityServerTlsPolicyServerCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkSecurityServerTlsPolicyServerCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkSecurityServerTlsPolicyServerCertificateOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityServerTlsPolicyServerCertificateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityServerTlsPolicy.NetworkSecurityServerTlsPolicyServerCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkSecurityServerTlsPolicyServerCertificateOutputReference_Override(n NetworkSecurityServerTlsPolicyServerCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityServerTlsPolicy.NetworkSecurityServerTlsPolicyServerCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetInternalValue(val *NetworkSecurityServerTlsPolicyServerCertificate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) PutCertificateProviderInstance(value *NetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance) {
	if err := n.validatePutCertificateProviderInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCertificateProviderInstance",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) PutGrpcEndpoint(value *NetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint) {
	if err := n.validatePutGrpcEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putGrpcEndpoint",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) ResetCertificateProviderInstance() {
	_jsii_.InvokeVoid(
		n,
		"resetCertificateProviderInstance",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) ResetGrpcEndpoint() {
	_jsii_.InvokeVoid(
		n,
		"resetGrpcEndpoint",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkSecurityServerTlsPolicyServerCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

