package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/configuration-api/mcp-server/config"
	"github.com/configuration-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Patch_transactionrules_transactionruleidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		transactionRuleIdVal, ok := args["transactionRuleId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: transactionRuleId"), nil
		}
		transactionRuleId, ok := transactionRuleIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: transactionRuleId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.TransactionRuleInfo
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/transactionRules/%s", cfg.BaseURL, transactionRuleId)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.TransactionRule
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePatch_transactionrules_transactionruleidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_transactionRules_transactionRuleId",
		mcp.WithDescription("Update a transaction rule"),
		mcp.WithString("transactionRuleId", mcp.Required(), mcp.Description("The unique identifier of the transaction rule.")),
		mcp.WithArray("mccs", mcp.Description("Input parameter: List of Merchant Category Codes (MCCs) to which the rule applies.")),
		mcp.WithString("balancePlatformId", mcp.Description("Input parameter: The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/get/balancePlatforms/{id}__queryParam_id) to which the rule applies.")),
		mcp.WithObject("interval", mcp.Required(), mcp.Description("")),
		mcp.WithString("reference", mcp.Required(), mcp.Description("Input parameter: Your reference for the transaction rule, maximum 150 characters.")),
		mcp.WithString("endDate", mcp.Description("Input parameter: The date when the rule will stop being evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.\n\nIf not provided, the rule will be evaluated until the rule status is set to **inactive**.")),
		mcp.WithNumber("maxTransactions", mcp.Description("Input parameter: The maximum number of transactions that a payment instrument can be used for in the period specified in the `interval.type`.")),
		mcp.WithString("paymentInstrumentId", mcp.Description("Input parameter: The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstruments__resParam_id) to which the rule applies.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Input parameter: Type of conditions provided in the rule.\n\nPossible values:\n  * **allowList**: The rule provides categories (such as country and MCC) where payments must be allowed.\n * **blockList**: The rule provides categories (such as country and MCC) where payments must be blocked.\n * **maxUsage**: The rule sets limits for the maximum amount or maximum number of transactions for the lifetime of the payment instrument.\n * **velocity**: The rule sets limits for the maximum amount or maximum number of transactions for a given time interval.\n")),
		mcp.WithObject("amount", mcp.Description("")),
		mcp.WithArray("countries", mcp.Description("Input parameter: List of countries to which the rule applies. For example: **[\"NL\",\"US\"]**.")),
		mcp.WithArray("entryModes", mcp.Description("Input parameter: List of point-of-sale entry modes to which the rule applies.\n\nPossible values: **manual**, **chip**, **magstripe**, **contactless**, **cof**, **token**, **server**, **barcode**, **ocr**.")),
		mcp.WithString("startDate", mcp.Description("Input parameter: The date when the rule will start to be evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.\n\nIf not provided when creating a transaction rule, the `startDate` is set to the date when the rule status is set to **active**. \n\n")),
		mcp.WithString("paymentInstrumentGroupId", mcp.Description("Input parameter: The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstrumentGroups__resParam_id) to which the rule applies.")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Input parameter: Your description for the transaction rule, maximum 300 characters.")),
		mcp.WithArray("processingTypes", mcp.Description("Input parameter: List of processing types to which the rule applies.\n\nPossible values: **atmWithdraw**, **pos**, **ecommerce**, **moto**, **recurring**, **balanceInquiry**.")),
		mcp.WithString("status", mcp.Description("Input parameter: The status of the transaction rule. If you provide a `startDate` in the request, the rule is automatically created \nwith an **active** status. \n\nPossible values: **active**, **inactive**.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_transactionrules_transactionruleidHandler(cfg),
	}
}
