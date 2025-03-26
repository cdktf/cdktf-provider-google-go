// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/databasemigrationservicemigrationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/database_migration_service_migration_job google_database_migration_service_migration_job}.
type DatabaseMigrationServiceMigrationJob interface {
	cdktf.TerraformResource
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
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DumpFlags() DatabaseMigrationServiceMigrationJobDumpFlagsOutputReference
	DumpFlagsInput() *DatabaseMigrationServiceMigrationJobDumpFlags
	DumpPath() *string
	SetDumpPath(val *string)
	DumpPathInput() *string
	DumpType() *string
	SetDumpType(val *string)
	DumpTypeInput() *string
	EffectiveLabels() cdktf.StringMap
	Error() DatabaseMigrationServiceMigrationJobErrorList
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MigrationJobId() *string
	SetMigrationJobId(val *string)
	MigrationJobIdInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PerformanceConfig() DatabaseMigrationServiceMigrationJobPerformanceConfigOutputReference
	PerformanceConfigInput() *DatabaseMigrationServiceMigrationJobPerformanceConfig
	Phase() *string
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
	ReverseSshConnectivity() DatabaseMigrationServiceMigrationJobReverseSshConnectivityOutputReference
	ReverseSshConnectivityInput() *DatabaseMigrationServiceMigrationJobReverseSshConnectivity
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	State() *string
	StaticIpConnectivity() DatabaseMigrationServiceMigrationJobStaticIpConnectivityOutputReference
	StaticIpConnectivityInput() *DatabaseMigrationServiceMigrationJobStaticIpConnectivity
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatabaseMigrationServiceMigrationJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VpcPeeringConnectivity() DatabaseMigrationServiceMigrationJobVpcPeeringConnectivityOutputReference
	VpcPeeringConnectivityInput() *DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity
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
	PutDumpFlags(value *DatabaseMigrationServiceMigrationJobDumpFlags)
	PutPerformanceConfig(value *DatabaseMigrationServiceMigrationJobPerformanceConfig)
	PutReverseSshConnectivity(value *DatabaseMigrationServiceMigrationJobReverseSshConnectivity)
	PutStaticIpConnectivity(value *DatabaseMigrationServiceMigrationJobStaticIpConnectivity)
	PutTimeouts(value *DatabaseMigrationServiceMigrationJobTimeouts)
	PutVpcPeeringConnectivity(value *DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity)
	ResetDisplayName()
	ResetDumpFlags()
	ResetDumpPath()
	ResetDumpType()
	ResetId()
	ResetLabels()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceConfig()
	ResetProject()
	ResetReverseSshConnectivity()
	ResetStaticIpConnectivity()
	ResetTimeouts()
	ResetVpcPeeringConnectivity()
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

// The jsii proxy struct for DatabaseMigrationServiceMigrationJob
type jsiiProxy_DatabaseMigrationServiceMigrationJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DumpFlags() DatabaseMigrationServiceMigrationJobDumpFlagsOutputReference {
	var returns DatabaseMigrationServiceMigrationJobDumpFlagsOutputReference
	_jsii_.Get(
		j,
		"dumpFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DumpFlagsInput() *DatabaseMigrationServiceMigrationJobDumpFlags {
	var returns *DatabaseMigrationServiceMigrationJobDumpFlags
	_jsii_.Get(
		j,
		"dumpFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DumpPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DumpPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DumpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) DumpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Error() DatabaseMigrationServiceMigrationJobErrorList {
	var returns DatabaseMigrationServiceMigrationJobErrorList
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) MigrationJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) MigrationJobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationJobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) PerformanceConfig() DatabaseMigrationServiceMigrationJobPerformanceConfigOutputReference {
	var returns DatabaseMigrationServiceMigrationJobPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"performanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) PerformanceConfigInput() *DatabaseMigrationServiceMigrationJobPerformanceConfig {
	var returns *DatabaseMigrationServiceMigrationJobPerformanceConfig
	_jsii_.Get(
		j,
		"performanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Phase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) ReverseSshConnectivity() DatabaseMigrationServiceMigrationJobReverseSshConnectivityOutputReference {
	var returns DatabaseMigrationServiceMigrationJobReverseSshConnectivityOutputReference
	_jsii_.Get(
		j,
		"reverseSshConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) ReverseSshConnectivityInput() *DatabaseMigrationServiceMigrationJobReverseSshConnectivity {
	var returns *DatabaseMigrationServiceMigrationJobReverseSshConnectivity
	_jsii_.Get(
		j,
		"reverseSshConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) StaticIpConnectivity() DatabaseMigrationServiceMigrationJobStaticIpConnectivityOutputReference {
	var returns DatabaseMigrationServiceMigrationJobStaticIpConnectivityOutputReference
	_jsii_.Get(
		j,
		"staticIpConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) StaticIpConnectivityInput() *DatabaseMigrationServiceMigrationJobStaticIpConnectivity {
	var returns *DatabaseMigrationServiceMigrationJobStaticIpConnectivity
	_jsii_.Get(
		j,
		"staticIpConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Timeouts() DatabaseMigrationServiceMigrationJobTimeoutsOutputReference {
	var returns DatabaseMigrationServiceMigrationJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) VpcPeeringConnectivity() DatabaseMigrationServiceMigrationJobVpcPeeringConnectivityOutputReference {
	var returns DatabaseMigrationServiceMigrationJobVpcPeeringConnectivityOutputReference
	_jsii_.Get(
		j,
		"vpcPeeringConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob) VpcPeeringConnectivityInput() *DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity {
	var returns *DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity
	_jsii_.Get(
		j,
		"vpcPeeringConnectivityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/database_migration_service_migration_job google_database_migration_service_migration_job} Resource.
func NewDatabaseMigrationServiceMigrationJob(scope constructs.Construct, id *string, config *DatabaseMigrationServiceMigrationJobConfig) DatabaseMigrationServiceMigrationJob {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceMigrationJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceMigrationJob{}

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/database_migration_service_migration_job google_database_migration_service_migration_job} Resource.
func NewDatabaseMigrationServiceMigrationJob_Override(d DatabaseMigrationServiceMigrationJob, scope constructs.Construct, id *string, config *DatabaseMigrationServiceMigrationJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetDumpPath(val *string) {
	if err := j.validateSetDumpPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dumpPath",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetDumpType(val *string) {
	if err := j.validateSetDumpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dumpType",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetMigrationJobId(val *string) {
	if err := j.validateSetMigrationJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationJobId",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJob)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseMigrationServiceMigrationJob resource upon running "cdktf plan <stack-name>".
func DatabaseMigrationServiceMigrationJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceMigrationJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
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
func DatabaseMigrationServiceMigrationJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceMigrationJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseMigrationServiceMigrationJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceMigrationJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseMigrationServiceMigrationJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceMigrationJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseMigrationServiceMigrationJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) PutDumpFlags(value *DatabaseMigrationServiceMigrationJobDumpFlags) {
	if err := d.validatePutDumpFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDumpFlags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) PutPerformanceConfig(value *DatabaseMigrationServiceMigrationJobPerformanceConfig) {
	if err := d.validatePutPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPerformanceConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) PutReverseSshConnectivity(value *DatabaseMigrationServiceMigrationJobReverseSshConnectivity) {
	if err := d.validatePutReverseSshConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReverseSshConnectivity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) PutStaticIpConnectivity(value *DatabaseMigrationServiceMigrationJobStaticIpConnectivity) {
	if err := d.validatePutStaticIpConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStaticIpConnectivity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) PutTimeouts(value *DatabaseMigrationServiceMigrationJobTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) PutVpcPeeringConnectivity(value *DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity) {
	if err := d.validatePutVpcPeeringConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVpcPeeringConnectivity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetDumpFlags() {
	_jsii_.InvokeVoid(
		d,
		"resetDumpFlags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetDumpPath() {
	_jsii_.InvokeVoid(
		d,
		"resetDumpPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetDumpType() {
	_jsii_.InvokeVoid(
		d,
		"resetDumpType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetPerformanceConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPerformanceConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetReverseSshConnectivity() {
	_jsii_.InvokeVoid(
		d,
		"resetReverseSshConnectivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetStaticIpConnectivity() {
	_jsii_.InvokeVoid(
		d,
		"resetStaticIpConnectivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ResetVpcPeeringConnectivity() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcPeeringConnectivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

