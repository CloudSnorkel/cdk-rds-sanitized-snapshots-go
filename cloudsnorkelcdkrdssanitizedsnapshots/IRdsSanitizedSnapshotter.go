package cloudsnorkelcdkrdssanitizedsnapshots

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
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
	// Default: 'postgres' for PostgreSQL and not set for MySQL.
	//
	// Experimental.
	DatabaseName() *string
	// VPC subnets to use for temporary databases.
	// Default: ec2.SubnetType.PRIVATE_ISOLATED
	//
	// Experimental.
	DbSubnets() *awsec2.SubnetSelection
	// Cluster where sanitization task will be executed.
	// Default: a new cluster running on given VPC.
	//
	// Experimental.
	FargateCluster() awsecs.ICluster
	// VPC subnets to use for sanitization task.
	// Default: ec2.SubnetType.PRIVATE_WITH_EGRESS
	//
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
	// Default: cluster identifier (which might be too long).
	//
	// Experimental.
	SnapshotPrefix() *string
	// Prefix for all temporary snapshots and databases.
	//
	// The step function execution id will be added to it.
	// Default: 'sanitize'.
	//
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

