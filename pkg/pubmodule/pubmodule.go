package pubmodule

import (
	"encoding/xml"
	"github.com/aldehir/S1000D/pkg/common"
)

// PublicationModule represents a complete S1000D Publication Module (Issue 5.0)
type PublicationModule struct {
	XMLName               xml.Name              `xml:"pm"`
	IdentAndStatusSection IdentAndStatusSection `xml:"identAndStatusSection"`
	Content               Content               `xml:"content"`
}

// IdentAndStatusSection contains identification and status information
type IdentAndStatusSection struct {
	XMLName   xml.Name  `xml:"identAndStatusSection"`
	PMAddress PMAddress `xml:"pmAddress"`
	PMStatus  PMStatus  `xml:"pmStatus"`
}

// PMAddress contains the publication module addressing information
type PMAddress struct {
	XMLName        xml.Name       `xml:"pmAddress"`
	PMIdent        common.PMIdent `xml:"pmIdent"`
	PMAddressItems PMAddressItems `xml:"pmAddressItems"`
}

// PMAddressItems contains additional addressing elements
type PMAddressItems struct {
	XMLName   xml.Name  `xml:"pmAddressItems"`
	IssueDate IssueDate `xml:"issueDate,omitempty"`
	PMTitle   PMTitle   `xml:"pmTitle"`
}

// IssueDate represents the issue date
type IssueDate struct {
	XMLName xml.Name `xml:"issueDate"`
	Year    string   `xml:"year,attr"`
	Month   string   `xml:"month,attr"`
	Day     string   `xml:"day,attr"`
}

// PMTitle represents the publication module title
type PMTitle struct {
	XMLName xml.Name `xml:"pmTitle"`
	Title   string   `xml:",chardata"`
}

// PMStatus contains status information about the publication module
type PMStatus struct {
	XMLName           xml.Name          `xml:"pmStatus"`
	Security          common.Security   `xml:"security"`
	ResponsiblePartnerCompany *ResponsiblePartnerCompany `xml:"responsiblePartnerCompany,omitempty"`
	Originator        *Originator       `xml:"originator,omitempty"`
	Applicability     *Applicability    `xml:"applic,omitempty"`
	BrExDmRef         []common.DMRef    `xml:"brexDmRef,omitempty"`
	QualityAssurance  *QualityAssurance `xml:"qualityAssurance,omitempty"`
}

// ResponsiblePartnerCompany represents the responsible partner company
type ResponsiblePartnerCompany struct {
	XMLName        xml.Name `xml:"responsiblePartnerCompany"`
	EnterpriseCode string   `xml:"enterpriseCode,attr"`
	EnterpriseName string   `xml:",chardata"`
}

// Originator represents the originator of the publication module
type Originator struct {
	XMLName        xml.Name `xml:"originator"`
	EnterpriseCode string   `xml:"enterpriseCode,attr,omitempty"`
	EnterpriseName string   `xml:",chardata"`
}

// Applicability represents applicability information
type Applicability struct {
	XMLName     xml.Name `xml:"applic"`
	DisplayText string   `xml:"displayText,omitempty"`
}

// QualityAssurance represents quality assurance information
type QualityAssurance struct {
	XMLName           xml.Name           `xml:"qualityAssurance"`
	Unverified        *Unverified        `xml:"unverified,omitempty"`
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

// Content represents the content section of a publication module
type Content struct {
	XMLName xml.Name `xml:"content"`
	PMEntry []PMEntry `xml:"pmEntry"`
}

// PMEntry represents an entry in the publication module
type PMEntry struct {
	XMLName     xml.Name     `xml:"pmEntry"`
	PMEntryTitle string      `xml:"pmEntryTitle,omitempty"`
	DMRef       *common.DMRef `xml:"dmRef,omitempty"`
	PMEntry     []PMEntry    `xml:"pmEntry,omitempty"`
}

// PMEntryTitle represents a title for a PM entry
type PMEntryTitle struct {
	XMLName xml.Name `xml:"pmEntryTitle"`
	Title   string   `xml:",chardata"`
}
