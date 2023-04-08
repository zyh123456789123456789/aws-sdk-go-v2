// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a list of inventory type names for the account, or return a list of
// attribute names for a specific Inventory item type.
func (c *Client) GetInventorySchema(ctx context.Context, params *GetInventorySchemaInput, optFns ...func(*Options)) (*GetInventorySchemaOutput, error) {
	if params == nil {
		params = &GetInventorySchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInventorySchema", params, optFns, c.addOperationGetInventorySchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInventorySchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInventorySchemaInput struct {

	// Returns inventory schemas that support aggregation. For example, this call
	// returns the AWS:InstanceInformation type, because it supports aggregation based
	// on the PlatformName , PlatformType , and PlatformVersion attributes.
	Aggregator bool

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Returns the sub-type schema for a specified inventory type.
	SubType *bool

	// The type of inventory item to return.
	TypeName *string

	noSmithyDocumentSerde
}

type GetInventorySchemaOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Inventory schemas returned by the request.
	Schemas []types.InventoryItemSchema

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInventorySchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInventorySchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInventorySchema{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInventorySchema(options.Region), middleware.Before); err != nil {
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

// GetInventorySchemaAPIClient is a client that implements the GetInventorySchema
// operation.
type GetInventorySchemaAPIClient interface {
	GetInventorySchema(context.Context, *GetInventorySchemaInput, ...func(*Options)) (*GetInventorySchemaOutput, error)
}

var _ GetInventorySchemaAPIClient = (*Client)(nil)

// GetInventorySchemaPaginatorOptions is the paginator options for
// GetInventorySchema
type GetInventorySchemaPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInventorySchemaPaginator is a paginator for GetInventorySchema
type GetInventorySchemaPaginator struct {
	options   GetInventorySchemaPaginatorOptions
	client    GetInventorySchemaAPIClient
	params    *GetInventorySchemaInput
	nextToken *string
	firstPage bool
}

// NewGetInventorySchemaPaginator returns a new GetInventorySchemaPaginator
func NewGetInventorySchemaPaginator(client GetInventorySchemaAPIClient, params *GetInventorySchemaInput, optFns ...func(*GetInventorySchemaPaginatorOptions)) *GetInventorySchemaPaginator {
	if params == nil {
		params = &GetInventorySchemaInput{}
	}

	options := GetInventorySchemaPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInventorySchemaPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInventorySchemaPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetInventorySchema page.
func (p *GetInventorySchemaPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInventorySchemaOutput, error) {
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

	result, err := p.client.GetInventorySchema(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInventorySchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetInventorySchema",
	}
}
