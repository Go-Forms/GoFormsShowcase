package gotest

import "github.com/Go-Forms/GoForms"

// GoTest-designer.go is the generated-looking half of the WinForms-style
// partial-class split: field declarations and initializeComponent() live
// here. Hand-written logic and event handler bodies belong in GoTest.go.
// The GoForms Designer regenerates only this file.
type GoTest struct {
	*goforms.Form
}

// NewGoTest mirrors "new GoTest()" - allocate, then initializeComponent.
func NewGoTest() *GoTest {
	f := &GoTest{Form: goforms.NewForm("GoTest", 500, 350)}
	f.initializeComponent()
	return f
}

// initializeComponent mirrors the WinForms designer's InitializeComponent():
// pure layout, no business logic.
func (f *GoTest) initializeComponent() {
	f.SetClientSize(500, 350)
	f.CenterOnScreen()
}
