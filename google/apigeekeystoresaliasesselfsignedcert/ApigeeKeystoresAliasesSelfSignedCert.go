// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeekeystoresaliasesselfsignedcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v12/apigeekeystoresaliasesselfsignedcert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/apigee_keystores_aliases_self_signed_cert google_apigee_keystores_aliases_self_signed_cert}.
type ApigeeKeystoresAliasesSelfSignedCert interface {
	cdktf.TerraformResource
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertsInfo() ApigeeKeystoresAliasesSelfSignedCertCertsInfoList
	CertValidityInDays() *float64
	SetCertValidityInDays(val *float64)
	CertValidityInDaysInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KeySize() *string
	SetKeySize(val *string)
	KeySizeInput() *string
	Keystore() *string
	SetKeystore(val *string)
	KeystoreInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OrgId() *string
	SetOrgId(val *string)
	OrgIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SigAlg() *string
	SetSigAlg(val *string)
	SigAlgInput() *string
	Subject() ApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference
	SubjectAlternativeDnsNames() ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference
	SubjectAlternativeDnsNamesInput() *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames
	SubjectInput() *ApigeeKeystoresAliasesSelfSignedCertSubject
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApigeeKeystoresAliasesSelfSignedCertTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutSubject(value *ApigeeKeystoresAliasesSelfSignedCertSubject)
	PutSubjectAlternativeDnsNames(value *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames)
	PutTimeouts(value *ApigeeKeystoresAliasesSelfSignedCertTimeouts)
	ResetCertValidityInDays()
	ResetId()
	ResetKeySize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubjectAlternativeDnsNames()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApigeeKeystoresAliasesSelfSignedCert
type jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) CertsInfo() ApigeeKeystoresAliasesSelfSignedCertCertsInfoList {
	var returns ApigeeKeystoresAliasesSelfSignedCertCertsInfoList
	_jsii_.Get(
		j,
		"certsInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) CertValidityInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certValidityInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) CertValidityInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certValidityInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) KeySize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) KeySizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Keystore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) KeystoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) SigAlg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) SigAlgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Subject() ApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference {
	var returns ApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) SubjectAlternativeDnsNames() ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference {
	var returns ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference
	_jsii_.Get(
		j,
		"subjectAlternativeDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) SubjectAlternativeDnsNamesInput() *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames {
	var returns *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames
	_jsii_.Get(
		j,
		"subjectAlternativeDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) SubjectInput() *ApigeeKeystoresAliasesSelfSignedCertSubject {
	var returns *ApigeeKeystoresAliasesSelfSignedCertSubject
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Timeouts() ApigeeKeystoresAliasesSelfSignedCertTimeoutsOutputReference {
	var returns ApigeeKeystoresAliasesSelfSignedCertTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/apigee_keystores_aliases_self_signed_cert google_apigee_keystores_aliases_self_signed_cert} Resource.
func NewApigeeKeystoresAliasesSelfSignedCert(scope constructs.Construct, id *string, config *ApigeeKeystoresAliasesSelfSignedCertConfig) ApigeeKeystoresAliasesSelfSignedCert {
	_init_.Initialize()

	if err := validateNewApigeeKeystoresAliasesSelfSignedCertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/apigee_keystores_aliases_self_signed_cert google_apigee_keystores_aliases_self_signed_cert} Resource.
func NewApigeeKeystoresAliasesSelfSignedCert_Override(a ApigeeKeystoresAliasesSelfSignedCert, scope constructs.Construct, id *string, config *ApigeeKeystoresAliasesSelfSignedCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetCertValidityInDays(val *float64) {
	if err := j.validateSetCertValidityInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certValidityInDays",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetKeySize(val *string) {
	if err := j.validateSetKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySize",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetKeystore(val *string) {
	if err := j.validateSetKeystoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystore",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert)SetSigAlg(val *string) {
	if err := j.validateSetSigAlgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigAlg",
		val,
	)
}

// Generates CDKTF code for importing a ApigeeKeystoresAliasesSelfSignedCert resource upon running "cdktf plan <stack-name>".
func ApigeeKeystoresAliasesSelfSignedCert_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigeeKeystoresAliasesSelfSignedCert_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApigeeKeystoresAliasesSelfSignedCert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeKeystoresAliasesSelfSignedCert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeKeystoresAliasesSelfSignedCert_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeKeystoresAliasesSelfSignedCert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeKeystoresAliasesSelfSignedCert_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeKeystoresAliasesSelfSignedCert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigeeKeystoresAliasesSelfSignedCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) PutSubject(value *ApigeeKeystoresAliasesSelfSignedCertSubject) {
	if err := a.validatePutSubjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubject",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) PutSubjectAlternativeDnsNames(value *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames) {
	if err := a.validatePutSubjectAlternativeDnsNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubjectAlternativeDnsNames",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) PutTimeouts(value *ApigeeKeystoresAliasesSelfSignedCertTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ResetCertValidityInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetCertValidityInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ResetKeySize() {
	_jsii_.InvokeVoid(
		a,
		"resetKeySize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ResetSubjectAlternativeDnsNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeDnsNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

