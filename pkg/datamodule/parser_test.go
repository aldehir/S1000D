package datamodule

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDataModule(t *testing.T) {
	// Get the examples directory path
	examplesDir := filepath.Join("..", "..", "examples")
	dmFile := filepath.Join(examplesDir, "example_datamodule.xml")

	// Check if file exists
	if _, err := os.Stat(dmFile); os.IsNotExist(err) {
		t.Skipf("Example file not found: %s", dmFile)
	}

	dm, err := ParseFile(dmFile)
	if err != nil {
		t.Fatalf("Failed to parse data module: %v", err)
	}

	// Test DMC
	expectedDMC := "MYAIRCRAFT-A-00-00-00-00A-040A-D"
	if got := dm.GetDMC(); got != expectedDMC {
		t.Errorf("GetDMC() = %q, want %q", got, expectedDMC)
	}

	// Test title
	expectedTitle := "Engine - Description"
	if got := dm.GetTitle(); got != expectedTitle {
		t.Errorf("GetTitle() = %q, want %q", got, expectedTitle)
	}

	// Test issue info
	expectedIssue := "001-00"
	if got := dm.GetIssueInfo(); got != expectedIssue {
		t.Errorf("GetIssueInfo() = %q, want %q", got, expectedIssue)
	}

	// Test language
	expectedLang := "en-US"
	if got := dm.GetLanguage(); got != expectedLang {
		t.Errorf("GetLanguage() = %q, want %q", got, expectedLang)
	}

	// Test content
	if dm.Content.Description == nil {
		t.Error("Expected description content, got nil")
	}

	if len(dm.Content.Description.LevelledPara) == 0 {
		t.Error("Expected levelled paragraphs, got none")
	}
}

func TestParseProcedure(t *testing.T) {
	examplesDir := filepath.Join("..", "..", "examples")
	dmFile := filepath.Join(examplesDir, "example_procedure.xml")

	// Check if file exists
	if _, err := os.Stat(dmFile); os.IsNotExist(err) {
		t.Skipf("Example file not found: %s", dmFile)
	}

	dm, err := ParseFile(dmFile)
	if err != nil {
		t.Fatalf("Failed to parse procedure: %v", err)
	}

	// Test DMC
	expectedDMC := "MYAIRCRAFT-A-72-10-00-00A-520A-C"
	if got := dm.GetDMC(); got != expectedDMC {
		t.Errorf("GetDMC() = %q, want %q", got, expectedDMC)
	}

	// Test that it has procedure content
	if dm.Content.Procedure == nil {
		t.Error("Expected procedure content, got nil")
	}

	if dm.Content.Procedure != nil {
		steps := dm.Content.Procedure.MainProcedure.ProcedureStep
		if len(steps) == 0 {
			t.Error("Expected procedure steps, got none")
		}
		if len(steps) != 8 {
			t.Errorf("Expected 8 procedure steps, got %d", len(steps))
		}
	}
}

func TestParseBytes(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<dmodule>
  <identAndStatusSection>
    <dmAddress>
      <dmIdent>
        <dmCode modelIdentCode="TEST" systemDiffCode="A" systemCode="00"
                subSystemCode="0" subSubSystemCode="0" assyCode="00"
                disassyCode="00" disassyCodeVariant="A" infoCode="040"
                infoCodeVariant="A" itemLocationCode="D"/>
        <issueInfo issueNumber="001"/>
        <language languageIsoCode="en" countryIsoCode="US"/>
      </dmIdent>
      <dmAddressItems>
        <dmTitle>
          <techName>Test Module</techName>
        </dmTitle>
      </dmAddressItems>
    </dmAddress>
    <dmStatus>
      <security securityClassification="01"/>
    </dmStatus>
  </identAndStatusSection>
  <content>
    <description>
      <levelledPara>
        <para>Test content</para>
      </levelledPara>
    </description>
  </content>
</dmodule>`)

	dm, err := ParseBytes(xmlData)
	if err != nil {
		t.Fatalf("Failed to parse bytes: %v", err)
	}

	expectedDMC := "TEST-A-00-00-00-00A-040A-D"
	if got := dm.GetDMC(); got != expectedDMC {
		t.Errorf("GetDMC() = %q, want %q", got, expectedDMC)
	}
}
