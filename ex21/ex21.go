package ex21

import (
	"L1/ex21/xmlService"
	"fmt"
	"time"
)

// Example was taken from: https://www.youtube.com/watch?v=o9sCFOv-uKE

type JsonDocument struct {
	data string
}

func NewJsonDocument(data string) *JsonDocument {
	return &JsonDocument{
		data: data,
	}
}

func (d *JsonDocument) ConvertToXml() string {
	return "<xmlService>convert to xmlService</xmlService>"
}

type JsonDocumentAdapter struct {
	d *JsonDocument
}

func NewJsonDocumentAdapter(d *JsonDocument) *JsonDocumentAdapter {
	return &JsonDocumentAdapter{
		d: d,
	}
}

func (a *JsonDocumentAdapter) SendXml() bool {
	fmt.Printf("[SendXml] {pkg ex21} converted\nconverted data: %s\n", a.d.ConvertToXml())
	// Imitation of working process
	time.Sleep(3 * time.Second)
	fmt.Println("[SendXml] {pkg ex21} done: data was sent")

	return true
}

func Test() {
	// xmlService foreign package example
	fmt.Println("[main] called xmlService methods")
	xml := xmlService.NewXmlDocument("...")
	xml.SendXml()

	// ex21 package example
	fmt.Println("[main] called ex21 methods")
	json := NewJsonDocument("...")
	adapter := NewJsonDocumentAdapter(json)
	adapter.SendXml()
}
