package main

import (
	"github.com/configuration-api/mcp-server/config"
	"github.com/configuration-api/mcp-server/models"
	tools_transaction_rules "github.com/configuration-api/mcp-server/tools/transaction_rules"
	tools_platform "github.com/configuration-api/mcp-server/tools/platform"
	tools_documents "github.com/configuration-api/mcp-server/tools/documents"
	tools_payment_instruments "github.com/configuration-api/mcp-server/tools/payment_instruments"
	tools_balance_accounts "github.com/configuration-api/mcp-server/tools/balance_accounts"
	tools_payment_instrument_groups "github.com/configuration-api/mcp-server/tools/payment_instrument_groups"
	tools_bank_account_validation "github.com/configuration-api/mcp-server/tools/bank_account_validation"
	tools_account_holders "github.com/configuration-api/mcp-server/tools/account_holders"
	tools_transfer_instruments "github.com/configuration-api/mcp-server/tools/transfer_instruments"
	tools_legal_entities "github.com/configuration-api/mcp-server/tools/legal_entities"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_transaction_rules.CreateDelete_transactionrules_transactionruleidTool(cfg),
		tools_transaction_rules.CreateGet_transactionrules_transactionruleidTool(cfg),
		tools_transaction_rules.CreatePatch_transactionrules_transactionruleidTool(cfg),
		tools_platform.CreateGet_balanceplatforms_idTool(cfg),
		tools_documents.CreateDelete_documents_idTool(cfg),
		tools_documents.CreateGet_documents_idTool(cfg),
		tools_documents.CreatePatch_documents_idTool(cfg),
		tools_payment_instruments.CreateGet_paymentinstruments_id_transactionrulesTool(cfg),
		tools_balance_accounts.CreateGet_balanceaccounts_idTool(cfg),
		tools_balance_accounts.CreatePatch_balanceaccounts_idTool(cfg),
		tools_payment_instrument_groups.CreateGet_paymentinstrumentgroups_idTool(cfg),
		tools_balance_accounts.CreatePost_balanceaccountsTool(cfg),
		tools_balance_accounts.CreateGet_balanceaccounts_id_paymentinstrumentsTool(cfg),
		tools_documents.CreatePost_documentsTool(cfg),
		tools_platform.CreateGet_balanceplatforms_id_accountholdersTool(cfg),
		tools_payment_instrument_groups.CreatePost_paymentinstrumentgroupsTool(cfg),
		tools_transaction_rules.CreatePost_transactionrulesTool(cfg),
		tools_payment_instrument_groups.CreateGet_paymentinstrumentgroups_id_transactionrulesTool(cfg),
		tools_bank_account_validation.CreatePost_validatebankaccountidentificationTool(cfg),
		tools_account_holders.CreatePost_accountholdersTool(cfg),
		tools_payment_instruments.CreateGet_paymentinstruments_id_revealTool(cfg),
		tools_transfer_instruments.CreateGet_transferinstruments_idTool(cfg),
		tools_transfer_instruments.CreatePatch_transferinstruments_idTool(cfg),
		tools_transfer_instruments.CreateDelete_transferinstruments_idTool(cfg),
		tools_transfer_instruments.CreatePost_transferinstrumentsTool(cfg),
		tools_account_holders.CreatePatch_accountholders_idTool(cfg),
		tools_account_holders.CreateGet_accountholders_idTool(cfg),
		tools_legal_entities.CreateGet_legalentities_idTool(cfg),
		tools_legal_entities.CreatePatch_legalentities_idTool(cfg),
		tools_payment_instruments.CreatePost_paymentinstrumentsTool(cfg),
		tools_legal_entities.CreatePost_legalentitiesTool(cfg),
		tools_account_holders.CreateGet_accountholders_id_balanceaccountsTool(cfg),
		tools_payment_instruments.CreateGet_paymentinstruments_idTool(cfg),
		tools_payment_instruments.CreatePatch_paymentinstruments_idTool(cfg),
	}
}
