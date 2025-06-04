// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/developerconnectconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerCredential() DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential
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
	InternalValue() *DeveloperConnectConnectionBitbucketDataCenterConfig
	SetInternalValue(val *DeveloperConnectConnectionBitbucketDataCenterConfig)
	ReadAuthorizerCredential() DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference
	ReadAuthorizerCredentialInput() *DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential
	ServerVersion() *string
	ServiceDirectoryConfig() DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig
	SslCaCertificate() *string
	SetSslCaCertificate(val *string)
	SslCaCertificateInput() *string
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
	PutAuthorizerCredential(value *DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential)
	PutReadAuthorizerCredential(value *DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential)
	PutServiceDirectoryConfig(value *DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig)
	ResetServiceDirectoryConfig()
	ResetSslCaCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference
type jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) AuthorizerCredential() DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference {
	var returns DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) AuthorizerCredentialInput() *DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential {
	var returns *DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) InternalValue() *DeveloperConnectConnectionBitbucketDataCenterConfig {
	var returns *DeveloperConnectConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ReadAuthorizerCredential() DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference {
	var returns DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"readAuthorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ReadAuthorizerCredentialInput() *DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential {
	var returns *DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential
	_jsii_.Get(
		j,
		"readAuthorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ServiceDirectoryConfig() DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference {
	var returns DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ServiceDirectoryConfigInput() *DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig {
	var returns *DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDeveloperConnectConnectionBitbucketDataCenterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference_Override(d DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetInternalValue(val *DeveloperConnectConnectionBitbucketDataCenterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) PutAuthorizerCredential(value *DeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential) {
	if err := d.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) PutReadAuthorizerCredential(value *DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential) {
	if err := d.validatePutReadAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadAuthorizerCredential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) PutServiceDirectoryConfig(value *DeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig) {
	if err := d.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

