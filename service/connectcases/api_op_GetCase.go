// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a specific case if it exists.
func (c *Client) GetCase(ctx context.Context, params *GetCaseInput, optFns ...func(*Options)) (*GetCaseOutput, error) {
	if params == nil {
		params = &GetCaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCase", params, optFns, c.addOperationGetCaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCaseInput struct {

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// A list of unique field identifiers.
	//
	// This member is required.
	Fields []types.FieldIdentifier

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCaseOutput struct {

	// A list of detailed field information.
	//
	// This member is required.
	Fields []types.FieldValue

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	// The token for the next set of results. This is null if there are no more
	// results to return.
	NextToken *string

	// A map of of key-value pairs that represent tags on a resource. Tags are used to
	// organize, track, or control access for this resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCase{}, middleware.After)
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
	if err = addOpGetCaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCase(options.Region), middleware.Before); err != nil {
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

// GetCaseAPIClient is a client that implements the GetCase operation.
type GetCaseAPIClient interface {
	GetCase(context.Context, *GetCaseInput, ...func(*Options)) (*GetCaseOutput, error)
}

var _ GetCaseAPIClient = (*Client)(nil)

// GetCasePaginatorOptions is the paginator options for GetCase
type GetCasePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCasePaginator is a paginator for GetCase
type GetCasePaginator struct {
	options   GetCasePaginatorOptions
	client    GetCaseAPIClient
	params    *GetCaseInput
	nextToken *string
	firstPage bool
}

// NewGetCasePaginator returns a new GetCasePaginator
func NewGetCasePaginator(client GetCaseAPIClient, params *GetCaseInput, optFns ...func(*GetCasePaginatorOptions)) *GetCasePaginator {
	if params == nil {
		params = &GetCaseInput{}
	}

	options := GetCasePaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCasePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCasePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCase page.
func (p *GetCasePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCaseOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetCase(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cases",
		OperationName: "GetCase",
	}
}
