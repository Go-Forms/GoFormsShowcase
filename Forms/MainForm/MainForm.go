package mainform

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"

	"goforms"
	aboutform "goformsshowcase/Forms/AboutForm"
	controlsform "goformsshowcase/Forms/ControlsForm"
	dataform "goformsshowcase/Forms/DataForm"
	dialogsform "goformsshowcase/Forms/DialogsForm"
	layoutform "goformsshowcase/Forms/LayoutForm"
	"goformsshowcase/eventlog"
)

// MainForm.go is the hand-written half: event handlers and logic. Layout
// lives in MainForm-designer.go.

// maxLogLines bounds the log so a long session can't grow the list without
// limit - the interesting entries are always the most recent ones.
const maxLogLines = 400

// open tracks the modeless demo forms so re-clicking a launcher button
// brings the existing window back rather than piling up duplicates, which
// is what WinForms code usually does by hand too.
type openForms struct {
	controls *controlsform.ControlsForm
	data     *dataform.DataForm
	layout   *layoutform.LayoutForm
	dialogs  *dialogsform.DialogsForm
}

var open openForms

var (
	logLines []string
	pending  []string
	paused   bool
	total    int
	lastMsg  string
)

func (mf *MainForm) MainForm_Load(sender any, e goforms.EventArgs) {
	mf.buildMenu()
	mf.buildToolStrip()
	mf.buildLogContextMenu()

	mf.statusStrip.AddPanel("Ready")
	mf.statusStrip.AddPanel("")

	// Every form in the showcase reports here.
	eventlog.Subscribe(mf.appendEntry)
	mf.clock.Start()
	mf.logTimer.Start()

	eventlog.Log("Main/Form", "Load", "showcase started")
}

func (mf *MainForm) MainForm_Closing(sender any, e *goforms.CancelEventArgs) {
	// Closing the shell closes the whole showcase, so shut the modeless
	// windows down with it rather than leaving orphans behind.
	mf.clock.Stop()
	mf.logTimer.Stop()
	for _, f := range []interface{ Close() }{
		formOrNil(open.controls), formOrNil(open.data),
		formOrNil(open.layout), formOrNil(open.dialogs),
	} {
		if f != nil {
			f.Close()
		}
	}
}

func (mf *MainForm) MainForm_Closed(sender any, e goforms.EventArgs) {
	goforms.Exit()
}

func (mf *MainForm) MainForm_Resize(sender any, e goforms.EventArgs) {
	size := mf.ClientSize()
	mf.setStatus(1, fmt.Sprintf("%.0f × %.0f", size.Width, size.Height))
}

func (mf *MainForm) MainForm_KeyDown(sender any, e goforms.KeyEventArgs) {
	eventlog.Logf("Main/Form", "KeyDown", "%s", e.KeyCode)
}

// formOrNil turns a possibly-nil concrete form into a nil interface, so the
// loop above doesn't call Close on a typed nil.
func formOrNil[T interface{ Close() }](f T) interface{ Close() } {
	var zero T
	if any(f) == any(zero) {
		return nil
	}
	return f
}

// --- menus, toolbar, context menu ---------------------------------------

func (mf *MainForm) buildMenu() {
	file := goforms.NewTopMenu("File",
		goforms.NewMenuItem("Clear log"),
		goforms.NewMenuSeparator(),
		goforms.NewMenuItem("Exit"),
	)
	file.Items[0].Click.Handle(mf.menuClear_Click)
	file.Items[2].Click.Handle(mf.menuExit_Click)

	windows := goforms.NewTopMenu("Windows",
		goforms.NewMenuItem("Controls"),
		goforms.NewMenuItem("Data & grids"),
		goforms.NewMenuItem("Layout & anchoring"),
		goforms.NewMenuItem("Dialogs"),
	)
	windows.Items[0].Click.Handle(mf.menuControls_Click)
	windows.Items[1].Click.Handle(mf.menuData_Click)
	windows.Items[2].Click.Handle(mf.menuLayout_Click)
	windows.Items[3].Click.Handle(mf.menuDialogs_Click)

	help := goforms.NewTopMenu("Help", goforms.NewMenuItem("About..."))
	help.Items[0].Click.Handle(mf.menuAbout_Click)

	mf.SetMainMenu(goforms.NewMenuStrip(file, windows, help))
}

func (mf *MainForm) buildToolStrip() {
	mf.toolStrip.AddButton("Controls", func() { mf.showControls() })
	mf.toolStrip.AddButton("Data", func() { mf.showData() })
	mf.toolStrip.AddButton("Layout", func() { mf.showLayout() })
	mf.toolStrip.AddButton("Dialogs", func() { mf.showDialogs() })
	mf.toolStrip.AddSeparator()
	mf.toolStrip.AddButton("Clear log", func() { mf.clearLog() })
}

// buildLogContextMenu shows off Control.SetContextMenu on a ListBox - one of
// the interactive controls that could never carry a context menu before the
// interaction overlay landed.
func (mf *MainForm) buildLogContextMenu() {
	clear := goforms.NewMenuItem("Clear log")
	clear.Click.Handle(func(any, goforms.EventArgs) { mf.clearLog() })

	pause := goforms.NewMenuItem("Pause / resume")
	pause.Click.Handle(func(any, goforms.EventArgs) {
		mf.chkPause.SetChecked(!mf.chkPause.Checked())
	})

	count := goforms.NewMenuItem("Show total")
	count.Click.Handle(func(any, goforms.EventArgs) {
		mf.setStatus(0, fmt.Sprintf("%d events since start", total))
	})

	mf.lstLog.SetContextMenu(goforms.NewContextMenu(
		clear, pause, goforms.NewMenuSeparator(), count,
	))
}

// --- log ------------------------------------------------------------------

// appendEntry runs on the UI goroutine (that is where events are raised), so
// it could touch controls directly - but it deliberately does not.
//
// Every Refresh marks the canvas dirty and costs a full repaint of the
// window. Redrawing the list once per event made the whole app crawl as soon
// as anything fired rapidly, so entries are only queued here and painted by
// flushLog on a timer. Under a burst that turns hundreds of repaints a
// second into ten.
func (mf *MainForm) appendEntry(e eventlog.Entry) {
	total++
	if paused {
		return
	}
	pending = append(pending, e.String())
	lastMsg = e.Source + " " + e.Event
}

// flushLog paints whatever has accumulated since the last tick.
func (mf *MainForm) flushLog() {
	if len(pending) == 0 {
		return
	}
	logLines = append(logLines, pending...)
	pending = pending[:0]
	if len(logLines) > maxLogLines {
		logLines = logLines[len(logLines)-maxLogLines:]
	}
	mf.lstLog.SetItems(logLines)
	mf.lblCount.SetText(fmt.Sprintf("%d events", total))
	mf.setStatus(0, lastMsg)
}

func (mf *MainForm) clearLog() {
	logLines = nil
	pending = pending[:0]
	mf.lstLog.SetItems(nil)
	mf.setStatus(0, "Log cleared")
}

func (mf *MainForm) setStatus(panel int, text string) {
	panels := mf.statusStrip.Panels()
	if panel < len(panels) {
		panels[panel].SetText(text)
	}
}

// --- launchers ------------------------------------------------------------

func (mf *MainForm) showControls() {
	if open.controls == nil {
		open.controls = controlsform.NewControlsForm()
		open.controls.Closed.Handle(func(any, goforms.EventArgs) { open.controls = nil })
	}
	open.controls.Show()
}

func (mf *MainForm) showData() {
	if open.data == nil {
		open.data = dataform.NewDataForm()
		open.data.Closed.Handle(func(any, goforms.EventArgs) { open.data = nil })
	}
	open.data.Show()
}

func (mf *MainForm) showLayout() {
	if open.layout == nil {
		open.layout = layoutform.NewLayoutForm()
		open.layout.Closed.Handle(func(any, goforms.EventArgs) { open.layout = nil })
	}
	open.layout.Show()
}

func (mf *MainForm) showDialogs() {
	if open.dialogs == nil {
		open.dialogs = dialogsform.NewDialogsForm()
		open.dialogs.Closed.Handle(func(any, goforms.EventArgs) { open.dialogs = nil })
	}
	open.dialogs.Show()
}

// showAbout opens a *modal* form and waits for its DialogResult, which is
// the other half of "switching forms".
//
// ShowDialog blocks, so it must run off the UI goroutine - otherwise it
// would block the very thread that has to draw the dialog. fyne.Do hops
// back to the UI thread with the answer.
func (mf *MainForm) showAbout() {
	go func() {
		about := aboutform.NewAboutForm()
		result := about.ShowDialog()
		fyne.Do(func() {
			eventlog.Logf("Main/About", "DialogResult", "%s", resultName(result))
		})
	}()
}

func resultName(r goforms.DialogResult) string {
	switch r {
	case goforms.DialogOK:
		return "OK"
	case goforms.DialogCancel:
		return "Cancel"
	case goforms.DialogYes:
		return "Yes"
	case goforms.DialogNo:
		return "No"
	}
	return "None"
}

// --- designer-wired handlers ---------------------------------------------

func (mf *MainForm) btnControls_Click(sender any, e goforms.MouseEventArgs) { mf.showControls() }
func (mf *MainForm) btnData_Click(sender any, e goforms.MouseEventArgs)     { mf.showData() }
func (mf *MainForm) btnLayout_Click(sender any, e goforms.MouseEventArgs)   { mf.showLayout() }
func (mf *MainForm) btnDialogs_Click(sender any, e goforms.MouseEventArgs)  { mf.showDialogs() }
func (mf *MainForm) btnAbout_Click(sender any, e goforms.MouseEventArgs)    { mf.showAbout() }
func (mf *MainForm) btnClear_Click(sender any, e goforms.MouseEventArgs)    { mf.clearLog() }

func (mf *MainForm) logTimer_Tick(sender any, e goforms.EventArgs) { mf.flushLog() }

func (mf *MainForm) chkMotion_CheckedChanged(sender any, e goforms.EventArgs) {
	eventlog.SetMotionEnabled(mf.chkMotion.Checked())
	if mf.chkMotion.Checked() {
		mf.setStatus(0, "Logging MouseMove/MouseWheel - expect a lot of entries")
	} else {
		mf.setStatus(0, "Motion logging off")
	}
}

func (mf *MainForm) chkPause_CheckedChanged(sender any, e goforms.EventArgs) {
	paused = mf.chkPause.Checked()
	if paused {
		mf.setStatus(0, "Logging paused")
	} else {
		mf.setStatus(0, "Logging resumed")
	}
}

func (mf *MainForm) lstLog_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	if item := mf.lstLog.SelectedItem(); item != "" {
		mf.setStatus(0, item)
	}
}

func (mf *MainForm) clock_Tick(sender any, e goforms.EventArgs) {
	mf.setStatus(1, time.Now().Format("15:04:05"))
}

func (mf *MainForm) menuClear_Click(sender any, e goforms.EventArgs)    { mf.clearLog() }
func (mf *MainForm) menuExit_Click(sender any, e goforms.EventArgs)     { mf.Close() }
func (mf *MainForm) menuControls_Click(sender any, e goforms.EventArgs) { mf.showControls() }
func (mf *MainForm) menuData_Click(sender any, e goforms.EventArgs)     { mf.showData() }
func (mf *MainForm) menuLayout_Click(sender any, e goforms.EventArgs)   { mf.showLayout() }
func (mf *MainForm) menuDialogs_Click(sender any, e goforms.EventArgs)  { mf.showDialogs() }
func (mf *MainForm) menuAbout_Click(sender any, e goforms.EventArgs)    { mf.showAbout() }

func (mf *MainForm) toolStrip_Button4() {
}
