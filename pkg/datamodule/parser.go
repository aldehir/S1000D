package datamodule

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Parse parses a Data Module from an io.Reader
func Parse(r io.Reader) (*DataModule, error) {
	decoder := xml.NewDecoder(r)
	var dm DataModule

	if err := decoder.Decode(&dm); err != nil {
		return nil, fmt.Errorf("failed to decode data module: %w", err)
	}

	return &dm, nil
}

// ParseFile parses a Data Module from a file
func ParseFile(filename string) (*DataModule, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return Parse(file)
}

// ParseBytes parses a Data Module from a byte slice
func ParseBytes(data []byte) (*DataModule, error) {
	var dm DataModule
	if err := xml.Unmarshal(data, &dm); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data module: %w", err)
	}
	return &dm, nil
}

// GetDMC returns the Data Module Code as a string
func (dm *DataModule) GetDMC() string {
	return dm.IdentAndStatusSection.DMAddress.DMIdent.DMCode.String()
}

// GetTitle returns the data module title
func (dm *DataModule) GetTitle() string {
	techName := dm.IdentAndStatusSection.DMAddress.DMAddressItems.DMTitle.TechName
	infoName := dm.IdentAndStatusSection.DMAddress.DMAddressItems.DMTitle.InfoName

	if infoName != "" {
		return techName + " - " + infoName
	}
	return techName
}

// GetIssueInfo returns the issue information
func (dm *DataModule) GetIssueInfo() string {
	issueInfo := dm.IdentAndStatusSection.DMAddress.DMIdent.IssueInfo
	result := issueInfo.IssueNumber
	if issueInfo.InWork != "" {
		result += "-" + issueInfo.InWork
	}
	return result
}

// GetLanguage returns the language code
func (dm *DataModule) GetLanguage() string {
	lang := dm.IdentAndStatusSection.DMAddress.DMIdent.Language
	return lang.LanguageIsoCode + "-" + lang.CountryIsoCode
}
