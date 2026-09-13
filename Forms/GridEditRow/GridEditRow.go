package grideditrow

import (
	"strconv"
	"strings"

	goforms "github.com/Go-Forms/GoForms"
)

// GridEditRow.go is the hand-written half of the partial-class split: event
// handlers and business logic go here. Layout lives in GridEditRow-designer.go.
//
// The dialog deliberately knows nothing about the DataGridView it edits. It
// takes a row in and hands a row back, so the caller keeps hold of the grid
// and of which row it asked about - see DataForm.EditRow_Click. That is what
// lets two of these be open at once on the desktop, where a modal form does
// not freeze its owner, without the second one overwriting the first one's
// row.

// The columns of the row, in the order the grid stores them. Load and Values
// index by these, so a caller passing df.Grid.Rows()[i] straight in needs no
// mapping of its own.
const (
	FieldOrder = iota
	FieldCustomer
	FieldProduct
	FieldQty
	FieldTotal
	FieldCount
)

// hintColor tints lblHint when a field is rejected. ForeColor is not a
// designer property, so the colour is applied here rather than in
// initializeComponent.
var hintColor = goforms.RGB(0xE5, 0x48, 0x4D)

// Load fills the boxes from one grid row. A short row leaves the remaining
// boxes empty, which is what "Add row" then "Edit row" would hand us.
func (f *GridEditRow) Load(row []string) {
	for i, box := range f.boxes() {
		if i < len(row) {
			box.SetText(row[i])
		} else {
			box.SetText("")
		}
	}
}

// Values returns what the boxes hold, in the same column order as Load. It
// is only meaningful after ShowDialog has returned goforms.DialogOK.
func (f *GridEditRow) Values() []string {
	boxes := f.boxes()
	out := make([]string, len(boxes))
	for i, box := range boxes {
		out[i] = strings.TrimSpace(box.Text())
	}
	return out
}

// SetCaption replaces the line above the fields, so the caller can name the
// row being edited rather than leaving the generic wording.
func (f *GridEditRow) SetCaption(text string) { f.lblCaption.SetText(text) }

// boxes lists the text boxes in column order - the single place that ties a
// Field* constant to a control, so adding a column is a one-line change here
// plus the layout in the designer file.
func (f *GridEditRow) boxes() []*goforms.TextBox {
	return []*goforms.TextBox{
		FieldOrder:    f.txtOrder,
		FieldCustomer: f.txtCustomer,
		FieldProduct:  f.txtProduct,
		FieldQty:      f.txtQty,
		FieldTotal:    f.txtTotal,
	}
}

func (f *GridEditRow) btnOK_Click(sender any, e goforms.MouseEventArgs) {
	if problem := f.validate(); problem != "" {
		f.lblHint.SetForeColor(hintColor)
		f.lblHint.SetText(problem)
		return
	}
	f.CloseWithResult(goforms.DialogOK)
}

func (f *GridEditRow) btnCancel_Click(sender any, e goforms.MouseEventArgs) {
	f.CloseWithResult(goforms.DialogCancel)
}

// validate reports the first thing wrong with the form, or "" when the row is
// fit to be written back. The dialog stays open on a problem, which is the
// WinForms way of saying no: the handler simply does not set DialogResult.
func (f *GridEditRow) validate() string {
	v := f.Values()
	if v[FieldOrder] == "" {
		return "Order cannot be empty - it is the row's key."
	}
	if v[FieldCustomer] == "" {
		return "Customer cannot be empty."
	}
	if _, err := strconv.Atoi(v[FieldQty]); err != nil {
		return "Qty must be a whole number, for example 12."
	}
	return ""
}
