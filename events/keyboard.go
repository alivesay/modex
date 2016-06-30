package events

import (
	"fmt"
	"strings"

	"github.com/go-gl/glfw/v3.1/glfw"
)

type Key int

const (
	KeyUnknown      Key = Key(glfw.KeyUnknown)
	KeySpace            = Key(glfw.KeySpace)
	KeyApostrophe       = Key(glfw.KeyApostrophe)
	KeyComma            = Key(glfw.KeyComma)
	KeyMinus            = Key(glfw.KeyMinus)
	KeyPeriod           = Key(glfw.KeyPeriod)
	KeySlash            = Key(glfw.KeySlash)
	Key0                = Key(glfw.Key0)
	Key1                = Key(glfw.Key1)
	Key2                = Key(glfw.Key2)
	Key3                = Key(glfw.Key3)
	Key4                = Key(glfw.Key4)
	Key5                = Key(glfw.Key5)
	Key6                = Key(glfw.Key6)
	Key7                = Key(glfw.Key7)
	Key8                = Key(glfw.Key8)
	Key9                = Key(glfw.Key9)
	KeySemicolon        = Key(glfw.KeySemicolon)
	KeyEqual            = Key(glfw.KeyEqual)
	KeyA                = Key(glfw.KeyA)
	KeyB                = Key(glfw.KeyB)
	KeyC                = Key(glfw.KeyC)
	KeyD                = Key(glfw.KeyD)
	KeyE                = Key(glfw.KeyE)
	KeyF                = Key(glfw.KeyF)
	KeyG                = Key(glfw.KeyG)
	KeyH                = Key(glfw.KeyH)
	KeyI                = Key(glfw.KeyI)
	KeyJ                = Key(glfw.KeyJ)
	KeyK                = Key(glfw.KeyK)
	KeyL                = Key(glfw.KeyL)
	KeyM                = Key(glfw.KeyM)
	KeyN                = Key(glfw.KeyN)
	KeyO                = Key(glfw.KeyO)
	KeyP                = Key(glfw.KeyP)
	KeyQ                = Key(glfw.KeyQ)
	KeyR                = Key(glfw.KeyR)
	KeyS                = Key(glfw.KeyS)
	KeyT                = Key(glfw.KeyT)
	KeyU                = Key(glfw.KeyU)
	KeyV                = Key(glfw.KeyV)
	KeyW                = Key(glfw.KeyW)
	KeyX                = Key(glfw.KeyX)
	KeyY                = Key(glfw.KeyY)
	KeyZ                = Key(glfw.KeyZ)
	KeyLeftBracket      = Key(glfw.KeyLeftBracket)
	KeyBackslash        = Key(glfw.KeyBackslash)
	KeyRightBracket     = Key(glfw.KeyRightBracket)
	KeyGraveAccent      = Key(glfw.KeyGraveAccent)
	KeyWorld1           = Key(glfw.KeyWorld1)
	KeyWorld2           = Key(glfw.KeyWorld2)
	KeyEscape           = Key(glfw.KeyEscape)
	KeyEnter            = Key(glfw.KeyEnter)
	KeyTab              = Key(glfw.KeyTab)
	KeyBackspace        = Key(glfw.KeyBackspace)
	KeyInsert           = Key(glfw.KeyInsert)
	KeyDelete           = Key(glfw.KeyDelete)
	KeyRight            = Key(glfw.KeyRight)
	KeyLeft             = Key(glfw.KeyLeft)
	KeyDown             = Key(glfw.KeyDown)
	KeyUp               = Key(glfw.KeyUp)
	KeyPageUp           = Key(glfw.KeyPageUp)
	KeyPageDown         = Key(glfw.KeyPageDown)
	KeyHome             = Key(glfw.KeyHome)
	KeyEnd              = Key(glfw.KeyEnd)
	KeyCapsLock         = Key(glfw.KeyCapsLock)
	KeyScrollLock       = Key(glfw.KeyScrollLock)
	KeyNumLock          = Key(glfw.KeyNumLock)
	KeyPrintScreen      = Key(glfw.KeyPrintScreen)
	KeyPause            = Key(glfw.KeyPause)
	KeyF1               = Key(glfw.KeyF1)
	KeyF2               = Key(glfw.KeyF2)
	KeyF3               = Key(glfw.KeyF3)
	KeyF4               = Key(glfw.KeyF4)
	KeyF5               = Key(glfw.KeyF5)
	KeyF6               = Key(glfw.KeyF6)
	KeyF7               = Key(glfw.KeyF7)
	KeyF8               = Key(glfw.KeyF8)
	KeyF9               = Key(glfw.KeyF9)
	KeyF10              = Key(glfw.KeyF10)
	KeyF11              = Key(glfw.KeyF11)
	KeyF12              = Key(glfw.KeyF12)
	KeyF13              = Key(glfw.KeyF13)
	KeyF14              = Key(glfw.KeyF14)
	KeyF15              = Key(glfw.KeyF15)
	KeyF16              = Key(glfw.KeyF16)
	KeyF17              = Key(glfw.KeyF17)
	KeyF18              = Key(glfw.KeyF18)
	KeyF19              = Key(glfw.KeyF19)
	KeyF20              = Key(glfw.KeyF20)
	KeyF21              = Key(glfw.KeyF21)
	KeyF22              = Key(glfw.KeyF22)
	KeyF23              = Key(glfw.KeyF23)
	KeyF24              = Key(glfw.KeyF24)
	KeyF25              = Key(glfw.KeyF25)
	KeyKPO              = Key(glfw.KeyKP0)
	KeyKP1              = Key(glfw.KeyKP1)
	KeyKP2              = Key(glfw.KeyKP2)
	KeyKP3              = Key(glfw.KeyKP3)
	KeyKP4              = Key(glfw.KeyKP4)
	KeyKP5              = Key(glfw.KeyKP5)
	KeyKP6              = Key(glfw.KeyKP6)
	KeyKP7              = Key(glfw.KeyKP7)
	KeyKP8              = Key(glfw.KeyKP8)
	KeyKP9              = Key(glfw.KeyKP9)
	KeyKPDecimal        = Key(glfw.KeyKPDecimal)
	KeyKPDivide         = Key(glfw.KeyKPDivide)
	KeyKPMultiply       = Key(glfw.KeyKPMultiply)
	KeyKPSubtract       = Key(glfw.KeyKPSubtract)
	KeyKPAdd            = Key(glfw.KeyKPAdd)
	KeyKPEnter          = Key(glfw.KeyKPEnter)
	KeyKPEqual          = Key(glfw.KeyKPEqual)
	KeyLeftShift        = Key(glfw.KeyLeftShift)
	KeyLeftControl      = Key(glfw.KeyLeftControl)
	KeyLeftAlt          = Key(glfw.KeyLeftAlt)
	KeyLeftSuper        = Key(glfw.KeyLeftSuper)
	KeyRightShift       = Key(glfw.KeyRightShift)
	KeyRightControl     = Key(glfw.KeyRightControl)
	KeyRightAlt         = Key(glfw.KeyRightAlt)
	KeyRightSuper       = Key(glfw.KeyRightSuper)
	KeyMenu             = Key(glfw.KeyMenu)
	KeyLast             = Key(glfw.KeyLast)
)

var KeyNames = map[Key]string{
	KeyUnknown:      `KeyUnknown`,
	KeySpace:        `KeySpace`,
	KeyApostrophe:   `KeyApostrophe`,
	KeyComma:        `KeyComma`,
	KeyMinus:        `KeyMinus`,
	KeyPeriod:       `KeyPeriod`,
	KeySlash:        `KeySlash`,
	Key0:            `Key0`,
	Key1:            `Key1`,
	Key2:            `Key2`,
	Key3:            `Key3`,
	Key4:            `Key4`,
	Key5:            `Key5`,
	Key6:            `Key6`,
	Key7:            `Key7`,
	Key8:            `Key8`,
	Key9:            `Key9`,
	KeySemicolon:    `KeySemicolon`,
	KeyEqual:        `KeyEqual`,
	KeyA:            `KeyA`,
	KeyB:            `KeyB`,
	KeyC:            `KeyC`,
	KeyD:            `KeyD`,
	KeyE:            `KeyE`,
	KeyF:            `KeyF`,
	KeyG:            `KeyG`,
	KeyH:            `KeyH`,
	KeyI:            `KeyI`,
	KeyJ:            `KeyJ`,
	KeyK:            `KeyK`,
	KeyL:            `KeyL`,
	KeyM:            `KeyM`,
	KeyN:            `KeyN`,
	KeyO:            `KeyO`,
	KeyP:            `KeyP`,
	KeyQ:            `KeyQ`,
	KeyR:            `KeyR`,
	KeyS:            `KeyS`,
	KeyT:            `KeyT`,
	KeyU:            `KeyU`,
	KeyV:            `KeyV`,
	KeyW:            `KeyW`,
	KeyX:            `KeyX`,
	KeyY:            `KeyY`,
	KeyZ:            `KeyZ`,
	KeyLeftBracket:  `KeyLeftBracket`,
	KeyBackslash:    `KeyBackslash`,
	KeyRightBracket: `KeyRightBracket`,
	KeyGraveAccent:  `KeyGraveAccent`,
	KeyWorld1:       `KeyWorld1`,
	KeyWorld2:       `KeyWorld2`,
	KeyEscape:       `KeyEscape`,
	KeyEnter:        `KeyEnter`,
	KeyTab:          `KeyTab`,
	KeyBackspace:    `KeyBackspace`,
	KeyInsert:       `KeyInsert`,
	KeyDelete:       `KeyDelete`,
	KeyRight:        `KeyRight`,
	KeyLeft:         `KeyLeft`,
	KeyDown:         `KeyDown`,
	KeyUp:           `KeyUp`,
	KeyPageUp:       `KeyPageUp`,
	KeyPageDown:     `KeyPageDown`,
	KeyHome:         `KeyHome`,
	KeyEnd:          `KeyEnd`,
	KeyCapsLock:     `KeyCapsLock`,
	KeyScrollLock:   `KeyScrolLock`,
	KeyNumLock:      `KeyNumLock`,
	KeyPrintScreen:  `KeyPrintScreen`,
	KeyPause:        `KeyPause`,
	KeyF1:           `KeyF1`,
	KeyF2:           `KeyF2`,
	KeyF3:           `KeyF3`,
	KeyF4:           `KetF4`,
	KeyF5:           `KeyF5`,
	KeyF6:           `KeyF6`,
	KeyF7:           `KeyF7`,
	KeyF8:           `KeyF8`,
	KeyF9:           `KeyF9`,
	KeyF10:          `KeyF10`,
	KeyF11:          `KeyF11`,
	KeyF12:          `KeyF12`,
	KeyF13:          `KeyF13`,
	KeyF14:          `KeyF14`,
	KeyF15:          `KeyF15`,
	KeyF16:          `KeyF16`,
	KeyF17:          `KeyF17`,
	KeyF18:          `KeyF18`,
	KeyF19:          `KeyF19`,
	KeyF20:          `KeyF20`,
	KeyF21:          `KeyF21`,
	KeyF22:          `KeyF22`,
	KeyF23:          `KeyF23`,
	KeyF24:          `KeyF24`,
	KeyF25:          `KeyF25`,
	KeyKPO:          `KeyKPO`,
	KeyKP1:          `KeyKP1`,
	KeyKP2:          `KeyKP2`,
	KeyKP3:          `KeyKP3`,
	KeyKP4:          `KeyKP4`,
	KeyKP5:          `KeyKP5`,
	KeyKP6:          `KeyKP6`,
	KeyKP7:          `KeyKP7`,
	KeyKP8:          `KeyKP8`,
	KeyKP9:          `KeyKP9`,
	KeyKPDecimal:    `KeyKPDecimal`,
	KeyKPDivide:     `KeyKPDivide`,
	KeyKPMultiply:   `KeyKPMultiply`,
	KeyKPSubtract:   `KeyKPSubtract`,
	KeyKPAdd:        `KeyKPAdd`,
	KeyKPEnter:      `KeyKPEnter`,
	KeyKPEqual:      `KeyKPEqual`,
	KeyLeftShift:    `KeyLeftShift`,
	KeyLeftControl:  `KeyLeftControl`,
	KeyLeftAlt:      `KeyLeftAlt`,
	KeyLeftSuper:    `KeyLeftSuper`,
	KeyRightShift:   `KeyRightShift`,
	KeyRightControl: `KeyRightControl`,
	KeyRightAlt:     `KeyRightAlt`,
	KeyRightSuper:   `KeyRightSuper`,
	KeyMenu:         `KeyMenu`,
}

type ModifierKey int

const (
	ModShift   ModifierKey = ModifierKey(glfw.ModShift)
	ModControl             = ModifierKey(glfw.ModControl)
	ModAlt                 = ModifierKey(glfw.ModAlt)
	ModSuper               = ModifierKey(glfw.ModSuper)
)

var ModifierKeyNames = map[ModifierKey]string{
	ModShift:   `ModShift`,
	ModControl: `ModControl`,
	ModAlt:     `ModAlt`,
	ModSuper:   `ModSuper`,
}

type Action int

const (
	Release Action = Action(glfw.Release)
	Press          = Action(glfw.Press)
	Repeat         = Action(glfw.Repeat)
)

var ActionNames = map[Action]string{
	Release: `Release`,
	Press:   `Press`,
	Repeat:  `Repeat`,
}

type KeyEvent struct {
	Window   *glfw.Window
	Key      Key
	Scancode int
	Action   Action
	Mods     ModifierKey
}

func (keyEvent *KeyEvent) String() string {
	mods := make([]string, 0, len(ModifierKeyNames))
	for k, v := range ModifierKeyNames {
		if (keyEvent.Mods & k) != 0 {
			mods = append(mods, v)
		}
	}

	return fmt.Sprintf(
		"KeyEvent: %s: <%s> <%s> %#x",
		ActionNames[keyEvent.Action],
		strings.Join(mods, "+"),
		KeyNames[keyEvent.Key],
		keyEvent.Scancode)
}

type KeyCallback func(event *KeyEvent)
