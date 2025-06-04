// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildv2connection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/cloudbuildv2connection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudbuildv2ConnectionGitlabConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerCredential() Cloudbuildv2ConnectionGitlabConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential
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
	InternalValue() *Cloudbuildv2ConnectionGitlabConfig
	SetInternalValue(val *Cloudbuildv2ConnectionGitlabConfig)
	ReadAuthorizerCredential() Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialOutputReference
	ReadAuthorizerCredentialInput() *Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential
	ServerVersion() *string
	ServiceDirectoryConfig() Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig
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
	PutAuthorizerCredential(value *Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential)
	PutReadAuthorizerCredential(value *Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential)
	PutServiceDirectoryConfig(value *Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig)
	ResetHostUri()
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

// The jsii proxy struct for Cloudbuildv2ConnectionGitlabConfigOutputReference
type jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) AuthorizerCredential() Cloudbuildv2ConnectionGitlabConfigAuthorizerCredentialOutputReference {
	var returns Cloudbuildv2ConnectionGitlabConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) AuthorizerCredentialInput() *Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential {
	var returns *Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) InternalValue() *Cloudbuildv2ConnectionGitlabConfig {
	var returns *Cloudbuildv2ConnectionGitlabConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ReadAuthorizerCredential() Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialOutputReference {
	var returns Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"readAuthorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ReadAuthorizerCredentialInput() *Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential {
	var returns *Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential
	_jsii_.Get(
		j,
		"readAuthorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ServiceDirectoryConfig() Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigOutputReference {
	var returns Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ServiceDirectoryConfigInput() *Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig {
	var returns *Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) SslCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) SslCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewCloudbuildv2ConnectionGitlabConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Cloudbuildv2ConnectionGitlabConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildv2ConnectionGitlabConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2ConnectionGitlabConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildv2ConnectionGitlabConfigOutputReference_Override(c Cloudbuildv2ConnectionGitlabConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2ConnectionGitlabConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetInternalValue(val *Cloudbuildv2ConnectionGitlabConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetSslCa(val *string) {
	if err := j.validateSetSslCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCa",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) PutAuthorizerCredential(value *Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential) {
	if err := c.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) PutReadAuthorizerCredential(value *Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential) {
	if err := c.validatePutReadAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadAuthorizerCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) PutServiceDirectoryConfig(value *Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig) {
	if err := c.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ResetHostUri() {
	_jsii_.InvokeVoid(
		c,
		"resetHostUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ResetSslCa() {
	_jsii_.InvokeVoid(
		c,
		"resetSslCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_Cloudbuildv2ConnectionGitlabConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

