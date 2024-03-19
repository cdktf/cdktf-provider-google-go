// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappactivedirectory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/netappactivedirectory/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/netapp_active_directory google_netapp_active_directory}.
type NetappActiveDirectory interface {
	cdktf.TerraformResource
	AesEncryption() interface{}
	SetAesEncryption(val interface{})
	AesEncryptionInput() interface{}
	BackupOperators() *[]*string
	SetBackupOperators(val *[]*string)
	BackupOperatorsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dns() *string
	SetDns(val *string)
	DnsInput() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptDcConnections() interface{}
	SetEncryptDcConnections(val interface{})
	EncryptDcConnectionsInput() interface{}
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
	KdcHostname() *string
	SetKdcHostname(val *string)
	KdcHostnameInput() *string
	KdcIp() *string
	SetKdcIp(val *string)
	KdcIpInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LdapSigning() interface{}
	SetLdapSigning(val interface{})
	LdapSigningInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetBiosPrefix() *string
	SetNetBiosPrefix(val *string)
	NetBiosPrefixInput() *string
	NfsUsersWithLdap() interface{}
	SetNfsUsersWithLdap(val interface{})
	NfsUsersWithLdapInput() interface{}
	// The tree node.
	Node() constructs.Node
	OrganizationalUnit() *string
	SetOrganizationalUnit(val *string)
	OrganizationalUnitInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	SecurityOperators() *[]*string
	SetSecurityOperators(val *[]*string)
	SecurityOperatorsInput() *[]*string
	Site() *string
	SetSite(val *string)
	SiteInput() *string
	State() *string
	StateDetails() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetappActiveDirectoryTimeoutsOutputReference
	TimeoutsInput() interface{}
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *NetappActiveDirectoryTimeouts)
	ResetAesEncryption()
	ResetBackupOperators()
	ResetDescription()
	ResetEncryptDcConnections()
	ResetId()
	ResetKdcHostname()
	ResetKdcIp()
	ResetLabels()
	ResetLdapSigning()
	ResetNfsUsersWithLdap()
	ResetOrganizationalUnit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSecurityOperators()
	ResetSite()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NetappActiveDirectory
type jsiiProxy_NetappActiveDirectory struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetappActiveDirectory) AesEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aesEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) AesEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aesEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) BackupOperators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupOperators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) BackupOperatorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupOperatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Dns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) DnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) EncryptDcConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptDcConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) EncryptDcConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptDcConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) KdcHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) KdcHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) KdcIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) KdcIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) LdapSigning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) LdapSigningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapSigningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) NetBiosPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netBiosPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) NetBiosPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netBiosPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) NfsUsersWithLdap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsUsersWithLdap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) NfsUsersWithLdapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsUsersWithLdapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) OrganizationalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) OrganizationalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) SecurityOperators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityOperators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) SecurityOperatorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityOperatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) SiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Timeouts() NetappActiveDirectoryTimeoutsOutputReference {
	var returns NetappActiveDirectoryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappActiveDirectory) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/netapp_active_directory google_netapp_active_directory} Resource.
func NewNetappActiveDirectory(scope constructs.Construct, id *string, config *NetappActiveDirectoryConfig) NetappActiveDirectory {
	_init_.Initialize()

	if err := validateNewNetappActiveDirectoryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappActiveDirectory{}

	_jsii_.Create(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/netapp_active_directory google_netapp_active_directory} Resource.
func NewNetappActiveDirectory_Override(n NetappActiveDirectory, scope constructs.Construct, id *string, config *NetappActiveDirectoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetAesEncryption(val interface{}) {
	if err := j.validateSetAesEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aesEncryption",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetBackupOperators(val *[]*string) {
	if err := j.validateSetBackupOperatorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupOperators",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetDns(val *string) {
	if err := j.validateSetDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetEncryptDcConnections(val interface{}) {
	if err := j.validateSetEncryptDcConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptDcConnections",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetKdcHostname(val *string) {
	if err := j.validateSetKdcHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcHostname",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetKdcIp(val *string) {
	if err := j.validateSetKdcIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcIp",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetLdapSigning(val interface{}) {
	if err := j.validateSetLdapSigningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapSigning",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetNetBiosPrefix(val *string) {
	if err := j.validateSetNetBiosPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netBiosPrefix",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetNfsUsersWithLdap(val interface{}) {
	if err := j.validateSetNfsUsersWithLdapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsUsersWithLdap",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetOrganizationalUnit(val *string) {
	if err := j.validateSetOrganizationalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnit",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetSecurityOperators(val *[]*string) {
	if err := j.validateSetSecurityOperatorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityOperators",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetSite(val *string) {
	if err := j.validateSetSiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"site",
		val,
	)
}

func (j *jsiiProxy_NetappActiveDirectory)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a NetappActiveDirectory resource upon running "cdktf plan <stack-name>".
func NetappActiveDirectory_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetappActiveDirectory_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
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
func NetappActiveDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappActiveDirectory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappActiveDirectory_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappActiveDirectory_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappActiveDirectory_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappActiveDirectory_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetappActiveDirectory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.netappActiveDirectory.NetappActiveDirectory",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) PutTimeouts(value *NetappActiveDirectoryTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetAesEncryption() {
	_jsii_.InvokeVoid(
		n,
		"resetAesEncryption",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetBackupOperators() {
	_jsii_.InvokeVoid(
		n,
		"resetBackupOperators",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetEncryptDcConnections() {
	_jsii_.InvokeVoid(
		n,
		"resetEncryptDcConnections",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetKdcHostname() {
	_jsii_.InvokeVoid(
		n,
		"resetKdcHostname",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetKdcIp() {
	_jsii_.InvokeVoid(
		n,
		"resetKdcIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetLdapSigning() {
	_jsii_.InvokeVoid(
		n,
		"resetLdapSigning",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetNfsUsersWithLdap() {
	_jsii_.InvokeVoid(
		n,
		"resetNfsUsersWithLdap",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetOrganizationalUnit() {
	_jsii_.InvokeVoid(
		n,
		"resetOrganizationalUnit",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetSecurityOperators() {
	_jsii_.InvokeVoid(
		n,
		"resetSecurityOperators",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetSite() {
	_jsii_.InvokeVoid(
		n,
		"resetSite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappActiveDirectory) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappActiveDirectory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

