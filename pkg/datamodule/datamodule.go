package datamodule

import (
	"encoding/xml"
	"github.com/aldehir/S1000D/pkg/common"
)

// DataModule represents a complete S1000D Data Module (Issue 5.0)
type DataModule struct {
	XMLName               xml.Name              `xml:"dmodule"`
	IdentAndStatusSection IdentAndStatusSection `xml:"identAndStatusSection"`
	Content               Content               `xml:"content"`
}

// IdentAndStatusSection contains identification and status information
type IdentAndStatusSection struct {
	XMLName    xml.Name   `xml:"identAndStatusSection"`
	DMAddress  DMAddress  `xml:"dmAddress"`
	DMStatus   DMStatus   `xml:"dmStatus"`
}

// DMAddress contains the data module addressing information
type DMAddress struct {
	XMLName       xml.Name      `xml:"dmAddress"`
	DMIdent       common.DMIdent `xml:"dmIdent"`
	DMAddressItems DMAddressItems `xml:"dmAddressItems"`
}

// DMAddressItems contains additional addressing elements
type DMAddressItems struct {
	XMLName         xml.Name `xml:"dmAddressItems"`
	IssueDate       IssueDate `xml:"issueDate,omitempty"`
	DMTitle         DMTitle   `xml:"dmTitle"`
}

// IssueDate represents the issue date
type IssueDate struct {
	XMLName xml.Name `xml:"issueDate"`
	Year    string   `xml:"year,attr"`
	Month   string   `xml:"month,attr"`
	Day     string   `xml:"day,attr"`
}

// DMTitle represents the data module title
type DMTitle struct {
	XMLName    xml.Name `xml:"dmTitle"`
	TechName   string   `xml:"techName"`
	InfoName   string   `xml:"infoName,omitempty"`
}

// DMStatus contains status information about the data module
type DMStatus struct {
	XMLName           xml.Name          `xml:"dmStatus"`
	Security          common.Security   `xml:"security"`
	ResponsiblePartnerCompany *ResponsiblePartnerCompany `xml:"responsiblePartnerCompany,omitempty"`
	Originator        *Originator       `xml:"originator,omitempty"`
	AppLicability     *Applicability    `xml:"applic,omitempty"`
	BrExDmRef         []common.DMRef    `xml:"brexDmRef,omitempty"`
	QualityAssurance  *QualityAssurance `xml:"qualityAssurance,omitempty"`
}

// ResponsiblePartnerCompany represents the responsible partner company
type ResponsiblePartnerCompany struct {
	XMLName        xml.Name `xml:"responsiblePartnerCompany"`
	EnterpriseCode string   `xml:"enterpriseCode,attr"`
	EnterpriseName string   `xml:",chardata"`
}

// Originator represents the originator of the data module
type Originator struct {
	XMLName        xml.Name `xml:"originator"`
	EnterpriseCode string   `xml:"enterpriseCode,attr,omitempty"`
	EnterpriseName string   `xml:",chardata"`
}

// Applicability represents applicability information
type Applicability struct {
	XMLName xml.Name `xml:"applic"`
	// Simplified - full implementation would include display text and assertions
	DisplayText string `xml:"displayText,omitempty"`
}

// QualityAssurance represents quality assurance information
type QualityAssurance struct {
	XMLName       xml.Name `xml:"qualityAssurance"`
	Unverified    *Unverified `xml:"unverified,omitempty"`
	FirstVerification *FirstVerification `xml:"firstVerification,omitempty"`
}

// Unverified indicates the module is unverified
type Unverified struct {
	XMLName xml.Name `xml:"unverified"`
}

// FirstVerification represents first verification information
type FirstVerification struct {
	XMLName          xml.Name `xml:"firstVerification"`
	VerificationType string   `xml:"verificationType,attr,omitempty"`
}

// Content represents the content section of a data module
type Content struct {
	XMLName     xml.Name     `xml:"content"`
	Description *Description `xml:"description,omitempty"`
	Procedure   *Procedure   `xml:"procedure,omitempty"`
}

// Description represents descriptive content
type Description struct {
	XMLName        xml.Name        `xml:"description"`
	LevelledPara   []LevelledPara  `xml:"levelledPara"`
}

// Procedure represents procedural content
type Procedure struct {
	XMLName          xml.Name         `xml:"procedure"`
	MainProcedure    MainProcedure    `xml:"mainProcedure"`
}

// MainProcedure contains the main procedure steps
type MainProcedure struct {
	XMLName       xml.Name       `xml:"mainProcedure"`
	ProcedureStep []ProcedureStep `xml:"proceduralStep"`
}

// ProcedureStep represents a single procedure step
type ProcedureStep struct {
	XMLName xml.Name `xml:"proceduralStep"`
	Para    []Para   `xml:"para"`
}

// LevelledPara represents a levelled paragraph
type LevelledPara struct {
	XMLName      xml.Name       `xml:"levelledPara"`
	Title        string         `xml:"title,omitempty"`
	Para         []Para         `xml:"para"`
	LevelledPara []LevelledPara `xml:"levelledPara,omitempty"`
}

// Para represents a paragraph
type Para struct {
	XMLName xml.Name `xml:"para"`
	Content string   `xml:",innerxml"`
}
