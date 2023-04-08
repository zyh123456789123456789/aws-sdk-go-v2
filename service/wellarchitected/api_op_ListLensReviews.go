// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List lens reviews for a particular workload.
func (c *Client) ListLensReviews(ctx context.Context, params *ListLensReviewsInput, optFns ...func(*Options)) (*ListLensReviewsOutput, error) {
	if params == nil {
		params = &ListLensReviewsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLensReviews", params, optFns, c.addOperationListLensReviewsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLensReviewsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to list lens reviews.
type ListLensReviewsInput struct {

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults int32

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Output of a list lens reviews call.
type ListLensReviewsOutput struct {

	// List of lens summaries of lens reviews of a workload.
	LensReviewSummaries []types.LensReviewSummary

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLensReviewsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLensReviews{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLensReviews{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListLensReviewsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLensReviews(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListLensReviewsAPIClient is a client that implements the ListLensReviews
// operation.
type ListLensReviewsAPIClient interface {
	ListLensReviews(context.Context, *ListLensReviewsInput, ...func(*Options)) (*ListLensReviewsOutput, error)
}

var _ ListLensReviewsAPIClient = (*Client)(nil)

// ListLensReviewsPaginatorOptions is the paginator options for ListLensReviews
type ListLensReviewsPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLensReviewsPaginator is a paginator for ListLensReviews
type ListLensReviewsPaginator struct {
	options   ListLensReviewsPaginatorOptions
	client    ListLensReviewsAPIClient
	params    *ListLensReviewsInput
	nextToken *string
	firstPage bool
}

// NewListLensReviewsPaginator returns a new ListLensReviewsPaginator
func NewListLensReviewsPaginator(client ListLensReviewsAPIClient, params *ListLensReviewsInput, optFns ...func(*ListLensReviewsPaginatorOptions)) *ListLensReviewsPaginator {
	if params == nil {
		params = &ListLensReviewsInput{}
	}

	options := ListLensReviewsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLensReviewsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLensReviewsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLensReviews page.
func (p *ListLensReviewsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLensReviewsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLensReviews(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLensReviews(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListLensReviews",
	}
}
