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

func Post_paymentinstrumentsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.PaymentInstrumentInfo
		
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
		url := fmt.Sprintf("%s/paymentInstruments", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.PaymentInstrument
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

func CreatePost_paymentinstrumentsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_paymentInstruments",
		mcp.WithDescription("Create a payment instrument"),
		mcp.WithString("reference", mcp.Description("Input parameter: Your reference for the payment instrument, maximum 150 characters.")),
		mcp.WithString("status", mcp.Description("Input parameter: The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **Active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **Requested**.\n\nPossible values: \n\n * **Active**:  The payment instrument is active and can be used to make payments. \n\n * **Requested**: The payment instrument has been requested. This state is applicable for physical cards. \n\n* **Inactive**: The payment instrument is inactive and cannot be used to make payments. \n\n * **Suspended**: The payment instrument is temporarily suspended and cannot be used to make payments. \n\n * **Closed**: The payment instrument is permanently closed. This action cannot be undone. \n\n* **Stolen** \n\n * **Lost**\n\n ")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Input parameter: Type of payment instrument.\n\nPossible value: **card**, **bankAccount**. ")),
		mcp.WithString("balanceAccountId", mcp.Required(), mcp.Description("Input parameter: The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument.")),
		mcp.WithObject("card", mcp.Description("")),
		mcp.WithString("description", mcp.Description("Input parameter: Your description for the payment instrument, maximum 300 characters.")),
		mcp.WithString("issuingCountryCode", mcp.Required(), mcp.Description("Input parameter: The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**.")),
		mcp.WithString("paymentInstrumentGroupId", mcp.Description("Input parameter: The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_paymentinstrumentsHandler(cfg),
	}
}
