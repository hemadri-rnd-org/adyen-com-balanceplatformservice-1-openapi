package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/configuration-api/mcp-server/config"
	"github.com/configuration-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_accountholders_id_balanceaccountsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/accountHolders/%s/balanceAccounts%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.PaginatedBalanceAccountsResponse
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

func CreateGet_accountholders_id_balanceaccountsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_accountHolders_id_balanceAccounts",
		mcp.WithDescription("Get all balance accounts of an account holder"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the account holder.")),
		mcp.WithNumber("offset", mcp.Description("The number of items that you want to skip.")),
		mcp.WithNumber("limit", mcp.Description("The number of items returned per page, maximum 100 items. By default, the response returns 10 items per page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_accountholders_id_balanceaccountsHandler(cfg),
	}
}
