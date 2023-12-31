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

type FigiResult struct {
	Figi string `json:"figi,omitempty"`
	SecurityType string `json:"securityType,omitempty"`
	MarketSector string `json:"marketSector,omitempty"`
	Ticker string `json:"ticker,omitempty"`
	Name string `json:"name,omitempty"`
	ExchCode string `json:"exchCode,omitempty"`
	ShareClassFIGI string `json:"shareClassFIGI,omitempty"`
	CompositeFIGI string `json:"compositeFIGI,omitempty"`
	SecurityType2 string `json:"securityType2,omitempty"`
	SecurityDescription string `json:"securityDescription,omitempty"`
	// Exists when API is unable to show non-FIGI fields.
	Metadata string `json:"metadata,omitempty"`
}
