package dialogsform

import "github.com/Go-Forms/GoForms"

// DialogsForm shows every dialog GoForms wraps. Fyne parents a dialog to a
// window, and this form is that window - so each dialog appears over the
// button that opened it.
type DialogsForm struct {
	*goforms.Form

	grpMessage *goforms.GroupBox
	btnInfo    *goforms.Button
	btnOKCanc  *goforms.Button
	btnYesNo   *goforms.Button
	btnError   *goforms.Button
	btnInput   *goforms.Button

	grpFiles  *goforms.GroupBox
	btnOpen   *goforms.Button
	btnSave   *goforms.Button
	btnFolder *goforms.Button
	txtResult *goforms.TextBox

	grpColour *goforms.GroupBox
	btnColour *goforms.ColorPickerButton
	pnlSwatch *goforms.Panel

	lblLast *goforms.Label
}

// NewDialogsForm mirrors `new DialogsForm()`.
func NewDialogsForm() *DialogsForm {
	df := &DialogsForm{Form: goforms.NewForm("Showcase - Dialogs", 620, 470)}
	df.initializeComponent()
	return df
}

func (df *DialogsForm) initializeComponent() {
	df.SetClientSize(620, 470)

	df.grpMessage = goforms.NewGroupBox("Message boxes", 596, 130)
	df.grpMessage.SetBounds(12, 12, 596, 130)
	df.grpMessage.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	df.AddControl(df.grpMessage)

	df.btnInfo = goforms.NewButton("Information")
	df.btnInfo.SetBounds(16, 34, 130, 32)
	df.btnInfo.Click.Handle(df.btnInfo_Click)
	df.grpMessage.AddControl(df.btnInfo)

	df.btnOKCanc = goforms.NewButton("OK / Cancel")
	df.btnOKCanc.SetBounds(156, 34, 130, 32)
	df.btnOKCanc.Click.Handle(df.btnOKCanc_Click)
	df.grpMessage.AddControl(df.btnOKCanc)

	df.btnYesNo = goforms.NewButton("Yes / No")
	df.btnYesNo.SetBounds(296, 34, 130, 32)
	df.btnYesNo.Click.Handle(df.btnYesNo_Click)
	df.grpMessage.AddControl(df.btnYesNo)

	df.btnError = goforms.NewButton("Error")
	df.btnError.SetBounds(436, 34, 130, 32)
	df.btnError.Click.Handle(df.btnError_Click)
	df.grpMessage.AddControl(df.btnError)

	df.btnInput = goforms.NewButton("Input box...")
	df.btnInput.SetBounds(16, 76, 200, 32)
	df.btnInput.Click.Handle(df.btnInput_Click)
	df.grpMessage.AddControl(df.btnInput)

	df.grpFiles = goforms.NewGroupBox("File and folder pickers", 596, 150)
	df.grpFiles.SetBounds(12, 154, 596, 150)
	df.grpFiles.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	df.AddControl(df.grpFiles)

	df.btnOpen = goforms.NewButton("Open file...")
	df.btnOpen.SetBounds(16, 34, 150, 32)
	df.btnOpen.Click.Handle(df.btnOpen_Click)
	df.grpFiles.AddControl(df.btnOpen)

	df.btnSave = goforms.NewButton("Save file...")
	df.btnSave.SetBounds(176, 34, 150, 32)
	df.btnSave.Click.Handle(df.btnSave_Click)
	df.grpFiles.AddControl(df.btnSave)

	df.btnFolder = goforms.NewButton("Choose folder...")
	df.btnFolder.SetBounds(336, 34, 150, 32)
	df.btnFolder.Click.Handle(df.btnFolder_Click)
	df.grpFiles.AddControl(df.btnFolder)

	df.txtResult = goforms.NewTextBox()
	df.txtResult.SetBounds(16, 78, 564, 30)
	df.txtResult.SetPlaceholder("the chosen path appears here")
	df.txtResult.SetReadOnly(true)
	df.txtResult.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	df.grpFiles.AddControl(df.txtResult)

	df.grpColour = goforms.NewGroupBox("Colour picker", 596, 100)
	df.grpColour.SetBounds(12, 316, 596, 100)
	df.grpColour.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight)
	df.AddControl(df.grpColour)

	df.btnColour = goforms.NewColorPickerButton("Pick a colour", df.Form)
	df.btnColour.SetBounds(16, 34, 160, 32)
	df.btnColour.ColorChanged.Handle(df.btnColour_ColorChanged)
	df.grpColour.AddControl(df.btnColour)

	df.pnlSwatch = goforms.NewPanel(120, 32)
	df.pnlSwatch.SetBounds(190, 34, 120, 32)
	df.grpColour.AddControl(df.pnlSwatch)

	df.lblLast = goforms.NewLabel("Dialogs report through a callback, not a blocking return value.")
	df.lblLast.SetBounds(12, 428, 596, 24)
	df.lblLast.SetAnchor(goforms.AnchorBottom | goforms.AnchorLeft | goforms.AnchorRight)
	df.AddControl(df.lblLast)

	df.Load.Handle(df.DialogsForm_Load)
}
