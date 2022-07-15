// CDK construct to periodically take snapshots of RDS databases, sanitize them, and share with selected accounts.
package cloudsnorkelcdkrdssanitizedsnapshots

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-rds-sanitized-snapshots-go/cloudsnorkelcdkrdssanitizedsnapshots/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-rds-sanitized-snapshots-go/cloudsnorkelcdkrdssanitizedsnapshots/internal"
)

// Experimental.
type IRdsSanitizedSnapshotter interface {
	// Database cluster to snapshot and sanitize.
	//
	// Only one of `databaseCluster` and `databaseInstance` can be specified.
	// Experimental.
	DatabaseCluster() awsrds.IDatabaseCluster
	// Database instance to snapshot and sanitize.
	//
	// Only one of `databaseCluster` and `databaseInstance` can be specified.
	// Experimental.
	DatabaseInstance() awsrds.IDatabaseInstance
	// KMS key used to encrypt original database, if any.
	// Experimental.
	DatabaseKey() awskms.IKey
	// Name of database to connect to inside the RDS cluster or instance.
	//
	// This database will be used to execute the SQL script.
	// Experimental.
	DatabaseName() *string
	// VPC subnets to use for temporary databases.
	// Experimental.
	DbSubnets() *awsec2.SubnetSelection
	// Cluster where sanitization task will be executed.
	// Experimental.
	FargateCluster() awsecs.ICluster
	// VPC subnets to use for sanitization task.
	// Experimental.
	SanitizeSubnets() *awsec2.SubnetSelection
	// The schedule or rate (frequency) that determines when the sanitized snapshot runs automatically.
	// Experimental.
	Schedule() awsevents.Schedule
	// SQL script used to sanitize the database. It will be executed against the temporary database.
	//
	// You would usually want to start this with `USE mydatabase;`.
	// Experimental.
	Script() *string
	// List of accounts the sanitized snapshot should be shared with.
	// Experimental.
	ShareAccounts() *[]*string
	// Limit the number of snapshot history.
	//
	// Set this to delete old snapshots and only leave a certain number of snapshots.
	// Experimental.
	SnapshotHistoryLimit() *float64
	// Optional KMS key to encrypt target snapshot.
	// Experimental.
	SnapshotKey() awskms.IKey
	// Prefix for sanitized snapshot name.
	//
	// The current date and time will be added to it.
	// Experimental.
	SnapshotPrefix() *string
	// Prefix for all temporary snapshots and databases.
	//
	// The step function execution id will be added to it.
	// Experimental.
	TempPrefix() *string
	// VPC where temporary database and sanitizing task will be created.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IRdsSanitizedSnapshotter
type jsiiProxy_IRdsSanitizedSnapshotter struct {
	_ byte // padding
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) DatabaseCluster() awsrds.IDatabaseCluster {
	var returns awsrds.IDatabaseCluster
	_jsii_.Get(
		j,
		"databaseCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) DatabaseInstance() awsrds.IDatabaseInstance {
	var returns awsrds.IDatabaseInstance
	_jsii_.Get(
		j,
		"databaseInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) DatabaseKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"databaseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) DbSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"dbSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) FargateCluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"fargateCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) SanitizeSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"sanitizeSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) Schedule() awsevents.Schedule {
	var returns awsevents.Schedule
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) ShareAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"shareAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) SnapshotHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) SnapshotKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"snapshotKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) SnapshotPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) TempPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRdsSanitizedSnapshotter) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

// A process to create sanitized snapshots of RDS instance or cluster, optionally on a schedule.
//
// The process is handled by a step function.
//
// 1. Snapshot the source database
// 2. Optionally re-encrypt the snapshot with a different key in case you want to share it with an account that doesn't have access to the original key
// 3. Create a temporary database
// 4. Run a Fargate task to connect to the temporary database and execute an arbitrary SQL script to sanitize it
// 5. Snapshot the sanitized database
// 6. Clean-up temporary snapshots and databases
// Experimental.
type RdsSanitizedSnapshotter interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Props() IRdsSanitizedSnapshotter
	// Step function in charge of the entire process including snapshotting, sanitizing, and cleanup.
	//
	// Trigger this step function to get a new snapshot.
	// Experimental.
	Snapshotter() awsstepfunctions.StateMachine
	// Experimental.
	SetSnapshotter(val awsstepfunctions.StateMachine)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RdsSanitizedSnapshotter
type jsiiProxy_RdsSanitizedSnapshotter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_RdsSanitizedSnapshotter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsSanitizedSnapshotter) Props() IRdsSanitizedSnapshotter {
	var returns IRdsSanitizedSnapshotter
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsSanitizedSnapshotter) Snapshotter() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"snapshotter",
		&returns,
	)
	return returns
}


// Experimental.
func NewRdsSanitizedSnapshotter(scope constructs.Construct, id *string, props IRdsSanitizedSnapshotter) RdsSanitizedSnapshotter {
	_init_.Initialize()

	j := jsiiProxy_RdsSanitizedSnapshotter{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-rds-sanitized-snapshots.RdsSanitizedSnapshotter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRdsSanitizedSnapshotter_Override(r RdsSanitizedSnapshotter, scope constructs.Construct, id *string, props IRdsSanitizedSnapshotter) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-rds-sanitized-snapshots.RdsSanitizedSnapshotter",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RdsSanitizedSnapshotter) SetSnapshotter(val awsstepfunctions.StateMachine) {
	_jsii_.Set(
		j,
		"snapshotter",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RdsSanitizedSnapshotter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-rds-sanitized-snapshots.RdsSanitizedSnapshotter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsSanitizedSnapshotter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

