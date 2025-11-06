package common

import "encoding/xml"

// DMCode represents a Data Module Code (DMC) in S1000D
// Format: ModelIdentCode-SystemDiffCode-SystemCode-SubSystemCode-SubSubSystemCode-AssyCode-DisassyCode-DisassyCodeVariant-InfoCode-InfoCodeVariant-ItemLocationCode
type DMCode struct {
	XMLName             xml.Name `xml:"dmCode"`
	ModelIdentCode      string   `xml:"modelIdentCode,attr"`
	SystemDiffCode      string   `xml:"systemDiffCode,attr"`
	SystemCode          string   `xml:"systemCode,attr"`
	SubSystemCode       string   `xml:"subSystemCode,attr"`
	SubSubSystemCode    string   `xml:"subSubSystemCode,attr"`
	AssyCode            string   `xml:"assyCode,attr"`
	DisassyCode         string   `xml:"disassyCode,attr"`
	DisassyCodeVariant  string   `xml:"disassyCodeVariant,attr"`
	InfoCode            string   `xml:"infoCode,attr"`
	InfoCodeVariant     string   `xml:"infoCodeVariant,attr"`
	ItemLocationCode    string   `xml:"itemLocationCode,attr"`
	LearnCode           string   `xml:"learnCode,attr,omitempty"`
	LearnEventCode      string   `xml:"learnEventCode,attr,omitempty"`
}

// String returns the DMC as a formatted string
func (d *DMCode) String() string {
	return d.ModelIdentCode + "-" +
		d.SystemDiffCode + "-" +
		d.SystemCode + "-" +
		d.SubSystemCode + d.SubSubSystemCode + "-" +
		d.AssyCode + "-" +
		d.DisassyCode + d.DisassyCodeVariant + "-" +
		d.InfoCode + d.InfoCodeVariant + "-" +
		d.ItemLocationCode
}

// PMCode represents a Publication Module Code (PMC) in S1000D
type PMCode struct {
	XMLName         xml.Name `xml:"pmCode"`
	ModelIdentCode  string   `xml:"modelIdentCode,attr"`
	PmIssuer        string   `xml:"pmIssuer,attr"`
	PmNumber        string   `xml:"pmNumber,attr"`
	PmVolume        string   `xml:"pmVolume,attr"`
}

// String returns the PMC as a formatted string
func (p *PMCode) String() string {
	result := p.ModelIdentCode + "-" + p.PmIssuer + "-" + p.PmNumber
	if p.PmVolume != "" {
		result += "-" + p.PmVolume
	}
	return result
}

// IssueInfo represents issue information for S1000D modules
type IssueInfo struct {
	XMLName      xml.Name `xml:"issueInfo"`
	IssueNumber  string   `xml:"issueNumber,attr"`
	InWork       string   `xml:"inWork,attr,omitempty"`
}

// Language represents language information
type Language struct {
	XMLName      xml.Name `xml:"language"`
	LanguageIsoCode string `xml:"languageIsoCode,attr"`
	CountryIsoCode  string `xml:"countryIsoCode,attr"`
}

// DMIdent represents Data Module Identification
type DMIdent struct {
	XMLName   xml.Name  `xml:"dmIdent"`
	DMCode    DMCode    `xml:"dmCode"`
	IssueInfo IssueInfo `xml:"issueInfo,omitempty"`
	Language  Language  `xml:"language,omitempty"`
}

// PMIdent represents Publication Module Identification
type PMIdent struct {
	XMLName   xml.Name  `xml:"pmIdent"`
	PMCode    PMCode    `xml:"pmCode"`
	IssueInfo IssueInfo `xml:"issueInfo,omitempty"`
	Language  Language  `xml:"language,omitempty"`
}

// DMRef represents a reference to a Data Module
type DMRef struct {
	XMLName xml.Name `xml:"dmRef"`
	DMRefIdent DMRefIdent `xml:"dmRefIdent"`
}

// DMRefIdent represents the identification part of a DM reference
type DMRefIdent struct {
	XMLName   xml.Name `xml:"dmRefIdent"`
	DMCode    DMCode   `xml:"dmCode"`
	IssueInfo *IssueInfo `xml:"issueInfo,omitempty"`
}

// Security represents security classification
type Security struct {
	XMLName            xml.Name `xml:"security"`
	SecurityClassification string `xml:"securityClassification,attr"`
}
