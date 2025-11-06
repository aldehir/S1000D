package pubmodule

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Parse parses a Publication Module from an io.Reader
func Parse(r io.Reader) (*PublicationModule, error) {
	decoder := xml.NewDecoder(r)
	var pm PublicationModule

	if err := decoder.Decode(&pm); err != nil {
		return nil, fmt.Errorf("failed to decode publication module: %w", err)
	}

	return &pm, nil
}

// ParseFile parses a Publication Module from a file
func ParseFile(filename string) (*PublicationModule, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return Parse(file)
}

// ParseBytes parses a Publication Module from a byte slice
func ParseBytes(data []byte) (*PublicationModule, error) {
	var pm PublicationModule
	if err := xml.Unmarshal(data, &pm); err != nil {
		return nil, fmt.Errorf("failed to unmarshal publication module: %w", err)
	}
	return &pm, nil
}

// GetPMC returns the Publication Module Code as a string
func (pm *PublicationModule) GetPMC() string {
	return pm.IdentAndStatusSection.PMAddress.PMIdent.PMCode.String()
}

// GetTitle returns the publication module title
func (pm *PublicationModule) GetTitle() string {
	return pm.IdentAndStatusSection.PMAddress.PMAddressItems.PMTitle.Title
}

// GetIssueInfo returns the issue information
func (pm *PublicationModule) GetIssueInfo() string {
	issueInfo := pm.IdentAndStatusSection.PMAddress.PMIdent.IssueInfo
	result := issueInfo.IssueNumber
	if issueInfo.InWork != "" {
		result += "-" + issueInfo.InWork
	}
	return result
}

// GetLanguage returns the language code
func (pm *PublicationModule) GetLanguage() string {
	lang := pm.IdentAndStatusSection.PMAddress.PMIdent.Language
	return lang.LanguageIsoCode + "-" + lang.CountryIsoCode
}

// GetDataModuleRefs returns all data module references in the publication
func (pm *PublicationModule) GetDataModuleRefs() []string {
	var refs []string
	for _, entry := range pm.Content.PMEntry {
		refs = append(refs, collectDMRefs(entry)...)
	}
	return refs
}

// collectDMRefs recursively collects all DM references from PM entries
func collectDMRefs(entry PMEntry) []string {
	var refs []string

	if entry.DMRef != nil {
		refs = append(refs, entry.DMRef.DMRefIdent.DMCode.String())
	}

	for _, subEntry := range entry.PMEntry {
		refs = append(refs, collectDMRefs(subEntry)...)
	}

	return refs
}
