package cloudsnorkelcdkrdssanitizedsnapshots

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-rds-sanitized-snapshots-go/cloudsnorkelcdkrdssanitizedsnapshots/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-rds-sanitized-snapshots-go/cloudsnorkelcdkrdssanitizedsnapshots/internal"
)

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

	if err := validateNewRdsSanitizedSnapshotterParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_RdsSanitizedSnapshotter)SetSnapshotter(val awsstepfunctions.StateMachine) {
	if err := j.validateSetSnapshotterParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateRdsSanitizedSnapshotter_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

