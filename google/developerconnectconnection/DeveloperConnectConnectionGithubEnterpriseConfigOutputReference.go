// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/developerconnectconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeveloperConnectConnectionGithubEnterpriseConfigOutputReference interface {
	cdktf.ComplexObject
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AppInstallationId() *string
	SetAppInstallationId(val *string)
	AppInstallationIdInput() *string
	AppSlug() *string
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
	InstallationUri() *string
	InternalValue() *DeveloperConnectConnectionGithubEnterpriseConfig
	SetInternalValue(val *DeveloperConnectConnectionGithubEnterpriseConfig)
	PrivateKeySecretVersion() *string
	SetPrivateKeySecretVersion(val *string)
	PrivateKeySecretVersionInput() *string
	ServerVersion() *string
	ServiceDirectoryConfig() DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig
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
	PutServiceDirectoryConfig(value *DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig)
	ResetAppId()
	ResetAppInstallationId()
	ResetPrivateKeySecretVersion()
	ResetServiceDirectoryConfig()
	ResetSslCaCertificate()
	ResetWebhookSecretSecretVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeveloperConnectConnectionGithubEnterpriseConfigOutputReference
type jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppInstallationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppInstallationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InstallationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InternalValue() *DeveloperConnectConnectionGithubEnterpriseConfig {
	var returns *DeveloperConnectConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) PrivateKeySecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) PrivateKeySecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ServiceDirectoryConfig() DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference {
	var returns DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ServiceDirectoryConfigInput() *DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	var returns *DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewDeveloperConnectConnectionGithubEnterpriseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeveloperConnectConnectionGithubEnterpriseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDeveloperConnectConnectionGithubEnterpriseConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionGithubEnterpriseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeveloperConnectConnectionGithubEnterpriseConfigOutputReference_Override(d DeveloperConnectConnectionGithubEnterpriseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionGithubEnterpriseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetAppInstallationId(val *string) {
	if err := j.validateSetAppInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appInstallationId",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetInternalValue(val *DeveloperConnectConnectionGithubEnterpriseConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetPrivateKeySecretVersion(val *string) {
	if err := j.validateSetPrivateKeySecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeySecretVersion",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) PutServiceDirectoryConfig(value *DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig) {
	if err := d.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetAppId() {
	_jsii_.InvokeVoid(
		d,
		"resetAppId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetAppInstallationId() {
	_jsii_.InvokeVoid(
		d,
		"resetAppInstallationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetPrivateKeySecretVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateKeySecretVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetWebhookSecretSecretVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookSecretSecretVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

