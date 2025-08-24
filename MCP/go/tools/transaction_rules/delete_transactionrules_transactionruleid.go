package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/configuration-api/mcp-server/config"
	"github.com/configuration-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Delete_transactionrules_transactionruleidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/transactionRules/%s", cfg.BaseURL, transactionRuleId)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDelete_transactionrules_transactionruleidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_transactionRules_transactionRuleId",
		mcp.WithDescription("Delete a transaction rule"),
		mcp.WithString("transactionRuleId", mcp.Required(), mcp.Description("The unique identifier of the transaction rule.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_transactionrules_transactionruleidHandler(cfg),
	}
}
