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

// Retrieves the current effective patches (the patch and the approval state) for
// the specified patch baseline. Applies to patch baselines for Windows only.
func (c *Client) DescribeEffectivePatchesForPatchBaseline(ctx context.Context, params *DescribeEffectivePatchesForPatchBaselineInput, optFns ...func(*Options)) (*DescribeEffectivePatchesForPatchBaselineOutput, error) {
	if params == nil {
		params = &DescribeEffectivePatchesForPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEffectivePatchesForPatchBaseline", params, optFns, c.addOperationDescribeEffectivePatchesForPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEffectivePatchesForPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectivePatchesForPatchBaselineInput struct {

	// The ID of the patch baseline to retrieve the effective patches for.
	//
	// This member is required.
	BaselineId *string

	// The maximum number of patches to return (per page).
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEffectivePatchesForPatchBaselineOutput struct {

	// An array of patches and patch status.
	EffectivePatches []types.EffectivePatch

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEffectivePatchesForPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectivePatchesForPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectivePatchesForPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEffectivePatchesForPatchBaseline"); err != nil {
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
	if err = addOpDescribeEffectivePatchesForPatchBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectivePatchesForPatchBaseline(options.Region), middleware.Before); err != nil {
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

// DescribeEffectivePatchesForPatchBaselinePaginatorOptions is the paginator
// options for DescribeEffectivePatchesForPatchBaseline
type DescribeEffectivePatchesForPatchBaselinePaginatorOptions struct {
	// The maximum number of patches to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEffectivePatchesForPatchBaselinePaginator is a paginator for
// DescribeEffectivePatchesForPatchBaseline
type DescribeEffectivePatchesForPatchBaselinePaginator struct {
	options   DescribeEffectivePatchesForPatchBaselinePaginatorOptions
	client    DescribeEffectivePatchesForPatchBaselineAPIClient
	params    *DescribeEffectivePatchesForPatchBaselineInput
	nextToken *string
	firstPage bool
}

// NewDescribeEffectivePatchesForPatchBaselinePaginator returns a new
// DescribeEffectivePatchesForPatchBaselinePaginator
func NewDescribeEffectivePatchesForPatchBaselinePaginator(client DescribeEffectivePatchesForPatchBaselineAPIClient, params *DescribeEffectivePatchesForPatchBaselineInput, optFns ...func(*DescribeEffectivePatchesForPatchBaselinePaginatorOptions)) *DescribeEffectivePatchesForPatchBaselinePaginator {
	if params == nil {
		params = &DescribeEffectivePatchesForPatchBaselineInput{}
	}

	options := DescribeEffectivePatchesForPatchBaselinePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEffectivePatchesForPatchBaselinePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEffectivePatchesForPatchBaselinePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEffectivePatchesForPatchBaseline page.
func (p *DescribeEffectivePatchesForPatchBaselinePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEffectivePatchesForPatchBaselineOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeEffectivePatchesForPatchBaseline(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeEffectivePatchesForPatchBaselineAPIClient is a client that implements
// the DescribeEffectivePatchesForPatchBaseline operation.
type DescribeEffectivePatchesForPatchBaselineAPIClient interface {
	DescribeEffectivePatchesForPatchBaseline(context.Context, *DescribeEffectivePatchesForPatchBaselineInput, ...func(*Options)) (*DescribeEffectivePatchesForPatchBaselineOutput, error)
}

var _ DescribeEffectivePatchesForPatchBaselineAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeEffectivePatchesForPatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEffectivePatchesForPatchBaseline",
	}
}
