package main

import "fmt"

type Button interface {
	Click() string
}

type WinButton struct{}

func (b *WinButton) Click() string { return "Windows Button Clicked" }

type MacButton struct{}

func (b *MacButton) Click() string { return "Mac Button Clicked" }

type GUIFactory interface {
	CreateButton() Button
}

type WinFactory struct{}

func (w *WinFactory) CreateButton() Button { return &WinButton{} }

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button { return &MacButton{} }

func GetFactory(os string) GUIFactory {
	if os == "windows" {
		return &WinFactory{}
	} else if os == "mac" {
		return &MacFactory{}
	}
	return nil
}

func main() {
	factory := GetFactory("windows")
	btn := factory.CreateButton()
	fmt.Println(btn.Click())

	factory = GetFactory("mac")
	btn = factory.CreateButton()
	fmt.Println(btn.Click())
}
