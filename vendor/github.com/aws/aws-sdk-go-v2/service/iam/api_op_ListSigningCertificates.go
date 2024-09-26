// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the signing certificates associated with the
// specified IAM user. If none exists, the operation returns an empty list.
//
// Although each user is limited to a small number of signing certificates, you
// can still paginate the results using the MaxItems and Marker parameters.
//
// If the UserName field is not specified, the user name is determined implicitly
// based on the Amazon Web Services access key ID used to sign the request for this
// operation. This operation works for access keys under the Amazon Web Services
// account. Consequently, you can use this operation to manage Amazon Web Services
// account root user credentials even if the Amazon Web Services account has no
// associated users.
func (c *Client) ListSigningCertificates(ctx context.Context, params *ListSigningCertificatesInput, optFns ...func(*Options)) (*ListSigningCertificatesOutput, error) {
	if params == nil {
		params = &ListSigningCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSigningCertificates", params, optFns, c.addOperationListSigningCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSigningCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSigningCertificatesInput struct {

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	MaxItems *int32

	// The name of the IAM user whose signing certificates you want to examine.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	UserName *string

	noSmithyDocumentSerde
}

// Contains the response to a successful ListSigningCertificates request.
type ListSigningCertificatesOutput struct {

	// A list of the user's signing certificate information.
	//
	// This member is required.
	Certificates []types.SigningCertificate

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSigningCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSigningCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSigningCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSigningCertificates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSigningCertificates(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListSigningCertificatesPaginatorOptions is the paginator options for
// ListSigningCertificates
type ListSigningCertificatesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSigningCertificatesPaginator is a paginator for ListSigningCertificates
type ListSigningCertificatesPaginator struct {
	options   ListSigningCertificatesPaginatorOptions
	client    ListSigningCertificatesAPIClient
	params    *ListSigningCertificatesInput
	nextToken *string
	firstPage bool
}

// NewListSigningCertificatesPaginator returns a new
// ListSigningCertificatesPaginator
func NewListSigningCertificatesPaginator(client ListSigningCertificatesAPIClient, params *ListSigningCertificatesInput, optFns ...func(*ListSigningCertificatesPaginatorOptions)) *ListSigningCertificatesPaginator {
	if params == nil {
		params = &ListSigningCertificatesInput{}
	}

	options := ListSigningCertificatesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSigningCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSigningCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSigningCertificates page.
func (p *ListSigningCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSigningCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListSigningCertificates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListSigningCertificatesAPIClient is a client that implements the
// ListSigningCertificates operation.
type ListSigningCertificatesAPIClient interface {
	ListSigningCertificates(context.Context, *ListSigningCertificatesInput, ...func(*Options)) (*ListSigningCertificatesOutput, error)
}

var _ ListSigningCertificatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSigningCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSigningCertificates",
	}
}
