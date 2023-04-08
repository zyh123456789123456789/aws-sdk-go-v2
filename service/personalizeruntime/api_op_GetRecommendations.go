// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of recommended items. For campaigns, the campaign's Amazon
// Resource Name (ARN) is required and the required user and item input depends on
// the recipe type used to create the solution backing the campaign as follows:
//   - USER_PERSONALIZATION - userId required, itemId not used
//   - RELATED_ITEMS - itemId required, userId not used
//
// Campaigns that are backed by a solution created using a recipe of type
// PERSONALIZED_RANKING use the API. For recommenders, the recommender's ARN is
// required and the required item and user input depends on the use case
// (domain-based recipe) backing the recommender. For information on use case
// requirements see Choosing recommender use cases (https://docs.aws.amazon.com/personalize/latest/dg/domain-use-cases.html)
// .
func (c *Client) GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) {
	if params == nil {
		params = &GetRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendations", params, optFns, c.addOperationGetRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationsInput struct {

	// The Amazon Resource Name (ARN) of the campaign to use for getting
	// recommendations.
	CampaignArn *string

	// The contextual metadata to use when getting recommendations. Contextual
	// metadata includes any interaction information that might be relevant when
	// getting a user's recommendations, such as the user's current location or device
	// type.
	Context map[string]string

	// The ARN of the filter to apply to the returned recommendations. For more
	// information, see Filtering Recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter.html)
	// . When using this parameter, be sure the filter resource is ACTIVE .
	FilterArn *string

	// The values to use when filtering recommendations. For each placeholder
	// parameter in your filter expression, provide the parameter name (in matching
	// case) as a key and the filter value(s) as the corresponding value. Separate
	// multiple values for one parameter with a comma. For filter expressions that use
	// an INCLUDE element to include items, you must provide values for all parameters
	// that are defined in the expression. For filters with expressions that use an
	// EXCLUDE element to exclude items, you can omit the filter-values .In this case,
	// Amazon Personalize doesn't use that portion of the expression to filter
	// recommendations. For more information, see Filtering recommendations and user
	// segments (https://docs.aws.amazon.com/personalize/latest/dg/filter.html) .
	FilterValues map[string]string

	// The item ID to provide recommendations for. Required for RELATED_ITEMS recipe
	// type.
	ItemId *string

	// The number of results to return. The default is 25. The maximum is 500.
	NumResults int32

	// The promotions to apply to the recommendation request. A promotion defines
	// additional business rules that apply to a configurable subset of recommended
	// items.
	Promotions []types.Promotion

	// The Amazon Resource Name (ARN) of the recommender to use to get
	// recommendations. Provide a recommender ARN if you created a Domain dataset group
	// with a recommender for a domain use case.
	RecommenderArn *string

	// The user ID to provide recommendations for. Required for USER_PERSONALIZATION
	// recipe type.
	UserId *string

	noSmithyDocumentSerde
}

type GetRecommendationsOutput struct {

	// A list of recommendations sorted in descending order by prediction score. There
	// can be a maximum of 500 items in the list.
	ItemList []types.PredictedItem

	// The ID of the recommendation.
	RecommendationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecommendations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "GetRecommendations",
	}
}
