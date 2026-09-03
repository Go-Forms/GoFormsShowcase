package dataform

import "goforms"

// DataForm shows the data controls, with the DataGridView's sizing and
// scrolling flags switchable at runtime so their effect is visible rather
// than described.
type DataForm struct {
	*goforms.Form

	grpGrid     *goforms.GroupBox
	grid        *goforms.DataGridView
	lblColsMode *goforms.Label
	cboColsMode *goforms.ComboBox
	lblRowsMode *goforms.Label
	cboRowsMode *goforms.ComboBox
	lblScroll   *goforms.Label
	cboScroll   *goforms.ComboBox
	numRowH     *goforms.NumericUpDown
	lblRowH     *goforms.Label
	chkReadOnly *goforms.CheckBox
	chkLines    *goforms.CheckBox
	btnAddRow   *goforms.Button
	btnDelRow   *goforms.Button

	grpList  *goforms.GroupBox
	listView *goforms.ListView

	grpTree  *goforms.GroupBox
	treeView *goforms.TreeView
}

// NewDataForm mirrors `new DataForm()`.
func NewDataForm() *DataForm {
	df := &DataForm{Form: goforms.NewForm("Showcase - Data & grids", 1100, 720)}
	df.initializeComponent()
	return df
}

func (df *DataForm) initializeComponent() {
	df.SetClientSize(1100, 720)

	// --- DataGridView -----------------------------------------------------
	df.grpGrid = goforms.NewGroupBox("DataGridView - editable, with live sizing flags", 1076, 420)
	df.grpGrid.SetBounds(12, 12, 1076, 420)
	df.grpGrid.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	df.AddControl(df.grpGrid)

	df.grid = goforms.NewDataGridView("Order", "Customer", "Product", "Qty", "Total")
	df.grid.SetBounds(16, 34, 1044, 280)
	df.grid.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	df.grid.CellClick.Handle(df.grid_CellClick)
	df.grid.SelectionChanged.Handle(df.grid_SelectionChanged)
	df.grid.CellValueChanged.Handle(df.grid_CellValueChanged)
	df.grid.RowsChanged.Handle(df.grid_RowsChanged)
	df.grpGrid.AddControl(df.grid)

	df.lblColsMode = goforms.NewLabel("Columns:")
	df.lblColsMode.SetBounds(16, 328, 70, 24)
	df.grpGrid.AddControl(df.lblColsMode)

	df.cboColsMode = goforms.NewComboBox("SizeFixed", "SizeToContent", "SizeFill")
	df.cboColsMode.SetBounds(90, 324, 150, 30)
	df.cboColsMode.SelectedIndexChanged.Handle(df.cboColsMode_SelectedIndexChanged)
	df.grpGrid.AddControl(df.cboColsMode)

	df.lblRowsMode = goforms.NewLabel("Rows:")
	df.lblRowsMode.SetBounds(256, 328, 50, 24)
	df.grpGrid.AddControl(df.lblRowsMode)

	df.cboRowsMode = goforms.NewComboBox("SizeFixed", "SizeToContent")
	df.cboRowsMode.SetBounds(306, 324, 150, 30)
	df.cboRowsMode.SelectedIndexChanged.Handle(df.cboRowsMode_SelectedIndexChanged)
	df.grpGrid.AddControl(df.cboRowsMode)

	df.lblScroll = goforms.NewLabel("Scrollbars:")
	df.lblScroll.SetBounds(472, 328, 80, 24)
	df.grpGrid.AddControl(df.lblScroll)

	df.cboScroll = goforms.NewComboBox("ScrollBarsNone", "ScrollBarsHorizontal", "ScrollBarsVertical", "ScrollBarsBoth")
	df.cboScroll.SetBounds(556, 324, 190, 30)
	df.cboScroll.SelectedIndexChanged.Handle(df.cboScroll_SelectedIndexChanged)
	df.grpGrid.AddControl(df.cboScroll)

	df.lblRowH = goforms.NewLabel("Row height:")
	df.lblRowH.SetBounds(762, 328, 90, 24)
	df.grpGrid.AddControl(df.lblRowH)

	df.numRowH = goforms.NewNumericUpDown(16, 80, 30)
	df.numRowH.SetBounds(852, 324, 100, 30)
	df.numRowH.ValueChanged.Handle(df.numRowH_ValueChanged)
	df.grpGrid.AddControl(df.numRowH)

	df.chkReadOnly = goforms.NewCheckBox("Read only")
	df.chkReadOnly.SetBounds(16, 366, 120, 26)
	df.chkReadOnly.CheckedChanged.Handle(df.chkReadOnly_CheckedChanged)
	df.grpGrid.AddControl(df.chkReadOnly)

	df.chkLines = goforms.NewCheckBox("Grid lines")
	df.chkLines.SetBounds(144, 366, 120, 26)
	df.chkLines.SetChecked(true)
	df.chkLines.CheckedChanged.Handle(df.chkLines_CheckedChanged)
	df.grpGrid.AddControl(df.chkLines)

	df.btnAddRow = goforms.NewButton("Add row")
	df.btnAddRow.SetBounds(280, 362, 100, 32)
	df.btnAddRow.Click.Handle(df.btnAddRow_Click)
	df.grpGrid.AddControl(df.btnAddRow)

	df.btnDelRow = goforms.NewButton("Delete row")
	df.btnDelRow.SetBounds(388, 362, 110, 32)
	df.btnDelRow.Click.Handle(df.btnDelRow_Click)
	df.grpGrid.AddControl(df.btnDelRow)

	// --- ListView ----------------------------------------------------------
	df.grpList = goforms.NewGroupBox("ListView - read-only report view", 530, 250)
	df.grpList.SetBounds(12, 444, 530, 250)
	df.grpList.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorBottom)
	df.AddControl(df.grpList)

	df.listView = goforms.NewListView([]string{"File", "Size", "Modified"})
	df.listView.SetBounds(16, 34, 498, 196)
	df.listView.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	df.listView.SelectedIndexChanged.Handle(df.listView_SelectedIndexChanged)
	df.grpList.AddControl(df.listView)

	// --- TreeView -----------------------------------------------------------
	df.grpTree = goforms.NewGroupBox("TreeView", 534, 250)
	df.grpTree.SetBounds(554, 444, 534, 250)
	df.grpTree.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	df.AddControl(df.grpTree)

	df.treeView = goforms.NewTreeView()
	df.treeView.SetBounds(16, 34, 502, 196)
	df.treeView.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	df.treeView.NodeSelected.Handle(df.treeView_NodeSelected)
	treeViewNode1 := df.treeView.AddNode(nil, "Node 1")
	treeViewNode2 := df.treeView.AddNode(treeViewNode1, "Node 2")
	df.treeView.AddNode(treeViewNode2, "Node 3")
	df.grpTree.AddControl(df.treeView)

	df.Load.Handle(df.DataForm_Load)
}
