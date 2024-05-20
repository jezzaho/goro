package cli

import (
	"bufio"
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type CliInput struct {
	Line      string
	Beggining string
	Ending    string
}

func RenderMainScreen() CliInput {
	cliInput := CliInput{}

	pterm.DefaultCenter.Println("Program 'ROZKŁADACZ' służy do sprawdzania rozkładów linii lotniczych. \n Eryk Kiper 2024")
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("ROZKLADACZ")).Srender()
	pterm.DefaultCenter.Println(s)

	pterm.DefaultCenter.Println("Wybierz linię lotniczą dla której chcesz sprawdzić rozkład...")

	var optionsAirlines []string
	optionsAirlines = append(optionsAirlines, "LH - Lufthansa")
	optionsAirlines = append(optionsAirlines, "OS - Austrian Airlines")
	optionsAirlines = append(optionsAirlines, "LX - Swiss Airlines")
	optionsAirlines = append(optionsAirlines, "SN - Brussels Airways")

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	// LINIA LOTNICZA
	selectedAirline, _ := pterm.DefaultInteractiveSelect.WithOptions(optionsAirlines).Show()
	cliInput.Line = selectedAirline
	pterm.DefaultArea.Clear()
	// Display the selected option to the user with a green color for emphasis
	pterm.Info.Printfln("Wybrano: %s", pterm.Green(selectedAirline))
	pterm.DefaultCenter.Println("Wybierz przedział rozkładu...")
	var optionsTime []string
	optionsTime = append(optionsTime, "W23")
	optionsTime = append(optionsTime, "S24")
	optionsTime = append(optionsTime, "Własny")
	optionsTime = append(optionsTime, "W24")
	selectedTime, _ := pterm.DefaultInteractiveSelect.WithOptions(optionsTime).Show()
	if selectedTime == "W23" {
		cliInput.Beggining = "29NOV23"
		cliInput.Ending = "31MAR24"
	}
	if selectedTime == "S24" {
		cliInput.Beggining = "01APR24"
		cliInput.Ending = "27NOV24"
	}
	if selectedTime == "W24" {
		cliInput.Beggining = "28NOV24"
		cliInput.Ending = "" // FILL DATA
	}
	pterm.DefaultArea.Clear()
	// Display the selected option to the user with a green color for emphasis
	pterm.Info.Printfln("Wybrano: %s", pterm.Green(selectedTime))
	if selectedTime == "Własny" {
		// Poczatek rozkladu
		// Create an interactive text input with single line input mode and show it
		beginning, _ := pterm.DefaultInteractiveTextInput.Show("Wpisz początek przedziału w formacie DDMMMYY (np. 19JUL24)")
		// Print a blank line for better readability
		pterm.Println()
		// Print the user's answer with an info prefix
		pterm.Info.Printfln("Wybrano: %s", beginning)
		end, _ := pterm.DefaultInteractiveTextInput.Show("Wpisz koniec przedziału w formacie DDMMMYY (np. 19JUL24)")
		pterm.Info.Printfln("Wybrano: %s", end)
		cliInput.Beggining = beginning
		cliInput.Ending = end
		// Print a blank line for better readability
		pterm.Println()

		// Print the user's answer with an info prefix
		pterm.Info.Printfln("Wybrano: %s", end)
	}
	return cliInput
}

func RenderFinal() {
	pterm.DefaultCenter.Println("Program 'ROZKŁADACZ' służy do sprawdzania rozkładów linii lotniczych. \n Eryk Kiper 2024")
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("ROZKLADACZ")).Srender()
	pterm.DefaultCenter.Println(s)
	pterm.DefaultCenter.Println("Program zapisał plik csv jako \"output.csv\" w folderu bieżącym.")
	pterm.DefaultCenter.Println("Zadanie zakończone. Naciśnij Enter aby wyjść.")

	// Wait for user to press Enter before exiting
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
