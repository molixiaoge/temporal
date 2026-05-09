package frontend

import (
	"context"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
)

func (wh *WorkflowHandler) Echo(context.Context, *workflowservice.EchoRequest) (*workflowservice.EchoResponse, error) {
	return nil, serviceerror.NewUnimplemented("Echo is not enabled")
}

func (wh *WorkflowHandler) StartActivityExecution(context.Context, *workflowservice.StartActivityExecutionRequest) (*workflowservice.StartActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("StartActivityExecution is not enabled")
}

func (wh *WorkflowHandler) StartNexusOperationExecution(context.Context, *workflowservice.StartNexusOperationExecutionRequest) (*workflowservice.StartNexusOperationExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("StartNexusOperationExecution is not enabled")
}

func (wh *WorkflowHandler) DescribeActivityExecution(context.Context, *workflowservice.DescribeActivityExecutionRequest) (*workflowservice.DescribeActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("DescribeActivityExecution is not enabled")
}

func (wh *WorkflowHandler) DescribeNexusOperationExecution(context.Context, *workflowservice.DescribeNexusOperationExecutionRequest) (*workflowservice.DescribeNexusOperationExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("DescribeNexusOperationExecution is not enabled")
}

func (wh *WorkflowHandler) PollActivityExecution(context.Context, *workflowservice.PollActivityExecutionRequest) (*workflowservice.PollActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("PollActivityExecution is not enabled")
}

func (wh *WorkflowHandler) PollNexusOperationExecution(context.Context, *workflowservice.PollNexusOperationExecutionRequest) (*workflowservice.PollNexusOperationExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("PollNexusOperationExecution is not enabled")
}

func (wh *WorkflowHandler) ListActivityExecutions(context.Context, *workflowservice.ListActivityExecutionsRequest) (*workflowservice.ListActivityExecutionsResponse, error) {
	return nil, serviceerror.NewUnimplemented("ListActivityExecutions is not enabled")
}

func (wh *WorkflowHandler) ListNexusOperationExecutions(context.Context, *workflowservice.ListNexusOperationExecutionsRequest) (*workflowservice.ListNexusOperationExecutionsResponse, error) {
	return nil, serviceerror.NewUnimplemented("ListNexusOperationExecutions is not enabled")
}

func (wh *WorkflowHandler) CountActivityExecutions(context.Context, *workflowservice.CountActivityExecutionsRequest) (*workflowservice.CountActivityExecutionsResponse, error) {
	return nil, serviceerror.NewUnimplemented("CountActivityExecutions is not enabled")
}

func (wh *WorkflowHandler) CountNexusOperationExecutions(context.Context, *workflowservice.CountNexusOperationExecutionsRequest) (*workflowservice.CountNexusOperationExecutionsResponse, error) {
	return nil, serviceerror.NewUnimplemented("CountNexusOperationExecutions is not enabled")
}

func (wh *WorkflowHandler) RequestCancelActivityExecution(context.Context, *workflowservice.RequestCancelActivityExecutionRequest) (*workflowservice.RequestCancelActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("RequestCancelActivityExecution is not enabled")
}

func (wh *WorkflowHandler) RequestCancelNexusOperationExecution(context.Context, *workflowservice.RequestCancelNexusOperationExecutionRequest) (*workflowservice.RequestCancelNexusOperationExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("RequestCancelNexusOperationExecution is not enabled")
}

func (wh *WorkflowHandler) TerminateActivityExecution(context.Context, *workflowservice.TerminateActivityExecutionRequest) (*workflowservice.TerminateActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("TerminateActivityExecution is not enabled")
}

func (wh *WorkflowHandler) DeleteActivityExecution(context.Context, *workflowservice.DeleteActivityExecutionRequest) (*workflowservice.DeleteActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("DeleteActivityExecution is not enabled")
}

func (wh *WorkflowHandler) PauseActivityExecution(context.Context, *workflowservice.PauseActivityExecutionRequest) (*workflowservice.PauseActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("PauseActivityExecution is not enabled")
}

func (wh *WorkflowHandler) ResetActivityExecution(context.Context, *workflowservice.ResetActivityExecutionRequest) (*workflowservice.ResetActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("ResetActivityExecution is not enabled")
}

func (wh *WorkflowHandler) UnpauseActivityExecution(context.Context, *workflowservice.UnpauseActivityExecutionRequest) (*workflowservice.UnpauseActivityExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("UnpauseActivityExecution is not enabled")
}

func (wh *WorkflowHandler) UpdateActivityExecutionOptions(context.Context, *workflowservice.UpdateActivityExecutionOptionsRequest) (*workflowservice.UpdateActivityExecutionOptionsResponse, error) {
	return nil, serviceerror.NewUnimplemented("UpdateActivityExecutionOptions is not enabled")
}

func (wh *WorkflowHandler) TerminateNexusOperationExecution(context.Context, *workflowservice.TerminateNexusOperationExecutionRequest) (*workflowservice.TerminateNexusOperationExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("TerminateNexusOperationExecution is not enabled")
}

func (wh *WorkflowHandler) DeleteNexusOperationExecution(context.Context, *workflowservice.DeleteNexusOperationExecutionRequest) (*workflowservice.DeleteNexusOperationExecutionResponse, error) {
	return nil, serviceerror.NewUnimplemented("DeleteNexusOperationExecution is not enabled")
}
