package xmlService

import (
	"fmt"
	"time"
)

type XmlSender interface {
	SendXml() bool
}

type XmlDocument struct {
	data string
}

func NewXmlDocument(data string) *XmlDocument {
	return &XmlDocument{
		data: data,
	}
}

func (d *XmlDocument) SendXml() bool {
	// Imitation of working process
	time.Sleep(3 * time.Second)
	fmt.Println("[SendXml] {pkg xmlService} done: data was sent")

	return true
}
