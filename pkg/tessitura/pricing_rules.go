package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
)

type APIPricingRuleMessageTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
}

type APICriterionOperatorSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
	Symbol      *string `json:"Symbol"`
}

type APIPricingRuleCategorySummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
}

type APIPricingRuleTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
}

type APIPricingRuleMessage struct {
	Id                    *int                              `json:"Id"`
	Message               *string                          `json:"Message"`
	Inactive              bool                             `json:"Inactive"`
	IsFromMessageOnlyRule bool                             `json:"IsFromMessageOnlyRule"`
	MessageType           APIPricingRuleMessageTypeSummary `json:"MessageType"`
	PricingRule           APIEntitySummary                 `json:"PricingRule"`
	CreatedBy             *string                          `json:"CreatedBy"`
	CreatedDateTime       *string                          `json:"CreatedDateTime"`
	CreateLocation        *string                          `json:"CreateLocation"`
	UpdatedBy             *string                          `json:"UpdatedBy"`
	UpdatedDateTime       *string                          `json:"UpdatedDateTime"`
}

func (c *Client) GetPricingRuleMessages(ctx context.Context, pricingRuleID *int) ([]APIPricingRuleMessage, error) {
	path := "/TXN/PricingRuleMessage"
	if pricingRuleID != nil {
		path = fmt.Sprintf("%s?pricingRuleId=%d", path, *pricingRuleID)
	}
	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rule messages: %w", err)
	}
	var items []APIPricingRuleMessage
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rule messages: %w", err)
	}
	return items, nil
}

type APIPricingRule struct {
	Id                                             *int                          `json:"Id"`
	Description                                    *string                       `json:"Description"`
	Inactive                                       bool                          `json:"Inactive"`
	EditIndicator                                  bool                          `json:"EditIndicator"`
	DiscountAmount                                 float64                       `json:"DiscountAmount"`
	DiscountIsPercent                              bool                          `json:"DiscountIsPercent"`
	DiscountPercentRound                           int                           `json:"DiscountPercentRound"`
	DiscountPriceTypeId                            *int                          `json:"DiscountPriceTypeId"`
	RespectMinimumPrice                            bool                          `json:"RespectMinimumPrice"`
	ApplyOncePerOrder                              bool                          `json:"ApplyOncePerOrder"`
	ApplyToLowestCommonNumberOfSeats               bool                          `json:"ApplyToLowestCommonNumberOfSeats"`
	ExcludeGeneralPublic                           bool                          `json:"ExcludeGeneralPublic"`
	RuleAction                                     int                           `json:"RuleAction"`
	MaxSeats                                       *int                          `json:"MaxSeats"`
	StartDateTime                                  *string                       `json:"StartDateTime"`
	EndDateTime                                    *string                       `json:"EndDateTime"`
	OverTheLimitDateTime                           *string                       `json:"OverTheLimitDateTime"`
	ManualIndicator                                *string                       `json:"ManualIndicator"`
	AdditionalTime                                 *int                          `json:"AdditionalTime"`
	AdditionalTimeInDaysIndicator                  bool                          `json:"AdditionalTimeInDaysIndicator"`
	RuleCategory                                   APIPricingRuleCategorySummary `json:"RuleCategory"`
	RuleType                                       APIPricingRuleTypeSummary    `json:"RuleType"`
	ConstituentAttribute                           *int                          `json:"ConstituentAttribute"`
	ConstituentAttributeValue1                     *string                       `json:"ConstituentAttributeValue1"`
	ConstituentAttributeValue2                     *string                       `json:"ConstituentAttributeValue2"`
	ConstituentAttributeOperator                   APICriterionOperatorSummary  `json:"ConstituentAttributeOperator"`
	ConstituentListId                              *int                          `json:"ConstituentListId"`
	ConstituentExclusionListId                     *int                          `json:"ConstituentExclusionListId"`
	ConstituentRankType                            *int                          `json:"ConstituentRankType"`
	ConstituentRankingValue1                       *int                          `json:"ConstituentRankingValue1"`
	ConstituentRankingValue2                       *int                          `json:"ConstituentRankingValue2"`
	ConstituentRankingOperator                     APICriterionOperatorSummary  `json:"ConstituentRankingOperator"`
	Constituencies                                 *string                       `json:"Constituencies"`
	Appeals                                        *string                       `json:"Appeals"`
	PromotedAppeals                                *string                       `json:"PromotedAppeals"`
	PromotedSources                                *string                       `json:"PromotedSources"`
	Sources                                        *string                       `json:"Sources"`
	QualifyingPackage                              *string                       `json:"QualifyingPackage"`
	QualifyingPerformance                          *string                       `json:"QualifyingPerformance"`
	QualifyingPriceType1                           *string                       `json:"QualifyingPriceType1"`
	QualifyingPriceType2                           *string                       `json:"QualifyingPriceType2"`
	QualifyingProductionSeason                     *string                       `json:"QualifyingProductionSeason"`
	QualifyingSeasonPackageType                    *string                       `json:"QualifyingSeasonPackageType"`
	QualifyingSeatCount1Value1                     *int                          `json:"QualifyingSeatCount1Value1"`
	QualifyingSeatCount1Value2                     *int                          `json:"QualifyingSeatCount1Value2"`
	QualifyingSeatCount2Value1                     *int                          `json:"QualifyingSeatCount2Value1"`
	QualifyingSeatCount2Value2                     *int                          `json:"QualifyingSeatCount2Value2"`
	QualifyingZone                                 *string                       `json:"QualifyingZone"`
	QualifyingLevel                                *string                       `json:"QualifyingLevel"`
	QualifyingPeriod                               *string                       `json:"QualifyingPeriod"`
	ResultPackage                                  *string                       `json:"ResultPackage"`
	ResultPerformance                              *string                       `json:"ResultPerformance"`
	ResultPriceType                                *string                       `json:"ResultPriceType"`
	ResultProductionSeason                         *string                       `json:"ResultProductionSeason"`
	ResultSeasonPackageType                        *string                       `json:"ResultSeasonPackageType"`
	ResultSeatCount                                *int                          `json:"ResultSeatCount"`
	ResultMaximumSeats                             *int                          `json:"ResultMaximumSeats"`
	ResultMaximumSeats2                            *int                          `json:"ResultMaximumSeats2"`
	ResultZone                                     *string                       `json:"ResultZone"`
	ApplyToPerformancesWithinPackagesIndicator1    bool                          `json:"ApplyToPerformancesWithinPackagesIndicator1"`
	ApplyToPerformancesWithinPackagesIndicator2    bool                          `json:"ApplyToPerformancesWithinPackagesIndicator2"`
	Messages                                       []APIPricingRuleMessage       `json:"Messages"`
	CreatedBy                                      *string                       `json:"CreatedBy"`
	CreatedDateTime                                *string                       `json:"CreatedDateTime"`
	CreateLocation                                 *string                       `json:"CreateLocation"`
	UpdatedBy                                      *string                       `json:"UpdatedBy"`
	UpdatedDateTime                                *string                       `json:"UpdatedDateTime"`
}

type GetPricingRulesParams struct {
	PerformanceIDs string
	PackageIDs     string
	OrderDate      string
	ModeOfSaleID   *int
}

func (c *Client) GetPricingRules(ctx context.Context, params *GetPricingRulesParams) ([]APIPricingRule, error) {
	path := "/TXN/PricingRules"
	if params != nil {
		sep := "?"
		if params.PerformanceIDs != "" {
			path += sep + "performanceIds=" + params.PerformanceIDs
			sep = "&"
		}
		if params.PackageIDs != "" {
			path += sep + "packageIds=" + params.PackageIDs
			sep = "&"
		}
		if params.OrderDate != "" {
			path += sep + "orderDate=" + params.OrderDate
			sep = "&"
		}
		if params.ModeOfSaleID != nil {
			path += fmt.Sprintf("%smodeOfSaleId=%d", sep, *params.ModeOfSaleID)
		}
	}
	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rules: %w", err)
	}
	var items []APIPricingRule
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rules: %w", err)
	}
	return items, nil
}

type APIPricingRuleSummary struct {
	Id           *int                          `json:"Id"`
	Description  *string                       `json:"Description"`
	Inactive     bool                          `json:"Inactive"`
	RuleAction   int                           `json:"RuleAction"`
	StartDateTime *string                      `json:"StartDateTime"`
	EndDateTime  *string                       `json:"EndDateTime"`
	RuleType     APIPricingRuleTypeSummary     `json:"RuleType"`
	RuleCategory APIPricingRuleCategorySummary `json:"RuleCategory"`
}

type APIPricingRuleSetMap struct {
	Id              *int                  `json:"Id"`
	Rank            int                   `json:"Rank"`
	PricingRule     APIPricingRuleSummary `json:"PricingRule"`
	CreatedBy       *string              `json:"CreatedBy"`
	CreatedDateTime *string              `json:"CreatedDateTime"`
	CreateLocation  *string              `json:"CreateLocation"`
	UpdatedBy       *string              `json:"UpdatedBy"`
	UpdatedDateTime *string              `json:"UpdatedDateTime"`
}

func (c *Client) GetPricingRuleSetPricingRules(ctx context.Context, pricingRuleSetID *int) ([]APIPricingRuleSetMap, error) {
	path := "/TXN/PricingRuleSetPricingRules"
	if pricingRuleSetID != nil {
		path = fmt.Sprintf("%s?pricingRuleSetId=%d", path, *pricingRuleSetID)
	}
	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rule set pricing rules: %w", err)
	}
	var items []APIPricingRuleSetMap
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rule set pricing rules: %w", err)
	}
	return items, nil
}

type APIPricingRuleSet struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	EditIndicator   bool                `json:"EditIndicator"`
	ControlGroup    APIControlGroupRef  `json:"ControlGroup"`
	Rules           []APIPricingRuleSetMap `json:"Rules"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPricingRuleSets(ctx context.Context) ([]APIPricingRuleSet, error) {
	data, err := c.Get(ctx, "/TXN/PricingRuleSets")
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rule sets: %w", err)
	}
	var items []APIPricingRuleSet
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rule sets: %w", err)
	}
	return items, nil
}
