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

func Patch_documents_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.Document
		
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
		url := fmt.Sprintf("%s/documents/%s", cfg.BaseURL, id)
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
		var result models.Document
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

func CreatePatch_documents_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_documents_id",
		mcp.WithDescription("Update a document"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the document to be updated.")),
		mcp.WithString("expiryDate", mcp.Description("Input parameter: The expiry date of the document, in YYYY-MM-DD format.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Input parameter: The unique identifier of the document.")),
		mcp.WithString("issuerCountry", mcp.Description("Input parameter: The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**.")),
		mcp.WithString("modificationDate", mcp.Description("Input parameter: The modification date of the document.")),
		mcp.WithString("creationDate", mcp.Description("Input parameter: The creation date of the document.")),
		mcp.WithString("fileName", mcp.Description("Input parameter: The filename of the document.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Input parameter: Type of document, used when providing an ID number or uploading a document. The possible values depend on the legal entity type.\n\nWhen providing ID numbers:\n* For **individual**, the `type` values can be **driversLicense**, **identityCard**, **nationalIdNumber**, or **passport**.\n\nWhen uploading documents:\n* For **organization**, the `type` values can be **proofOfAddress**, **registrationDocument**, **taxDocument**, **proofOfOwnership**, or **proofOfIndustry**. \n\n* For **individual**, the `type` values can be **identityCard**, **driversLicense**, **proofOfNationalIdNumber**, **proofOfResidency**, or **proofOfIndustry**.\n\n* Use **bankStatement** to upload documents for a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).")),
		mcp.WithArray("attachments", mcp.Required(), mcp.Description("Input parameter: Array that contains the document. The array supports multiple attachments for uploading different sides or pages of a document.")),
		mcp.WithString("issuerState", mcp.Description("Input parameter: The state or province where the document was issued (AU only).")),
		mcp.WithString("number", mcp.Description("Input parameter: The number in the document.")),
		mcp.WithObject("owner", mcp.Required(), mcp.Description("")),
		mcp.WithObject("attachment", mcp.Description("")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Input parameter: Your description for the document.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_documents_idHandler(cfg),
	}
}
