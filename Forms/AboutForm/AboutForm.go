package aboutform

import "goforms"

func (af *AboutForm) AboutForm_Load(sender any, e goforms.EventArgs) {
	af.lblTitle.SetFont(goforms.NewFont("", 22, true, false))
	af.rtbInfo.SetText("A Windows Forms-style GUI framework for Go, rendered by **Fyne**.\n\n" +
		"- Absolute `X, Y, Width, Height` positioning\n" +
		"- C#-style multicast events on every control\n" +
		"- Anchor, Dock and layout panels\n" +
		"- A visual designer that edits `*-designer.go` in place\n\n" +
		"This window is **modal**: it was opened with `ShowDialog()` and\n" +
		"reports a `DialogResult` back to the caller.")
}

func (af *AboutForm) btnOK_Click(sender any, e goforms.MouseEventArgs) {
	af.CloseWithResult(goforms.DialogOK)
}

func (af *AboutForm) btnCanc_Click(sender any, e goforms.MouseEventArgs) {
	af.CloseWithResult(goforms.DialogCancel)
}
