package dialogsform

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Go-Forms/GoForms"
	"goformsshowcase/eventlog"
)

const src = "Dialogs"

func (df *DialogsForm) DialogsForm_Load(sender any, e goforms.EventArgs) {
	df.pnlSwatch.SetBackColor(goforms.RGB(80, 80, 90))
	eventlog.Log(src+"/Form", "Load", "dialogs form opened")
}

// parent is the window dialogs attach to: this form, the one the user
// clicked in. A dialog belongs to the form that raised it - parenting it to
// the showcase shell instead would pop it over a different window than the
// button that opened it, which reads as the dialog never appearing.
func (df *DialogsForm) parent() *goforms.Form { return df.Form }

func (df *DialogsForm) btnInfo_Click(sender any, e goforms.MouseEventArgs) {
	goforms.ShowMessageBox(df.parent(),
		"Every dialog reports through a callback instead of blocking.",
		"Information", goforms.MessageBoxOK, goforms.MessageBoxInformation,
		func(r goforms.DialogResult) {
			eventlog.Logf(src+"/MessageBox", "Result", "%s", resultName(r))
		})
}

func (df *DialogsForm) btnOKCanc_Click(sender any, e goforms.MouseEventArgs) {
	goforms.ShowMessageBox(df.parent(),
		"Apply these changes?", "Confirm",
		goforms.MessageBoxOKCancel, goforms.MessageBoxQuestion,
		func(r goforms.DialogResult) {
			eventlog.Logf(src+"/MessageBox", "Result", "%s", resultName(r))
		})
}

func (df *DialogsForm) btnYesNo_Click(sender any, e goforms.MouseEventArgs) {
	goforms.ShowMessageBox(df.parent(),
		"Discard the current row?", "Delete row",
		goforms.MessageBoxYesNo, goforms.MessageBoxQuestion,
		func(r goforms.DialogResult) {
			eventlog.Logf(src+"/MessageBox", "Result", "%s", resultName(r))
		})
}

func (df *DialogsForm) btnError_Click(sender any, e goforms.MouseEventArgs) {
	goforms.ShowMessageBox(df.parent(),
		"Could not reach the server.", "Error",
		goforms.MessageBoxOK, goforms.MessageBoxError,
		func(r goforms.DialogResult) {
			eventlog.Logf(src+"/MessageBox", "Result", "%s", resultName(r))
		})
}

func (df *DialogsForm) btnInput_Click(sender any, e goforms.MouseEventArgs) {
	goforms.ShowInputBox(df.parent(), "Rename", "New name",
		func(text string, ok bool) {
			if !ok {
				eventlog.Log(src+"/InputBox", "Cancelled", "")
				return
			}
			df.txtResult.SetText(text)
			eventlog.Logf(src+"/InputBox", "Text", "%q", text)
		})
}

func (df *DialogsForm) btnOpen_Click(sender any, e goforms.MouseEventArgs) {
	dlg := goforms.NewOpenFileDialog()
	dlg.Filter = goforms.FileDialogFilter{
		Description: "Go source",
		Extensions:  []string{".go", ".mod", ".md"},
	}
	dlg.Show(df.parent(), func(path string, ok bool) {
		if !ok {
			eventlog.Log(src+"/OpenFileDialog", "Cancelled", "")
			return
		}
		eventlog.Logf(src+"/OpenFileDialog", "FileName", "%s", path)

		// Reading it is the point of having opened it - and it is what
		// differs in a browser, where the file arrives as bytes with no
		// path behind them.
		data, err := dlg.ReadAll()
		if err != nil {
			df.txtResult.SetText(path + " (could not be read: " + err.Error() + ")")
			eventlog.Logf(src+"/OpenFileDialog", "ReadAll", "failed: %v", err)
			return
		}
		df.txtResult.SetText(fmt.Sprintf("%s - %d bytes, first line: %s", path, len(data), firstLine(data)))
		eventlog.Logf(src+"/OpenFileDialog", "ReadAll", "%d bytes", len(data))
	})
}

// firstLine is a readable sample of a file, for showing that it really was
// read rather than just named.
func firstLine(data []byte) string {
	line := string(data)
	if i := strings.IndexAny(line, "\r\n"); i >= 0 {
		line = line[:i]
	}
	if len(line) > 60 {
		line = line[:60] + "..."
	}
	if line == "" {
		return "(empty)"
	}
	return strconv.Quote(line)
}

func (df *DialogsForm) btnSave_Click(sender any, e goforms.MouseEventArgs) {
	dlg := goforms.NewSaveFileDialog()
	dlg.DefaultFileName = "showcase.txt"
	dlg.Show(df.parent(), func(path string, ok bool) {
		if !ok {
			eventlog.Log(src+"/SaveFileDialog", "Cancelled", "")
			return
		}
		eventlog.Logf(src+"/SaveFileDialog", "FileName", "%s", path)

		// And writing it: on the desktop this lands on disk, in a browser
		// it arrives as a download.
		body := []byte("Saved from the GoForms showcase at " + time.Now().Format(time.RFC3339) + "\n")
		if err := dlg.WriteAll(body); err != nil {
			df.txtResult.SetText(path + " (could not be written: " + err.Error() + ")")
			eventlog.Logf(src+"/SaveFileDialog", "WriteAll", "failed: %v", err)
			return
		}
		df.txtResult.SetText(fmt.Sprintf("%s - %d bytes written", path, len(body)))
		eventlog.Logf(src+"/SaveFileDialog", "WriteAll", "%d bytes", len(body))
	})
}

func (df *DialogsForm) btnFolder_Click(sender any, e goforms.MouseEventArgs) {
	dlg := goforms.NewFolderBrowserDialog()
	dlg.Show(df.parent(), func(path string, ok bool) {
		if !ok {
			eventlog.Log(src+"/FolderBrowserDialog", "Cancelled", "")
			return
		}
		df.txtResult.SetText(path)
		eventlog.Logf(src+"/FolderBrowserDialog", "SelectedPath", "%s", path)
	})
}

func (df *DialogsForm) btnColour_ColorChanged(sender any, e goforms.EventArgs) {
	df.pnlSwatch.SetBackColor(df.btnColour.SelectedColor)
	eventlog.Log(src+"/ColorDialog", "ColorChanged", "swatch updated")
}

func resultName(r goforms.DialogResult) string {
	switch r {
	case goforms.DialogOK:
		return "OK"
	case goforms.DialogCancel:
		return "Cancel"
	case goforms.DialogYes:
		return "Yes"
	case goforms.DialogNo:
		return "No"
	}
	return "None"
}
