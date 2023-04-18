package containerawscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/containerawscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAwsClusterControlPlaneOutputReference interface {
	cdktf.ComplexObject
	AwsServicesAuthentication() ContainerAwsClusterControlPlaneAwsServicesAuthenticationOutputReference
	AwsServicesAuthenticationInput() *ContainerAwsClusterControlPlaneAwsServicesAuthentication
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
	ConfigEncryption() ContainerAwsClusterControlPlaneConfigEncryptionOutputReference
	ConfigEncryptionInput() *ContainerAwsClusterControlPlaneConfigEncryption
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseEncryption() ContainerAwsClusterControlPlaneDatabaseEncryptionOutputReference
	DatabaseEncryptionInput() *ContainerAwsClusterControlPlaneDatabaseEncryption
	// Experimental.
	Fqn() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *ContainerAwsClusterControlPlane
	SetInternalValue(val *ContainerAwsClusterControlPlane)
	MainVolume() ContainerAwsClusterControlPlaneMainVolumeOutputReference
	MainVolumeInput() *ContainerAwsClusterControlPlaneMainVolume
	ProxyConfig() ContainerAwsClusterControlPlaneProxyConfigOutputReference
	ProxyConfigInput() *ContainerAwsClusterControlPlaneProxyConfig
	RootVolume() ContainerAwsClusterControlPlaneRootVolumeOutputReference
	RootVolumeInput() *ContainerAwsClusterControlPlaneRootVolume
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SshConfig() ContainerAwsClusterControlPlaneSshConfigOutputReference
	SshConfigInput() *ContainerAwsClusterControlPlaneSshConfig
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	PutAwsServicesAuthentication(value *ContainerAwsClusterControlPlaneAwsServicesAuthentication)
	PutConfigEncryption(value *ContainerAwsClusterControlPlaneConfigEncryption)
	PutDatabaseEncryption(value *ContainerAwsClusterControlPlaneDatabaseEncryption)
	PutMainVolume(value *ContainerAwsClusterControlPlaneMainVolume)
	PutProxyConfig(value *ContainerAwsClusterControlPlaneProxyConfig)
	PutRootVolume(value *ContainerAwsClusterControlPlaneRootVolume)
	PutSshConfig(value *ContainerAwsClusterControlPlaneSshConfig)
	ResetInstanceType()
	ResetMainVolume()
	ResetProxyConfig()
	ResetRootVolume()
	ResetSecurityGroupIds()
	ResetSshConfig()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerAwsClusterControlPlaneOutputReference
type jsiiProxy_ContainerAwsClusterControlPlaneOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) AwsServicesAuthentication() ContainerAwsClusterControlPlaneAwsServicesAuthenticationOutputReference {
	var returns ContainerAwsClusterControlPlaneAwsServicesAuthenticationOutputReference
	_jsii_.Get(
		j,
		"awsServicesAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) AwsServicesAuthenticationInput() *ContainerAwsClusterControlPlaneAwsServicesAuthentication {
	var returns *ContainerAwsClusterControlPlaneAwsServicesAuthentication
	_jsii_.Get(
		j,
		"awsServicesAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ConfigEncryption() ContainerAwsClusterControlPlaneConfigEncryptionOutputReference {
	var returns ContainerAwsClusterControlPlaneConfigEncryptionOutputReference
	_jsii_.Get(
		j,
		"configEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ConfigEncryptionInput() *ContainerAwsClusterControlPlaneConfigEncryption {
	var returns *ContainerAwsClusterControlPlaneConfigEncryption
	_jsii_.Get(
		j,
		"configEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) DatabaseEncryption() ContainerAwsClusterControlPlaneDatabaseEncryptionOutputReference {
	var returns ContainerAwsClusterControlPlaneDatabaseEncryptionOutputReference
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) DatabaseEncryptionInput() *ContainerAwsClusterControlPlaneDatabaseEncryption {
	var returns *ContainerAwsClusterControlPlaneDatabaseEncryption
	_jsii_.Get(
		j,
		"databaseEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) InternalValue() *ContainerAwsClusterControlPlane {
	var returns *ContainerAwsClusterControlPlane
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) MainVolume() ContainerAwsClusterControlPlaneMainVolumeOutputReference {
	var returns ContainerAwsClusterControlPlaneMainVolumeOutputReference
	_jsii_.Get(
		j,
		"mainVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) MainVolumeInput() *ContainerAwsClusterControlPlaneMainVolume {
	var returns *ContainerAwsClusterControlPlaneMainVolume
	_jsii_.Get(
		j,
		"mainVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ProxyConfig() ContainerAwsClusterControlPlaneProxyConfigOutputReference {
	var returns ContainerAwsClusterControlPlaneProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ProxyConfigInput() *ContainerAwsClusterControlPlaneProxyConfig {
	var returns *ContainerAwsClusterControlPlaneProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) RootVolume() ContainerAwsClusterControlPlaneRootVolumeOutputReference {
	var returns ContainerAwsClusterControlPlaneRootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) RootVolumeInput() *ContainerAwsClusterControlPlaneRootVolume {
	var returns *ContainerAwsClusterControlPlaneRootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) SshConfig() ContainerAwsClusterControlPlaneSshConfigOutputReference {
	var returns ContainerAwsClusterControlPlaneSshConfigOutputReference
	_jsii_.Get(
		j,
		"sshConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) SshConfigInput() *ContainerAwsClusterControlPlaneSshConfig {
	var returns *ContainerAwsClusterControlPlaneSshConfig
	_jsii_.Get(
		j,
		"sshConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewContainerAwsClusterControlPlaneOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerAwsClusterControlPlaneOutputReference {
	_init_.Initialize()

	if err := validateNewContainerAwsClusterControlPlaneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAwsClusterControlPlaneOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerAwsCluster.ContainerAwsClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerAwsClusterControlPlaneOutputReference_Override(c ContainerAwsClusterControlPlaneOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerAwsCluster.ContainerAwsClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetInternalValue(val *ContainerAwsClusterControlPlane) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutAwsServicesAuthentication(value *ContainerAwsClusterControlPlaneAwsServicesAuthentication) {
	if err := c.validatePutAwsServicesAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsServicesAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutConfigEncryption(value *ContainerAwsClusterControlPlaneConfigEncryption) {
	if err := c.validatePutConfigEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfigEncryption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutDatabaseEncryption(value *ContainerAwsClusterControlPlaneDatabaseEncryption) {
	if err := c.validatePutDatabaseEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDatabaseEncryption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutMainVolume(value *ContainerAwsClusterControlPlaneMainVolume) {
	if err := c.validatePutMainVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMainVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutProxyConfig(value *ContainerAwsClusterControlPlaneProxyConfig) {
	if err := c.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutRootVolume(value *ContainerAwsClusterControlPlaneRootVolume) {
	if err := c.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) PutSshConfig(value *ContainerAwsClusterControlPlaneSshConfig) {
	if err := c.validatePutSshConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSshConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetMainVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetMainVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetRootVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetRootVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetSshConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSshConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerAwsClusterControlPlaneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

