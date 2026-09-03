package dataform

import (
	"fmt"
	"time"

	"goforms"
	"goformsshowcase/eventlog"
)

const src = "Data"

var sampleOrders = [][]string{
	{"1001", "Acme Corp", "Widget", "12", "$1,240.00"},
	{"1002", "Globex", "Gadget", "3", "$87.50"},
	{"1003", "Initech", "Stapler\n(red)", "1", "$3,999.99"},
	{"1004", "Umbrella", "Vial", "40", "$12.00"},
	{"1005", "Soylent", "Ration", "250", "$540.25"},
	{"1006", "Hooli", "Box", "7", "$64.10"},
	{"1007", "Stark Industries", "Reactor", "1", "$25,000.00"},
}

func (df *DataForm) DataForm_Load(sender any, e goforms.EventArgs) {
	for _, row := range sampleOrders {
		df.grid.AddRow(row...)
	}
	df.grid.SetRowHeight(30)

	// Start from the defaults so the combo boxes agree with the grid.
	df.cboColsMode.SetSelectedItem("SizeFixed")
	df.cboRowsMode.SetSelectedItem("SizeFixed")
	df.cboScroll.SetSelectedItem("ScrollBarsBoth")

	df.fillListView()
	df.fillTree()
	df.buildContextMenus()

	eventlog.WireAll(src+"/grid", df.grid)
	eventlog.Log(src+"/Form", "Load", "data form opened")
}

func (df *DataForm) fillListView() {
	now := time.Now()
	files := []struct {
		name string
		size string
	}{
		{"README.md", "11 KB"},
		{"main.go", "344 B"},
		{"go.mod", "1.9 KB"},
		{"datagridview.go", "24 KB"},
		{"interaction.go", "9 KB"},
	}
	for i, f := range files {
		df.listView.AddRow(f.name, f.size, now.AddDate(0, 0, -i).Format("2006-01-02"))
	}
	df.listView.SetColumnWidth(0, 220)
	df.listView.SetColumnWidth(1, 100)
	df.listView.SetColumnWidth(2, 140)
}

func (df *DataForm) fillTree() {
	forms := df.treeView.AddNode(nil, "Forms")
	df.treeView.AddNode(forms, "MainForm")
	controls := df.treeView.AddNode(forms, "ControlsForm")
	df.treeView.AddNode(controls, "ControlsForm.go")
	df.treeView.AddNode(controls, "ControlsForm-designer.go")
	df.treeView.AddNode(forms, "DataForm")

	pkg := df.treeView.AddNode(nil, "eventlog")
	df.treeView.AddNode(pkg, "eventlog.go")
	df.treeView.AddNode(pkg, "wire.go")
}

// buildContextMenus puts right-click menus on the grid and the list - both
// interactive controls, both impossible to give a context menu before the
// interaction overlay.
func (df *DataForm) buildContextMenus() {
	dup := goforms.NewMenuItem("Duplicate selected row")
	dup.Click.Handle(func(any, goforms.EventArgs) { df.duplicateRow() })
	del := goforms.NewMenuItem("Delete selected row")
	del.Click.Handle(func(any, goforms.EventArgs) { df.deleteRow() })
	fit := goforms.NewMenuItem("Fit columns to content")
	fit.Click.Handle(func(any, goforms.EventArgs) {
		df.cboColsMode.SetSelectedItem("SizeToContent")
	})
	df.grid.SetContextMenu(goforms.NewContextMenu(dup, del, goforms.NewMenuSeparator(), fit))

	clear := goforms.NewMenuItem("Clear list")
	clear.Click.Handle(func(any, goforms.EventArgs) {
		df.listView.Clear()
		eventlog.Log(src+"/listView", "ContextMenu", "cleared")
	})
	refill := goforms.NewMenuItem("Refill list")
	refill.Click.Handle(func(any, goforms.EventArgs) {
		df.listView.Clear()
		df.fillListView()
		eventlog.Log(src+"/listView", "ContextMenu", "refilled")
	})
	df.listView.SetContextMenu(goforms.NewContextMenu(clear, refill))
}

func (df *DataForm) duplicateRow() {
	r := df.grid.SelectedRow()
	if r < 0 || r >= df.grid.RowCount() {
		return
	}
	row := df.grid.Rows()[r]
	df.grid.InsertRow(r+1, append([]string(nil), row...)...)
}

func (df *DataForm) deleteRow() {
	if r := df.grid.SelectedRow(); r >= 0 {
		df.grid.RemoveRow(r)
	}
}

// --- designer-wired handlers ---------------------------------------------

func (df *DataForm) grid_CellClick(sender any, e goforms.GridCellEventArgs) {
	eventlog.Logf(src+"/grid", "CellClick", "row %d col %d", e.Row, e.Col)
}

func (df *DataForm) grid_SelectionChanged(sender any, e goforms.GridCellEventArgs) {
	eventlog.Logf(src+"/grid", "SelectionChanged", "row %d col %d", e.Row, e.Col)
}

func (df *DataForm) grid_CellValueChanged(sender any, e goforms.GridCellValueEventArgs) {
	eventlog.Logf(src+"/grid", "CellValueChanged", "[%d,%d] %q -> %q", e.Row, e.Col, e.OldValue, e.NewValue)
}

func (df *DataForm) grid_RowsChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/grid", "RowsChanged", "%d rows", df.grid.RowCount())
}

func (df *DataForm) cboColsMode_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	mode := map[string]goforms.SizeMode{
		"SizeFixed":     goforms.SizeFixed,
		"SizeToContent": goforms.SizeToContent,
		"SizeFill":      goforms.SizeFill,
	}[df.cboColsMode.SelectedItem()]
	df.grid.SetAutoSizeColumnsMode(mode)
	eventlog.Logf(src+"/grid", "ColumnsMode", "%s", df.cboColsMode.SelectedItem())
}

func (df *DataForm) cboRowsMode_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	mode := goforms.SizeFixed
	if df.cboRowsMode.SelectedItem() == "SizeToContent" {
		mode = goforms.SizeToContent
	}
	df.grid.SetAutoSizeRowsMode(mode)
	eventlog.Logf(src+"/grid", "RowsMode", "%s", df.cboRowsMode.SelectedItem())
}

func (df *DataForm) cboScroll_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	bars := map[string]goforms.ScrollBars{
		"ScrollBarsNone":       goforms.ScrollBarsNone,
		"ScrollBarsHorizontal": goforms.ScrollBarsHorizontal,
		"ScrollBarsVertical":   goforms.ScrollBarsVertical,
		"ScrollBarsBoth":       goforms.ScrollBarsBoth,
	}[df.cboScroll.SelectedItem()]
	df.grid.SetScrollBars(bars)
	eventlog.Logf(src+"/grid", "ScrollBars", "%s", df.cboScroll.SelectedItem())
}

func (df *DataForm) numRowH_ValueChanged(sender any, e goforms.EventArgs) {
	h := float32(df.numRowH.Value())
	df.grid.SetRowHeight(h)
	// SetRowHeight also switches rows back to the fixed mode, so keep the
	// combo box honest rather than letting it claim otherwise.
	df.cboRowsMode.SetSelectedItem("SizeFixed")
	eventlog.Logf(src+"/grid", "RowHeight", "%.0f", h)
}

func (df *DataForm) chkReadOnly_CheckedChanged(sender any, e goforms.EventArgs) {
	df.grid.SetReadOnly(df.chkReadOnly.Checked())
	eventlog.Logf(src+"/grid", "ReadOnly", "%v", df.chkReadOnly.Checked())
}

func (df *DataForm) chkLines_CheckedChanged(sender any, e goforms.EventArgs) {
	df.grid.SetGridLines(df.chkLines.Checked())
	eventlog.Logf(src+"/grid", "GridLines", "%v", df.chkLines.Checked())
}

func (df *DataForm) btnAddRow_Click(sender any, e goforms.MouseEventArgs) {
	n := df.grid.RowCount() + 1001
	df.grid.AddRow(fmt.Sprintf("%d", n), "New customer", "New product", "1", "$0.00")
}

func (df *DataForm) btnDelRow_Click(sender any, e goforms.MouseEventArgs) {
	df.deleteRow()
}

func (df *DataForm) listView_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/listView", "SelectedIndexChanged", "row %d", df.listView.SelectedRow())
}

func (df *DataForm) treeView_NodeSelected(sender any, e goforms.EventArgs) {
	if n := df.treeView.SelectedNode(); n != nil {
		eventlog.Logf(src+"/treeView", "NodeSelected", "%s", n.Text)
	}
}
