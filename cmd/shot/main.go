// Command shot opens one showcase form on its own, so the documentation's
// screenshots can be produced without clicking through the shell.
//
// It exists because a screenshot has to be reproducible: driving the real
// MainForm with synthetic key presses to reach a particular window is fragile
// and depends on focus order, whereas naming the form is not.
//
// Usage:
//
//	go run ./cmd/shot <form>
//
// where <form> is one of: main, controls, data, layout, dialogs, about.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"goforms"

	aboutform "goformsshowcase/Forms/AboutForm"
	controlsform "goformsshowcase/Forms/ControlsForm"
	dataform "goformsshowcase/Forms/DataForm"
	dialogsform "goformsshowcase/Forms/DialogsForm"
	layoutform "goformsshowcase/Forms/LayoutForm"
	mainform "goformsshowcase/Forms/MainForm"
)

var forms = map[string]func() *goforms.Form{
	"main":     func() *goforms.Form { return mainform.NewMainForm().Form },
	"controls": func() *goforms.Form { return controlsform.NewControlsForm().Form },
	"data":     func() *goforms.Form { return dataform.NewDataForm().Form },
	"layout":   func() *goforms.Form { return layoutform.NewLayoutForm().Form },
	"dialogs":  func() *goforms.Form { return dialogsform.NewDialogsForm().Form },
	"about":    func() *goforms.Form { return aboutform.NewAboutForm().Form },
}

func main() {
	names := make([]string, 0, len(forms))
	for name := range forms {
		names = append(names, name)
	}
	sort.Strings(names)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: shot <%s>\n", strings.Join(names, "|"))
		os.Exit(2)
	}
	build, ok := forms[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown form %q; expected one of %s\n", os.Args[1], strings.Join(names, ", "))
		os.Exit(2)
	}

	goforms.NewApplication("com.example.goformsshowcase.shot")
	goforms.Run(build())
}
