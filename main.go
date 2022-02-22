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
var searchEntry = widget.NewEntry()

var radioBtn = widget.NewRadioGroup([]string{},
	func(val string) {
		if val != "" {
			wifpasword.Set(wifiPasswords[val])
		} else {
			wifpasword.Set("*************")
			//fmt.Println("Emptyyy")
		}

	})

//function for button (get wifi name and wifi passwords)
func searchWifis() {

	wifpasword.Set("Searching.....")
	wifiPasswords = pck.WifiPasswords(pck.GetWifiNames())

	for wifiName, _ := range wifiPasswords {
		radioBtn.Options = append(radioBtn.Options, wifiName)
	}
	radioBtn.Refresh()
	wifpasword.Set("Done! (************) Click on name of Wifi ")
	searchEntry.PlaceHolder = "Just write the name of the wifi here! Not need Click anything"
	searchEntry.Enable()

}

//make label on center
func makeObjCenter(obj fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), obj, layout.NewSpacer())
}

func main() {
	a := app.New()
	w := a.NewWindow("Show me Wifi Password")
	w.Resize(fyne.NewSize(500, 430))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	wifpasword.Set("See Connected wifi Passwords")
	name_lbl := makeObjCenter(widget.NewLabelWithData(wifpasword))

	//Search button
	btnSearch := widget.NewButton("Search", nil)
	changedBtn := func() {
		if btnSearch.Text == "Search" {
			btnSearch.Disable()
			searchWifis()
		}

		btnSearch.Refresh()
	}
	btnSearch.OnTapped = changedBtn

	//Search Entry
	searchEntry.PlaceHolder = "No WIFI, Search them!"
	searchEntry.OnChanged = func(s string) {
		wifpasword.Set(wifiPasswords[s])

		if s != "" {
			btnSearch.Disable()
		} else {
			btnSearch.Enable()
		}
	}
	searchEntry.Disable()

	// Wifi name / list
	var wifiNameList = container.NewVScroll(radioBtn)
	wifiNameList.SetMinSize(fyne.NewSize(300, 300))

	w.SetContent(container.NewVBox(container.NewVBox(name_lbl), searchEntry, wifiNameList, btnSearch))
	w.ShowAndRun()
}
