package wintrustapi

// StateAction identifies the type of verification action to be taken.
type StateAction uint32

// Possible state actions.
const (
	StateActionIgnore         StateAction = iota // WTD_STATEACTION_IGNORE
	StateActionVerify                            // WTD_STATEACTION_VERIFY
	StateActionClose                             // WTD_STATEACTION_CLOSE
	StateActionAutoCache                         // WTD_STATEACTION_AUTO_CACHE
	StateActionAutoCacheFlush                    // WTD_STATEACTION_AUTO_CACHE_FLUSH
)
