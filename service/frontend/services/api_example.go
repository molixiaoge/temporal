package services

import (
	"context"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// exampleWorkflowWrapper intercepts stable RPCs to add experimental behaviour.
// Embedding WorkflowServiceServer means every other method passes through.
type exampleWorkflowWrapper struct {
	workflowservice.WorkflowServiceServer
}

func (w *exampleWorkflowWrapper) Echo(_ context.Context, req *workflowservice.EchoRequest) (*workflowservice.EchoResponse, error) {
	return &workflowservice.EchoResponse{Payload: req.GetPayload()}, nil
}

func (w *exampleWorkflowWrapper) StartWorkflowExecution(
	ctx context.Context,
	req *workflowservice.StartWorkflowExecutionRequest,
) (*workflowservice.StartWorkflowExecutionResponse, error) {
	// experimental_enum_value: react to the FOO conflict policy.
	if req.GetWorkflowIdConflictPolicy() == enums.WORKFLOW_ID_CONFLICT_POLICY_FOO {
		return nil, status.Error(codes.InvalidArgument, "WORKFLOW_ID_CONFLICT_POLICY_FOO is not yet supported")
	}

	// experimental_field: read foo_text directly from the api_next request.
	if req.GetFooText() != "" {
		_ = req.GetFooText() // a real feature would use this
	}

	return w.WorkflowServiceServer.StartWorkflowExecution(ctx, req)
}

func init() {
	register("example", variant{registerWorkflow: func(server *grpc.Server, workflow workflowservice.WorkflowServiceServer) {
		wrapped := &exampleWorkflowWrapper{workflow}
		workflowservice.RegisterWorkflowServiceServer(server, wrapped)
	}})
}
