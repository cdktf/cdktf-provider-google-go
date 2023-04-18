package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference interface {
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
	CrossRealmTrustSharedPassword() *string
	SetCrossRealmTrustSharedPassword(val *string)
	CrossRealmTrustSharedPasswordInput() *string
	EnableKerberos() interface{}
	SetEnableKerberos(val interface{})
	EnableKerberosInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig
	SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig)
	KdcDbKey() *string
	SetKdcDbKey(val *string)
	KdcDbKeyInput() *string
	KeyPassword() *string
	SetKeyPassword(val *string)
	KeyPasswordInput() *string
	Keystore() *string
	SetKeystore(val *string)
	KeystoreInput() *string
	KeystorePassword() *string
	SetKeystorePassword(val *string)
	KeystorePasswordInput() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	RootPrincipalPassword() *string
	SetRootPrincipalPassword(val *string)
	RootPrincipalPasswordInput() *string
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
	Truststore() *string
	SetTruststore(val *string)
	TruststoreInput() *string
	TruststorePassword() *string
	SetTruststorePassword(val *string)
	TruststorePasswordInput() *string
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
	ResetCrossRealmTrustSharedPassword()
	ResetEnableKerberos()
	ResetKdcDbKey()
	ResetKeyPassword()
	ResetKeystore()
	ResetKeystorePassword()
	ResetKmsKey()
	ResetRealm()
	ResetRootPrincipalPassword()
	ResetTgtLifetimeHours()
	ResetTruststore()
	ResetTruststorePassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference
type jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Keystore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeystoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Truststore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TruststoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePasswordInput",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference_Override(d DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustAdminServer(val *string) {
	if err := j.validateSetCrossRealmTrustAdminServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustAdminServer",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustKdc(val *string) {
	if err := j.validateSetCrossRealmTrustKdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustKdc",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustRealm(val *string) {
	if err := j.validateSetCrossRealmTrustRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustRealm",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustSharedPassword(val *string) {
	if err := j.validateSetCrossRealmTrustSharedPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustSharedPassword",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetEnableKerberos(val interface{}) {
	if err := j.validateSetEnableKerberosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKerberos",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKdcDbKey(val *string) {
	if err := j.validateSetKdcDbKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcDbKey",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeyPassword(val *string) {
	if err := j.validateSetKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPassword",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystore(val *string) {
	if err := j.validateSetKeystoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystore",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystorePassword(val *string) {
	if err := j.validateSetKeystorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystorePassword",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetRootPrincipalPassword(val *string) {
	if err := j.validateSetRootPrincipalPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPrincipalPassword",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTgtLifetimeHours(val *float64) {
	if err := j.validateSetTgtLifetimeHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tgtLifetimeHours",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststore(val *string) {
	if err := j.validateSetTruststoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststore",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststorePassword(val *string) {
	if err := j.validateSetTruststorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststorePassword",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustAdminServer() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustAdminServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustKdc() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustKdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustRealm() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustRealm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustSharedPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRealmTrustSharedPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetEnableKerberos() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableKerberos",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKdcDbKey() {
	_jsii_.InvokeVoid(
		d,
		"resetKdcDbKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeyPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystore() {
	_jsii_.InvokeVoid(
		d,
		"resetKeystore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystorePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetKeystorePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetRealm() {
	_jsii_.InvokeVoid(
		d,
		"resetRealm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetRootPrincipalPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetRootPrincipalPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTgtLifetimeHours() {
	_jsii_.InvokeVoid(
		d,
		"resetTgtLifetimeHours",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststore() {
	_jsii_.InvokeVoid(
		d,
		"resetTruststore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststorePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetTruststorePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

