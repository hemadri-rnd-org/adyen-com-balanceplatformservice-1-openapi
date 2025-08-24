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

func Patch_balanceaccounts_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.BalanceAccountUpdateRequest
		
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
		url := fmt.Sprintf("%s/balanceAccounts/%s", cfg.BaseURL, id)
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
		var result models.BalanceAccount
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

func CreatePatch_balanceaccounts_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_balanceAccounts_id",
		mcp.WithDescription("Update a balance account"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the balance account.")),
		mcp.WithString("accountHolderId", mcp.Description("Input parameter: The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account.")),
		mcp.WithString("defaultCurrencyCode", mcp.Description("Input parameter: The default currency code of this balance account, in three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) format. \nThe default value is **EUR**.")),
		mcp.WithString("description", mcp.Description("Input parameter: A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder.")),
		mcp.WithString("reference", mcp.Description("Input parameter: Your reference to the balance account, maximum 150 characters.")),
		mcp.WithString("status", mcp.Description("Input parameter: The status of the balance account. Payment instruments linked to the balance account can only be used if the balance account status is **Active**.\n\nPossible values: **Active**, **Inactive**, **Closed**, **Suspended**.")),
		mcp.WithObject("sweepConfigurations", mcp.Description("Input parameter: Contains key-value pairs that specify [balance sweep per currency code](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__reqParam_sweepConfigurations).\n\nYou can update the balance account to add, update, or delete sweeps.\n\n* To add a sweep, send the currency code as a key and the configuration as the object.\n\n * To update a sweep, send the whole configuration with your updates.\n\n* To delete a sweep, set the value to **null**. For example, `\"EUR\": null`.")),
		mcp.WithString("timeZone", mcp.Description("Input parameter: The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**.\nDefaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_balanceaccounts_idHandler(cfg),
	}
}
