// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a target with a maintenance window.
func (c *Client) RegisterTargetWithMaintenanceWindow(ctx context.Context, params *RegisterTargetWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTargetWithMaintenanceWindowOutput, error) {
	if params == nil {
		params = &RegisterTargetWithMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTargetWithMaintenanceWindow", params, optFns, c.addOperationRegisterTargetWithMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTargetWithMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTargetWithMaintenanceWindowInput struct {

	// The type of target being registered with the maintenance window.
	//
	// This member is required.
	ResourceType types.MaintenanceWindowResourceType

	// The targets to register with the maintenance window. In other words, the
	// managed nodes to run commands on when the maintenance window runs.
	//
	// If a single maintenance window task is registered with multiple targets, its
	// task invocations occur sequentially and not in parallel. If your task must run
	// on multiple targets at the same time, register a task for each target
	// individually and assign each task the same priority level.
	//
	// You can specify targets using managed node IDs, resource group names, or tags
	// that have been applied to managed nodes.
	//
	// Example 1: Specify managed node IDs
	//
	//     Key=InstanceIds,Values=,,
	//
	// Example 2: Use tag key-pairs applied to managed nodes
	//
	//     Key=tag:,Values=,
	//
	// Example 3: Use tag-keys applied to managed nodes
	//
	//     Key=tag-key,Values=,
	//
	// Example 4: Use resource group names
	//
	//     Key=resource-groups:Name,Values=
	//
	// Example 5: Use filters for resource group types
	//
	//     Key=resource-groups:ResourceTypeFilters,Values=,
	//
	// For Key=resource-groups:ResourceTypeFilters , specify resource types in the
	// following format
	//
	//     Key=resource-groups:ResourceTypeFilters,Values=AWS::EC2::INSTANCE,AWS::EC2::VPC
	//
	// For more information about these examples formats, including the best use case
	// for each one, see [Examples: Register targets with a maintenance window]in the Amazon Web Services Systems Manager User Guide.
	//
	// [Examples: Register targets with a maintenance window]: https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html
	//
	// This member is required.
	Targets []types.Target

	// The ID of the maintenance window the target should be registered with.
	//
	// This member is required.
	WindowId *string

	// User-provided idempotency token.
	ClientToken *string

	// An optional description for the target.
	Description *string

	// An optional name for the target.
	Name *string

	// User-provided value that will be included in any Amazon CloudWatch Events
	// events raised while running tasks for these targets in this maintenance window.
	OwnerInformation *string

	noSmithyDocumentSerde
}

type RegisterTargetWithMaintenanceWindowOutput struct {

	// The ID of the target definition in this maintenance window.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterTargetWithMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterTargetWithMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterTargetWithMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterTargetWithMaintenanceWindow"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opRegisterTargetWithMaintenanceWindowMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterTargetWithMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTargetWithMaintenanceWindow(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterTargetWithMaintenanceWindowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterTargetWithMaintenanceWindowInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterTargetWithMaintenanceWindowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterTargetWithMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterTargetWithMaintenanceWindow",
	}
}
