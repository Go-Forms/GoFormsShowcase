package eventlog

import "goforms"

// eventful is anything carrying the inherited ControlEvents group - which is
// every GoForms control, through ControlBase.
type eventful interface {
	Events() *goforms.ControlEvents
}

// WireAll subscribes every inherited Control event to the log, so opening a
// form and poking at it shows the whole interaction surface firing live.
//
// The high-frequency ones (MouseMove, MouseWheel) are deliberately left out:
// they would bury everything else within a second of moving the pointer.
// WireMotion adds them for the one control that demonstrates them.
func WireAll(source string, c eventful) {
	ev := c.Events()

	ev.Click.Handle(func(_ any, e goforms.MouseEventArgs) {
		Logf(source, "Click", "at %.0f,%.0f button=%s clicks=%d", e.X, e.Y, buttonName(e.Button), e.Clicks)
	})
	ev.DoubleClick.Handle(func(_ any, e goforms.MouseEventArgs) {
		Logf(source, "DoubleClick", "at %.0f,%.0f", e.X, e.Y)
	})
	ev.MouseDown.Handle(func(_ any, e goforms.MouseEventArgs) {
		Logf(source, "MouseDown", "%s at %.0f,%.0f", buttonName(e.Button), e.X, e.Y)
	})
	ev.MouseUp.Handle(func(_ any, e goforms.MouseEventArgs) {
		Logf(source, "MouseUp", "%s at %.0f,%.0f", buttonName(e.Button), e.X, e.Y)
	})
	ev.MouseEnter.Handle(func(any, goforms.EventArgs) { Log(source, "MouseEnter", "") })
	ev.MouseLeave.Handle(func(any, goforms.EventArgs) { Log(source, "MouseLeave", "") })

	ev.KeyDown.Handle(func(_ any, e goforms.KeyEventArgs) {
		Logf(source, "KeyDown", "%s%s", modifiers(e), e.KeyCode)
	})
	ev.KeyUp.Handle(func(_ any, e goforms.KeyEventArgs) {
		Logf(source, "KeyUp", "%s%s", modifiers(e), e.KeyCode)
	})
	ev.KeyPress.Handle(func(_ any, e goforms.KeyPressEventArgs) {
		Logf(source, "KeyPress", "%q", e.KeyChar)
	})

	ev.GotFocus.Handle(func(any, goforms.EventArgs) { Log(source, "GotFocus", "") })
	ev.LostFocus.Handle(func(any, goforms.EventArgs) { Log(source, "LostFocus", "") })

	ev.Resize.Handle(func(any, goforms.EventArgs) { Log(source, "Resize", "") })
	ev.Move.Handle(func(any, goforms.EventArgs) { Log(source, "Move", "") })
	ev.VisibleChanged.Handle(func(any, goforms.EventArgs) { Log(source, "VisibleChanged", "") })
	ev.EnabledChanged.Handle(func(any, goforms.EventArgs) { Log(source, "EnabledChanged", "") })
}

// WireMotion adds the events WireAll leaves out because they fire constantly.
//
// They are gated behind MotionEnabled, off by default: a single MouseMove
// subscriber turns every pixel of pointer travel into a log entry, and the
// UI work behind each entry then dominates the frame budget.
func WireMotion(source string, c eventful) {
	ev := c.Events()
	ev.MouseMove.Handle(func(_ any, e goforms.MouseEventArgs) {
		if MotionEnabled() {
			Logf(source, "MouseMove", "%.0f,%.0f", e.X, e.Y)
		}
	})
	ev.MouseWheel.Handle(func(_ any, e goforms.MouseEventArgs) {
		if MotionEnabled() {
			Logf(source, "MouseWheel", "delta %.1f", e.Delta)
		}
	})
}

func buttonName(b goforms.MouseButton) string {
	switch b {
	case goforms.MouseButtonLeft:
		return "left"
	case goforms.MouseButtonRight:
		return "right"
	case goforms.MouseButtonMiddle:
		return "middle"
	}
	return "none"
}

// modifiers renders the modifier state as a "Ctrl+Shift+" style prefix.
// Fyne only reports modifiers with mouse events, so this reflects the last
// state it told us about rather than the instant of the keystroke.
func modifiers(e goforms.KeyEventArgs) string {
	s := ""
	for _, m := range []struct {
		on   bool
		name string
	}{{e.Control, "Ctrl"}, {e.Alt, "Alt"}, {e.Shift, "Shift"}, {e.Super, "Super"}} {
		if m.on {
			s += m.name + "+"
		}
	}
	return s
}
