// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all labels or a specific label if name is provided. This is a paginated
// API. If you provide a null maxResults , this action retrieves a maximum of 50
// records per page. If you provide a maxResults , the value must be between 10 and
// 50. To get the next page results, provide the pagination token from the
// GetGetLabelsResponse as part of your request. A null pagination token fetches
// the records from the beginning.
func (c *Client) GetLabels(ctx context.Context, params *GetLabelsInput, optFns ...func(*Options)) (*GetLabelsOutput, error) {
	if params == nil {
		params = &GetLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLabels", params, optFns, c.addOperationGetLabelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLabelsInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name of the label or labels to get.
	Name *string

	// The next token for the subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetLabelsOutput struct {

	// An array of labels.
	Labels []types.Label

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLabelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLabels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLabels(options.Region), middleware.Before); err != nil {
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

// GetLabelsAPIClient is a client that implements the GetLabels operation.
type GetLabelsAPIClient interface {
	GetLabels(context.Context, *GetLabelsInput, ...func(*Options)) (*GetLabelsOutput, error)
}

var _ GetLabelsAPIClient = (*Client)(nil)

// GetLabelsPaginatorOptions is the paginator options for GetLabels
type GetLabelsPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetLabelsPaginator is a paginator for GetLabels
type GetLabelsPaginator struct {
	options   GetLabelsPaginatorOptions
	client    GetLabelsAPIClient
	params    *GetLabelsInput
	nextToken *string
	firstPage bool
}

// NewGetLabelsPaginator returns a new GetLabelsPaginator
func NewGetLabelsPaginator(client GetLabelsAPIClient, params *GetLabelsInput, optFns ...func(*GetLabelsPaginatorOptions)) *GetLabelsPaginator {
	if params == nil {
		params = &GetLabelsInput{}
	}

	options := GetLabelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetLabelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetLabelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetLabels page.
func (p *GetLabelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetLabelsOutput, error) {
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

	result, err := p.client.GetLabels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetLabels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetLabels",
	}
}
