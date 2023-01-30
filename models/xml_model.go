package models

import (
	"encoding/xml"
)

type PurchaseOrder struct {
	XMLName             xml.Name `xml:"PurchaseOrder" json:"xml_name"`
	PurchaseOrderNumber string   `xml:"PurchaseOrderNumber,attr" json:"purchase_order_number"`
	OrderDate           string   `xml:"OrderDate,attr" json:"order_date"`
	Address             Address  `xml:"Address" json:"Address" validate:"required"`
	DeliveryNotes       string   `xml:"DeliveryNotes" json:"delivery_notes"`
	Items               Items    `xml:"Items" json:"items" validate:"required"`
}

type Address struct {
	Type    string `xml:"Type,attr" json:"type"`
	Name    string `xml:"Name" json:"name"`
	Street  string `xml:"Street" json:"street"`
	City    string `xml:"City" json:"city"`
	State   string `xml:"State" json:"state"`
	Zip     string `xml:"Zip" json:"zip"`
	Country string `xml:"Country" json:"country"`
}

type Items struct {
	XMLName xml.Name `xml:"Items" json:"XMLName"`
	Items   []Item   `xml:"Item" json:"Items"`
}

type Item struct {
	PartNumber   string `xml:"PartNumber,attr"`
	ProductName  string `xml:"ProductName"`
	Quantity     int    `xml:"Quantity"`
	Price        string `xml:"USPrice"`
	Comment      string `xml:"Comment"`
	ShipDate     string `xml:"ShipDate"`
	DeliveryDate string `xml:"DeliveryDate"`
}
