// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/certificatemanagercertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateManagerCertificateManagedOutputReference interface {
	cdktf.ComplexObject
	AuthorizationAttemptInfo() CertificateManagerCertificateManagedAuthorizationAttemptInfoList
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
	DnsAuthorizations() *[]*string
	SetDnsAuthorizations(val *[]*string)
	DnsAuthorizationsInput() *[]*string
	Domains() *[]*string
	SetDomains(val *[]*string)
	DomainsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CertificateManagerCertificateManaged
	SetInternalValue(val *CertificateManagerCertificateManaged)
	IssuanceConfig() *string
	SetIssuanceConfig(val *string)
	IssuanceConfigInput() *string
	ProvisioningIssue() CertificateManagerCertificateManagedProvisioningIssueList
	State() *string
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
	ResetDnsAuthorizations()
	ResetDomains()
	ResetIssuanceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CertificateManagerCertificateManagedOutputReference
type jsiiProxy_CertificateManagerCertificateManagedOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) AuthorizationAttemptInfo() CertificateManagerCertificateManagedAuthorizationAttemptInfoList {
	var returns CertificateManagerCertificateManagedAuthorizationAttemptInfoList
	_jsii_.Get(
		j,
		"authorizationAttemptInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) DnsAuthorizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAuthorizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) DnsAuthorizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAuthorizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) DomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) InternalValue() *CertificateManagerCertificateManaged {
	var returns *CertificateManagerCertificateManaged
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) IssuanceConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) IssuanceConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ProvisioningIssue() CertificateManagerCertificateManagedProvisioningIssueList {
	var returns CertificateManagerCertificateManagedProvisioningIssueList
	_jsii_.Get(
		j,
		"provisioningIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCertificateManagerCertificateManagedOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CertificateManagerCertificateManagedOutputReference {
	_init_.Initialize()

	if err := validateNewCertificateManagerCertificateManagedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertificateManagerCertificateManagedOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.certificateManagerCertificate.CertificateManagerCertificateManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCertificateManagerCertificateManagedOutputReference_Override(c CertificateManagerCertificateManagedOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.certificateManagerCertificate.CertificateManagerCertificateManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetDnsAuthorizations(val *[]*string) {
	if err := j.validateSetDnsAuthorizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsAuthorizations",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetDomains(val *[]*string) {
	if err := j.validateSetDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetInternalValue(val *CertificateManagerCertificateManaged) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetIssuanceConfig(val *string) {
	if err := j.validateSetIssuanceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuanceConfig",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CertificateManagerCertificateManagedOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ResetDnsAuthorizations() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsAuthorizations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ResetDomains() {
	_jsii_.InvokeVoid(
		c,
		"resetDomains",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ResetIssuanceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetIssuanceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CertificateManagerCertificateManagedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

