package pubmodule

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePublicationModule(t *testing.T) {
	// Get the examples directory path
	examplesDir := filepath.Join("..", "..", "examples")
	pmFile := filepath.Join(examplesDir, "example_pubmodule.xml")

	// Check if file exists
	if _, err := os.Stat(pmFile); os.IsNotExist(err) {
		t.Skipf("Example file not found: %s", pmFile)
	}

	pm, err := ParseFile(pmFile)
	if err != nil {
		t.Fatalf("Failed to parse publication module: %v", err)
	}

	// Test PMC
	expectedPMC := "MYAIRCRAFT-12345-00001-01"
	if got := pm.GetPMC(); got != expectedPMC {
		t.Errorf("GetPMC() = %q, want %q", got, expectedPMC)
	}

	// Test title
	expectedTitle := "Aircraft Engine Maintenance Manual"
	if got := pm.GetTitle(); got != expectedTitle {
		t.Errorf("GetTitle() = %q, want %q", got, expectedTitle)
	}

	// Test issue info
	expectedIssue := "001-00"
	if got := pm.GetIssueInfo(); got != expectedIssue {
		t.Errorf("GetIssueInfo() = %q, want %q", got, expectedIssue)
	}

	// Test language
	expectedLang := "en-US"
	if got := pm.GetLanguage(); got != expectedLang {
		t.Errorf("GetLanguage() = %q, want %q", got, expectedLang)
	}

	// Test PM entries
	if len(pm.Content.PMEntry) == 0 {
		t.Error("Expected PM entries, got none")
	}

	// Test data module references
	dmRefs := pm.GetDataModuleRefs()
	if len(dmRefs) != 2 {
		t.Errorf("Expected 2 data module references, got %d", len(dmRefs))
	}

	expectedRefs := []string{
		"MYAIRCRAFT-A-00-00-00-00A-040A-D",
		"MYAIRCRAFT-A-72-10-00-00A-520A-C",
	}

	for i, expected := range expectedRefs {
		if i >= len(dmRefs) {
			t.Errorf("Missing expected DM reference: %s", expected)
			continue
		}
		if dmRefs[i] != expected {
			t.Errorf("DM reference %d = %q, want %q", i, dmRefs[i], expected)
		}
	}
}

func TestParseBytes(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<pm>
  <identAndStatusSection>
    <pmAddress>
      <pmIdent>
        <pmCode modelIdentCode="TEST" pmIssuer="00000" pmNumber="99999" pmVolume="01"/>
        <issueInfo issueNumber="001"/>
        <language languageIsoCode="en" countryIsoCode="US"/>
      </pmIdent>
      <pmAddressItems>
        <pmTitle>Test Publication</pmTitle>
      </pmAddressItems>
    </pmAddress>
    <pmStatus>
      <security securityClassification="01"/>
    </pmStatus>
  </identAndStatusSection>
  <content>
    <pmEntry>
      <pmEntryTitle>Test Entry</pmEntryTitle>
    </pmEntry>
  </content>
</pm>`)

	pm, err := ParseBytes(xmlData)
	if err != nil {
		t.Fatalf("Failed to parse bytes: %v", err)
	}

	expectedPMC := "TEST-00000-99999-01"
	if got := pm.GetPMC(); got != expectedPMC {
		t.Errorf("GetPMC() = %q, want %q", got, expectedPMC)
	}

	expectedTitle := "Test Publication"
	if got := pm.GetTitle(); got != expectedTitle {
		t.Errorf("GetTitle() = %q, want %q", got, expectedTitle)
	}
}
