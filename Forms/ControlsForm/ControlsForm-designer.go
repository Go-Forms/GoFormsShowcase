package controlsform

import (
	"time"

	"github.com/Go-Forms/GoForms"
)

// ControlsForm shows every input and display control GoForms has, each one
// wired to the shared event log.
type ControlsForm struct {
	*goforms.Form

	// buttons & links
	grpButtons  *goforms.GroupBox
	btnPlain    *goforms.Button
	btnDisabled *goforms.Button
	lnkSite     *goforms.LinkLabel
	btnColor    *goforms.ColorPickerButton
	lblStyled   *goforms.Label

	// text entry
	grpText  *goforms.GroupBox
	txtName  *goforms.TextBox
	txtPass  *goforms.TextBox
	txtNotes *goforms.TextBox
	txtPhone *goforms.MaskedTextBox
	lblMask  *goforms.Label

	// choices
	grpChoice   *goforms.GroupBox
	chkAgree    *goforms.CheckBox
	rbSmall     *goforms.RadioButton
	rbMedium    *goforms.RadioButton
	rbLarge     *goforms.RadioButton
	sizeGroup   *goforms.RadioButtonGroup
	cboCountry  *goforms.ComboBox
	lstFruit    *goforms.ListBox
	clbToppings *goforms.CheckedListBox
	dudPriority *goforms.DomainUpDown

	// ranges
	grpRange   *goforms.GroupBox
	trkVolume  *goforms.TrackBar
	numQty     *goforms.NumericUpDown
	prgLoad    *goforms.ProgressBar
	sbHoriz    *goforms.ScrollBar
	lblRangeHi *goforms.Label

	// dates
	grpDates *goforms.GroupBox
	dtpWhen  *goforms.DateTimePicker
	calMonth *goforms.MonthCalendar

	// display
	grpDisplay *goforms.GroupBox
	picLogo    *goforms.PictureBox
	rtbNotes   *goforms.RichTextBox
}

// NewControlsForm mirrors `new ControlsForm()`.
func NewControlsForm() *ControlsForm {
	cf := &ControlsForm{Form: goforms.NewForm("Showcase - Controls", 1180, 800)}
	cf.initializeComponent()
	return cf
}

func (cf *ControlsForm) initializeComponent() {
	cf.SetClientSize(1180, 800)

	// --- buttons & links -------------------------------------------------
	cf.grpButtons = goforms.NewGroupBox("Buttons, links and styling", 370, 236)
	cf.grpButtons.SetBounds(12, 12, 370, 236)
	cf.AddControl(cf.grpButtons)

	cf.btnPlain = goforms.NewButton("Click me")
	cf.btnPlain.SetBounds(16, 32, 120, 32)
	cf.grpButtons.AddControl(cf.btnPlain)

	cf.btnDisabled = goforms.NewButton("Toggle enabled")
	cf.btnDisabled.SetBounds(148, 32, 150, 32)
	cf.btnDisabled.Click.Handle(cf.btnDisabled_Click)
	cf.grpButtons.AddControl(cf.btnDisabled)

	cf.lnkSite = goforms.NewLinkLabelWithURL("fyne.io (LinkLabel)", "https://fyne.io")
	cf.lnkSite.SetBounds(16, 76, 220, 24)
	cf.grpButtons.AddControl(cf.lnkSite)

	cf.btnColor = goforms.NewColorPickerButton("Pick colour", cf.Form)
	cf.btnColor.SetBounds(16, 110, 140, 32)
	cf.btnColor.SetDialogTitle("Label colour")
	cf.btnColor.ColorChanged.Handle(cf.btnColor_ColorChanged)
	cf.grpButtons.AddControl(cf.btnColor)

	cf.lblStyled = goforms.NewLabel("Styled label (Font + ForeColor)")
	cf.lblStyled.SetBounds(16, 160, 336, 60)
	cf.grpButtons.AddControl(cf.lblStyled)

	// --- text entry -------------------------------------------------------
	cf.grpText = goforms.NewGroupBox("Text entry", 380, 236)
	cf.grpText.SetBounds(394, 12, 380, 236)
	cf.AddControl(cf.grpText)

	cf.txtName = goforms.NewTextBox()
	cf.txtName.SetBounds(16, 32, 200, 30)
	cf.txtName.SetPlaceholder("your name")
	cf.txtName.TextChanged.Handle(cf.txtName_TextChanged)
	cf.grpText.AddControl(cf.txtName)

	cf.txtPass = goforms.NewPasswordTextBox()
	cf.txtPass.SetBounds(16, 68, 200, 30)
	cf.txtPass.SetPlaceholder("password")
	cf.grpText.AddControl(cf.txtPass)

	cf.txtPhone = goforms.NewMaskedTextBox("(000) 000-0000")
	cf.txtPhone.SetBounds(16, 104, 200, 30)
	cf.txtPhone.SetPlaceholder("(000) 000-0000")
	cf.txtPhone.MaskCompleted.Handle(cf.txtPhone_MaskCompleted)
	cf.grpText.AddControl(cf.txtPhone)

	cf.lblMask = goforms.NewLabel("MaskedTextBox: (000) 000-0000")
	cf.lblMask.SetBounds(232, 106, 136, 44)
	cf.grpText.AddControl(cf.lblMask)

	cf.txtNotes = goforms.NewMultilineTextBox()
	cf.txtNotes.SetBounds(16, 158, 348, 64)
	cf.txtNotes.SetPlaceholder("multiline notes")
	cf.grpText.AddControl(cf.txtNotes)

	// --- choices ----------------------------------------------------------
	cf.grpChoice = goforms.NewGroupBox("Choices", 762, 300)
	cf.grpChoice.SetBounds(12, 260, 762, 300)
	cf.AddControl(cf.grpChoice)

	cf.chkAgree = goforms.NewCheckBox("I agree to the terms")
	cf.chkAgree.SetBounds(16, 32, 200, 26)
	cf.chkAgree.CheckedChanged.Handle(cf.chkAgree_CheckedChanged)
	cf.grpChoice.AddControl(cf.chkAgree)

	cf.rbSmall = goforms.NewRadioButton("Small")
	cf.rbSmall.SetBounds(16, 64, 110, 28)
	cf.rbSmall.CheckedChanged.Handle(cf.size_CheckedChanged)
	cf.grpChoice.AddControl(cf.rbSmall)

	cf.rbMedium = goforms.NewRadioButton("Medium")
	cf.rbMedium.SetBounds(136, 64, 110, 28)
	cf.rbMedium.CheckedChanged.Handle(cf.size_CheckedChanged)
	cf.grpChoice.AddControl(cf.rbMedium)

	cf.rbLarge = goforms.NewRadioButton("Large")
	cf.rbLarge.SetBounds(256, 64, 110, 28)
	cf.rbLarge.CheckedChanged.Handle(cf.size_CheckedChanged)
	cf.grpChoice.AddControl(cf.rbLarge)

	cf.sizeGroup = goforms.NewRadioButtonGroup()
	cf.sizeGroup.Add(cf.rbSmall)
	cf.sizeGroup.Add(cf.rbMedium)
	cf.sizeGroup.Add(cf.rbLarge)

	cf.cboCountry = goforms.NewComboBox("Germany", "Japan", "Norway", "Peru")
	cf.cboCountry.SetBounds(16, 102, 200, 30)
	cf.cboCountry.SetPlaceholder("(select a country)")
	cf.cboCountry.SelectedIndexChanged.Handle(cf.cboCountry_SelectedIndexChanged)
	cf.grpChoice.AddControl(cf.cboCountry)

	cf.dudPriority = goforms.NewDomainUpDown("Low", "Normal", "High", "Critical")
	cf.dudPriority.SetBounds(16, 142, 200, 48)
	cf.dudPriority.SelectedItemChanged.Handle(cf.dudPriority_SelectedItemChanged)
	cf.grpChoice.AddControl(cf.dudPriority)

	cf.lstFruit = goforms.NewListBox("Apple", "Banana", "Cherry", "Dragonfruit", "Elderberry")
	cf.lstFruit.SetBounds(240, 32, 220, 240)
	cf.lstFruit.SelectedIndexChanged.Handle(cf.lstFruit_SelectedIndexChanged)
	cf.grpChoice.AddControl(cf.lstFruit)

	cf.clbToppings = goforms.NewCheckedListBox("Cheese", "Basil", "Olives", "Chilli", "Anchovies")
	cf.clbToppings.SetBounds(478, 32, 220, 240)
	cf.clbToppings.ItemCheck.Handle(cf.clbToppings_ItemCheck)
	cf.grpChoice.AddControl(cf.clbToppings)

	// --- ranges -----------------------------------------------------------
	cf.grpRange = goforms.NewGroupBox("Ranges and progress", 380, 236)
	cf.grpRange.SetBounds(786, 12, 380, 236)
	cf.AddControl(cf.grpRange)

	cf.trkVolume = goforms.NewTrackBar(0, 100)
	cf.trkVolume.SetBounds(16, 36, 340, 32)
	cf.trkVolume.ValueChanged.Handle(cf.trkVolume_ValueChanged)
	cf.grpRange.AddControl(cf.trkVolume)

	cf.numQty = goforms.NewNumericUpDown(0, 99, 1)
	cf.numQty.SetBounds(16, 78, 120, 32)
	cf.numQty.ValueChanged.Handle(cf.numQty_ValueChanged)
	cf.grpRange.AddControl(cf.numQty)

	cf.prgLoad = goforms.NewProgressBar()
	cf.prgLoad.SetBounds(16, 122, 340, 24)
	cf.grpRange.AddControl(cf.prgLoad)

	cf.sbHoriz = goforms.NewHScrollBar(0, 100)
	cf.sbHoriz.SetBounds(16, 156, 340, 22)
	cf.sbHoriz.ValueChanged.Handle(cf.sbHoriz_ValueChanged)
	cf.grpRange.AddControl(cf.sbHoriz)

	cf.lblRangeHi = goforms.NewLabel("TrackBar drives the ProgressBar")
	cf.lblRangeHi.SetBounds(16, 190, 340, 24)
	cf.grpRange.AddControl(cf.lblRangeHi)

	// --- dates -------------------------------------------------------------
	cf.grpDates = goforms.NewGroupBox("Dates", 380, 300)
	cf.grpDates.SetBounds(786, 260, 380, 300)
	cf.AddControl(cf.grpDates)

	cf.dtpWhen = goforms.NewDateTimePicker(time.Now())
	cf.dtpWhen.SetBounds(16, 32, 180, 30)
	cf.dtpWhen.ValueChanged.Handle(cf.dtpWhen_ValueChanged)
	cf.grpDates.AddControl(cf.dtpWhen)

	cf.calMonth = goforms.NewMonthCalendar(time.Now())
	cf.calMonth.SetBounds(16, 72, 348, 210)
	cf.calMonth.DateChanged.Handle(cf.calMonth_DateChanged)
	cf.grpDates.AddControl(cf.calMonth)

	// --- display ------------------------------------------------------------
	cf.grpDisplay = goforms.NewGroupBox("Display", 1154, 200)
	cf.grpDisplay.SetBounds(12, 572, 1154, 200)
	cf.AddControl(cf.grpDisplay)

	cf.picLogo = goforms.NewPictureBox(160, 150)
	cf.picLogo.SetBounds(16, 34, 160, 150)
	cf.grpDisplay.AddControl(cf.picLogo)

	cf.rtbNotes = goforms.NewRichTextBox("")
	cf.rtbNotes.SetBounds(190, 34, 946, 150)
	cf.grpDisplay.AddControl(cf.rtbNotes)

	cf.Load.Handle(cf.ControlsForm_Load)
	cf.Closing.Handle(cf.ControlsForm_Closing)
}
