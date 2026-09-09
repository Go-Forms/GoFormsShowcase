package controlsform

import (
	"image"
	"image/color"

	"github.com/Go-Forms/GoForms"
	"goformsshowcase/eventlog"
)

// ControlsForm.go is the hand-written half: handlers and logic.

const src = "Controls"

func (cf *ControlsForm) ControlsForm_Load(sender any, e goforms.EventArgs) {
	// Every control's inherited events go to the shared log, which is what
	// makes "all controls, all events" visible rather than just claimed.
	eventlog.WireAll(src+"/btnPlain", cf.btnPlain)
	eventlog.WireAll(src+"/txtName", cf.txtName)
	eventlog.WireAll(src+"/chkAgree", cf.chkAgree)
	eventlog.WireAll(src+"/lstFruit", cf.lstFruit)
	eventlog.WireAll(src+"/cboCountry", cf.cboCountry)
	eventlog.WireAll(src+"/trkVolume", cf.trkVolume)
	eventlog.WireAll(src+"/picLogo", cf.picLogo)

	// The picture box additionally reports the high-frequency events, so
	// they are demonstrated somewhere without drowning everything else.
	eventlog.WireMotion(src+"/picLogo", cf.picLogo)

	cf.styleTheLabel()
	cf.fillPicture()
	cf.fillRichText()
	cf.buildContextMenus()

	cf.prgLoad.SetRange(0, 100)
	cf.prgLoad.SetValue(30)
	cf.trkVolume.SetValue(30)

	eventlog.Log(src+"/Form", "Load", "controls form opened")
}

func (cf *ControlsForm) ControlsForm_Closing(sender any, e *goforms.CancelEventArgs) {
	eventlog.Log(src+"/Form", "Closing", "")
}

// styleTheLabel shows Control.Font and Control.ForeColor actually taking
// effect - Label renders its own text, so it can honour both.
func (cf *ControlsForm) styleTheLabel() {
	cf.lblStyled.SetFont(goforms.NewFont("", 18, true, true))
	cf.lblStyled.SetForeColor(goforms.RGB(90, 170, 255))
}

// fillPicture draws a gradient into the PictureBox so it has something to
// show without shipping an image file.
func (cf *ControlsForm) fillPicture() {
	const size = 150
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(x * 255 / size),
				G: uint8(y * 255 / size),
				B: 200,
				A: 255,
			})
		}
	}
	cf.picLogo.LoadImage(img)
}

func (cf *ControlsForm) fillRichText() {
	cf.rtbNotes.SetText("## RichTextBox\n" +
		"Formatting is **Markdown**, not RTF - that is what Fyne's `RichText` parses.\n\n" +
		"- *Italic*, **bold** and `code`\n" +
		"- Hover, click and type anywhere on this form and watch the event log\n" +
		"- Right-click the list, the picture or the notes for a context menu")
}

// buildContextMenus attaches right-click menus to three very different
// controls - one interactive, one passive, one text entry - because that
// combination is exactly what used not to work.
func (cf *ControlsForm) buildContextMenus() {
	clear := goforms.NewMenuItem("Clear selection")
	clear.Click.Handle(func(any, goforms.EventArgs) {
		cf.lstFruit.SetSelectedIndex(-1)
		eventlog.Log(src+"/lstFruit", "ContextMenu", "Clear selection")
	})
	add := goforms.NewMenuItem("Add 'Fig'")
	add.Click.Handle(func(any, goforms.EventArgs) {
		cf.lstFruit.SetItems(append(cf.lstFruit.Items(), "Fig"))
		eventlog.Log(src+"/lstFruit", "ContextMenu", "Add 'Fig'")
	})
	cf.lstFruit.SetContextMenu(goforms.NewContextMenu(clear, goforms.NewMenuSeparator(), add))

	redraw := goforms.NewMenuItem("Redraw gradient")
	redraw.Click.Handle(func(any, goforms.EventArgs) {
		cf.fillPicture()
		eventlog.Log(src+"/picLogo", "ContextMenu", "Redraw gradient")
	})
	cf.picLogo.SetContextMenu(goforms.NewContextMenu(redraw))

	upper := goforms.NewMenuItem("Shout the name")
	upper.Click.Handle(func(any, goforms.EventArgs) {
		cf.txtName.SetText(cf.txtName.Text() + "!")
		eventlog.Log(src+"/txtName", "ContextMenu", "Shout")
	})
	cf.txtName.SetContextMenu(goforms.NewContextMenu(upper))
}

// --- designer-wired handlers ---------------------------------------------

func (cf *ControlsForm) btnDisabled_Click(sender any, e goforms.MouseEventArgs) {
	// EnabledChanged is one of the inherited lifecycle events, so this shows
	// up in the log twice: once as the Click, once as the state change.
	cf.btnPlain.SetEnabled(!cf.btnPlain.Enabled())
	eventlog.Logf(src+"/btnDisabled", "Click", "btnPlain enabled=%v", cf.btnPlain.Enabled())
}

func (cf *ControlsForm) btnColor_ColorChanged(sender any, e goforms.EventArgs) {
	cf.lblStyled.SetForeColor(cf.btnColor.SelectedColor)
	eventlog.Log(src+"/btnColor", "ColorChanged", "label recoloured")
}

func (cf *ControlsForm) txtName_TextChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/txtName", "TextChanged", "%q", cf.txtName.Text())
}

func (cf *ControlsForm) txtPhone_MaskCompleted(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/txtPhone", "MaskCompleted", "%s", cf.txtPhone.Text())
}

func (cf *ControlsForm) chkAgree_CheckedChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/chkAgree", "CheckedChanged", "%v", cf.chkAgree.Checked())
}

// size_CheckedChanged serves all three radio buttons; sender says which one.
func (cf *ControlsForm) size_CheckedChanged(sender any, e goforms.EventArgs) {
	rb, ok := sender.(*goforms.RadioButton)
	if !ok {
		return
	}
	eventlog.Logf(src+"/"+rb.Text(), "CheckedChanged", "%v", rb.Checked())
}

func (cf *ControlsForm) cboCountry_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/cboCountry", "SelectedIndexChanged", "%s", cf.cboCountry.SelectedItem())
}

func (cf *ControlsForm) dudPriority_SelectedItemChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/dudPriority", "SelectedItemChanged", "%s", cf.dudPriority.SelectedItem())
}

func (cf *ControlsForm) lstFruit_SelectedIndexChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/lstFruit", "SelectedIndexChanged", "%d %s",
		cf.lstFruit.SelectedIndex(), cf.lstFruit.SelectedItem())
}

func (cf *ControlsForm) clbToppings_ItemCheck(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/clbToppings", "ItemCheck", "%v", cf.clbToppings.CheckedItems())
}

func (cf *ControlsForm) trkVolume_ValueChanged(sender any, e goforms.EventArgs) {
	v := cf.trkVolume.Value()
	cf.prgLoad.SetValue(int(v))
	eventlog.Logf(src+"/trkVolume", "ValueChanged", "%.0f", v)
}

func (cf *ControlsForm) numQty_ValueChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/numQty", "ValueChanged", "%.0f", cf.numQty.Value())
}

func (cf *ControlsForm) sbHoriz_ValueChanged(sender any, e goforms.EventArgs) {
	eventlog.Logf(src+"/sbHoriz", "ValueChanged", "%.0f", cf.sbHoriz.Value())
}

func (cf *ControlsForm) dtpWhen_ValueChanged(sender any, e goforms.EventArgs) {
	eventlog.Log(src+"/dtpWhen", "ValueChanged", cf.dtpWhen.Value().Format("2006-01-02"))
}

func (cf *ControlsForm) calMonth_DateChanged(sender any, e goforms.EventArgs) {
	eventlog.Log(src+"/calMonth", "DateChanged", cf.calMonth.Value().Format("2006-01-02"))
}
