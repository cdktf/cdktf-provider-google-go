// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/datagooglecomposerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComposerEnvironmentConfigAOutputReference interface {
	cdktf.ComplexObject
	AirflowUri() *string
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
	DagGcsPrefix() *string
	DatabaseConfig() DataGoogleComposerEnvironmentConfigDatabaseConfigList
	DataRetentionConfig() DataGoogleComposerEnvironmentConfigDataRetentionConfigList
	EncryptionConfig() DataGoogleComposerEnvironmentConfigEncryptionConfigList
	EnvironmentSize() *string
	// Experimental.
	Fqn() *string
	GkeCluster() *string
	InternalValue() *DataGoogleComposerEnvironmentConfigA
	SetInternalValue(val *DataGoogleComposerEnvironmentConfigA)
	MaintenanceWindow() DataGoogleComposerEnvironmentConfigMaintenanceWindowList
	MasterAuthorizedNetworksConfig() DataGoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigList
	NodeConfig() DataGoogleComposerEnvironmentConfigNodeConfigList
	NodeCount() *float64
	PrivateEnvironmentConfig() DataGoogleComposerEnvironmentConfigPrivateEnvironmentConfigList
	RecoveryConfig() DataGoogleComposerEnvironmentConfigRecoveryConfigList
	ResilienceMode() *string
	SoftwareConfig() DataGoogleComposerEnvironmentConfigSoftwareConfigList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebServerConfig() DataGoogleComposerEnvironmentConfigWebServerConfigList
	WebServerNetworkAccessControl() DataGoogleComposerEnvironmentConfigWebServerNetworkAccessControlList
	WorkloadsConfig() DataGoogleComposerEnvironmentConfigWorkloadsConfigList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComposerEnvironmentConfigAOutputReference
type jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) AirflowUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) DagGcsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagGcsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) DatabaseConfig() DataGoogleComposerEnvironmentConfigDatabaseConfigList {
	var returns DataGoogleComposerEnvironmentConfigDatabaseConfigList
	_jsii_.Get(
		j,
		"databaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) DataRetentionConfig() DataGoogleComposerEnvironmentConfigDataRetentionConfigList {
	var returns DataGoogleComposerEnvironmentConfigDataRetentionConfigList
	_jsii_.Get(
		j,
		"dataRetentionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) EncryptionConfig() DataGoogleComposerEnvironmentConfigEncryptionConfigList {
	var returns DataGoogleComposerEnvironmentConfigEncryptionConfigList
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) EnvironmentSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GkeCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) InternalValue() *DataGoogleComposerEnvironmentConfigA {
	var returns *DataGoogleComposerEnvironmentConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) MaintenanceWindow() DataGoogleComposerEnvironmentConfigMaintenanceWindowList {
	var returns DataGoogleComposerEnvironmentConfigMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) MasterAuthorizedNetworksConfig() DataGoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigList {
	var returns DataGoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigList
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) NodeConfig() DataGoogleComposerEnvironmentConfigNodeConfigList {
	var returns DataGoogleComposerEnvironmentConfigNodeConfigList
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) PrivateEnvironmentConfig() DataGoogleComposerEnvironmentConfigPrivateEnvironmentConfigList {
	var returns DataGoogleComposerEnvironmentConfigPrivateEnvironmentConfigList
	_jsii_.Get(
		j,
		"privateEnvironmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) RecoveryConfig() DataGoogleComposerEnvironmentConfigRecoveryConfigList {
	var returns DataGoogleComposerEnvironmentConfigRecoveryConfigList
	_jsii_.Get(
		j,
		"recoveryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) ResilienceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resilienceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) SoftwareConfig() DataGoogleComposerEnvironmentConfigSoftwareConfigList {
	var returns DataGoogleComposerEnvironmentConfigSoftwareConfigList
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) WebServerConfig() DataGoogleComposerEnvironmentConfigWebServerConfigList {
	var returns DataGoogleComposerEnvironmentConfigWebServerConfigList
	_jsii_.Get(
		j,
		"webServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) WebServerNetworkAccessControl() DataGoogleComposerEnvironmentConfigWebServerNetworkAccessControlList {
	var returns DataGoogleComposerEnvironmentConfigWebServerNetworkAccessControlList
	_jsii_.Get(
		j,
		"webServerNetworkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) WorkloadsConfig() DataGoogleComposerEnvironmentConfigWorkloadsConfigList {
	var returns DataGoogleComposerEnvironmentConfigWorkloadsConfigList
	_jsii_.Get(
		j,
		"workloadsConfig",
		&returns,
	)
	return returns
}


func NewDataGoogleComposerEnvironmentConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComposerEnvironmentConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComposerEnvironmentConfigAOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComposerEnvironment.DataGoogleComposerEnvironmentConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComposerEnvironmentConfigAOutputReference_Override(d DataGoogleComposerEnvironmentConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComposerEnvironment.DataGoogleComposerEnvironmentConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference)SetInternalValue(val *DataGoogleComposerEnvironmentConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

