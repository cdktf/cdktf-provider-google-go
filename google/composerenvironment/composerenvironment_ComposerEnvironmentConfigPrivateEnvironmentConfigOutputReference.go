package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference interface {
	cdktf.ComplexObject
	CloudComposerConnectionSubnetwork() *string
	SetCloudComposerConnectionSubnetwork(val *string)
	CloudComposerConnectionSubnetworkInput() *string
	CloudComposerNetworkIpv4CidrBlock() *string
	SetCloudComposerNetworkIpv4CidrBlock(val *string)
	CloudComposerNetworkIpv4CidrBlockInput() *string
	CloudSqlIpv4CidrBlock() *string
	SetCloudSqlIpv4CidrBlock(val *string)
	CloudSqlIpv4CidrBlockInput() *string
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
	EnablePrivatelyUsedPublicIps() interface{}
	SetEnablePrivatelyUsedPublicIps(val interface{})
	EnablePrivatelyUsedPublicIpsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComposerEnvironmentConfigPrivateEnvironmentConfig
	SetInternalValue(val *ComposerEnvironmentConfigPrivateEnvironmentConfig)
	MasterIpv4CidrBlock() *string
	SetMasterIpv4CidrBlock(val *string)
	MasterIpv4CidrBlockInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebServerIpv4CidrBlock() *string
	SetWebServerIpv4CidrBlock(val *string)
	WebServerIpv4CidrBlockInput() *string
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
	ResetCloudComposerConnectionSubnetwork()
	ResetCloudComposerNetworkIpv4CidrBlock()
	ResetCloudSqlIpv4CidrBlock()
	ResetEnablePrivateEndpoint()
	ResetEnablePrivatelyUsedPublicIps()
	ResetMasterIpv4CidrBlock()
	ResetWebServerIpv4CidrBlock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference
type jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerConnectionSubnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerConnectionSubnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerConnectionSubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerConnectionSubnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerNetworkIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerNetworkIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerNetworkIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerNetworkIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudSqlIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudSqlIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivateEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivatelyUsedPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivatelyUsedPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivatelyUsedPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivatelyUsedPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) InternalValue() *ComposerEnvironmentConfigPrivateEnvironmentConfig {
	var returns *ComposerEnvironmentConfigPrivateEnvironmentConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) MasterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) MasterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) WebServerIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webServerIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) WebServerIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webServerIpv4CidrBlockInput",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigPrivateEnvironmentConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference_Override(c ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetCloudComposerConnectionSubnetwork(val *string) {
	if err := j.validateSetCloudComposerConnectionSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudComposerConnectionSubnetwork",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetCloudComposerNetworkIpv4CidrBlock(val *string) {
	if err := j.validateSetCloudComposerNetworkIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudComposerNetworkIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetCloudSqlIpv4CidrBlock(val *string) {
	if err := j.validateSetCloudSqlIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudSqlIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetEnablePrivateEndpoint(val interface{}) {
	if err := j.validateSetEnablePrivateEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetEnablePrivatelyUsedPublicIps(val interface{}) {
	if err := j.validateSetEnablePrivatelyUsedPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivatelyUsedPublicIps",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetInternalValue(val *ComposerEnvironmentConfigPrivateEnvironmentConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetMasterIpv4CidrBlock(val *string) {
	if err := j.validateSetMasterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetWebServerIpv4CidrBlock(val *string) {
	if err := j.validateSetWebServerIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webServerIpv4CidrBlock",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetCloudComposerConnectionSubnetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudComposerConnectionSubnetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetCloudComposerNetworkIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudComposerNetworkIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetCloudSqlIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudSqlIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetEnablePrivateEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePrivateEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetEnablePrivatelyUsedPublicIps() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePrivatelyUsedPublicIps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetMasterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetMasterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetWebServerIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetWebServerIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

