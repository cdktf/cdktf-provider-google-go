// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesedgecacheorigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/networkservicesedgecacheorigin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference interface {
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
	HeaderAction() NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionOutputReference
	HeaderActionInput() *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction
	InternalValue() *NetworkServicesEdgeCacheOriginOriginOverrideAction
	SetInternalValue(val *NetworkServicesEdgeCacheOriginOriginOverrideAction)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRewrite() NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewriteOutputReference
	UrlRewriteInput() *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite
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
	PutHeaderAction(value *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction)
	PutUrlRewrite(value *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite)
	ResetHeaderAction()
	ResetUrlRewrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference
type jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) HeaderAction() NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionOutputReference {
	var returns NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) HeaderActionInput() *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction {
	var returns *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) InternalValue() *NetworkServicesEdgeCacheOriginOriginOverrideAction {
	var returns *NetworkServicesEdgeCacheOriginOriginOverrideAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) UrlRewrite() NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewriteOutputReference {
	var returns NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) UrlRewriteInput() *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite {
	var returns *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}


func NewNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference_Override(n NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetInternalValue(val *NetworkServicesEdgeCacheOriginOriginOverrideAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) PutHeaderAction(value *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction) {
	if err := n.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) PutUrlRewrite(value *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite) {
	if err := n.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		n,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		n,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

