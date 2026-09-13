package grideditrow

import goforms "github.com/Go-Forms/GoForms"

// GridEditRow-designer.go is the generated-looking half of the WinForms-style
// partial-class split: field declarations and initializeComponent() live
// here. Hand-written logic and event handler bodies belong in GridEditRow.go.
// The GoForms Designer regenerates only this file.
//
// The five field rows stand for the five columns of the DataForm's grid, in
// the grid's own order - see the Field* constants in GridEditRow.go, which
// are what the caller indexes by.
type GridEditRow struct {
	*goforms.Form
	lblCaption  *goforms.Label
	lblOrder    *goforms.Label
	txtOrder    *goforms.TextBox
	lblCustomer *goforms.Label
	txtCustomer *goforms.TextBox
	lblProduct  *goforms.Label
	txtProduct  *goforms.TextBox
	lblQty      *goforms.Label
	txtQty      *goforms.TextBox
	lblTotal    *goforms.Label
	txtTotal    *goforms.TextBox
	lblHint     *goforms.Label
	btnOK       *goforms.Button
	btnCancel   *goforms.Button
}

// NewGridEditRow mirrors "new GridEditRow()" - allocate, then initializeComponent.
func NewGridEditRow() *GridEditRow {
	f := &GridEditRow{Form: goforms.NewForm("Edit row", 340, 356)}
	f.initializeComponent()
	return f
}

// initializeComponent mirrors the WinForms designer's InitializeComponent():
// pure layout, no business logic.
func (f *GridEditRow) initializeComponent() {
	f.SetClientSize(340, 356)
	f.CenterOnScreen()

	f.lblCaption = goforms.NewLabel("Edit the selected order")
	f.lblCaption.SetBounds(16, 14, 308, 24)
	f.lblCaption.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.lblCaption)

	f.lblOrder = goforms.NewLabel("Order:")
	f.lblOrder.SetBounds(16, 52, 84, 24)
	f.lblOrder.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	f.AddControl(f.lblOrder)

	f.txtOrder = goforms.NewTextBox()
	f.txtOrder.SetBounds(108, 48, 216, 30)
	f.txtOrder.SetPlaceholder("1008")
	f.txtOrder.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.txtOrder)

	f.lblCustomer = goforms.NewLabel("Customer:")
	f.lblCustomer.SetBounds(16, 86, 84, 24)
	f.lblCustomer.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	f.AddControl(f.lblCustomer)

	f.txtCustomer = goforms.NewTextBox()
	f.txtCustomer.SetBounds(108, 82, 216, 30)
	f.txtCustomer.SetPlaceholder("Acme Corp")
	f.txtCustomer.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.txtCustomer)

	f.lblProduct = goforms.NewLabel("Product:")
	f.lblProduct.SetBounds(16, 120, 84, 24)
	f.lblProduct.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	f.AddControl(f.lblProduct)

	// Multiline, because a Product cell may hold more than one line - the
	// sample data's "Stapler\n(red)" is exactly that - and a single-line box
	// would show only the first line of it.
	f.txtProduct = goforms.NewMultilineTextBox()
	f.txtProduct.SetBounds(108, 116, 216, 56)
	f.txtProduct.SetPlaceholder("Widget")
	f.txtProduct.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.txtProduct)

	f.lblQty = goforms.NewLabel("Qty:")
	f.lblQty.SetBounds(16, 180, 84, 24)
	f.lblQty.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	f.AddControl(f.lblQty)

	f.txtQty = goforms.NewTextBox()
	f.txtQty.SetBounds(108, 176, 216, 30)
	f.txtQty.SetPlaceholder("1")
	f.txtQty.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.txtQty)

	f.lblTotal = goforms.NewLabel("Total:")
	f.lblTotal.SetBounds(16, 214, 84, 24)
	f.lblTotal.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft)
	f.AddControl(f.lblTotal)

	f.txtTotal = goforms.NewTextBox()
	f.txtTotal.SetBounds(108, 210, 216, 30)
	f.txtTotal.SetPlaceholder("$0.00")
	f.txtTotal.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.txtTotal)

	f.lblHint = goforms.NewLabel("Save writes these values back into the grid.")
	f.lblHint.SetBounds(16, 248, 308, 44)
	f.lblHint.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	f.AddControl(f.lblHint)

	f.btnOK = goforms.NewButton("Save")
	f.btnOK.SetBounds(136, 306, 90, 32)
	f.btnOK.SetAnchor(goforms.AnchorBottom | goforms.AnchorRight)
	f.btnOK.Click.Handle(f.btnOK_Click)
	f.AddControl(f.btnOK)

	f.btnCancel = goforms.NewButton("Cancel")
	f.btnCancel.SetBounds(234, 306, 90, 32)
	f.btnCancel.SetAnchor(goforms.AnchorBottom | goforms.AnchorRight)
	f.btnCancel.Click.Handle(f.btnCancel_Click)
	f.AddControl(f.btnCancel)
}
