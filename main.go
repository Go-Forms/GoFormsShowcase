package main

import (
	"github.com/Go-Forms/GoForms"

	mainform "goformsshowcase/Forms/MainForm"
)

// main.go wires up the application by hand, the same way a WinForms
// Program.cs Main() calls Application.Run(new MainForm()).
func main() { 
	goforms.NewApplication("com.example.goformsshowcase")
	goforms.Run(mainform.NewMainForm().Form)
}
