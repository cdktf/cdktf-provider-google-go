// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/developerconnectconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeveloperConnectConnectionGithubConfigOutputReference interface {
	cdktf.ComplexObject
	AppInstallationId() *string
	SetAppInstallationId(val *string)
	AppInstallationIdInput() *string
	AuthorizerCredential() DeveloperConnectConnectionGithubConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *DeveloperConnectConnectionGithubConfigAuthorizerCredential
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
	GithubApp() *string
	SetGithubApp(val *string)
	GithubAppInput() *string
	InstallationUri() *string
	InternalValue() *DeveloperConnectConnectionGithubConfig
	SetInternalValue(val *DeveloperConnectConnectionGithubConfig)
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
	PutAuthorizerCredential(value *DeveloperConnectConnectionGithubConfigAuthorizerCredential)
	ResetAppInstallationId()
	ResetAuthorizerCredential()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeveloperConnectConnectionGithubConfigOutputReference
type jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) AppInstallationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) AppInstallationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) AuthorizerCredential() DeveloperConnectConnectionGithubConfigAuthorizerCredentialOutputReference {
	var returns DeveloperConnectConnectionGithubConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) AuthorizerCredentialInput() *DeveloperConnectConnectionGithubConfigAuthorizerCredential {
	var returns *DeveloperConnectConnectionGithubConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GithubApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GithubAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) InstallationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) InternalValue() *DeveloperConnectConnectionGithubConfig {
	var returns *DeveloperConnectConnectionGithubConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeveloperConnectConnectionGithubConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeveloperConnectConnectionGithubConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDeveloperConnectConnectionGithubConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionGithubConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeveloperConnectConnectionGithubConfigOutputReference_Override(d DeveloperConnectConnectionGithubConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionGithubConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetAppInstallationId(val *string) {
	if err := j.validateSetAppInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appInstallationId",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetGithubApp(val *string) {
	if err := j.validateSetGithubAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubApp",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetInternalValue(val *DeveloperConnectConnectionGithubConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) PutAuthorizerCredential(value *DeveloperConnectConnectionGithubConfigAuthorizerCredential) {
	if err := d.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) ResetAppInstallationId() {
	_jsii_.InvokeVoid(
		d,
		"resetAppInstallationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) ResetAuthorizerCredential() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorizerCredential",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeveloperConnectConnectionGithubConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

