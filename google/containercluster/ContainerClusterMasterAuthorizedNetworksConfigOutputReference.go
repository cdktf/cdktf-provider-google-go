// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterMasterAuthorizedNetworksConfigOutputReference interface {
	cdktf.ComplexObject
	CidrBlocks() ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksList
	CidrBlocksInput() interface{}
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
	GcpPublicCidrsAccessEnabled() interface{}
	SetGcpPublicCidrsAccessEnabled(val interface{})
	GcpPublicCidrsAccessEnabledInput() interface{}
	InternalValue() *ContainerClusterMasterAuthorizedNetworksConfig
	SetInternalValue(val *ContainerClusterMasterAuthorizedNetworksConfig)
	PrivateEndpointEnforcementEnabled() interface{}
	SetPrivateEndpointEnforcementEnabled(val interface{})
	PrivateEndpointEnforcementEnabledInput() interface{}
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
	PutCidrBlocks(value interface{})
	ResetCidrBlocks()
	ResetGcpPublicCidrsAccessEnabled()
	ResetPrivateEndpointEnforcementEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterMasterAuthorizedNetworksConfigOutputReference
type jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) CidrBlocks() ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksList {
	var returns ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksList
	_jsii_.Get(
		j,
		"cidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) CidrBlocksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GcpPublicCidrsAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpPublicCidrsAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GcpPublicCidrsAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpPublicCidrsAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) InternalValue() *ContainerClusterMasterAuthorizedNetworksConfig {
	var returns *ContainerClusterMasterAuthorizedNetworksConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) PrivateEndpointEnforcementEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointEnforcementEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) PrivateEndpointEnforcementEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointEnforcementEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterMasterAuthorizedNetworksConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterMasterAuthorizedNetworksConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterMasterAuthorizedNetworksConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterMasterAuthorizedNetworksConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterMasterAuthorizedNetworksConfigOutputReference_Override(c ContainerClusterMasterAuthorizedNetworksConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterMasterAuthorizedNetworksConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetGcpPublicCidrsAccessEnabled(val interface{}) {
	if err := j.validateSetGcpPublicCidrsAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpPublicCidrsAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetInternalValue(val *ContainerClusterMasterAuthorizedNetworksConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetPrivateEndpointEnforcementEnabled(val interface{}) {
	if err := j.validateSetPrivateEndpointEnforcementEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointEnforcementEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) PutCidrBlocks(value interface{}) {
	if err := c.validatePutCidrBlocksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCidrBlocks",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ResetCidrBlocks() {
	_jsii_.InvokeVoid(
		c,
		"resetCidrBlocks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ResetGcpPublicCidrsAccessEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetGcpPublicCidrsAccessEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ResetPrivateEndpointEnforcementEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateEndpointEnforcementEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterMasterAuthorizedNetworksConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

