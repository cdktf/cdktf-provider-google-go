// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabasecloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/datagoogleoracledatabasecloudexadatainfrastructure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/oracle_database_cloud_exadata_infrastructure google_oracle_database_cloud_exadata_infrastructure}.
type DataGoogleOracleDatabaseCloudExadataInfrastructure interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudExadataInfrastructureId() *string
	SetCloudExadataInfrastructureId(val *string)
	CloudExadataInfrastructureIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DeletionProtection() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	EffectiveLabels() cdktf.StringMap
	EntitlementId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpOracleZone() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() cdktf.StringMap
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Properties() DataGoogleOracleDatabaseCloudExadataInfrastructurePropertiesList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataGoogleOracleDatabaseCloudExadataInfrastructure
type jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) CloudExadataInfrastructureIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) DeletionProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) EntitlementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GcpOracleZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpOracleZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Properties() DataGoogleOracleDatabaseCloudExadataInfrastructurePropertiesList {
	var returns DataGoogleOracleDatabaseCloudExadataInfrastructurePropertiesList
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/oracle_database_cloud_exadata_infrastructure google_oracle_database_cloud_exadata_infrastructure} Data Source.
func NewDataGoogleOracleDatabaseCloudExadataInfrastructure(scope constructs.Construct, id *string, config *DataGoogleOracleDatabaseCloudExadataInfrastructureConfig) DataGoogleOracleDatabaseCloudExadataInfrastructure {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseCloudExadataInfrastructureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/oracle_database_cloud_exadata_infrastructure google_oracle_database_cloud_exadata_infrastructure} Data Source.
func NewDataGoogleOracleDatabaseCloudExadataInfrastructure_Override(d DataGoogleOracleDatabaseCloudExadataInfrastructure, scope constructs.Construct, id *string, config *DataGoogleOracleDatabaseCloudExadataInfrastructureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetCloudExadataInfrastructureId(val *string) {
	if err := j.validateSetCloudExadataInfrastructureIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleOracleDatabaseCloudExadataInfrastructure resource upon running "cdktf plan <stack-name>".
func DataGoogleOracleDatabaseCloudExadataInfrastructure_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleOracleDatabaseCloudExadataInfrastructure_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
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
func DataGoogleOracleDatabaseCloudExadataInfrastructure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleOracleDatabaseCloudExadataInfrastructure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleOracleDatabaseCloudExadataInfrastructure_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleOracleDatabaseCloudExadataInfrastructure_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleOracleDatabaseCloudExadataInfrastructure_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleOracleDatabaseCloudExadataInfrastructure_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleOracleDatabaseCloudExadataInfrastructure_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructure.DataGoogleOracleDatabaseCloudExadataInfrastructure",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructure) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

