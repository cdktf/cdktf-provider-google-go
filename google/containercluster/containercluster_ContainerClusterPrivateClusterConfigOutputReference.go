package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v3/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v3/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterPrivateClusterConfigOutputReference interface {
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
	EnablePrivateEndpoint() interface{}
	SetEnablePrivateEndpoint(val interface{})
	EnablePrivateEndpointInput() interface{}
	EnablePrivateNodes() interface{}
	SetEnablePrivateNodes(val interface{})
	EnablePrivateNodesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerClusterPrivateClusterConfig
	SetInternalValue(val *ContainerClusterPrivateClusterConfig)
	MasterGlobalAccessConfig() ContainerClusterPrivateClusterConfigMasterGlobalAccessConfigOutputReference
	MasterGlobalAccessConfigInput() *ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig
	MasterIpv4CidrBlock() *string
	SetMasterIpv4CidrBlock(val *string)
	MasterIpv4CidrBlockInput() *string
	PeeringName() *string
	PrivateEndpoint() *string
	PublicEndpoint() *string
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
	PutMasterGlobalAccessConfig(value *ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig)
	ResetEnablePrivateNodes()
	ResetMasterGlobalAccessConfig()
	ResetMasterIpv4CidrBlock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterPrivateClusterConfigOutputReference
type jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) EnablePrivateEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) EnablePrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) EnablePrivateNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) EnablePrivateNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) InternalValue() *ContainerClusterPrivateClusterConfig {
	var returns *ContainerClusterPrivateClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) MasterGlobalAccessConfig() ContainerClusterPrivateClusterConfigMasterGlobalAccessConfigOutputReference {
	var returns ContainerClusterPrivateClusterConfigMasterGlobalAccessConfigOutputReference
	_jsii_.Get(
		j,
		"masterGlobalAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) MasterGlobalAccessConfigInput() *ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig {
	var returns *ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig
	_jsii_.Get(
		j,
		"masterGlobalAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) MasterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) MasterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) PeeringName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) PublicEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterPrivateClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterPrivateClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterPrivateClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterPrivateClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterPrivateClusterConfigOutputReference_Override(c ContainerClusterPrivateClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterPrivateClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetEnablePrivateEndpoint(val interface{}) {
	if err := j.validateSetEnablePrivateEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetEnablePrivateNodes(val interface{}) {
	if err := j.validateSetEnablePrivateNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateNodes",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetInternalValue(val *ContainerClusterPrivateClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetMasterIpv4CidrBlock(val *string) {
	if err := j.validateSetMasterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) PutMasterGlobalAccessConfig(value *ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig) {
	if err := c.validatePutMasterGlobalAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMasterGlobalAccessConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ResetEnablePrivateNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePrivateNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ResetMasterGlobalAccessConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetMasterGlobalAccessConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ResetMasterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetMasterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterPrivateClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

