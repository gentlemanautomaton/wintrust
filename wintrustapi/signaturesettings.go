package wintrustapi

// SignatureSettings can be used to communicate signature verification
// requirements in wintrust API calls.
type SignatureSettings struct {
	Size             uint32
	Index            uint32
	Flags            uint32
	SecondarySigs    uint32
	VerifiedSigIndex uint32
	CryptoPolicy     uintptr
}
