package layoutform

import "github.com/Go-Forms/GoForms"

// LayoutForm demonstrates the parts of the framework you can only see by
// resizing the window: anchoring, docking, and the layout panels.
type LayoutForm struct {
	*goforms.Form

	lblHint *goforms.Label

	// docking: four slabs plus a fill, all inside one panel
	grpDock   *goforms.GroupBox
	dockHost  *goforms.Panel
	dockTop   *goforms.Panel
	dockLeft  *goforms.Panel
	dockRight *goforms.Panel
	dockFoot  *goforms.Panel
	dockFill  *goforms.Panel

	// anchoring: one control per rule
	grpAnchor    *goforms.GroupBox
	anchTopLeft  *goforms.Button
	anchTopRight *goforms.Button
	anchBottom   *goforms.Button
	anchStretch  *goforms.TextBox
	anchCentre   *goforms.Button

	// flow and table panels
	grpFlow    *goforms.GroupBox
	flow       *goforms.FlowLayoutPanel
	cboFlowDir *goforms.ComboBox
	chkWrap    *goforms.CheckBox

	grpTable *goforms.GroupBox
	table    *goforms.TableLayoutPanel

	// split container and tabs
	split *goforms.SplitContainer
	tabs  *goforms.TabControl

	// btnDockFill fills the DockFill panel above, and is the proof that a
	// docked control takes input where it is drawn rather than where its
	// stored bounds would put it.
	btnDockFill *goforms.Button
}

// NewLayoutForm mirrors `new LayoutForm()`.
func NewLayoutForm() *LayoutForm {
	lf := &LayoutForm{Form: goforms.NewForm("Showcase - Layout & anchoring", 1080, 760)}
	lf.initializeComponent()
	return lf
}

func (lf *LayoutForm) initializeComponent() {
	lf.SetClientSize(1080, 760)

	lf.lblHint = goforms.NewLabel("Resize this window - everything below reacts.")
	lf.lblHint.SetBounds(12, 10, 600, 24)
	lf.lblHint.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	lf.AddControl(lf.lblHint)

	// --- docking -----------------------------------------------------------
	lf.grpDock = goforms.NewGroupBox("Dock", 520, 260)
	lf.grpDock.SetBounds(12, 40, 520, 260)
	lf.grpDock.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	lf.AddControl(lf.grpDock)

	lf.dockHost = goforms.NewPanel(488, 210)
	lf.dockHost.SetBounds(10, 31, 488, 210)
	lf.dockHost.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	lf.grpDock.AddControl(lf.dockHost)

	// Docked controls carve up the client area in the order they are added:
	// the top slab takes the full width, then the sides take what is left.
	lf.dockTop = goforms.NewPanel(10, 30)
	lf.dockTop.SetBounds(0, 0, 10, 30)
	lf.dockTop.SetDock(goforms.DockTop)
	lf.dockHost.AddControl(lf.dockTop)

	lf.dockFoot = goforms.NewPanel(10, 26)
	lf.dockFoot.SetAnchor(goforms.AnchorTop)
	lf.dockFoot.SetBounds(0, 0, 10, 26)
	lf.dockFoot.SetDock(goforms.DockBottom)
	lf.dockHost.AddControl(lf.dockFoot)

	lf.dockLeft = goforms.NewPanel(90, 10)
	lf.dockLeft.SetBounds(0, 0, 90, 10)
	lf.dockLeft.SetDock(goforms.DockLeft)
	lf.dockHost.AddControl(lf.dockLeft)

	lf.dockRight = goforms.NewPanel(90, 10)
	lf.dockRight.SetBounds(0, 0, 90, 10)
	lf.dockRight.SetDock(goforms.DockRight)
	lf.dockHost.AddControl(lf.dockRight)

	lf.dockFill = goforms.NewPanel(10, 10)
	lf.dockFill.SetBounds(95, 0, 10, 10)
	lf.dockFill.SetDock(goforms.DockFill)
	lf.dockHost.AddControl(lf.dockFill)

	// --- anchoring ----------------------------------------------------------
	lf.grpAnchor = goforms.NewGroupBox("Anchor", 520, 260)
	lf.grpAnchor.SetBounds(544, 40, 520, 260)
	lf.grpAnchor.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	lf.AddControl(lf.grpAnchor)

	lf.anchTopLeft = goforms.NewButton("Top|Left")
	lf.anchTopLeft.SetBounds(16, 34, 110, 30)
	lf.anchTopLeft.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	lf.grpAnchor.AddControl(lf.anchTopLeft)

	lf.anchTopRight = goforms.NewButton("Top|Right")
	lf.anchTopRight.SetBounds(394, 34, 110, 30)
	lf.anchTopRight.SetAnchor(goforms.AnchorTop | goforms.AnchorRight)
	lf.grpAnchor.AddControl(lf.anchTopRight)

	lf.anchStretch = goforms.NewTextBox()
	lf.anchStretch.SetBounds(16, 76, 488, 30)
	lf.anchStretch.SetText("Left|Right - stretches")
	lf.anchStretch.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	lf.grpAnchor.AddControl(lf.anchStretch)

	lf.anchCentre = goforms.NewButton("AnchorNone - floats")
	lf.anchCentre.SetBounds(180, 130, 160, 30)
	lf.anchCentre.SetAnchor(goforms.AnchorNone)
	lf.grpAnchor.AddControl(lf.anchCentre)

	lf.anchBottom = goforms.NewButton("Bottom|Right")
	lf.anchBottom.SetBounds(361, 216, 130, 30)
	lf.anchBottom.SetAnchor(goforms.AnchorBottom | goforms.AnchorRight)
	lf.grpAnchor.AddControl(lf.anchBottom)

	// --- flow panel -----------------------------------------------------------
	lf.grpFlow = goforms.NewGroupBox("FlowLayoutPanel", 520, 220)
	lf.grpFlow.SetBounds(12, 312, 520, 220)
	lf.grpFlow.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	lf.AddControl(lf.grpFlow)

	lf.cboFlowDir = goforms.NewComboBox("FlowLeftToRight", "FlowTopDown", "FlowRightToLeft", "FlowBottomUp")
	lf.cboFlowDir.SetBounds(16, 32, 200, 30)
	lf.cboFlowDir.SelectedIndexChanged.Handle(lf.cboFlowDir_SelectedIndexChanged)
	lf.grpFlow.AddControl(lf.cboFlowDir)

	lf.chkWrap = goforms.NewCheckBox("Wrap contents")
	lf.chkWrap.SetBounds(228, 34, 150, 26)
	lf.chkWrap.SetChecked(true)
	lf.chkWrap.CheckedChanged.Handle(lf.chkWrap_CheckedChanged)
	lf.grpFlow.AddControl(lf.chkWrap)

	lf.flow = goforms.NewFlowLayoutPanel(488, 140)
	lf.flow.SetBounds(16, 70, 488, 140)
	lf.flow.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	lf.grpFlow.AddControl(lf.flow)

	// --- table panel -----------------------------------------------------------
	lf.grpTable = goforms.NewGroupBox("TableLayoutPanel - Absolute / Percent / AutoSize tracks", 520, 220)
	lf.grpTable.SetBounds(544, 312, 520, 220)
	lf.grpTable.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	lf.AddControl(lf.grpTable)

	lf.table = goforms.NewTableLayoutPanel(488, 176, 3, 3)
	lf.table.SetBounds(16, 34, 488, 176)
	lf.table.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	lf.grpTable.AddControl(lf.table)

	// --- split container + tabs -------------------------------------------------
	lf.split = goforms.NewSplitContainer(1052, 200, true)
	lf.split.SetDock(goforms.DockBottom)
	lf.split.SetBounds(0, 556, 1052, 200)
	lf.split.SetAnchor(goforms.AnchorNone)
	lf.AddControl(lf.split)

	lf.tabs = goforms.NewTabControl(500, 190)
	lf.tabs.SetDock(goforms.DockNone)
	lf.tabs.SetBounds(11, 7, 500, 190)
	lf.tabs.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	lf.tabs.SelectedIndexChanged.Handle(lf.tabs_SelectedIndexChanged)
	lf.split.Panel2().AddControl(lf.tabs)

	// The pages themselves are added in LayoutForm.fillSplitAndTabs.

	lf.btnDockFill = goforms.NewButton("Docked to fill - click me")
	lf.btnDockFill.SetBounds(20, 20, 90, 30)
	lf.btnDockFill.SetDock(goforms.DockFill)
	lf.btnDockFill.SetAnchor(goforms.AnchorTop | goforms.AnchorBottom | goforms.AnchorLeft | goforms.AnchorRight)
	lf.btnDockFill.Click.Handle(lf.btnDockFill_Click)
	lf.dockFill.AddControl(lf.btnDockFill)

	lf.Load.Handle(lf.LayoutForm_Load)
	lf.Resize.Handle(lf.LayoutForm_Resize)
}
