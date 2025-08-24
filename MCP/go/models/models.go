package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PaginatedBalanceAccountsResponse represents the PaginatedBalanceAccountsResponse schema from the OpenAPI specification
type PaginatedBalanceAccountsResponse struct {
	Balanceaccounts []BalanceAccount `json:"balanceAccounts"` // List of balance accounts.
	Hasnext bool `json:"hasNext"` // Indicates whether there are more items on the next page.
	Hasprevious bool `json:"hasPrevious"` // Indicates whether there are more items on the previous page.
}

// PaymentInstrumentReference represents the PaymentInstrumentReference schema from the OpenAPI specification
type PaymentInstrumentReference struct {
	Id string `json:"id"` // The unique identifier of the payment instrument.
}

// Attachment represents the Attachment schema from the OpenAPI specification
type Attachment struct {
	Pagetype string `json:"pageType,omitempty"` // Specifies which side of the ID card is uploaded. * When `type` is **driversLicense** or **identityCard**, set this to **front** or **back**. * When omitted, we infer the page number based on the order of attachments.
	Content string `json:"content"` // The document in Base64-encoded string format.
	Contenttype string `json:"contentType,omitempty"` // The file format. Possible values: **application/pdf**, **image/jpg**, **image/jpeg**, **image/png**.
	Filename string `json:"filename,omitempty"` // The name of the file including the file extension.
}

// NOLocalAccountIdentification represents the NOLocalAccountIdentification schema from the OpenAPI specification
type NOLocalAccountIdentification struct {
	Accountnumber string `json:"accountNumber"` // The 11-digit bank account number, without separators or whitespace.
	TypeField string `json:"type"` // **noLocal**
}

// Name represents the Name schema from the OpenAPI specification
type Name struct {
	Firstname string `json:"firstName"` // The first name.
	Lastname string `json:"lastName"` // The last name.
}

// AULocalAccountIdentification represents the AULocalAccountIdentification schema from the OpenAPI specification
type AULocalAccountIdentification struct {
	TypeField string `json:"type"` // **auLocal**
	Accountnumber string `json:"accountNumber"` // The bank account number, without separators or whitespace.
	Bsbcode string `json:"bsbCode"` // The 6-digit [Bank State Branch (BSB) code](https://en.wikipedia.org/wiki/Bank_state_branch), without separators or whitespace.
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Personaldata PersonalData `json:"personalData,omitempty"`
	Phonenumber PhoneNumber `json:"phoneNumber,omitempty"`
	Webaddress string `json:"webAddress,omitempty"` // The URL of the website of the contact.
	Address Address `json:"address,omitempty"`
	Email string `json:"email,omitempty"` // The e-mail address of the contact.
	Fullphonenumber string `json:"fullPhoneNumber,omitempty"` // The phone number of the contact provided as a single string. It will be handled as a landline phone. **Examples:** "0031 6 11 22 33 44", "+316/1122-3344", "(0031) 611223344"
	Name Name `json:"name,omitempty"`
}

// LegalEntityCapability represents the LegalEntityCapability schema from the OpenAPI specification
type LegalEntityCapability struct {
	Problems []CapabilityProblem `json:"problems,omitempty"` // Contains verification errors and the actions that you can take to resolve them.
	Requestedlevel string `json:"requestedLevel,omitempty"` // The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring. Possible values: **notApplicable**, **low**, **medium**, **high**.
	Transferinstruments []SupportingEntityCapability `json:"transferInstruments,omitempty"` // Capability status for transfer instruments associated with legal entity
	Requested bool `json:"requested,omitempty"` // Indicates whether the capability is requested. To check whether the Legal Entity is permitted to use the capability,
	Requestedsettings CapabilitySettings `json:"requestedSettings,omitempty"`
	Allowed bool `json:"allowed,omitempty"` // Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful
	Allowedlevel string `json:"allowedLevel,omitempty"` // The capability level that is allowed for the legal entity. Possible values: **notApplicable**, **low**, **medium**, **high**.
	Verificationstatus string `json:"verificationStatus,omitempty"` // The status of the verification checks for the capability. Possible values: * **pending**: Adyen is running the verification. * **invalid**: The verification failed. Check if the `errors` array contains more information. * **valid**: The verification has been successfully completed. * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	Allowedsettings CapabilitySettings `json:"allowedSettings,omitempty"`
}

// AccountHolderCapability represents the AccountHolderCapability schema from the OpenAPI specification
type AccountHolderCapability struct {
	Requested bool `json:"requested,omitempty"` // Indicates whether the capability is requested. To check whether the account holder is permitted to use the capability, refer to the `allowed` field.
	Requestedsettings JSONObject `json:"requestedSettings,omitempty"`
	Verificationstatus string `json:"verificationStatus,omitempty"` // The status of the verification checks for the capability. Possible values: * **pending**: Adyen is running the verification. * **invalid**: The verification failed. Check if the `errors` array contains more information. * **valid**: The verification has been successfully completed. * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	Allowedsettings JSONObject `json:"allowedSettings,omitempty"`
	Requestedlevel string `json:"requestedLevel,omitempty"` // The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring. Possible values: **notApplicable**, **low**, **medium**, **high**.
	Allowedlevel string `json:"allowedLevel,omitempty"` // The capability level that is allowed for the account holder. Possible values: **notApplicable**, **low**, **medium**, **high**.
	Problems []CapabilityProblem `json:"problems,omitempty"` // Contains verification errors and the actions that you can take to resolve them.
	Transferinstruments []AccountSupportingEntityCapability `json:"transferInstruments,omitempty"` // Contains the status of the transfer instruments associated with this capability.
	Enabled bool `json:"enabled,omitempty"` // Indicates whether the capability is enabled. If **false**, the capability is temporarily disabled for the account holder.
	Allowed bool `json:"allowed,omitempty"` // Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful and the account holder is permitted to use the capability.
}

// PaymentInstrumentUpdateRequest represents the PaymentInstrumentUpdateRequest schema from the OpenAPI specification
type PaymentInstrumentUpdateRequest struct {
	Card CardInfo `json:"card,omitempty"`
	Status string `json:"status,omitempty"` // The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **Active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **Requested**. Possible values: * **Active**: The payment instrument is active and can be used to make payments. * **Requested**: The payment instrument has been requested. This state is applicable for physical cards. * **Inactive**: The payment instrument is inactive and cannot be used to make payments. * **Suspended**: The payment instrument is temporarily suspended and cannot be used to make payments. * **Closed**: The payment instrument is permanently closed. This action cannot be undone. * **Stolen** * **Lost**
	Statuscomment string `json:"statusComment,omitempty"` // Comment for the status of the payment instrument.
	Balanceaccountid string `json:"balanceAccountId,omitempty"` // The unique identifier of the balance account associated with this payment instrument. >You can only change the balance account ID if the payment instrument has **Requested** or **Inactive** status.
}

// BalancePlatform represents the BalancePlatform schema from the OpenAPI specification
type BalancePlatform struct {
	Description string `json:"description,omitempty"` // Your description of the balance platform, maximum 300 characters.
	Id string `json:"id"` // The unique identifier of the balance platform.
	Status string `json:"status,omitempty"` // The status of the balance platform. Possible values: **Active**, **Inactive**, **Closed**, **Suspended**.
}

// BirthData represents the BirthData schema from the OpenAPI specification
type BirthData struct {
	Dateofbirth string `json:"dateOfBirth,omitempty"` // The individual's date of birth, in YYYY-MM-DD format.
}

// BulkAddress represents the BulkAddress schema from the OpenAPI specification
type BulkAddress struct {
	Company string `json:"company,omitempty"` // The name of the company.
	Country string `json:"country"` // The two-character ISO-3166-1 alpha-2 country code. For example, **US**.
	Mobile string `json:"mobile,omitempty"` // The full telephone number.
	Stateorprovince string `json:"stateOrProvince,omitempty"` // The two-letter ISO 3166-2 state or province code. Maximum length: 2 characters for addresses in the US.
	City string `json:"city,omitempty"` // The name of the city.
	Postalcode string `json:"postalCode,omitempty"` // The postal code. Maximum length: * 5 digits for addresses in the US. * 10 characters for all other countries.
	Street string `json:"street,omitempty"` // The streetname of the house.
	Email string `json:"email,omitempty"` // The email address.
	Housenumberorname string `json:"houseNumberOrName,omitempty"` // The house number or name.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// IdentificationData represents the IdentificationData schema from the OpenAPI specification
type IdentificationData struct {
	Number string `json:"number,omitempty"` // The number in the document.
	TypeField string `json:"type"` // Type of document, used when providing an ID number or uploading a document. The possible values depend on the legal entity type. When providing ID numbers: * For **individual**, the `type` values can be **driversLicense**, **identityCard**, **nationalIdNumber**, or **passport**. When uploading documents: * For **organization**, the `type` values can be **proofOfAddress**, **registrationDocument**, **taxDocument**, **proofOfOwnership**, or **proofOfIndustry**. * For **individual**, the `type` values can be **identityCard**, **driversLicense**, **proofOfNationalIdNumber**, **proofOfResidency**, or **proofOfIndustry**. * Use **bankStatement** to upload documents for a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).
	Cardnumber string `json:"cardNumber,omitempty"` // The card number of the document that was issued (AU only).
	Expirydate string `json:"expiryDate,omitempty"` // The expiry date of the document, in YYYY-MM-DD format.
	Issuercountry string `json:"issuerCountry,omitempty"` // The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**.
	Issuerstate string `json:"issuerState,omitempty"` // The state or province where the document was issued (AU only).
	Nationalidexempt bool `json:"nationalIdExempt,omitempty"` // Applies only to individuals in the US. Set to **true** if the individual does not have an SSN. To verify their identity, Adyen will require them to upload an ID document.
}

// PersonalData represents the PersonalData schema from the OpenAPI specification
type PersonalData struct {
	Dateofbirth string `json:"dateOfBirth,omitempty"` // The date of birth of the person. The date should be in ISO-8601 format yyyy-mm-dd (e.g. 2000-01-31).
	Idnumber string `json:"idNumber,omitempty"` // An ID number of the person.
	Nationality string `json:"nationality,omitempty"` // The nationality of the person represented by a two-character country code. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	Registeredaddress GeneratedType `json:"registeredAddress"`
	Registrationnumber string `json:"registrationNumber,omitempty"` // The organization's registration number.
	Taxexempt bool `json:"taxExempt,omitempty"` // Indicates whether the legal entity is exempt from tax. When **true**, the `taxIdAbsenceReason` must be provided.
	Taxidabsencereason string `json:"taxIdAbsenceReason,omitempty"` // The reason the organization has not provided a tax identifier. Possible values: **industryExemption**, **belowTaxThreshold**.
	Webdata WebData `json:"webData,omitempty"`
	Email string `json:"email,omitempty"` // The email address of the legal entity.
	Legalname string `json:"legalName"` // The organization's legal name.
	Principalplaceofbusiness GeneratedType `json:"principalPlaceOfBusiness,omitempty"`
	Stockdata StockData `json:"stockData,omitempty"`
	Taxid string `json:"taxId,omitempty"` // The organization's tax identifier.
	Description string `json:"description,omitempty"` // Your description for the organization.
	Doingbusinessas string `json:"doingBusinessAs,omitempty"` // The organization's trading name, if different from the registered legal name.
	Phone GeneratedType `json:"phone,omitempty"`
	TypeField string `json:"type,omitempty"` // Type of organization. Possible values: **associationIncorporated**, **governmentalOrganization**, **listedPublicCompany**, **nonProfit**, **partnershipIncorporated**, **privateCompany**.
}

// CapabilityProblem represents the CapabilityProblem schema from the OpenAPI specification
type CapabilityProblem struct {
}

// AccountSupportingEntityCapability represents the AccountSupportingEntityCapability schema from the OpenAPI specification
type AccountSupportingEntityCapability struct {
	Id string `json:"id,omitempty"` // The ID of the supporting entity.
	Requested bool `json:"requested,omitempty"` // Indicates whether the capability is requested. To check whether the account holder is permitted to use the capability, refer to the `allowed` field.
	Requestedlevel string `json:"requestedLevel,omitempty"` // The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring. Possible values: **notApplicable**, **low**, **medium**, **high**.
	Verificationstatus string `json:"verificationStatus,omitempty"` // The status of the verification checks for the supporting entity capability. Possible values: * **pending**: Adyen is running the verification. * **invalid**: The verification failed. Check if the `errors` array contains more information. * **valid**: The verification has been successfully completed. * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	Allowed bool `json:"allowed,omitempty"` // Indicates whether the supporting entity capability is allowed. Adyen sets this to **true** if the verification is successful and the account holder is permitted to use the capability.
	Allowedlevel string `json:"allowedLevel,omitempty"` // The capability level that is allowed for the account holder. Possible values: **notApplicable**, **low**, **medium**, **high**.
	Enabled bool `json:"enabled,omitempty"` // Indicates whether the capability is enabled. If **false**, the capability is temporarily disabled for the account holder.
}

// WebData represents the WebData schema from the OpenAPI specification
type WebData struct {
	Webaddress string `json:"webAddress,omitempty"` // The URL of the website or the app store URL.
	Webaddressid string `json:"webAddressId,omitempty"` // The unique identifier of the web address.
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
	Attachments []Attachment `json:"attachments"` // Array that contains the document. The array supports multiple attachments for uploading different sides or pages of a document.
	Issuerstate string `json:"issuerState,omitempty"` // The state or province where the document was issued (AU only).
	Number string `json:"number,omitempty"` // The number in the document.
	Owner OwnerEntity `json:"owner"`
	Attachment Attachment `json:"attachment,omitempty"`
	Description string `json:"description"` // Your description for the document.
	Expirydate string `json:"expiryDate,omitempty"` // The expiry date of the document, in YYYY-MM-DD format.
	Id string `json:"id"` // The unique identifier of the document.
	Issuercountry string `json:"issuerCountry,omitempty"` // The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**.
	Modificationdate string `json:"modificationDate,omitempty"` // The modification date of the document.
	Creationdate string `json:"creationDate,omitempty"` // The creation date of the document.
	Filename string `json:"fileName,omitempty"` // The filename of the document.
	TypeField string `json:"type"` // Type of document, used when providing an ID number or uploading a document. The possible values depend on the legal entity type. When providing ID numbers: * For **individual**, the `type` values can be **driversLicense**, **identityCard**, **nationalIdNumber**, or **passport**. When uploading documents: * For **organization**, the `type` values can be **proofOfAddress**, **registrationDocument**, **taxDocument**, **proofOfOwnership**, or **proofOfIndustry**. * For **individual**, the `type` values can be **identityCard**, **driversLicense**, **proofOfNationalIdNumber**, **proofOfResidency**, or **proofOfIndustry**. * Use **bankStatement** to upload documents for a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).
}

// LegalEntityInfoRequiredType represents the LegalEntityInfoRequiredType schema from the OpenAPI specification
type LegalEntityInfoRequiredType struct {
	Reference string `json:"reference,omitempty"` // Your reference for the legal entity, maximum 150 characters.
	TypeField string `json:"type"` // The type of legal entity. Possible values: **individual** or **organization**
	Capabilities map[string]interface{} `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Entityassociations []LegalEntityAssociation `json:"entityAssociations,omitempty"` // List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories.
	Individual Individual `json:"individual,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number string `json:"number"` // The phone number.
	TypeField string `json:"type"` // The type of phone number. Possible values: **mobile**, **landline**, **sip**, **fax.**
	Countrycode string `json:"countryCode,omitempty"` // The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code. For example, **US** or **NL**.
}

// TransactionRule represents the TransactionRule schema from the OpenAPI specification
type TransactionRule struct {
	Processingtypes []string `json:"processingTypes,omitempty"` // List of processing types to which the rule applies. Possible values: **atmWithdraw**, **pos**, **ecommerce**, **moto**, **recurring**, **balanceInquiry**.
	Interval TransactionRuleInterval `json:"interval"`
	TypeField string `json:"type"` // Type of conditions provided in the rule. Possible values: * **allowList**: The rule provides categories (such as country and MCC) where payments must be allowed. * **blockList**: The rule provides categories (such as country and MCC) where payments must be blocked. * **maxUsage**: The rule sets limits for the maximum amount or maximum number of transactions for the lifetime of the payment instrument. * **velocity**: The rule sets limits for the maximum amount or maximum number of transactions for a given time interval.
	Mccs []string `json:"mccs,omitempty"` // List of Merchant Category Codes (MCCs) to which the rule applies.
	Reference string `json:"reference"` // Your reference for the transaction rule, maximum 150 characters.
	Startdate string `json:"startDate,omitempty"` // The date when the rule will start to be evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**. If not provided when creating a transaction rule, the `startDate` is set to the date when the rule status is set to **active**.
	Paymentinstrumentid string `json:"paymentInstrumentId,omitempty"` // The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstruments__resParam_id) to which the rule applies.
	Balanceplatformid string `json:"balancePlatformId,omitempty"` // The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/get/balancePlatforms/{id}__queryParam_id) to which the rule applies.
	Countries []string `json:"countries,omitempty"` // List of countries to which the rule applies. For example: **["NL","US"]**.
	Description string `json:"description"` // Your description for the transaction rule, maximum 300 characters.
	Id string `json:"id,omitempty"` // The unique identifier of the transaction rule.
	Amount Amount `json:"amount,omitempty"`
	Paymentinstrumentgroupid string `json:"paymentInstrumentGroupId,omitempty"` // The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstrumentGroups__resParam_id) to which the rule applies.
	Enddate string `json:"endDate,omitempty"` // The date when the rule will stop being evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**. If not provided, the rule will be evaluated until the rule status is set to **inactive**.
	Entrymodes []string `json:"entryModes,omitempty"` // List of point-of-sale entry modes to which the rule applies. Possible values: **manual**, **chip**, **magstripe**, **contactless**, **cof**, **token**, **server**, **barcode**, **ocr**.
	Maxtransactions int `json:"maxTransactions,omitempty"` // The maximum number of transactions that a payment instrument can be used for in the period specified in the `interval.type`.
	Status string `json:"status,omitempty"` // The status of the transaction rule. If you provide a `startDate` in the request, the rule is automatically created with an **active** status. Possible values: **active**, **inactive**.
}

// CZLocalAccountIdentification represents the CZLocalAccountIdentification schema from the OpenAPI specification
type CZLocalAccountIdentification struct {
	TypeField string `json:"type"` // **czLocal**
	Accountnumber string `json:"accountNumber"` // The 2- to 16-digit bank account number (Číslo účtu) in the following format: - The optional prefix (předčíslí). - The required second part (základní část) which must be at least two non-zero digits. Examples: - **19-123457** (with prefix) - **123457** (without prefix) - **000019-0000123457** (with prefix, normalized) - **000000-0000123457** (without prefix, normalized)
	Bankcode string `json:"bankCode"` // The 4-digit bank code (Kód banky), without separators or whitespace.
}

// AdditionalBankIdentification represents the AdditionalBankIdentification schema from the OpenAPI specification
type AdditionalBankIdentification struct {
	Code string `json:"code,omitempty"` // The value of the additional bank identification.
	TypeField string `json:"type,omitempty"` // The type of additional bank identification, depending on the country. Possible values: * **gbSortCode**: The 6-digit [UK sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or spaces * **usRoutingNumber**: The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or spaces.
}

// VoidResponse represents the VoidResponse schema from the OpenAPI specification
type VoidResponse struct {
}

// SweepConfiguration represents the SweepConfiguration schema from the OpenAPI specification
type SweepConfiguration struct {
	Sweepamount Amount `json:"sweepAmount,omitempty"`
	Targetamount Amount `json:"targetAmount,omitempty"`
	TypeField string `json:"type,omitempty"` // The direction of sweep, whether pushing out or pulling in funds to the balance account. If not provided, by default, this is set to **push**. Possible values: * **push**: _push out funds_ to a destination balance account or transfer instrument. * **pull**: _pull in funds_ from a source merchant account, transfer instrument, or balance account.
	Balanceaccountid string `json:"balanceAccountId,omitempty"` // The unique identifier of the destination or source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). You can only use this for periodic sweep schedules such as `schedule.type` **daily** or **monthly**.
	Id string `json:"id"` // The unique identifier of the sweep.
	Merchantaccount string `json:"merchantAccount,omitempty"` // The merchant account that will be the source of funds. You can only use this if you are processing payments with Adyen. This can only be used for sweeps of `type` **pull** and `schedule.type` **balance**.
	Schedule interface{} `json:"schedule"` // The schedule when the `triggerAmount` is evaluated. If the balance meets the threshold, funds are pushed out of or pulled in to the balance account.
	Transferinstrumentid string `json:"transferInstrumentId,omitempty"` // The unique identifier of the destination or source [transfer instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/transferInstruments__resParam_id). You can also use this in combination with a `merchantAccount` and a `type` **pull** to start a direct debit request from the source transfer instrument. To use this feature, reach out to your Adyen contact.
	Status string `json:"status,omitempty"` // The status of the sweep. If not provided, by default, this is set to **active**. Possible values: * **active**: the sweep is enabled and funds will be pulled in or pushed out based on the defined configuration. * **inactive**: the sweep is disabled and cannot be triggered.
	Triggeramount Amount `json:"triggerAmount,omitempty"`
}

// ContactDetails represents the ContactDetails schema from the OpenAPI specification
type ContactDetails struct {
	Address Address `json:"address"`
	Email string `json:"email"` // The email address of the account holder.
	Phone Phone `json:"phone"`
	Webaddress string `json:"webAddress,omitempty"` // The URL of the account holder's website.
}

// ServiceError represents the ServiceError schema from the OpenAPI specification
type ServiceError struct {
	Errortype string `json:"errorType,omitempty"` // The category of the error.
	Message string `json:"message,omitempty"` // A short explanation of the issue.
	Pspreference string `json:"pspReference,omitempty"` // The PSP reference of the payment.
	Status int `json:"status,omitempty"` // The HTTP response status.
	Errorcode string `json:"errorCode,omitempty"` // The error code mapped to the error message.
}

// Phone represents the Phone schema from the OpenAPI specification
type Phone struct {
	Number string `json:"number"` // The full phone number provided as a single string. For example, **"0031 6 11 22 33 44"**, **"+316/1122-3344"**, or **"(0031) 611223344"**.
	TypeField string `json:"type"` // Type of phone number. Possible values: **Landline**, **Mobile**.
}

// PaymentInstrument represents the PaymentInstrument schema from the OpenAPI specification
type PaymentInstrument struct {
	Card Card `json:"card,omitempty"`
	Reference string `json:"reference,omitempty"` // Your reference for the payment instrument, maximum 150 characters.
	TypeField string `json:"type"` // Type of payment instrument. Possible value: **card**, **bankAccount**.
	Description string `json:"description,omitempty"` // Your description for the payment instrument, maximum 300 characters.
	Paymentinstrumentgroupid string `json:"paymentInstrumentGroupId,omitempty"` // The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs.
	Id string `json:"id"` // The unique identifier of the payment instrument.
	Issuingcountrycode string `json:"issuingCountryCode"` // The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**.
	Balanceaccountid string `json:"balanceAccountId"` // The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument.
	Bankaccount interface{} `json:"bankAccount,omitempty"` // Contains the business account details. Returned when you create a payment instrument with `type` **bankAccount**.
	Status string `json:"status,omitempty"` // The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **Active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **Requested**. Possible values: * **Active**: The payment instrument is active and can be used to make payments. * **Requested**: The payment instrument has been requested. This state is applicable for physical cards. * **Inactive**: The payment instrument is inactive and cannot be used to make payments. * **Suspended**: The payment instrument is temporarily suspended and cannot be used to make payments. * **Closed**: The payment instrument is permanently closed. This action cannot be undone. * **Stolen** * **Lost**
}

// TransactionRuleInfo represents the TransactionRuleInfo schema from the OpenAPI specification
type TransactionRuleInfo struct {
	Amount Amount `json:"amount,omitempty"`
	Countries []string `json:"countries,omitempty"` // List of countries to which the rule applies. For example: **["NL","US"]**.
	Entrymodes []string `json:"entryModes,omitempty"` // List of point-of-sale entry modes to which the rule applies. Possible values: **manual**, **chip**, **magstripe**, **contactless**, **cof**, **token**, **server**, **barcode**, **ocr**.
	Startdate string `json:"startDate,omitempty"` // The date when the rule will start to be evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**. If not provided when creating a transaction rule, the `startDate` is set to the date when the rule status is set to **active**.
	Paymentinstrumentgroupid string `json:"paymentInstrumentGroupId,omitempty"` // The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstrumentGroups__resParam_id) to which the rule applies.
	Description string `json:"description"` // Your description for the transaction rule, maximum 300 characters.
	Processingtypes []string `json:"processingTypes,omitempty"` // List of processing types to which the rule applies. Possible values: **atmWithdraw**, **pos**, **ecommerce**, **moto**, **recurring**, **balanceInquiry**.
	Status string `json:"status,omitempty"` // The status of the transaction rule. If you provide a `startDate` in the request, the rule is automatically created with an **active** status. Possible values: **active**, **inactive**.
	Mccs []string `json:"mccs,omitempty"` // List of Merchant Category Codes (MCCs) to which the rule applies.
	Balanceplatformid string `json:"balancePlatformId,omitempty"` // The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/get/balancePlatforms/{id}__queryParam_id) to which the rule applies.
	Interval TransactionRuleInterval `json:"interval"`
	Reference string `json:"reference"` // Your reference for the transaction rule, maximum 150 characters.
	Enddate string `json:"endDate,omitempty"` // The date when the rule will stop being evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**. If not provided, the rule will be evaluated until the rule status is set to **inactive**.
	Maxtransactions int `json:"maxTransactions,omitempty"` // The maximum number of transactions that a payment instrument can be used for in the period specified in the `interval.type`.
	Paymentinstrumentid string `json:"paymentInstrumentId,omitempty"` // The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstruments__resParam_id) to which the rule applies.
	TypeField string `json:"type"` // Type of conditions provided in the rule. Possible values: * **allowList**: The rule provides categories (such as country and MCC) where payments must be allowed. * **blockList**: The rule provides categories (such as country and MCC) where payments must be blocked. * **maxUsage**: The rule sets limits for the maximum amount or maximum number of transactions for the lifetime of the payment instrument. * **velocity**: The rule sets limits for the maximum amount or maximum number of transactions for a given time interval.
}

// CapabilityProblemEntity represents the CapabilityProblemEntity schema from the OpenAPI specification
type CapabilityProblemEntity struct {
	Id string `json:"id,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// CardConfiguration represents the CardConfiguration schema from the OpenAPI specification
type CardConfiguration struct {
	Bulkaddress BulkAddress `json:"bulkAddress,omitempty"`
	Configurationprofileid string `json:"configurationProfileId"` // The ID of the card configuration profile that contains the settings of the card. For example, the envelope and PIN mailer designs or the logistics company handling the shipment. All the settings in the profile are applied to the card, unless you provide other fields to override them. For example, send the `shipmentMethod` to override the logistics company defined in the card configuration profile.
	Insert string `json:"insert,omitempty"` // Overrides the insert design ID defined in the `configurationProfileId`. An insert is any additional material, such as marketing materials, that are shipped together with the card.
	Language string `json:"language,omitempty"` // The two-letter [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code of the card. For example, **en**.
	Logoimageid string `json:"logoImageId,omitempty"` // The ID of the logo image. This is the image that will be printed on the partial front of the card, such as a logo on the upper right corner.
	Activation string `json:"activation,omitempty"` // Overrides the activation label design ID defined in the `configurationProfileId`. The activation label is attached to the card and contains the activation instructions.
	Cardimageid string `json:"cardImageId,omitempty"` // The ID of the card image. This is the image that will be printed on the full front of the card.
	Carrier string `json:"carrier,omitempty"` // Overrides the carrier design ID defined in the `configurationProfileId`. The carrier is the letter or packaging to which the card is attached.
	Carrierimageid string `json:"carrierImageId,omitempty"` // The ID of the carrier image. This is the image that will printed on the letter to which the card is attached.
	Currency string `json:"currency,omitempty"` // The three-letter [ISO-4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the card. For example, **EUR**.
	Envelope string `json:"envelope,omitempty"` // Overrides the envelope design ID defined in the `configurationProfileId`.
	Pinmailer string `json:"pinMailer,omitempty"` // Overrides the PIN mailer design ID defined in the `configurationProfileId`. The PIN mailer is the letter on which the PIN is printed.
	Shipmentmethod string `json:"shipmentMethod,omitempty"` // Overrides the logistics company defined in the `configurationProfileId`.
	Activationurl string `json:"activationUrl,omitempty"` // Your app's URL, if you want to activate cards through your app. For example, **my-app://ref1236a7d**. A QR code is created based on this URL, and is included in the carrier. Before you use this field, reach out to your Adyen contact to set up the QR code process. Maximum length: 255 characters.
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	Phonecountrycode string `json:"phoneCountryCode,omitempty"` // The two-character ISO-3166-1 alpha-2 country code of the phone number. For example, **US** or **NL**.
	Phonenumber string `json:"phoneNumber,omitempty"` // The phone number. The inclusion of the phone number country code is not necessary.
	Phonetype string `json:"phoneType,omitempty"` // The type of the phone number. Possible values: **Landline**, **Mobile**, **SIP**, **Fax**.
}

// JSONPath represents the JSONPath schema from the OpenAPI specification
type JSONPath struct {
	Content []string `json:"content,omitempty"`
}

// SELocalAccountIdentification represents the SELocalAccountIdentification schema from the OpenAPI specification
type SELocalAccountIdentification struct {
	Accountnumber string `json:"accountNumber"` // The 7- to 10-digit bank account number ([Bankkontonummer](https://sv.wikipedia.org/wiki/Bankkonto)), without the clearing number, separators, or whitespace.
	Clearingnumber string `json:"clearingNumber"` // The 4- to 5-digit clearing number ([Clearingnummer](https://sv.wikipedia.org/wiki/Clearingnummer)), without separators or whitespace.
	TypeField string `json:"type"` // **seLocal**
}

// TransferInstrument represents the TransferInstrument schema from the OpenAPI specification
type TransferInstrument struct {
	Legalentityid string `json:"legalEntityId"` // The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the transfer instrument.
	TypeField string `json:"type"` // The type of transfer instrument. Possible value: **bankAccount**.
	Bankaccount BankAccountInfo `json:"bankAccount"`
	Documents []EntityReference `json:"documents,omitempty"` // List of documents uploaded for the transfer instrument.
	Id string `json:"id"` // The unique identifier of the transfer instrument.
}

// LegalEntity represents the LegalEntity schema from the OpenAPI specification
type LegalEntity struct {
	Individual Individual `json:"individual,omitempty"`
	Organization Organization `json:"organization,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of legal entity. Possible values: **individual** or **organization**
	Id string `json:"id"` // The unique identifier of the legal entity.
	Transferinstruments []TransferInstrumentReference `json:"transferInstruments,omitempty"` // List of transfer instruments that the legal entity owns.
	Capabilities map[string]interface{} `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Reference string `json:"reference,omitempty"` // Your reference for the legal entity, maximum 150 characters.
	Documents []EntityReference `json:"documents,omitempty"` // List of documents uploaded for the legal entity.
	Entityassociations []LegalEntityAssociation `json:"entityAssociations,omitempty"` // List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories.
}

// JSONObject represents the JSONObject schema from the OpenAPI specification
type JSONObject struct {
	Paths []JSONPath `json:"paths,omitempty"`
	Rootpath JSONPath `json:"rootPath,omitempty"`
}

// CardInfo represents the CardInfo schema from the OpenAPI specification
type CardInfo struct {
	Formfactor string `json:"formFactor"` // The form factor of the card. Possible values: **virtual**, **physical**.
	Authentication Authentication `json:"authentication,omitempty"`
	Brand string `json:"brand"` // The brand of the physical or the virtual card. Possible values: **visa**, **mc**.
	Brandvariant string `json:"brandVariant"` // The brand variant of the physical or the virtual card. >Contact your Adyen Implementation Manager to get the values that are relevant to your integration. Examples: **visadebit**, **mcprepaid**.
	Cardholdername string `json:"cardholderName"` // The name of the cardholder. Maximum length: 26 characters.
	Configuration CardConfiguration `json:"configuration,omitempty"`
	Deliverycontact Contact `json:"deliveryContact,omitempty"`
}

// BalanceAccountUpdateRequest represents the BalanceAccountUpdateRequest schema from the OpenAPI specification
type BalanceAccountUpdateRequest struct {
	Description string `json:"description,omitempty"` // A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder.
	Reference string `json:"reference,omitempty"` // Your reference to the balance account, maximum 150 characters.
	Status string `json:"status,omitempty"` // The status of the balance account. Payment instruments linked to the balance account can only be used if the balance account status is **Active**. Possible values: **Active**, **Inactive**, **Closed**, **Suspended**.
	Sweepconfigurations map[string]interface{} `json:"sweepConfigurations,omitempty"` // Contains key-value pairs that specify [balance sweep per currency code](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__reqParam_sweepConfigurations). You can update the balance account to add, update, or delete sweeps. * To add a sweep, send the currency code as a key and the configuration as the object. * To update a sweep, send the whole configuration with your updates. * To delete a sweep, set the value to **null**. For example, `"EUR": null`.
	Timezone string `json:"timeZone,omitempty"` // The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	Accountholderid string `json:"accountHolderId,omitempty"` // The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account.
	Defaultcurrencycode string `json:"defaultCurrencyCode,omitempty"` // The default currency code of this balance account, in three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) format. The default value is **EUR**.
}

// PaymentInstrumentInfo represents the PaymentInstrumentInfo schema from the OpenAPI specification
type PaymentInstrumentInfo struct {
	Status string `json:"status,omitempty"` // The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **Active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **Requested**. Possible values: * **Active**: The payment instrument is active and can be used to make payments. * **Requested**: The payment instrument has been requested. This state is applicable for physical cards. * **Inactive**: The payment instrument is inactive and cannot be used to make payments. * **Suspended**: The payment instrument is temporarily suspended and cannot be used to make payments. * **Closed**: The payment instrument is permanently closed. This action cannot be undone. * **Stolen** * **Lost**
	TypeField string `json:"type"` // Type of payment instrument. Possible value: **card**, **bankAccount**.
	Balanceaccountid string `json:"balanceAccountId"` // The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument.
	Card CardInfo `json:"card,omitempty"`
	Description string `json:"description,omitempty"` // Your description for the payment instrument, maximum 300 characters.
	Issuingcountrycode string `json:"issuingCountryCode"` // The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**.
	Paymentinstrumentgroupid string `json:"paymentInstrumentGroupId,omitempty"` // The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs.
	Reference string `json:"reference,omitempty"` // Your reference for the payment instrument, maximum 150 characters.
}

// CapabilitySettings represents the CapabilitySettings schema from the OpenAPI specification
type CapabilitySettings struct {
	Authorizedcardusers bool `json:"authorizedCardUsers,omitempty"` // The number of card holders who can use the card.
	Fundingsource []string `json:"fundingSource,omitempty"` // The funding source of the card, for example **debit**.
	Interval string `json:"interval,omitempty"` // The period when the rule conditions apply.
	Maxamount Amount `json:"maxAmount,omitempty"`
	Amountperindustry map[string]interface{} `json:"amountPerIndustry,omitempty"` // The maximum amount a card holder can spend per industry.
}

// AccountHolderInfo represents the AccountHolderInfo schema from the OpenAPI specification
type AccountHolderInfo struct {
	Timezone string `json:"timeZone,omitempty"` // The [time zone](https://www.iana.org/time-zones) of the account holder. For example, **Europe/Amsterdam**. Defaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	Balanceplatform string `json:"balancePlatform,omitempty"` // The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms.
	Capabilities map[string]interface{} `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability.
	Contactdetails ContactDetails `json:"contactDetails,omitempty"`
	Description string `json:"description,omitempty"` // Your description for the account holder, maximum 300 characters.
	Legalentityid string `json:"legalEntityId"` // The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder.
	Reference string `json:"reference,omitempty"` // Your reference for the account holder, maximum 150 characters.
}

// Card represents the Card schema from the OpenAPI specification
type Card struct {
	Configuration CardConfiguration `json:"configuration,omitempty"`
	Deliverycontact Contact `json:"deliveryContact,omitempty"`
	Cvc string `json:"cvc,omitempty"` // The CVC2 value of the card. > The CVC2 is not sent by default. This is only returned in the `POST` response for single-use virtual cards.
	Cardholdername string `json:"cardholderName"` // The name of the cardholder. Maximum length: 26 characters.
	Formfactor string `json:"formFactor"` // The form factor of the card. Possible values: **virtual**, **physical**.
	Lastfour string `json:"lastFour,omitempty"` // Last last four digits of the card number.
	Number string `json:"number"` // The primary account number (PAN) of the card. > The PAN is masked by default and returned only for single-use virtual cards.
	Authentication Authentication `json:"authentication,omitempty"`
	Bin string `json:"bin,omitempty"` // The bank identification number (BIN) of the card number.
	Brandvariant string `json:"brandVariant"` // The brand variant of the physical or the virtual card. >Contact your Adyen Implementation Manager to get the values that are relevant to your integration. Examples: **visadebit**, **mcprepaid**.
	Expiration Expiry `json:"expiration,omitempty"`
	Brand string `json:"brand"` // The brand of the physical or the virtual card. Possible values: **visa**, **mc**.
}

// Individual represents the Individual schema from the OpenAPI specification
type Individual struct {
	Nationality string `json:"nationality,omitempty"` // The individual's nationality.
	Phone GeneratedType `json:"phone,omitempty"`
	Residentialaddress GeneratedType `json:"residentialAddress"`
	Webdata WebData `json:"webData,omitempty"`
	Birthdata BirthData `json:"birthData,omitempty"`
	Email string `json:"email,omitempty"` // The email address of the legal entity.
	Identificationdata IdentificationData `json:"identificationData,omitempty"`
	Name GeneratedType `json:"name"`
}

// InvalidField represents the InvalidField schema from the OpenAPI specification
type InvalidField struct {
	Name string `json:"name"` // The field that has an invalid value.
	Value string `json:"value"` // The invalid value.
	Message string `json:"message"` // Description of the validation error.
}

// StockData represents the StockData schema from the OpenAPI specification
type StockData struct {
	Marketidentifier string `json:"marketIdentifier,omitempty"` // The four-digit [Market Identifier Code](https://en.wikipedia.org/wiki/Market_Identifier_Code) of the stock market where the organization's stocks are traded.
	Stocknumber string `json:"stockNumber,omitempty"` // The 12-digit International Securities Identification Number (ISIN) of the company, without dashes (-).
	Tickersymbol string `json:"tickerSymbol,omitempty"` // The stock ticker symbol.
}

// Balance represents the Balance schema from the OpenAPI specification
type Balance struct {
	Available int64 `json:"available"` // The remaining amount available for spending.
	Balance int64 `json:"balance"` // The total amount in the balance.
	Currency string `json:"currency"` // The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance.
	Reserved int64 `json:"reserved"` // The amount reserved for payments that have been authorised, but have not been captured yet.
}

// TransferInstrumentReference represents the TransferInstrumentReference schema from the OpenAPI specification
type TransferInstrumentReference struct {
	Accountidentifier string `json:"accountIdentifier"` // The masked IBAN or bank account number.
	Id string `json:"id"` // The unique identifier of the resource.
}

// VerificationError represents the VerificationError schema from the OpenAPI specification
type VerificationError struct {
	Capabilities []string `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that the legal entity can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Code string `json:"code,omitempty"` // The general error code.
	Message string `json:"message,omitempty"` // The general error message.
	Remediatingactions []RemediatingAction `json:"remediatingActions,omitempty"` // An object containing possible solutions to fix a verification error.
	Suberrors []GeneratedType `json:"subErrors,omitempty"` // An array containing more granular information about the cause of the verification error.
	TypeField string `json:"type,omitempty"` // The type of error.
}

// TransactionRulesResponse represents the TransactionRulesResponse schema from the OpenAPI specification
type TransactionRulesResponse struct {
	Transactionrules []TransactionRule `json:"transactionRules,omitempty"` // List of transaction rules.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Capabilities []string `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that the legal entity can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Code string `json:"code,omitempty"` // The general error code.
	Message string `json:"message,omitempty"` // The general error message.
	Remediatingactions []RemediatingAction `json:"remediatingActions,omitempty"` // An object containing possible solutions to fix a verification error.
	TypeField string `json:"type,omitempty"` // The type of error.
}

// PLLocalAccountIdentification represents the PLLocalAccountIdentification schema from the OpenAPI specification
type PLLocalAccountIdentification struct {
	Accountnumber string `json:"accountNumber"` // The 26-digit bank account number ([Numer rachunku](https://pl.wikipedia.org/wiki/Numer_Rachunku_Bankowego)), without separators or whitespace.
	TypeField string `json:"type"` // **plLocal**
}

// PaymentInstrumentGroupInfo represents the PaymentInstrumentGroupInfo schema from the OpenAPI specification
type PaymentInstrumentGroupInfo struct {
	Balanceplatform string `json:"balancePlatform"` // The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the payment instrument group belongs.
	Description string `json:"description,omitempty"` // Your description for the payment instrument group, maximum 300 characters.
	Properties map[string]interface{} `json:"properties,omitempty"` // Properties of the payment instrument group.
	Reference string `json:"reference,omitempty"` // Your reference for the payment instrument group, maximum 150 characters.
	Txvariant string `json:"txVariant"` // The tx variant of the payment instrument group.
}

// RestServiceError represents the RestServiceError schema from the OpenAPI specification
type RestServiceError struct {
	TypeField string `json:"type"` // A URI that identifies the problem type, pointing to human-readable documentation on this problem type.
	Response JSONObject `json:"response,omitempty"`
	Errorcode string `json:"errorCode"` // A code that identifies the problem type.
	Invalidfields []InvalidField `json:"invalidFields,omitempty"` // Detailed explanation of each validation error, when applicable.
	Requestid string `json:"requestId,omitempty"` // A unique reference for the request, essentially the same as `pspReference`.
	Detail string `json:"detail"` // A human-readable explanation specific to this occurrence of the problem.
	Instance string `json:"instance,omitempty"` // A unique URI that identifies the specific occurrence of the problem.
	Status int `json:"status"` // The HTTP status code.
	Title string `json:"title"` // A short, human-readable summary of the problem type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Postalcode string `json:"postalCode,omitempty"` // Postal code. Required if `stateOrProvince` and/or `city` is provided.
	Stateorprovince string `json:"stateOrProvince,omitempty"` // The two-letter ISO 3166-2 state or province code. For example, **CA** in the US. If you specify the state or province, you must also send `city`, `postalCode`, and `street`.
	Street string `json:"street,omitempty"` // The name of the street, and the house or building number. Required if `stateOrProvince` and/or `city` is provided.
	Street2 string `json:"street2,omitempty"` // The apartment, unit, or suite number.
	City string `json:"city,omitempty"` // The name of the city. Required if `stateOrProvince` is provided. If you specify the city, you must also send `postalCode` and `street`.
	Country string `json:"country"` // The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
}

// TransactionRuleResponse represents the TransactionRuleResponse schema from the OpenAPI specification
type TransactionRuleResponse struct {
	Transactionrule TransactionRule `json:"transactionRule,omitempty"`
}

// OwnerEntity represents the OwnerEntity schema from the OpenAPI specification
type OwnerEntity struct {
	Id string `json:"id"` // Unique identifier of the resource that owns the document. For `type` **legalEntity**, this value is the unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id). For `type` **bankAccount**, this value is the unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).
	TypeField string `json:"type"` // Type of resource that owns the document. Possible values: **legalEntity**, **bankAccount**.
}

// Expiry represents the Expiry schema from the OpenAPI specification
type Expiry struct {
	Month string `json:"month,omitempty"` // The month in which the card will expire.
	Year string `json:"year,omitempty"` // The year in which the card will expire.
}

// IbanAccountIdentification represents the IbanAccountIdentification schema from the OpenAPI specification
type IbanAccountIdentification struct {
	Iban string `json:"iban"` // The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard.
	TypeField string `json:"type"` // **iban**
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Firstname string `json:"firstName"` // The individual's first name.
	Infix string `json:"infix,omitempty"` // The infix in the individual's name, if any.
	Lastname string `json:"lastName"` // The individual's last name.
}

// PaginatedPaymentInstrumentsResponse represents the PaginatedPaymentInstrumentsResponse schema from the OpenAPI specification
type PaginatedPaymentInstrumentsResponse struct {
	Hasnext bool `json:"hasNext"` // Indicates whether there are more items on the next page.
	Hasprevious bool `json:"hasPrevious"` // Indicates whether there are more items on the previous page.
	Paymentinstruments []PaymentInstrument `json:"paymentInstruments"` // List of payment instruments associated with the balance account.
}

// BalanceAccount represents the BalanceAccount schema from the OpenAPI specification
type BalanceAccount struct {
	Balances []Balance `json:"balances,omitempty"` // List of balances with the amount and currency.
	Status string `json:"status,omitempty"` // The status of the balance account, set to **Active** by default.
	Id string `json:"id"` // The unique identifier of the balance account.
	Reference string `json:"reference,omitempty"` // Your reference for the balance account, maximum 150 characters.
	Description string `json:"description,omitempty"` // A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder.
	Paymentinstruments []PaymentInstrumentReference `json:"paymentInstruments,omitempty"` // List of [payment instruments](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/paymentInstruments) associated with the balance account.
	Sweepconfigurations map[string]interface{} `json:"sweepConfigurations,omitempty"` // Contains key-value pairs that specify configurations for balance sweeps per currency code. A sweep pulls in or pushes out funds based on a defined schedule, amount, and a source (for pulling funds) or a destination (for pushing funds). The key must be a three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) in uppercase. For example, **EUR**. The value must be an object containing the sweep configuration.
	Accountholderid string `json:"accountHolderId"` // The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account.
	Defaultcurrencycode string `json:"defaultCurrencyCode,omitempty"` // The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**.
	Timezone string `json:"timeZone,omitempty"` // The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
}

// Amount represents the Amount schema from the OpenAPI specification
type Amount struct {
	Currency string `json:"currency"` // The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"` // The amount of the transaction, in [minor units](https://docs.adyen.com/development-resources/currency-codes).
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Country string `json:"country"` // The two-character ISO-3166-1 alpha-2 country code. For example, **US**. > If you don't know the country or are not collecting the country from the shopper, provide `country` as `ZZ`.
	Housenumberorname string `json:"houseNumberOrName"` // The number or name of the house. Maximum length: 3000 characters.
	Postalcode string `json:"postalCode"` // A maximum of five digits for an address in the US, or a maximum of ten characters for an address in all other countries.
	Stateorprovince string `json:"stateOrProvince,omitempty"` // The two-character ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. > Required for the US and Canada.
	Street string `json:"street"` // The name of the street. Maximum length: 3000 characters. > The house number should not be included in this field; it should be separately provided via `houseNumberOrName`.
	City string `json:"city"` // The name of the city. Maximum length: 3000 characters.
}

// SGLocalAccountIdentification represents the SGLocalAccountIdentification schema from the OpenAPI specification
type SGLocalAccountIdentification struct {
	Bic string `json:"bic"` // The bank's 8- or 11-character BIC or SWIFT code.
	TypeField string `json:"type,omitempty"` // **sgLocal**
	Accountnumber string `json:"accountNumber"` // The 4- to 19-digit bank account number, without separators or whitespace.
}

// AccountHolder represents the AccountHolder schema from the OpenAPI specification
type AccountHolder struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability.
	Description string `json:"description,omitempty"` // Your description for the account holder, maximum 300 characters.
	Contactdetails ContactDetails `json:"contactDetails,omitempty"`
	Legalentityid string `json:"legalEntityId"` // The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder.
	Reference string `json:"reference,omitempty"` // Your reference for the account holder, maximum 150 characters.
	Status string `json:"status,omitempty"` // The status of the account holder. Possible values: * **Active**: The account holder is active. This is the default status when creating an account holder. * **Inactive**: The account holder is temporarily inactive due to missing KYC details. You can set the account back to active by providing the missing KYC details. * **Suspended**: The account holder is permanently deactivated by Adyen. This action cannot be undone. * **Closed**: The account holder is permanently deactivated by you. This action cannot be undone.
	Balanceplatform string `json:"balancePlatform,omitempty"` // The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms.
	Id string `json:"id"` // The unique identifier of the account holder.
	Primarybalanceaccount string `json:"primaryBalanceAccount,omitempty"` // The ID of the account holder's primary balance account. By default, this is set to the first balance account that you create for the account holder. To assign a different balance account, send a PATCH request.
	Timezone string `json:"timeZone,omitempty"` // The [time zone](https://www.iana.org/time-zones) of the account holder. For example, **Europe/Amsterdam**. Defaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
}

// Authentication represents the Authentication schema from the OpenAPI specification
type Authentication struct {
	Email string `json:"email,omitempty"` // The email address where the one-time password (OTP) is sent.
	Password string `json:"password,omitempty"` // The password used for 3D Secure password-based authentication. The value must be between 1 to 30 characters and must only contain the following supported characters. * Characters between **a-z**, **A-Z**, and **0-9** * Special characters: **äöüßÄÖÜ+-*/ç%()=?!~#'",;:$&àùòâôûáúó**
	Phone Phone `json:"phone,omitempty"`
}

// Void represents the Void schema from the OpenAPI specification
type Void struct {
}

// LegalEntityInfo represents the LegalEntityInfo schema from the OpenAPI specification
type LegalEntityInfo struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty"` // Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Entityassociations []LegalEntityAssociation `json:"entityAssociations,omitempty"` // List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories.
	Individual Individual `json:"individual,omitempty"`
	Organization Organization `json:"organization,omitempty"`
	Reference string `json:"reference,omitempty"` // Your reference for the legal entity, maximum 150 characters.
	TypeField string `json:"type,omitempty"` // The type of legal entity. Possible values: **individual** or **organization**
}

// TransactionRuleInterval represents the TransactionRuleInterval schema from the OpenAPI specification
type TransactionRuleInterval struct {
	TypeField string `json:"type"` // The [type of interval](https://docs.adyen.com/issuing/transaction-rules#time-intervals) during which the rule conditions and limits apply, and how often counters are reset. Possible values: * **perTransaction**: conditions are evaluated and the counters are reset for every transaction. * **daily**: the counters are reset daily at 00:00:00 UTC. * **weekly**: the counters are reset every Monday at 00:00:00 UTC. * **monthly**: the counters reset every first day of the month at 00:00:00 UTC. * **lifetime**: conditions are applied to the lifetime of the payment instrument.
}

// TransferInstrumentInfo represents the TransferInstrumentInfo schema from the OpenAPI specification
type TransferInstrumentInfo struct {
	Bankaccount BankAccountInfo `json:"bankAccount"`
	Legalentityid string `json:"legalEntityId"` // The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the transfer instrument.
	TypeField string `json:"type"` // The type of transfer instrument. Possible value: **bankAccount**.
}

// EntityReference represents the EntityReference schema from the OpenAPI specification
type EntityReference struct {
	Id string `json:"id,omitempty"` // The unique identifier of the resource.
}

// CALocalAccountIdentification represents the CALocalAccountIdentification schema from the OpenAPI specification
type CALocalAccountIdentification struct {
	Institutionnumber string `json:"institutionNumber"` // The 3-digit institution number, without separators or whitespace.
	Transitnumber string `json:"transitNumber"` // The 5-digit transit number, without separators or whitespace.
	TypeField string `json:"type"` // **caLocal**
	Accountnumber string `json:"accountNumber"` // The 5- to 12-digit bank account number, without separators or whitespace.
}

// BalanceAccountInfo represents the BalanceAccountInfo schema from the OpenAPI specification
type BalanceAccountInfo struct {
	Timezone string `json:"timeZone,omitempty"` // The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	Accountholderid string `json:"accountHolderId"` // The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account.
	Defaultcurrencycode string `json:"defaultCurrencyCode,omitempty"` // The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**.
	Description string `json:"description,omitempty"` // A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder.
	Reference string `json:"reference,omitempty"` // Your reference for the balance account, maximum 150 characters.
	Sweepconfigurations map[string]interface{} `json:"sweepConfigurations,omitempty"` // Contains key-value pairs that specify configurations for balance sweeps per currency code. A sweep pulls in or pushes out funds based on a defined schedule, amount, and a source (for pulling funds) or a destination (for pushing funds). The key must be a three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) in uppercase. For example, **EUR**. The value must be an object containing the sweep configuration.
}

// CronSweepSchedule represents the CronSweepSchedule schema from the OpenAPI specification
type CronSweepSchedule struct {
	TypeField string `json:"type,omitempty"` // The schedule type. Possible values: * **cron**: push out funds based on a cron expression. * **daily**: push out funds daily at 07:00 AM CET. * **weekly**: push out funds every Monday at 07:00 AM CET. * **monthly**: push out funds every first of the month at 07:00 AM CET. * **balance**: pull in funds instantly if the balance is less than or equal to the `triggerAmount`. You can only use this for sweeps of `type` **pull** and when the source is a `merchantAccount` or `transferInstrument`.
	Cronexpression string `json:"cronExpression"` // A [cron expression](https://en.wikipedia.org/wiki/Cron#CRON_expression) that is used to set the sweep schedule. The schedule uses the time zone of the balance account. For example, **30 17 * * MON** schedules a sweep every Monday at 17:30. The expression must have five values separated by a single space in the following order: * Minute: **0-59** * Hour: **0-23** * Day of the month: **1-31** * Month: **1-12** or **JAN-DEC** * Day of the week: **0-7** (0 and 7 are Sunday) or **MON-SUN**. The following non-standard characters are supported: **&ast;**, **L**, **#**, **W** and **/**. See [crontab guru](https://crontab.guru/) for more examples.
}

// BankAccountInfo represents the BankAccountInfo schema from the OpenAPI specification
type BankAccountInfo struct {
	Bankcity string `json:"bankCity,omitempty"` // The city where the bank is located.
	Branchcode string `json:"branchCode,omitempty"` // The branch code of the branch under which the bank account is registered. Required when you provide an `accountNumber`. In the following countries, this value corresponds to: * United States: routing number * United Kingdom: sort code * Germany: Bankleitzahl
	Checkcode string `json:"checkCode,omitempty"` // The check code of the bank account.
	Accountnumber string `json:"accountNumber,omitempty"` // The bank account number (without separators). When this is provided, the `branchCode` is also required.
	Accounttype string `json:"accountType,omitempty"` // The type of bank account.
	Countrycode string `json:"countryCode,omitempty"` // The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the bank account is registered. For example, **NL**.
	Iban string `json:"iban,omitempty"` // The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard.
	Bankbicswift string `json:"bankBicSwift,omitempty"` // The bank's BIC or SWIFT code.
	Bankcode string `json:"bankCode,omitempty"` // The bank code of the banking institution with which the bank account is registered.
	Bankname string `json:"bankName,omitempty"` // The name of the banking institution where the bank account is held.
	Currencycode string `json:"currencyCode"` // The account's three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). For example, **EUR**.
}

// SweepSchedule represents the SweepSchedule schema from the OpenAPI specification
type SweepSchedule struct {
	TypeField string `json:"type,omitempty"` // The schedule type. Possible values: * **cron**: push out funds based on a cron expression. * **daily**: push out funds daily at 07:00 AM CET. * **weekly**: push out funds every Monday at 07:00 AM CET. * **monthly**: push out funds every first of the month at 07:00 AM CET. * **balance**: pull in funds instantly if the balance is less than or equal to the `triggerAmount`. You can only use this for sweeps of `type` **pull** and when the source is a `merchantAccount` or `transferInstrument`.
}

// NumberAndBicAccountIdentification represents the NumberAndBicAccountIdentification schema from the OpenAPI specification
type NumberAndBicAccountIdentification struct {
	Additionalbankidentification AdditionalBankIdentification `json:"additionalBankIdentification,omitempty"`
	Bic string `json:"bic"` // The bank's 8- or 11-character BIC or SWIFT code.
	TypeField string `json:"type"` // **numberAndBic**
	Accountnumber string `json:"accountNumber"` // The bank account number, without separators or whitespace. The length and format depends on the bank or country.
}

// PaymentInstrumentRevealInfo represents the PaymentInstrumentRevealInfo schema from the OpenAPI specification
type PaymentInstrumentRevealInfo struct {
	Pan string `json:"pan"` // The primary account number (PAN) of the card.
	Cvc string `json:"cvc"` // The CVC2 value of the card.
	Expiration Expiry `json:"expiration"`
}

// BankAccountIdentificationValidationRequest represents the BankAccountIdentificationValidationRequest schema from the OpenAPI specification
type BankAccountIdentificationValidationRequest struct {
	Accountidentification interface{} `json:"accountIdentification"` // Bank account identification.
}

// UpdatePaymentInstrument represents the UpdatePaymentInstrument schema from the OpenAPI specification
type UpdatePaymentInstrument struct {
	TypeField string `json:"type"` // Type of payment instrument. Possible value: **card**, **bankAccount**.
	Id string `json:"id"` // The unique identifier of the payment instrument.
	Paymentinstrumentgroupid string `json:"paymentInstrumentGroupId,omitempty"` // The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs.
	Card Card `json:"card,omitempty"`
	Issuingcountrycode string `json:"issuingCountryCode"` // The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**.
	Reference string `json:"reference,omitempty"` // Your reference for the payment instrument, maximum 150 characters.
	Statuscomment string `json:"statusComment,omitempty"` // Comment for the status of the payment instrument.
	Balanceaccountid string `json:"balanceAccountId"` // The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument.
	Bankaccount interface{} `json:"bankAccount,omitempty"` // Contains the business account details. Returned when you create a payment instrument with `type` **bankAccount**.
	Description string `json:"description,omitempty"` // Your description for the payment instrument, maximum 300 characters.
	Status string `json:"status,omitempty"` // The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **Active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **Requested**. Possible values: * **Active**: The payment instrument is active and can be used to make payments. * **Requested**: The payment instrument has been requested. This state is applicable for physical cards. * **Inactive**: The payment instrument is inactive and cannot be used to make payments. * **Suspended**: The payment instrument is temporarily suspended and cannot be used to make payments. * **Closed**: The payment instrument is permanently closed. This action cannot be undone. * **Stolen** * **Lost**
}

// PaginatedAccountHoldersResponse represents the PaginatedAccountHoldersResponse schema from the OpenAPI specification
type PaginatedAccountHoldersResponse struct {
	Hasprevious bool `json:"hasPrevious"` // Indicates whether there are more items on the previous page.
	Accountholders []AccountHolder `json:"accountHolders"` // List of account holders.
	Hasnext bool `json:"hasNext"` // Indicates whether there are more items on the next page.
}

// PaymentInstrumentGroup represents the PaymentInstrumentGroup schema from the OpenAPI specification
type PaymentInstrumentGroup struct {
	Reference string `json:"reference,omitempty"` // Your reference for the payment instrument group, maximum 150 characters.
	Txvariant string `json:"txVariant"` // The tx variant of the payment instrument group.
	Balanceplatform string `json:"balancePlatform"` // The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the payment instrument group belongs.
	Description string `json:"description,omitempty"` // Your description for the payment instrument group, maximum 300 characters.
	Id string `json:"id,omitempty"` // The unique identifier of the payment instrument group.
	Properties map[string]interface{} `json:"properties,omitempty"` // Properties of the payment instrument group.
}

// HULocalAccountIdentification represents the HULocalAccountIdentification schema from the OpenAPI specification
type HULocalAccountIdentification struct {
	Accountnumber string `json:"accountNumber"` // The 24-digit bank account number, without separators or whitespace.
	TypeField string `json:"type"` // **huLocal**
}

// RemediatingAction represents the RemediatingAction schema from the OpenAPI specification
type RemediatingAction struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// USLocalAccountIdentification represents the USLocalAccountIdentification schema from the OpenAPI specification
type USLocalAccountIdentification struct {
	Routingnumber string `json:"routingNumber"` // The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace.
	TypeField string `json:"type"` // **usLocal**
	Accountnumber string `json:"accountNumber"` // The bank account number, without separators or whitespace.
	Accounttype string `json:"accountType,omitempty"` // The bank account type. Possible values: **checking** or **savings**. Defaults to **checking**.
}

// UKLocalAccountIdentification represents the UKLocalAccountIdentification schema from the OpenAPI specification
type UKLocalAccountIdentification struct {
	Accountnumber string `json:"accountNumber"` // The 8-digit bank account number, without separators or whitespace.
	Sortcode string `json:"sortCode"` // The 6-digit [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace.
	TypeField string `json:"type"` // **ukLocal**
}

// LegalEntityAssociation represents the LegalEntityAssociation schema from the OpenAPI specification
type LegalEntityAssociation struct {
	Associatorid string `json:"associatorId,omitempty"` // The unique identifier of another legal entity with which the `legalEntityId` is associated. When the `legalEntityId` is associated to legal entities other than the current one, the response returns all the associations.
	Entitytype string `json:"entityType,omitempty"` // The legal entity type of associated legal entity. For example, **organization**, **soleProprietorship** or **individual**.
	Jobtitle string `json:"jobTitle,omitempty"` // The individual's job title if the `type` is **uboThroughControl** or **signatory**.
	Legalentityid string `json:"legalEntityId"` // The unique identifier of the associated [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id).
	Name string `json:"name,omitempty"` // The name of the associated [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id). - For **individual**, `name.firstName` and `name.lastName`. - For **organization**, `legalName`. - For **soleProprietorship**, `name`.
	TypeField string `json:"type"` // Defines the relationship of the legal entity to the current legal entity. For example, **uboThroughOwnership**, **uboThroughControl**, **signatory**, or **soleProprietorship**.
}

// SupportingEntityCapability represents the SupportingEntityCapability schema from the OpenAPI specification
type SupportingEntityCapability struct {
	Allowed bool `json:"allowed,omitempty"` // Indicates whether the supporting entity capability is allowed. If a supporting entity is allowed but its parent legal entity is not, it means there are other supporting entities that failed validation. **The allowed supporting entity can still be used**
	Id string `json:"id,omitempty"` // Supporting entity reference
	Requested bool `json:"requested,omitempty"` // Indicates whether the supporting entity capability is requested.
	Verificationstatus string `json:"verificationStatus,omitempty"` // The status of the verification checks for the supporting entity capability. Possible values: * **pending**: Adyen is running the verification. * **invalid**: The verification failed. Check if the `errors` array contains more information. * **valid**: The verification has been successfully completed. * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
}
