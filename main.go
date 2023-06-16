package main

import (
    "github.com/rivo/tview"
)

func main() {
    app := tview.NewApplication()
	
    chatHistory := tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetWrap(true).
	ScrollToEnd(true).
	SetChangedFunc(func() {
	    app.Draw()
	})

    inputField := tview.NewInputField().
        SetFieldBackgroundColor(tcell.ColorBlack).
	SetFieldTextColor(tcell.ColorWhite)

    // Handle user input
    inputField.SetDoneFunc(func(key tcell.key) {
	if key == tcell.KeyEnter {
	    message := inputField.GetText()
	    // Display the user is chat history
	    chatHistory.Write([]byte("> " + message + "\n"))
	    inputField.SetText("") // Clear the input field
	    // TODO: Add code to send message to OpenAI API and display the response
	}
    })

    flex := tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(chatHistory, 0, 1, false).
	AddItem(inputField, 1, 0, true)

    app.SetRoot(flex, true).Run()

