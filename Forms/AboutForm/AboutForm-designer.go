package aboutform

import "github.com/Go-Forms/GoForms"

// AboutForm is the modal half of "switching forms": it is opened with
// ShowDialog and reports an OK/Cancel DialogResult back to its caller.
type AboutForm struct {
	*goforms.Form

	lblTitle *goforms.Label
	rtbInfo  *goforms.RichTextBox
	btnOK    *goforms.Button
	btnCanc  *goforms.Button
}

// NewAboutForm mirrors `new AboutForm()`.
func NewAboutForm() *AboutForm {
	af := &AboutForm{Form: goforms.NewForm("About GoForms", 480, 330)}
	af.initializeComponent()
	return af
}

func (af *AboutForm) initializeComponent() {
	af.SetClientSize(480, 330)
	af.CenterOnScreen()

	af.lblTitle = goforms.NewLabel("GoForms Showcase")
	af.lblTitle.SetBounds(20, 16, 440, 34)
	af.AddControl(af.lblTitle)

	af.rtbInfo = goforms.NewRichTextBox("")
	af.rtbInfo.SetBounds(20, 58, 440, 200)
	af.rtbInfo.SetAnchor(goforms.AnchorTop | goforms.AnchorLeft | goforms.AnchorRight | goforms.AnchorBottom)
	af.AddControl(af.rtbInfo)

	af.btnOK = goforms.NewButton("OK")
	af.btnOK.SetBounds(268, 274, 90, 34)
	af.btnOK.SetAnchor(goforms.AnchorBottom | goforms.AnchorRight)
	af.btnOK.Click.Handle(af.btnOK_Click)
	af.AddControl(af.btnOK)

	af.btnCanc = goforms.NewButton("Cancel")
	af.btnCanc.SetBounds(370, 274, 90, 34)
	af.btnCanc.SetAnchor(goforms.AnchorBottom | goforms.AnchorRight)
	af.btnCanc.Click.Handle(af.btnCanc_Click)
	af.AddControl(af.btnCanc)

	af.Load.Handle(af.AboutForm_Load)
}
