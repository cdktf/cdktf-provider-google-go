package containerazurecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v3/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v3/containerazurecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAzureClusterControlPlaneOutputReference interface {
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
	DatabaseEncryption() ContainerAzureClusterControlPlaneDatabaseEncryptionOutputReference
	DatabaseEncryptionInput() *ContainerAzureClusterControlPlaneDatabaseEncryption
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerAzureClusterControlPlane
	SetInternalValue(val *ContainerAzureClusterControlPlane)
	MainVolume() ContainerAzureClusterControlPlaneMainVolumeOutputReference
	MainVolumeInput() *ContainerAzureClusterControlPlaneMainVolume
	ProxyConfig() ContainerAzureClusterControlPlaneProxyConfigOutputReference
	ProxyConfigInput() *ContainerAzureClusterControlPlaneProxyConfig
	ReplicaPlacements() ContainerAzureClusterControlPlaneReplicaPlacementsList
	ReplicaPlacementsInput() interface{}
	RootVolume() ContainerAzureClusterControlPlaneRootVolumeOutputReference
	RootVolumeInput() *ContainerAzureClusterControlPlaneRootVolume
	SshConfig() ContainerAzureClusterControlPlaneSshConfigOutputReference
	SshConfigInput() *ContainerAzureClusterControlPlaneSshConfig
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutDatabaseEncryption(value *ContainerAzureClusterControlPlaneDatabaseEncryption)
	PutMainVolume(value *ContainerAzureClusterControlPlaneMainVolume)
	PutProxyConfig(value *ContainerAzureClusterControlPlaneProxyConfig)
	PutReplicaPlacements(value interface{})
	PutRootVolume(value *ContainerAzureClusterControlPlaneRootVolume)
	PutSshConfig(value *ContainerAzureClusterControlPlaneSshConfig)
	ResetDatabaseEncryption()
	ResetMainVolume()
	ResetProxyConfig()
	ResetReplicaPlacements()
	ResetRootVolume()
	ResetTags()
	ResetVmSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerAzureClusterControlPlaneOutputReference
type jsiiProxy_ContainerAzureClusterControlPlaneOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) DatabaseEncryption() ContainerAzureClusterControlPlaneDatabaseEncryptionOutputReference {
	var returns ContainerAzureClusterControlPlaneDatabaseEncryptionOutputReference
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) DatabaseEncryptionInput() *ContainerAzureClusterControlPlaneDatabaseEncryption {
	var returns *ContainerAzureClusterControlPlaneDatabaseEncryption
	_jsii_.Get(
		j,
		"databaseEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) InternalValue() *ContainerAzureClusterControlPlane {
	var returns *ContainerAzureClusterControlPlane
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) MainVolume() ContainerAzureClusterControlPlaneMainVolumeOutputReference {
	var returns ContainerAzureClusterControlPlaneMainVolumeOutputReference
	_jsii_.Get(
		j,
		"mainVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) MainVolumeInput() *ContainerAzureClusterControlPlaneMainVolume {
	var returns *ContainerAzureClusterControlPlaneMainVolume
	_jsii_.Get(
		j,
		"mainVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ProxyConfig() ContainerAzureClusterControlPlaneProxyConfigOutputReference {
	var returns ContainerAzureClusterControlPlaneProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ProxyConfigInput() *ContainerAzureClusterControlPlaneProxyConfig {
	var returns *ContainerAzureClusterControlPlaneProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ReplicaPlacements() ContainerAzureClusterControlPlaneReplicaPlacementsList {
	var returns ContainerAzureClusterControlPlaneReplicaPlacementsList
	_jsii_.Get(
		j,
		"replicaPlacements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ReplicaPlacementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaPlacementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) RootVolume() ContainerAzureClusterControlPlaneRootVolumeOutputReference {
	var returns ContainerAzureClusterControlPlaneRootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) RootVolumeInput() *ContainerAzureClusterControlPlaneRootVolume {
	var returns *ContainerAzureClusterControlPlaneRootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) SshConfig() ContainerAzureClusterControlPlaneSshConfigOutputReference {
	var returns ContainerAzureClusterControlPlaneSshConfigOutputReference
	_jsii_.Get(
		j,
		"sshConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) SshConfigInput() *ContainerAzureClusterControlPlaneSshConfig {
	var returns *ContainerAzureClusterControlPlaneSshConfig
	_jsii_.Get(
		j,
		"sshConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewContainerAzureClusterControlPlaneOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerAzureClusterControlPlaneOutputReference {
	_init_.Initialize()

	if err := validateNewContainerAzureClusterControlPlaneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAzureClusterControlPlaneOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerAzureCluster.ContainerAzureClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerAzureClusterControlPlaneOutputReference_Override(c ContainerAzureClusterControlPlaneOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerAzureCluster.ContainerAzureClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetInternalValue(val *ContainerAzureClusterControlPlane) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) PutDatabaseEncryption(value *ContainerAzureClusterControlPlaneDatabaseEncryption) {
	if err := c.validatePutDatabaseEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDatabaseEncryption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) PutMainVolume(value *ContainerAzureClusterControlPlaneMainVolume) {
	if err := c.validatePutMainVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMainVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) PutProxyConfig(value *ContainerAzureClusterControlPlaneProxyConfig) {
	if err := c.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) PutReplicaPlacements(value interface{}) {
	if err := c.validatePutReplicaPlacementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReplicaPlacements",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) PutRootVolume(value *ContainerAzureClusterControlPlaneRootVolume) {
	if err := c.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) PutSshConfig(value *ContainerAzureClusterControlPlaneSshConfig) {
	if err := c.validatePutSshConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSshConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetDatabaseEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetDatabaseEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetMainVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetMainVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetReplicaPlacements() {
	_jsii_.InvokeVoid(
		c,
		"resetReplicaPlacements",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetRootVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetRootVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ResetVmSize() {
	_jsii_.InvokeVoid(
		c,
		"resetVmSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerAzureClusterControlPlaneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

