// Functional test for the api_next incubator. go.temporal.io/api is replaced
// with api_next in this branch, so experimental fields and RPCs are available
// on the regular generated API packages.
package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	taskqueuepb "go.temporal.io/api/taskqueue/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/server/common/dynamicconfig"
	"go.temporal.io/server/tests/testcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestExperimentalApi_Example(t *testing.T) {
	env := testcore.NewEnv(t,
		testcore.WithDynamicConfig(dynamicconfig.FrontendAPIVariant, "example"))

	ctx := env.Context()
	frontend := env.ExperimentalFrontend()
	workflowClient := workflowservice.NewWorkflowServiceClient(frontend.Workflow())

	t.Run("echo responds", func(t *testing.T) {
		resp, err := workflowClient.Echo(ctx, &workflowservice.EchoRequest{
			Namespace: env.Namespace().String(),
			Payload:   "hello",
		})
		require.NoError(t, err)
		require.Equal(t, "hello", resp.GetPayload())
	})

	t.Run("stable RPCs still work", func(t *testing.T) {
		_, err := env.FrontendClient().DescribeNamespace(ctx, &workflowservice.DescribeNamespaceRequest{
			Namespace: env.Namespace().String(),
		})
		require.NoError(t, err)
	})

	t.Run("experimental enum value handled by server", func(t *testing.T) {
		_, err := env.FrontendClient().StartWorkflowExecution(ctx, &workflowservice.StartWorkflowExecutionRequest{
			Namespace:                env.Namespace().String(),
			WorkflowId:               "test-experimental-enum",
			WorkflowType:             &commonpb.WorkflowType{Name: "test"},
			TaskQueue:                &taskqueuepb.TaskQueue{Name: "test"},
			WorkflowIdConflictPolicy: enumspb.WORKFLOW_ID_CONFLICT_POLICY_FOO,
		})
		require.NotEqual(t, codes.Unimplemented, status.Code(err))
	})

	t.Run("experimental field handled by server", func(t *testing.T) {
		req := &workflowservice.StartWorkflowExecutionRequest{
			Namespace:    env.Namespace().String(),
			WorkflowId:   "test-experimental-overlay",
			WorkflowType: &commonpb.WorkflowType{Name: "test"},
			TaskQueue:    &taskqueuepb.TaskQueue{Name: "test"},
			FooText:      "hello api_next",
		}
		_, err := env.FrontendClient().StartWorkflowExecution(ctx, req)
		require.NotEqual(t, codes.Unimplemented, status.Code(err))
	})
}

func TestExperimentalApi_Stable(t *testing.T) {
	env := testcore.NewEnv(t)

	ctx := env.Context()
	frontend := env.ExperimentalFrontend()
	workflowClient := workflowservice.NewWorkflowServiceClient(frontend.Workflow())

	t.Run("echo not available", func(t *testing.T) {
		_, err := workflowClient.Echo(ctx, &workflowservice.EchoRequest{})
		require.Error(t, err)
		require.Contains(t, []codes.Code{codes.InvalidArgument, codes.Unimplemented}, serviceerror.ToStatus(err).Code())
	})

	t.Run("stable RPCs work", func(t *testing.T) {
		_, err := env.FrontendClient().DescribeNamespace(ctx, &workflowservice.DescribeNamespaceRequest{
			Namespace: env.Namespace().String(),
		})
		if err != nil {
			var namespaceNotFound *serviceerror.NamespaceNotFound
			require.NotEqual(t, codes.Unimplemented, status.Code(err))
			require.ErrorAs(t, err, &namespaceNotFound)
		}
	})
}

func TestExperimentalApi_MessageField(t *testing.T) {
	req := &workflowservice.StartWorkflowExecutionRequest{
		FooText: "hello api_next",
	}

	payload, err := proto.Marshal(req)
	require.NoError(t, err)
	roundTrip := &workflowservice.StartWorkflowExecutionRequest{}
	require.NoError(t, proto.Unmarshal(payload, roundTrip))
	require.Equal(t, "hello api_next", roundTrip.GetFooText())
}

func TestExperimentalApi_EnumValue(t *testing.T) {
	req := &workflowservice.StartWorkflowExecutionRequest{
		WorkflowIdConflictPolicy: enumspb.WORKFLOW_ID_CONFLICT_POLICY_FOO,
	}

	require.Equal(t, enumspb.WorkflowIdConflictPolicy(1000), req.GetWorkflowIdConflictPolicy())

	payload, err := proto.Marshal(req)
	require.NoError(t, err)

	roundTrip := &workflowservice.StartWorkflowExecutionRequest{}
	require.NoError(t, proto.Unmarshal(payload, roundTrip))
	require.Equal(t, enumspb.WORKFLOW_ID_CONFLICT_POLICY_FOO, roundTrip.GetWorkflowIdConflictPolicy())
}
