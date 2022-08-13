package wintrustapi

// UIMode identifies a user interface mode when making wintrust API calls.
type UIMode uint32

// User-interface modes.
const (
	UserInterfaceAll    UIMode = iota + 1 // WTD_UI_ALL
	UserInterfaceNone                     // WTD_UI_NONE
	UserInterfaceNoBad                    // WTD_UI_NOBAD
	UserInterfaceNoGood                   // WTD_UI_NOGOOD
)

// UIContext identifiers a user interface context when making wintrust API
// calls.
type UIContext uint32

// User-interface contexts.
const (
	UIContextExecute UIContext = iota // WTD_UICONTEXT_EXECUTE
	UIContextInstall                  // WTD_UICONTEXT_INSTALL
)
