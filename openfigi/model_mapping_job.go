/*
 * OpenFIGI API
 *
 * A free & open API for FIGI discovery.
 *
 * API version: 1.4.0
 * Contact: support@openfigi.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package openfigi

// For V3: securityType2 is required when idType is BASE_TICKER or ID_EXCH_SYMBOL.  expiration is required when securityType2 is Option or Warrant.  maturity is required when securityType2 is Pool.
type MappingJob[IdValueType comparable] struct {
	IdType                  string      `json:"idType"`
	IdValue                 IdValueType `json:"idValue"`
	ExchCode                string      `json:"exchCode,omitempty"`
	MicCode                 string      `json:"micCode,omitempty"`
	Currency                string      `json:"currency,omitempty"`
	MarketSecDes            string      `json:"marketSecDes,omitempty"`
	SecurityType            string      `json:"securityType,omitempty"`
	SecurityType2           string      `json:"securityType2,omitempty"`
	IncludeUnlistedEquities bool        `json:"includeUnlistedEquities,omitempty"`
	OptionType              string      `json:"optionType,omitempty"`
	Strike                  *[]float64  `json:"strike,omitempty"`
	ContractSize            *[]float64  `json:"contractSize,omitempty"`
	Coupon                  *[]float64  `json:"coupon,omitempty"`
	Expiration              *[]string   `json:"expiration,omitempty"`
	Maturity                *[]string   `json:"maturity,omitempty"`
	StateCode               string      `json:"stateCode,omitempty"`
}
