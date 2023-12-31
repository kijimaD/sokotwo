package utils

import "github.com/hajimehoshi/ebiten/v2"

// KeyMap is a US keyboard mapping
var KeyMap = map[string]ebiten.Key{
	"A":              ebiten.KeyA,
	"B":              ebiten.KeyB,
	"C":              ebiten.KeyC,
	"D":              ebiten.KeyD,
	"E":              ebiten.KeyE,
	"F":              ebiten.KeyF,
	"G":              ebiten.KeyG,
	"H":              ebiten.KeyH,
	"I":              ebiten.KeyI,
	"J":              ebiten.KeyJ,
	"K":              ebiten.KeyK,
	"L":              ebiten.KeyL,
	"M":              ebiten.KeyM,
	"N":              ebiten.KeyN,
	"O":              ebiten.KeyO,
	"P":              ebiten.KeyP,
	"Q":              ebiten.KeyQ,
	"R":              ebiten.KeyR,
	"S":              ebiten.KeyS,
	"T":              ebiten.KeyT,
	"U":              ebiten.KeyU,
	"V":              ebiten.KeyV,
	"W":              ebiten.KeyW,
	"X":              ebiten.KeyX,
	"Y":              ebiten.KeyY,
	"Z":              ebiten.KeyZ,
	"AltLeft":        ebiten.KeyAltLeft,
	"AltRight":       ebiten.KeyAltRight,
	"Down":           ebiten.KeyArrowDown,
	"Left":           ebiten.KeyArrowLeft,
	"Right":          ebiten.KeyArrowRight,
	"Up":             ebiten.KeyArrowUp,
	"Backquote":      ebiten.KeyBackquote,
	"Backslash":      ebiten.KeyBackslash,
	"Backspace":      ebiten.KeyBackspace,
	"BracketLeft":    ebiten.KeyBracketLeft,
	"BracketRight":   ebiten.KeyBracketRight,
	"CapsLock":       ebiten.KeyCapsLock,
	"Comma":          ebiten.KeyComma,
	"ContextMenu":    ebiten.KeyContextMenu,
	"ControlLeft":    ebiten.KeyControlLeft,
	"ControlRight":   ebiten.KeyControlRight,
	"Delete":         ebiten.KeyDelete,
	"0":              ebiten.KeyDigit0,
	"1":              ebiten.KeyDigit1,
	"2":              ebiten.KeyDigit2,
	"3":              ebiten.KeyDigit3,
	"4":              ebiten.KeyDigit4,
	"5":              ebiten.KeyDigit5,
	"6":              ebiten.KeyDigit6,
	"7":              ebiten.KeyDigit7,
	"8":              ebiten.KeyDigit8,
	"9":              ebiten.KeyDigit9,
	"End":            ebiten.KeyEnd,
	"Enter":          ebiten.KeyEnter,
	"Equal":          ebiten.KeyEqual,
	"Escape":         ebiten.KeyEscape,
	"F1":             ebiten.KeyF1,
	"F2":             ebiten.KeyF2,
	"F3":             ebiten.KeyF3,
	"F4":             ebiten.KeyF4,
	"F5":             ebiten.KeyF5,
	"F6":             ebiten.KeyF6,
	"F7":             ebiten.KeyF7,
	"F8":             ebiten.KeyF8,
	"F9":             ebiten.KeyF9,
	"F10":            ebiten.KeyF10,
	"F11":            ebiten.KeyF11,
	"F12":            ebiten.KeyF12,
	"Home":           ebiten.KeyHome,
	"Insert":         ebiten.KeyInsert,
	"MetaLeft":       ebiten.KeyMetaLeft,
	"MetaRight":      ebiten.KeyMetaRight,
	"Minus":          ebiten.KeyMinus,
	"NumLock":        ebiten.KeyNumLock,
	"Numpad0":        ebiten.KeyNumpad0,
	"Numpad1":        ebiten.KeyNumpad1,
	"Numpad2":        ebiten.KeyNumpad2,
	"Numpad3":        ebiten.KeyNumpad3,
	"Numpad4":        ebiten.KeyNumpad4,
	"Numpad5":        ebiten.KeyNumpad5,
	"Numpad6":        ebiten.KeyNumpad6,
	"Numpad7":        ebiten.KeyNumpad7,
	"Numpad8":        ebiten.KeyNumpad8,
	"Numpad9":        ebiten.KeyNumpad9,
	"NumpadAdd":      ebiten.KeyNumpadAdd,
	"NumpadDecimal":  ebiten.KeyNumpadDecimal,
	"NumpadDivide":   ebiten.KeyNumpadDivide,
	"NumpadEnter":    ebiten.KeyNumpadEnter,
	"NumpadEqual":    ebiten.KeyNumpadEqual,
	"NumpadMultiply": ebiten.KeyNumpadMultiply,
	"NumpadSubtract": ebiten.KeyNumpadSubtract,
	"PageDown":       ebiten.KeyPageDown,
	"PageUp":         ebiten.KeyPageUp,
	"Pause":          ebiten.KeyPause,
	"Period":         ebiten.KeyPeriod,
	"PrintScreen":    ebiten.KeyPrintScreen,
	"Quote":          ebiten.KeyQuote,
	"ScrollLock":     ebiten.KeyScrollLock,
	"Semicolon":      ebiten.KeySemicolon,
	"ShiftLeft":      ebiten.KeyShiftLeft,
	"ShiftRight":     ebiten.KeyShiftRight,
	"Slash":          ebiten.KeySlash,
	"Space":          ebiten.KeySpace,
	"Tab":            ebiten.KeyTab,
}

// GamepadButtonMap is a gamepad button mapping
var GamepadButtonMap = map[string]ebiten.GamepadButton{
	"GamepadButton0":  ebiten.GamepadButton0,
	"GamepadButton1":  ebiten.GamepadButton1,
	"GamepadButton2":  ebiten.GamepadButton2,
	"GamepadButton3":  ebiten.GamepadButton3,
	"GamepadButton4":  ebiten.GamepadButton4,
	"GamepadButton5":  ebiten.GamepadButton5,
	"GamepadButton6":  ebiten.GamepadButton6,
	"GamepadButton7":  ebiten.GamepadButton7,
	"GamepadButton8":  ebiten.GamepadButton8,
	"GamepadButton9":  ebiten.GamepadButton9,
	"GamepadButton10": ebiten.GamepadButton10,
	"GamepadButton11": ebiten.GamepadButton11,
	"GamepadButton12": ebiten.GamepadButton12,
	"GamepadButton13": ebiten.GamepadButton13,
	"GamepadButton14": ebiten.GamepadButton14,
	"GamepadButton15": ebiten.GamepadButton15,
	"GamepadButton16": ebiten.GamepadButton16,
	"GamepadButton17": ebiten.GamepadButton17,
	"GamepadButton18": ebiten.GamepadButton18,
	"GamepadButton19": ebiten.GamepadButton19,
	"GamepadButton20": ebiten.GamepadButton20,
	"GamepadButton21": ebiten.GamepadButton21,
	"GamepadButton22": ebiten.GamepadButton22,
	"GamepadButton23": ebiten.GamepadButton23,
	"GamepadButton24": ebiten.GamepadButton24,
	"GamepadButton25": ebiten.GamepadButton25,
	"GamepadButton26": ebiten.GamepadButton26,
	"GamepadButton27": ebiten.GamepadButton27,
	"GamepadButton28": ebiten.GamepadButton28,
	"GamepadButton29": ebiten.GamepadButton29,
	"GamepadButton30": ebiten.GamepadButton30,
	"GamepadButton31": ebiten.GamepadButton31,
}

// MouseButtonMap is a mouse button mapping
var MouseButtonMap = map[string]ebiten.MouseButton{
	"MouseButtonLeft":   ebiten.MouseButtonLeft,
	"MouseButtonRight":  ebiten.MouseButtonRight,
	"MouseButtonMiddle": ebiten.MouseButtonMiddle,
}
