// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/dataproccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference interface {
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
	CrossRealmTrustAdminServer() *string
	SetCrossRealmTrustAdminServer(val *string)
	CrossRealmTrustAdminServerInput() *string
	CrossRealmTrustKdc() *string
	SetCrossRealmTrustKdc(val *string)
	CrossRealmTrustKdcInput() *string
	CrossRealmTrustRealm() *string
	SetCrossRealmTrustRealm(val *string)
	CrossRealmTrustRealmInput() *string
	CrossRealmTrustSharedPasswordUri() *string
	SetCrossRealmTrustSharedPasswordUri(val *string)
	CrossRealmTrustSharedPasswordUriInput() *string
	EnableKerberos() interface{}
	SetEnableKerberos(val interface{})
	EnableKerberosInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocClusterClusterConfigSecurityConfigKerberosConfig
	SetInternalValue(val *DataprocClusterClusterConfigSecurityConfigKerberosConfig)
	KdcDbKeyUri() *string
	SetKdcDbKeyUri(val *string)
	KdcDbKeyUriInput() *string
	KeyPasswordUri() *string
	SetKeyPasswordUri(val *string)
	KeyPasswordUriInput() *string
	KeystorePasswordUri() *string
	SetKeystorePasswordUri(val *string)
	KeystorePasswordUriInput() *string
	KeystoreUri() *string
	SetKeystoreUri(val *string)
	KeystoreUriInput() *string
	KmsKeyUri() *string
	SetKmsKeyUri(val *string)
	KmsKeyUriInput() *string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	RootPrincipalPasswordUri() *string
	SetRootPrincipalPasswordUri(val *string)
	RootPrincipalPasswordUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TgtLifetimeHours() *float64
	SetTgtLifetimeHours(val *float64)
	TgtLifetimeHoursInput() *float64
	TruststorePasswordUri() *string
	SetTruststorePasswordUri(val *string)
	TruststorePasswordUriInput() *string
	TruststoreUri() *string
	SetTruststoreUri(val *string)
	TruststoreUriInput() *string
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
	ResetCrossRealmTrustAdminServer()
	ResetCrossRealmTrustKdc()
	ResetCrossRealmTrustRealm()
	ResetCrossRealmTrustSharedPasswordUri()
	ResetEnableKerberos()
	ResetKdcDbKeyUri()
	ResetKeyPasswordUri()
	ResetKeystorePasswordUri()
	ResetKeystoreUri()
	ResetRealm()
	ResetTgtLifetimeHours()
	ResetTruststorePasswordUri()
	ResetTruststoreUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) InternalValue() *DataprocClusterClusterConfigSecurityConfigKerberosConfig {
	var returns *DataprocClusterClusterConfigSecurityConfigKerberosConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystoreUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystoreUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststoreUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststoreUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUriInput",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference_Override(d DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustAdminServer(val *string) {
	if err := j.validateSetCrossRealmTrustAdminServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustAdminServer",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustKdc(val *string) {
	if err := j.validateSetCrossRealmTrustKdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustKdc",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustRealm(val *string) {
	if err := j.validateSetCrossRealmTrustRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustRealm",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustSharedPasswordUri(val *string) {
	if err := j.validateSetCrossRealmTrustSharedPasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustSharedPasswordUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetEnableKerberos(val interface{}) {
	if err := j.validateSetEnableKerberosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKerberos",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetInternalValue(val *DataprocClusterClusterConfigSecurityConfigKerberosConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKdcDbKeyUri(val *string) {
	if err := j.validateSetKdcDbKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcDbKeyUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeyPasswordUri(val *string) {
	if err := j.validateSetKeyPasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPasswordUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystorePasswordUri(val *string) {
	if err := j.validateSetKeystorePasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystorePasswordUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystoreUri(val *string) {
	if err := j.validateSetKeystoreUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystoreUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKmsKeyUri(val *string) {
	if err := j.validateSetKmsKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetRootPrincipalPasswordUri(val *string) {
	if err := j.validateSetRootPrincipalPasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPrincipalPasswordUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTgtLifetimeHours(val *float64) {
	if err := j.validateSetTgtLifetimeHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tgtLifetimeHours",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststorePasswordUri(val *string) {
	if err := j.validateSetTruststorePasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststorePasswordUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststoreUri(val *string) {
	if err := j.validateSetTruststoreUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststoreUri",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustAdminServer() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustAdminServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustKdc() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustKdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustRealm() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustRealm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustSharedPasswordUri() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustSharedPasswordUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetEnableKerberos() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableKerberos",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKdcDbKeyUri() {
	_jsii_.InvokeVoid(
		d,
		"resetKdcDbKeyUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeyPasswordUri() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyPasswordUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystorePasswordUri() {
	_jsii_.InvokeVoid(
		d,
		"resetKeystorePasswordUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystoreUri() {
	_jsii_.InvokeVoid(
		d,
		"resetKeystoreUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetRealm() {
	_jsii_.InvokeVoid(
		d,
		"resetRealm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTgtLifetimeHours() {
	_jsii_.InvokeVoid(
		d,
		"resetTgtLifetimeHours",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststorePasswordUri() {
	_jsii_.InvokeVoid(
		d,
		"resetTruststorePasswordUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststoreUri() {
	_jsii_.InvokeVoid(
		d,
		"resetTruststoreUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

