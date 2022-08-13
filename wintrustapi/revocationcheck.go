package wintrustapi

// RevocationCheck describes possible options for revocation checking when
// making wintrust API calls.
type RevocationCheck uint32

// Revocation checking options.
const (
	RevocationCheckNone       RevocationCheck = iota // WTD_REVOKE_NONE
	RevocationCheckWholeChain                        // WTD_REVOKE_WHOLECHAIN
)
