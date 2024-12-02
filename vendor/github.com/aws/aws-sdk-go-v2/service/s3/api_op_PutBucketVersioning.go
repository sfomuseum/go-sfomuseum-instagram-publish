// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported for directory buckets.
//
// When you enable versioning on a bucket for the first time, it might take a
// short amount of time for the change to be fully propagated. We recommend that
// you wait for 15 minutes after enabling versioning before issuing write
// operations ( PUT or DELETE ) on objects in the bucket.
//
// Sets the versioning state of an existing bucket.
//
// You can set the versioning state with one of the following values:
//
// Enabled—Enables versioning for the objects in the bucket. All objects added to
// the bucket receive a unique version ID.
//
// Suspended—Disables versioning for the objects in the bucket. All objects added
// to the bucket receive the version ID null.
//
// If the versioning state has never been set on a bucket, it has no versioning
// state; a [GetBucketVersioning]request does not return a versioning state value.
//
// In order to enable MFA Delete, you must be the bucket owner. If you are the
// bucket owner and want to enable MFA Delete in the bucket versioning
// configuration, you must include the x-amz-mfa request header and the Status and
// the MfaDelete request elements in a request to set the versioning state of the
// bucket.
//
// If you have an object expiration lifecycle configuration in your non-versioned
// bucket and you want to maintain the same permanent delete behavior when you
// enable versioning, you must add a noncurrent expiration policy. The noncurrent
// expiration lifecycle configuration will manage the deletes of the noncurrent
// object versions in the version-enabled bucket. (A version-enabled bucket
// maintains one current and zero or more noncurrent object versions.) For more
// information, see [Lifecycle and Versioning].
//
// The following operations are related to PutBucketVersioning :
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [GetBucketVersioning]
//
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Lifecycle and Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-and-other-bucket-config
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html
func (c *Client) PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error) {
	if params == nil {
		params = &PutBucketVersioningInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketVersioning", params, optFns, c.addOperationPutBucketVersioningMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketVersioningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketVersioningInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// Container for setting the versioning state.
	//
	// This member is required.
	VersioningConfiguration *types.VersioningConfiguration

	// Indicates the algorithm used to create the checksum for the object when you use
	// the SDK. This header will not provide any additional functionality if you don't
	// use the SDK. When you send this header, there must be a corresponding
	// x-amz-checksum or x-amz-trailer header sent. Otherwise, Amazon S3 fails the
	// request with the HTTP status code 400 Bad Request . For more information, see [Checking object integrity]
	// in the Amazon S3 User Guide.
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided
	// ChecksumAlgorithm parameter.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumAlgorithm types.ChecksumAlgorithm

	// >The base64-encoded 128-bit MD5 digest of the data. You must use this header as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see [RFC 1864].
	//
	// For requests made using the Amazon Web Services Command Line Interface (CLI) or
	// Amazon Web Services SDKs, this field is calculated automatically.
	//
	// [RFC 1864]: http://www.ietf.org/rfc/rfc1864.txt
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// The concatenation of the authentication device's serial number, a space, and
	// the value that is displayed on your authentication device.
	MFA *string

	noSmithyDocumentSerde
}

func (in *PutBucketVersioningInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketVersioningOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketVersioningMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketVersioning{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketVersioning{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketVersioning"); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIsExpressUserAgent(stack); err != nil {
		return err
	}
	if err = addOpPutBucketVersioningValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketVersioning(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketVersioningInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketVersioningUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3cust.AddExpressDefaultChecksumMiddleware(stack); err != nil {
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

func (v *PutBucketVersioningInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketVersioning(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketVersioning",
	}
}

// getPutBucketVersioningRequestAlgorithmMember gets the request checksum
// algorithm value provided as input.
func getPutBucketVersioningRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketVersioningInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketVersioningInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketVersioningRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketVersioningBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getPutBucketVersioningBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketVersioningInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketVersioningUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketVersioningBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
