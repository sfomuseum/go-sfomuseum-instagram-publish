// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a connection to a target (for example, a managed node) for a Session
// Manager session. Returns a URL and token that can be used to open a WebSocket
// connection for sending input and receiving outputs.
//
// Amazon Web Services CLI usage: start-session is an interactive command that
// requires the Session Manager plugin to be installed on the client machine making
// the call. For information, see [Install the Session Manager plugin for the Amazon Web Services CLI]in the Amazon Web Services Systems Manager User
// Guide.
//
// Amazon Web Services Tools for PowerShell usage: Start-SSMSession isn't
// currently supported by Amazon Web Services Tools for PowerShell on Windows local
// machines.
//
// [Install the Session Manager plugin for the Amazon Web Services CLI]: https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html
func (c *Client) StartSession(ctx context.Context, params *StartSessionInput, optFns ...func(*Options)) (*StartSessionOutput, error) {
	if params == nil {
		params = &StartSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSession", params, optFns, c.addOperationStartSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSessionInput struct {

	// The managed node to connect to for the session.
	//
	// This member is required.
	Target *string

	// The name of the SSM document you want to use to define the type of session,
	// input parameters, or preferences for the session. For example,
	// SSM-SessionManagerRunShell . You can call the GetDocument API to verify the document
	// exists before attempting to start a session. If no document name is provided, a
	// shell to the managed node is launched by default. For more information, see [Start a session]in
	// the Amazon Web Services Systems Manager User Guide.
	//
	// [Start a session]: https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html
	DocumentName *string

	// The values you want to specify for the parameters defined in the Session
	// document.
	Parameters map[string][]string

	// The reason for connecting to the instance. This value is included in the
	// details for the Amazon CloudWatch Events event created when you start the
	// session.
	Reason *string

	noSmithyDocumentSerde
}

type StartSessionOutput struct {

	// The ID of the session.
	SessionId *string

	// A URL back to SSM Agent on the managed node that the Session Manager client
	// uses to send commands and receive output from the node. Format:
	// wss://ssmmessages.region.amazonaws.com/v1/data-channel/session-id?stream=(input|output)
	//
	// region represents the Region identifier for an Amazon Web Services Region
	// supported by Amazon Web Services Systems Manager, such as us-east-2 for the US
	// East (Ohio) Region. For a list of supported region values, see the Region column
	// in [Systems Manager service endpoints]in the Amazon Web Services General Reference.
	//
	// session-id represents the ID of a Session Manager session, such as
	// 1a2b3c4dEXAMPLE .
	//
	// [Systems Manager service endpoints]: https://docs.aws.amazon.com/general/latest/gr/ssm.html#ssm_region
	StreamUrl *string

	// An encrypted token value containing session and caller information. This token
	// is used to authenticate the connection to the managed node, and is valid only
	// long enough to ensure the connection is successful. Never share your session's
	// token.
	TokenValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSession"); err != nil {
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
	if err = addOpStartSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSession",
	}
}
