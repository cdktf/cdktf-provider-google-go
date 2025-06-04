// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildv2connection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/cloudbuildv2connection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerCredential() Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredential
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
	HostUri() *string
	SetHostUri(val *string)
	HostUriInput() *string
	InternalValue() *Cloudbuildv2ConnectionBitbucketDataCenterConfig
	SetInternalValue(val *Cloudbuildv2ConnectionBitbucketDataCenterConfig)
	ReadAuthorizerCredential() Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference
	ReadAuthorizerCredentialInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredential
	ServerVersion() *string
	ServiceDirectoryConfig() Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig
	SslCa() *string
	SetSslCa(val *string)
	SslCaInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebhookSecretSecretVersion() *string
	SetWebhookSecretSecretVersion(val *string)
	WebhookSecretSecretVersionInput() *string
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
	PutAuthorizerCredential(value *Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredential)
	PutReadAuthorizerCredential(value *Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredential)
	PutServiceDirectoryConfig(value *Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig)
	ResetServiceDirectoryConfig()
	ResetSslCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference
type jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) AuthorizerCredential() Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference {
	var returns Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) AuthorizerCredentialInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredential {
	var returns *Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) InternalValue() *Cloudbuildv2ConnectionBitbucketDataCenterConfig {
	var returns *Cloudbuildv2ConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ReadAuthorizerCredential() Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference {
	var returns Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"readAuthorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ReadAuthorizerCredentialInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredential {
	var returns *Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredential
	_jsii_.Get(
		j,
		"readAuthorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ServiceDirectoryConfig() Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference {
	var returns Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ServiceDirectoryConfigInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig {
	var returns *Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) SslCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) SslCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewCloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildv2ConnectionBitbucketDataCenterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference_Override(c Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetInternalValue(val *Cloudbuildv2ConnectionBitbucketDataCenterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetSslCa(val *string) {
	if err := j.validateSetSslCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCa",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) PutAuthorizerCredential(value *Cloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredential) {
	if err := c.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) PutReadAuthorizerCredential(value *Cloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredential) {
	if err := c.validatePutReadAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadAuthorizerCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) PutServiceDirectoryConfig(value *Cloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig) {
	if err := c.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ResetSslCa() {
	_jsii_.InvokeVoid(
		c,
		"resetSslCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

