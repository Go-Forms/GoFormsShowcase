package grideditrow

import (
	"testing"

	"fyne.io/fyne/v2/test"
	goforms "github.com/Go-Forms/GoForms"
)

// newDialog builds the form under Fyne's test driver, which is all
// NewGridEditRow needs - it never shows a window.
func newDialog(t *testing.T) *GridEditRow {
	t.Helper()
	test.NewApp()
	return NewGridEditRow()
}

// A row handed to the dialog must come back out of it unchanged when the user
// types nothing. The newline matters: the sample data's "Stapler\n(red)" is a
// two-line cell, and a single-line box would show only "Stapler" and put only
// that back.
func TestLoadValuesRoundTrip(t *testing.T) {
	f := newDialog(t)
	row := []string{"1003", "Initech", "Stapler\n(red)", "1", "$3,999.99"}

	f.Load(row)
	got := f.Values()

	if len(got) != FieldCount {
		t.Fatalf("Values() returned %d fields, want %d", len(got), FieldCount)
	}
	for i, want := range row {
		if got[i] != want {
			t.Errorf("field %d: got %q, want %q", i, got[i], want)
		}
	}
}

// Values trims what the user typed, so a stray space does not become part of
// a cell.
func TestValuesTrimsSurroundingSpace(t *testing.T) {
	f := newDialog(t)
	f.Load([]string{"  1003  ", " Initech ", "Widget", " 12 ", "$1.00"})

	got := f.Values()
	if got[FieldOrder] != "1003" || got[FieldCustomer] != "Initech" || got[FieldQty] != "12" {
		t.Errorf("Values() did not trim: %q", got)
	}
}

// "Add row" makes a row with every cell filled, but a caller with a shorter
// row must not leave the previous order's values sitting in the spare boxes.
func TestLoadClearsTheBoxesAShortRowDoesNotReach(t *testing.T) {
	f := newDialog(t)
	f.Load([]string{"1003", "Initech", "Stapler", "1", "$3,999.99"})

	f.Load([]string{"1008", "Acme"})

	got := f.Values()
	for i, want := range []string{"1008", "Acme", "", "", ""} {
		if got[i] != want {
			t.Errorf("field %d: got %q, want %q", i, got[i], want)
		}
	}
}

func TestValidateAcceptsAWholeRow(t *testing.T) {
	f := newDialog(t)
	f.Load([]string{"1003", "Initech", "Stapler", "12", "$3,999.99"})

	if problem := f.validate(); problem != "" {
		t.Errorf("a complete row was refused: %s", problem)
	}
}

// Save must not write a row the grid cannot use. Each of these leaves the
// dialog open instead, with the reason in lblHint.
func TestValidateRefusesAnUnusableRow(t *testing.T) {
	cases := []struct {
		name string
		row  []string
	}{
		{"no order", []string{"", "Initech", "Stapler", "12", "$1.00"}},
		{"blank order", []string{"   ", "Initech", "Stapler", "12", "$1.00"}},
		{"no customer", []string{"1003", "", "Stapler", "12", "$1.00"}},
		{"qty in words", []string{"1003", "Initech", "Stapler", "two", "$1.00"}},
		{"fractional qty", []string{"1003", "Initech", "Stapler", "1.5", "$1.00"}},
		{"empty qty", []string{"1003", "Initech", "Stapler", "", "$1.00"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newDialog(t)
			f.Load(c.row)
			if f.validate() == "" {
				t.Errorf("%v was accepted", c.row)
			}
		})
	}
}

// A refused Save explains itself in place rather than closing the form.
func TestSaveExplainsWhyItRefused(t *testing.T) {
	f := newDialog(t)
	f.Load([]string{"1003", "Initech", "Stapler", "two", "$1.00"})
	before := f.lblHint.Text()

	f.btnOK_Click(f.btnOK, goforms.MouseEventArgs{})

	if f.lblHint.Text() == before {
		t.Error("the hint still shows the idle wording after a refused Save")
	}
	if got := f.DialogResult(); got != goforms.DialogNone {
		t.Errorf("a refused Save set DialogResult to %v; the form should still be open", got)
	}
}

// Cancel closes with a result the caller can tell apart from Save.
func TestCancelReportsCancel(t *testing.T) {
	f := newDialog(t)
	f.Load([]string{"1003", "Initech", "Stapler", "1", "$1.00"})

	f.btnCancel_Click(f.btnCancel, goforms.MouseEventArgs{})

	if got := f.DialogResult(); got != goforms.DialogCancel {
		t.Errorf("Cancel closed with %v, want DialogCancel", got)
	}
}
