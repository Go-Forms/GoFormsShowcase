package mainform

import (
	"goforms"
)

// MainForm-designer.go is the generated-looking half of the WinForms-style
// partial-class split: field declarations and initializeComponent() live
// here. Hand-written logic and event handler bodies belong in MainForm.go.
type MainForm struct {
	*goforms.Form

	toolStrip *goforms.ToolStrip

	grpForms     *goforms.GroupBox
	btnControls  *goforms.Button
	btnData      *goforms.Button
	btnLayout    *goforms.Button
	btnDialogs   *goforms.Button
	btnAbout     *goforms.Button
	lblFormsHint *goforms.Label

	grpLog   *goforms.GroupBox
	lstLog   *goforms.ListBox
	btnClear *goforms.Button
	chkPause *goforms.CheckBox
	lblCount *goforms.Label

	chkMotion *goforms.CheckBox

	statusStrip *goforms.StatusStrip
	clock       *goforms.Timer
	// logTimer batches the event log's repaints; see MainForm.appendEntry.
	logTimer *goforms.Timer
	button1  *goforms.Button
	button2  *goforms.Button
}

// NewMainForm mirrors `new MainForm()` - allocate, then initializeComponent.
func NewMainForm() *MainForm {
	mf := &MainForm{Form: goforms.NewForm("GoForms Showcase", 1000, 680)}
	mf.initializeComponent()
	return mf
}

// initializeComponent mirrors the WinForms designer's InitializeComponent():
// pure layout, no business logic.
func (mf *MainForm) initializeComponent() {
	mf.SetClientSize(1000, 680)
	mf.CenterOnScreen()

	mf.toolStrip = goforms.NewToolStrip()
	mf.toolStrip.SetBounds(0, 0, 1000, 34)
	mf.toolStrip.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	mf.toolStrip.AddButton("Button 4", mf.toolStrip_Button4)
	mf.toolStrip.AddSeparator()
	mf.toolStrip.AddSeparator()
	mf.toolStrip.AddButton("Button 1", nil)
	mf.AddControl(mf.toolStrip)

	// --- demo form launcher ---------------------------------------------
	mf.grpForms = goforms.NewGroupBox("Demo forms", 230, 300)
	mf.grpForms.SetBounds(12, 46, 230, 300)
	mf.grpForms.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	mf.AddControl(mf.grpForms)

	mf.btnControls = goforms.NewButton("Controls...")
	mf.btnControls.SetBounds(16, 34, 198, 32)
	mf.btnControls.Click.Handle(mf.btnControls_Click)
	mf.grpForms.AddControl(mf.btnControls)

	mf.btnData = goforms.NewButton("Data & grids...")
	mf.btnData.SetBounds(16, 72, 198, 32)
	mf.btnData.Click.Handle(mf.btnData_Click)
	mf.grpForms.AddControl(mf.btnData)

	mf.btnLayout = goforms.NewButton("Layout & anchoring...")
	mf.btnLayout.SetBounds(16, 110, 198, 32)
	mf.btnLayout.Click.Handle(mf.btnLayout_Click)
	mf.grpForms.AddControl(mf.btnLayout)

	mf.btnDialogs = goforms.NewButton("Dialogs...")
	mf.btnDialogs.SetBounds(16, 148, 198, 32)
	mf.btnDialogs.Click.Handle(mf.btnDialogs_Click)
	mf.grpForms.AddControl(mf.btnDialogs)

	mf.btnAbout = goforms.NewButton("About (modal)...")
	mf.btnAbout.SetBounds(16, 194, 198, 32)
	mf.btnAbout.Click.Handle(mf.btnAbout_Click)
	mf.grpForms.AddControl(mf.btnAbout)

	mf.lblFormsHint = goforms.NewLabel("The first four open modeless,\nso several can be open at once.")
	mf.lblFormsHint.SetBounds(16, 236, 198, 48)
	mf.grpForms.AddControl(mf.lblFormsHint)

	// --- live event log --------------------------------------------------
	mf.grpLog = goforms.NewGroupBox("Event log", 734, 546)
	mf.grpLog.SetBounds(254, 46, 734, 546)
	mf.grpLog.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	mf.AddControl(mf.grpLog)

	mf.lstLog = goforms.NewListBox()
	mf.lstLog.SetTabStop(true)
	mf.lstLog.SetBounds(12, 32, 710, 466)
	mf.lstLog.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	mf.lstLog.SelectedIndexChanged.Handle(mf.lstLog_SelectedIndexChanged)
	mf.grpLog.AddControl(mf.lstLog)

	mf.btnClear = goforms.NewButton("Clear")
	mf.btnClear.SetBounds(12, 506, 90, 30)
	mf.btnClear.SetAnchor(goforms.AnchorBottom | goforms.AnchorLeft)
	mf.btnClear.Click.Handle(mf.btnClear_Click)
	mf.grpLog.AddControl(mf.btnClear)

	mf.chkPause = goforms.NewCheckBox("Pause logging")
	mf.chkPause.SetChecked(false)
	mf.chkPause.SetBounds(114, 508, 150, 26)
	mf.chkPause.SetAnchor(goforms.AnchorBottom | goforms.AnchorLeft)
	mf.chkPause.CheckedChanged.Handle(mf.chkPause_CheckedChanged)
	mf.grpLog.AddControl(mf.chkPause)

	mf.chkMotion = goforms.NewCheckBox("Log mouse motion")
	mf.chkMotion.SetChecked(false)
	mf.chkMotion.SetBounds(272, 508, 170, 26)
	mf.chkMotion.SetAnchor(goforms.AnchorBottom | goforms.AnchorLeft)
	mf.chkMotion.CheckedChanged.Handle(mf.chkMotion_CheckedChanged)
	mf.grpLog.AddControl(mf.chkMotion)

	mf.lblCount = goforms.NewLabel("0 events")
	mf.lblCount.SetBounds(456, 510, 200, 24)
	mf.lblCount.SetAnchor(goforms.AnchorBottom | goforms.AnchorLeft)
	mf.grpLog.AddControl(mf.lblCount)

	// --- status bar -------------------------------------------------------
	mf.statusStrip = goforms.NewStatusStrip(1000, 28)
	mf.statusStrip.SetBounds(0, 652, 1000, 28)
	mf.statusStrip.SetAnchor(goforms.AnchorBottom | goforms.AnchorLeft | goforms.AnchorRight)
	mf.AddControl(mf.statusStrip)

	mf.clock = goforms.NewTimer(1000)
	mf.clock.Tick.Handle(mf.clock_Tick)

	mf.logTimer = goforms.NewTimer(100)
	mf.logTimer.Tick.Handle(mf.logTimer_Tick)

	mf.Load.Handle(mf.MainForm_Load)
	mf.Closing.Handle(mf.MainForm_Closing)
	mf.Closed.Handle(mf.MainForm_Closed)
	mf.Resize.Handle(mf.MainForm_Resize)
	mf.KeyDown.Handle(mf.MainForm_KeyDown)
	mf.button1 = goforms.NewButton("Button")
	mf.button1.SetBounds(92, 401, 90, 30)
	mf.AddControl(mf.button1)
	mf.button2 = goforms.NewButton("Button")
	mf.button2.SetBounds(92, 456, 90, 30)
	mf.AddControl(mf.button2)
}
