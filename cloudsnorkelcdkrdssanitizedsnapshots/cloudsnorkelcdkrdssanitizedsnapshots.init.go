package cloudsnorkelcdkrdssanitizedsnapshots

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@cloudsnorkel/cdk-rds-sanitized-snapshots.IRdsSanitizedSnapshotter",
		reflect.TypeOf((*IRdsSanitizedSnapshotter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "databaseCluster", GoGetter: "DatabaseCluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInstance", GoGetter: "DatabaseInstance"},
			_jsii_.MemberProperty{JsiiProperty: "databaseKey", GoGetter: "DatabaseKey"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnets", GoGetter: "DbSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "fargateCluster", GoGetter: "FargateCluster"},
			_jsii_.MemberProperty{JsiiProperty: "sanitizeSubnets", GoGetter: "SanitizeSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberProperty{JsiiProperty: "script", GoGetter: "Script"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotHistoryLimit", GoGetter: "SnapshotHistoryLimit"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotKey", GoGetter: "SnapshotKey"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotPrefix", GoGetter: "SnapshotPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "tempPrefix", GoGetter: "TempPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			return &jsiiProxy_IRdsSanitizedSnapshotter{}
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-rds-sanitized-snapshots.RdsSanitizedSnapshotter",
		reflect.TypeOf((*RdsSanitizedSnapshotter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotter", GoGetter: "Snapshotter"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RdsSanitizedSnapshotter{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
}
