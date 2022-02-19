package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Muhammedhuseynov/golang_wifi_password/pck"
)

// GUI PART
var wifpasword = binding.NewString()
var wifiPasswords map[string]string
var radioBtn = widget.NewRadioGroup([]string{},
	func(val string) {
		wifpasword.Set(wifiPasswords[val])
	})

//function for button (get wifi name and wifi passwords)
func searchWifis() {
	wifi_names := pck.GetWifiNames()
	wifpasword.Set("Searching.....")
	wifiPasswords = pck.WifiPasswords(wifi_names)
	for _, wifi := range wifi_names {
		radioBtn.Options = append(radioBtn.Options, wifi)
	}
	radioBtn.Refresh()
	wifpasword.Set("Done! Click on name of Wifi")
}

//make label on center
func makeObjCenter(obj fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), obj, layout.NewSpacer())
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500, 389))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	wifpasword.Set("See Connected wifi Passwords")
	name_lbl := makeObjCenter(widget.NewLabelWithData(wifpasword))

	// Wifi name / list
	var wifiNameList = container.NewVScroll(radioBtn)
	wifiNameList.SetMinSize(fyne.NewSize(300, 300))

	//Search button
	btnSearch := widget.NewButton("Search", searchWifis)
	w.SetContent(container.NewVBox(container.NewVBox(name_lbl), wifiNameList, btnSearch))
	w.ShowAndRun()
}
