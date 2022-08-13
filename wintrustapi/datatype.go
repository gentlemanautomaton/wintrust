package wintrustapi

// DataType identifies the type of data provided when making wintrust API
// calls.
type DataType uint32

// Possible data types.
const (
	DataTypeFile    DataType = iota + 1 // WTD_CHOICE_FILE
	DataTypeCatalog                     // WTD_CHOICE_CATALOG
	DataTypeBlob                        // WTD_CHOICE_BLOB
	DataTypeSigner                      // WTD_CHOICE_SIGNER
	DataTypeCert                        // WTD_CHOICE_CERT
)
