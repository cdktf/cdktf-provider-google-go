// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/cloudtasksqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudTasksQueueHttpTargetUriOverrideOutputReference interface {
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
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *CloudTasksQueueHttpTargetUriOverride
	SetInternalValue(val *CloudTasksQueueHttpTargetUriOverride)
	PathOverride() CloudTasksQueueHttpTargetUriOverridePathOverrideOutputReference
	PathOverrideInput() *CloudTasksQueueHttpTargetUriOverridePathOverride
	Port() *string
	SetPort(val *string)
	PortInput() *string
	QueryOverride() CloudTasksQueueHttpTargetUriOverrideQueryOverrideOutputReference
	QueryOverrideInput() *CloudTasksQueueHttpTargetUriOverrideQueryOverride
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriOverrideEnforceMode() *string
	SetUriOverrideEnforceMode(val *string)
	UriOverrideEnforceModeInput() *string
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
	PutPathOverride(value *CloudTasksQueueHttpTargetUriOverridePathOverride)
	PutQueryOverride(value *CloudTasksQueueHttpTargetUriOverrideQueryOverride)
	ResetHost()
	ResetPathOverride()
	ResetPort()
	ResetQueryOverride()
	ResetScheme()
	ResetUriOverrideEnforceMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudTasksQueueHttpTargetUriOverrideOutputReference
type jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) InternalValue() *CloudTasksQueueHttpTargetUriOverride {
	var returns *CloudTasksQueueHttpTargetUriOverride
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) PathOverride() CloudTasksQueueHttpTargetUriOverridePathOverrideOutputReference {
	var returns CloudTasksQueueHttpTargetUriOverridePathOverrideOutputReference
	_jsii_.Get(
		j,
		"pathOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) PathOverrideInput() *CloudTasksQueueHttpTargetUriOverridePathOverride {
	var returns *CloudTasksQueueHttpTargetUriOverridePathOverride
	_jsii_.Get(
		j,
		"pathOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) QueryOverride() CloudTasksQueueHttpTargetUriOverrideQueryOverrideOutputReference {
	var returns CloudTasksQueueHttpTargetUriOverrideQueryOverrideOutputReference
	_jsii_.Get(
		j,
		"queryOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) QueryOverrideInput() *CloudTasksQueueHttpTargetUriOverrideQueryOverride {
	var returns *CloudTasksQueueHttpTargetUriOverrideQueryOverride
	_jsii_.Get(
		j,
		"queryOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) UriOverrideEnforceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriOverrideEnforceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) UriOverrideEnforceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriOverrideEnforceModeInput",
		&returns,
	)
	return returns
}


func NewCloudTasksQueueHttpTargetUriOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudTasksQueueHttpTargetUriOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewCloudTasksQueueHttpTargetUriOverrideOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudTasksQueue.CloudTasksQueueHttpTargetUriOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudTasksQueueHttpTargetUriOverrideOutputReference_Override(c CloudTasksQueueHttpTargetUriOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudTasksQueue.CloudTasksQueueHttpTargetUriOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetInternalValue(val *CloudTasksQueueHttpTargetUriOverride) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference)SetUriOverrideEnforceMode(val *string) {
	if err := j.validateSetUriOverrideEnforceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriOverrideEnforceMode",
		val,
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) PutPathOverride(value *CloudTasksQueueHttpTargetUriOverridePathOverride) {
	if err := c.validatePutPathOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPathOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) PutQueryOverride(value *CloudTasksQueueHttpTargetUriOverrideQueryOverride) {
	if err := c.validatePutQueryOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		c,
		"resetHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ResetPathOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetPathOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ResetQueryOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ResetScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ResetUriOverrideEnforceMode() {
	_jsii_.InvokeVoid(
		c,
		"resetUriOverrideEnforceMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetUriOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

