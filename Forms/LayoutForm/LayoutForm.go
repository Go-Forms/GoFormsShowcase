package layoutform

import (
	"fmt"
	"github.com/Go-Forms/GoForms"

	"goformsshowcase/eventlog"
)

const src = "Layout"

func (lf *LayoutForm) LayoutForm_Load(sender any, e goforms.EventArgs) {
	lf.paintDockPanels()
	lf.fillFlow()
	lf.fillTable()
	lf.fillSplitAndTabs()

	lf.cboFlowDir.SetSelectedItem("FlowLeftToRight")

	eventlog.WireAll(src+"/dockFill", lf.dockFill)
	eventlog.Log(src+"/Form", "Load", "layout form opened")
}

func (lf *LayoutForm) LayoutForm_Resize(sender any, e goforms.EventArgs) {
	size := lf.ClientSize()
	eventlog.Logf(src+"/Form", "Resize", "%.0f x %.0f", size.Width, size.Height)
}

// paintDockPanels gives each docked slab its own colour, so which one took
// which part of the client area is obvious at a glance.
func (lf *LayoutForm) paintDockPanels() {
	lf.dockTop.SetBackColor(goforms.RGB(60, 110, 180))
	lf.dockFoot.SetBackColor(goforms.RGB(60, 140, 120))
	lf.dockLeft.SetBackColor(goforms.RGB(160, 90, 60))
	lf.dockRight.SetBackColor(goforms.RGB(140, 70, 150))
	lf.dockFill.SetBackColor(goforms.RGB(70, 70, 80))

	label := func(p *goforms.Panel, text string) {
		l := goforms.NewLabel(text)
		l.SetBounds(6, 4, 140, 22)
		l.SetForeColor(goforms.RGB(255, 255, 255))
		p.AddControl(l)
	}
	label(lf.dockTop, "DockTop")
	label(lf.dockFoot, "DockBottom")
	label(lf.dockLeft, "DockLeft")
	label(lf.dockRight, "DockRight")
	label(lf.dockFill, "DockFill")
}

// fillFlow drops a handful of differently sized buttons into the flow panel
// so wrapping and direction changes are visible.
func (lf *LayoutForm) fillFlow() {
	widths := []float32{70, 110, 90, 130, 80, 100, 120}
	for i, w := range widths {
		b := goforms.NewButton(fmt.Sprintf("Item %d", i+1))
		b.SetBounds(0, 0, w, 30)
		lf.flow.AddControl(b)
	}
}

// fillTable styles the tracks and drops one labelled panel per cell, plus a
// spanning cell across the bottom row.
func (lf *LayoutForm) fillTable() {
	lf.table.SetColumnStyle(0, goforms.Absolute(110))
	lf.table.SetColumnStyle(1, goforms.Percent(40))
	lf.table.SetColumnStyle(2, goforms.AutoSize())
	lf.table.SetRowStyle(0, goforms.Absolute(44))
	lf.table.SetRowStyle(1, goforms.AutoSize())
	lf.table.SetRowStyle(2, goforms.AutoSize())

	titles := []string{"Absolute 110", "Percent 40%", "AutoSize", "r1c1", "r1c2", "r1c3"}
	for _, t := range titles {
		p := goforms.NewPanel(10, 10)
		p.SetBounds(0, 0, 10, 10)
		p.SetBackColor(goforms.RGB(70, 78, 92))
		l := goforms.NewLabel(t)
		l.SetBounds(6, 6, 130, 22)
		p.AddControl(l)
		lf.table.AddControl(p)
	}

	wide := goforms.NewPanel(10, 10)
	wide.SetBounds(0, 0, 10, 10)
	wide.SetBackColor(goforms.RGB(90, 70, 60))
	l := goforms.NewLabel("spans all three columns")
	l.SetBounds(6, 6, 300, 22)
	wide.AddControl(l)
	lf.table.AddControlSpanning(wide, 0, 2, 3, 1)
}

func (lf *LayoutForm) fillSplitAndTabs() {
	lf.split.SetSplitterDistance(0.42)

	notes := goforms.NewRichTextBox("### SplitContainer\n" +
		"Drag the bar. Both halves are real `Panel`s, so they take child\n" +
		"controls like any other container - the tabs on the right live in\n" +
		"`Panel2`.")
	notes.SetBounds(8, 8, 400, 180)
	notes.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	lf.split.Panel1().AddControl(notes)

	general := lf.tabs.AddTab("General")
	chk := goforms.NewCheckBox("A checkbox on a tab page")
	chk.SetBounds(16, 16, 260, 26)
	general.AddControl(chk)

	sliders := lf.tabs.AddTab("Sliders")
	tb := goforms.NewTrackBar(0, 100)
	tb.SetBounds(16, 20, 300, 30)
	sliders.AddControl(tb)

	about := lf.tabs.AddTab("About")
	lbl := goforms.NewLabel("TabControl pages are containers too.")
	lbl.SetBounds(16, 16, 320, 24)
	about.AddControl(lbl)
}

// --- designer-wired handlers ---------------------------------------------

func (lf *LayoutForm) cboFlowDir_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	dir := map[string]goforms.FlowDirection{
		"FlowLeftToRight": goforms.FlowLeftToRight,
		"FlowTopDown":     goforms.FlowTopDown,
		"FlowRightToLeft": goforms.FlowRightToLeft,
		"FlowBottomUp":    goforms.FlowBottomUp,
	}[lf.cboFlowDir.SelectedItem()]
	lf.flow.SetFlowDirection(dir)
	eventlog.Logf(src+"/flow", "FlowDirection", "%s", lf.cboFlowDir.SelectedItem())
}

func (lf *LayoutForm) chkWrap_CheckedChanged(sender any, e goforms.EventArgs) {
	lf.flow.SetWrapContents(lf.chkWrap.Checked())
	eventlog.Logf(src+"/flow", "WrapContents", "%v", lf.chkWrap.Checked())
}

func (lf *LayoutForm) tabs_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	if p := lf.tabs.SelectedTab(); p != nil {
		eventlog.Logf(src+"/tabs", "SelectedIndexChanged", "%s", p.Text())
	}
}

// btnDockFill_Click is wired to the button docked into the DockFill panel, so
// clicking it shows that a docked control takes input where it is drawn,
// not where its stored bounds would put it.
func (lf *LayoutForm) btnDockFill_Click(sender any, e goforms.MouseEventArgs) {
	eventlog.Log(src+"/btnDockFill", "Click", "the DockFill button")
}
