package g810

const (
	KG_LOGO = iota
	KG_INDICATORS = iota
	KG_MULTIMEDIA = iota
	KG_GKEYS = iota
	KG_KEYS = iota
)

const (
	KB_UNKNOWN = iota
	KB_G213 = iota
	KB_G410 = iota
	KB_G413 = iota
	KB_G513 = iota
	KB_G610 = iota
	KB_G810 = iota
	KB_G910 = iota
	KB_GPRO = iota
)

type KeyGroup uint8

const (
	GroupLogo KeyGroup = iota
	GroupIndicators
	GroupMultimedia
	GroupGKeys
	GroupFKeys
	GroupModifiers
	GroupFunctions
	GroupArrows
	GroupNumeric
	GroupKeys
)

type KeyGroupAddress uint8

const (
	GroupAddressLogo KeyGroupAddress = iota
	GroupAddressIndicators
	GroupAddressMultimedia
	GroupAddressGKeys
	GroupAddressKeys
)

type Key uint16

const (
	KeyLogo Key = iota + (Key)(GroupAddressLogo) << 8 | 0x01
	KeyLogo2
)

const (
	KeyBacklight Key = iota + (Key)(GroupAddressIndicators) << 8 | 0x01
	KeyGame
	KeyCaps
	KeyScroll
	KeyNum
)

const (
	KeyNext = iota + (Key)(GroupAddressMultimedia) << 8 | 0xB5
	KeyPrev
	KeyStop
	KeyPlay = (Key)(GroupAddressMultimedia) << 8 | 0xCD
	KeyMute = (Key)(GroupAddressMultimedia) << 8 | 0xE2
)

const (
	KeyG1 = iota + (Key)(GroupAddressGKeys) << 8 | 0x01
	KeyG2
	KeyG3
	KeyG4
	KeyG5
	KeyG6
	KeyG7
	KeyG8
	KeyG9
)

const (
	KeyA = iota + (Key)(GroupAddressKeys) << 8 | 0x04
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	KeyN1
	KeyN2
	KeyN3
	KeyN4
	KeyN5
	KeyN6
	KeyN7
	KeyN8
	KeyN9
	KeyN0
)

const (
	KeyEnter Key = iota + KeyN0 + 1
	KeyEsc
	KeyBackspace
	KeyTab
	KeySpace
	KeyMinus
	KeyEqual
	KeyOpenBracket
	KeyCloseBracket
	KeyBackslash
	KeyDollar
	KeySemicolon
	KeyQuote
	KeyTilde
	KeyComma
	KeyPeriod
	KeySlash
	KeyCapsLock

	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12

	KeyPrintScreen
	KeyScrollLock
	KeyPauseBreak

	KeyInsert
	KeyHome
	KeyPageUp
	KeyDelete
	KeyEnd
	KeyPageDown

	KeyArrowRight
	KeyArrowLeft
	KeyArrowBottom
	KeyArrowUp

	KeyNumLock
	KeyNumSlash
	KeyNumAsterisk
	KeyNumMinus
	KeyNumPlus
	KeyNumEnter

	KeyNum1
	KeyNum2
	KeyNum3
	KeyNum4
	KeyNum5
	KeyNum6
	KeyNum7
	KeyNum8
	KeyNum9
	KeyNum0
	KeyNumDot

	KeyIntlBackslash
	KeyMenu
)

const (
	KeyCtrlLeft Key = iota + (Key)(GroupAddressKeys) << 8 | 0xE0
	KeyShiftLeft
	KeyAltLeft
	KeyWinLeft

	KeyCtrlRight
	KeyShiftRight
	KeyAltRight
	KeyWinRight
)
