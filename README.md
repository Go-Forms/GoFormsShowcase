# GoForms Showcase

[![CI](https://github.com/Go-Forms/GoFormsShowcase/actions/workflows/ci.yml/badge.svg)](https://github.com/Go-Forms/GoFormsShowcase/actions/workflows/ci.yml)
![License](https://img.shields.io/badge/license-MIT-blue)

A single application that exercises the whole framework: every control, the
full event surface, form switching (modeless and modal), context menus,
dialogs, anchoring, docking and the layout panels.

Run it with:

    go run .

## What is where

| Form | Shows |
|---|---|
| **MainForm** | The shell: menu bar, tool strip, status bar with a clock, and a live **event log** every other form reports into. Right-click the log for a context menu. |
| **ControlsForm** | Buttons, links, colour picker, styled `Label`, text boxes (plain / password / multiline / masked), check boxes, radio group, combo, list, checked list, domain spinner, track bar, numeric spinner, progress bar, scroll bar, date picker, month calendar, picture box, rich text. |
| **DataForm** | `DataGridView` with its sizing and scrolling flags switchable **at runtime**, plus `ListView` and `TreeView`. |
| **LayoutForm** | `Dock` on all five edges, the four `Anchor` rules, `FlowLayoutPanel`, `TableLayoutPanel` with Absolute/Percent/AutoSize tracks and spanning, `SplitContainer` and `TabControl`. Resize the window to watch it all react. |
| **DialogsForm** | Message boxes (OK / OK-Cancel / Yes-No / error), input box, open/save file and folder pickers, colour dialog. |
| **AboutForm** | A **modal** form: opened with `ShowDialog()`, returns a `DialogResult`. |

## The event log

`eventlog.WireAll(name, control)` subscribes a control's whole inherited
event surface at once - `Click`, `DoubleClick`, `MouseDown/Up/Enter/Leave`,
`KeyDown/Up/Press`, `GotFocus/LostFocus`, `Resize`, `Move`,
`VisibleChanged`, `EnabledChanged` - so "all controls, all events" is
something you can watch rather than take on trust.

`MouseMove` and `MouseWheel` are behind the **Log mouse motion** checkbox.
They fire continuously, and every entry costs a repaint, so leaving them on
by default made the whole app crawl.

The log itself is painted on a 100 ms timer rather than per event, for the
same reason: in Fyne every `Refresh` marks the canvas dirty, so redrawing
once per event turns a burst into hundreds of full-window repaints a second.
 