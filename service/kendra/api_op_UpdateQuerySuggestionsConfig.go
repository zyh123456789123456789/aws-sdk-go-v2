// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the settings of query suggestions for an index. Amazon Kendra supports
// partial updates, so you only need to provide the fields you want to update. If
// an update is currently processing (i.e. 'happening'), you need to wait for the
// update to finish before making another update. Updates to query suggestions
// settings might not take effect right away. The time for your updated settings to
// take effect depends on the updates made and the number of search queries in your
// index. You can still enable/disable query suggestions at any time.
// UpdateQuerySuggestionsConfig is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func (c *Client) UpdateQuerySuggestionsConfig(ctx context.Context, params *UpdateQuerySuggestionsConfigInput, optFns ...func(*Options)) (*UpdateQuerySuggestionsConfigOutput, error) {
	if params == nil {
		params = &UpdateQuerySuggestionsConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQuerySuggestionsConfig", params, optFns, c.addOperationUpdateQuerySuggestionsConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQuerySuggestionsConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQuerySuggestionsConfigInput struct {

	// The identifier of the index with query suggestions you want to update.
	//
	// This member is required.
	IndexId *string

	// TRUE to include queries without user information (i.e. all queries,
	// irrespective of the user), otherwise FALSE to only include queries with user
	// information. If you pass user information to Amazon Kendra along with the
	// queries, you can set this flag to FALSE and instruct Amazon Kendra to only
	// consider queries with user information. If you set to FALSE , Amazon Kendra only
	// considers queries searched at least MinimumQueryCount times across
	// MinimumNumberOfQueryingUsers unique users for suggestions. If you set to TRUE ,
	// Amazon Kendra ignores all user information and learns from all queries.
	IncludeQueriesWithoutUserInformation *bool

	// The minimum number of unique users who must search a query in order for the
	// query to be eligible to suggest to your users. Increasing this number might
	// decrease the number of suggestions. However, this ensures a query is searched by
	// many users and is truly popular to suggest to users. How you tune this setting
	// depends on your specific needs.
	MinimumNumberOfQueryingUsers *int32

	// The the minimum number of times a query must be searched in order to be
	// eligible to suggest to your users. Decreasing this number increases the number
	// of suggestions. However, this affects the quality of suggestions as it sets a
	// low bar for a query to be considered popular to suggest to users. How you tune
	// this setting depends on your specific needs.
	MinimumQueryCount *int32

	// Set the mode to ENABLED or LEARN_ONLY . By default, Amazon Kendra enables query
	// suggestions. LEARN_ONLY mode allows you to turn off query suggestions. You can
	// to update this at any time. In LEARN_ONLY mode, Amazon Kendra continues to
	// learn from new queries to keep suggestions up to date for when you are ready to
	// switch to ENABLED mode again.
	Mode types.Mode

	// How recent your queries are in your query log time window. The time window is
	// the number of days from current day to past days. By default, Amazon Kendra sets
	// this to 180.
	QueryLogLookBackWindowInDays *int32

	noSmithyDocumentSerde
}

type UpdateQuerySuggestionsConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateQuerySuggestionsConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateQuerySuggestionsConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateQuerySuggestionsConfig{}, middleware.After)
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
	if err = addOpUpdateQuerySuggestionsConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQuerySuggestionsConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQuerySuggestionsConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "UpdateQuerySuggestionsConfig",
	}
}
