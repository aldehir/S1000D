package markdown

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aldehir/S1000D/pkg/datamodule"
	"github.com/aldehir/S1000D/pkg/pubmodule"
)

func TestRenderDataModule_Descriptive(t *testing.T) {
	// Parse the example descriptive data module
	dm, err := datamodule.ParseFile("../../examples/example_datamodule.xml")
	if err != nil {
		t.Fatalf("Failed to parse data module: %v", err)
	}

	// Create a temporary directory for output
	tmpDir := t.TempDir()

	// Create renderer with custom options
	opts := &RendererOptions{
		OutputDir:     tmpDir,
		LinkExtension: ".md",
	}
	renderer := NewRenderer(opts)

	// Render the data module
	filepath, err := renderer.RenderDataModule(dm)
	if err != nil {
		t.Fatalf("Failed to render data module: %v", err)
	}

	// Read the generated file
	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	contentStr := string(content)

	// Verify YAML frontmatter exists
	if !strings.HasPrefix(contentStr, "---\n") {
		t.Error("File should start with YAML frontmatter")
	}

	// Verify key metadata fields are present
	expectedFields := []string{
		"dmc: MYAIRCRAFT-A-00-00-00-00A-040A-D",
		"title: Engine - Description",
		"issue: 001-00",
		"language: en-US",
		"issue_date: 2024-01-15",
		"security_classification: 01",
	}

	for _, field := range expectedFields {
		if !strings.Contains(contentStr, field) {
			t.Errorf("Expected field not found: %s", field)
		}
	}

	// Verify responsible party and originator are present (format may vary)
	if !strings.Contains(contentStr, "responsible_party:") {
		t.Error("Responsible party field not found")
	}
	if !strings.Contains(contentStr, "originator:") {
		t.Error("Originator field not found")
	}

	// Verify main title
	if !strings.Contains(contentStr, "# Engine - Description") {
		t.Error("Main title not found in content")
	}

	// Verify section headings
	if !strings.Contains(contentStr, "## General") {
		t.Error("Section heading 'General' not found")
	}

	if !strings.Contains(contentStr, "### Engine Components") {
		t.Error("Subsection heading 'Engine Components' not found")
	}

	if !strings.Contains(contentStr, "### Engine Specifications") {
		t.Error("Subsection heading 'Engine Specifications' not found")
	}

	// Verify content
	if !strings.Contains(contentStr, "This data module provides a description of the aircraft engine system.") {
		t.Error("Expected content not found")
	}

	if !strings.Contains(contentStr, "Maximum thrust: 150 kN") {
		t.Error("Expected specification not found")
	}
}

func TestRenderDataModule_Procedural(t *testing.T) {
	// Parse the example procedural data module
	dm, err := datamodule.ParseFile("../../examples/example_procedure.xml")
	if err != nil {
		t.Fatalf("Failed to parse data module: %v", err)
	}

	// Create a temporary directory for output
	tmpDir := t.TempDir()

	// Create renderer
	opts := &RendererOptions{
		OutputDir: tmpDir,
	}
	renderer := NewRenderer(opts)

	// Render the data module
	filepath, err := renderer.RenderDataModule(dm)
	if err != nil {
		t.Fatalf("Failed to render data module: %v", err)
	}

	// Read the generated file
	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	contentStr := string(content)

	// Verify metadata
	if !strings.Contains(contentStr, "dmc: MYAIRCRAFT-A-72-10-00-00A-520A-C") {
		t.Error("DMC not found in frontmatter")
	}

	if !strings.Contains(contentStr, "title: Engine Oil System - Oil Change Procedure") {
		t.Error("Title not found in frontmatter")
	}

	// Verify procedure heading
	if !strings.Contains(contentStr, "## Procedure") {
		t.Error("Procedure heading not found")
	}

	// Verify procedure steps are numbered
	if !strings.Contains(contentStr, "1. ") {
		t.Error("Numbered step 1 not found")
	}

	if !strings.Contains(contentStr, "8. ") {
		t.Error("Numbered step 8 not found")
	}

	// Verify specific step content
	if !strings.Contains(contentStr, "Ensure the engine has been shut down for at least 30 minutes") {
		t.Error("Expected procedure step content not found")
	}

	if !strings.Contains(contentStr, "Run the engine for 5 minutes and recheck the oil level") {
		t.Error("Expected final step content not found")
	}
}

func TestRenderPublicationModule(t *testing.T) {
	// Parse the example publication module
	pm, err := pubmodule.ParseFile("../../examples/example_pubmodule.xml")
	if err != nil {
		t.Fatalf("Failed to parse publication module: %v", err)
	}

	// Create a temporary directory for output
	tmpDir := t.TempDir()

	// Create renderer
	opts := &RendererOptions{
		OutputDir: tmpDir,
	}
	renderer := NewRenderer(opts)

	// Render the publication module
	files, err := renderer.RenderPublicationModule(pm)
	if err != nil {
		t.Fatalf("Failed to render publication module: %v", err)
	}

	// Verify at least one file was created
	if len(files) == 0 {
		t.Fatal("No files were created")
	}

	// Get the main publication file
	var pmFilepath string
	for _, path := range files {
		pmFilepath = path
		break
	}

	// Read the generated file
	content, err := os.ReadFile(pmFilepath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	contentStr := string(content)

	// Verify PMC in frontmatter
	if !strings.Contains(contentStr, "pmc: MYAIRCRAFT-12345-00001-01") {
		t.Error("PMC not found in frontmatter")
	}

	// Verify title
	if !strings.Contains(contentStr, "title: Aircraft Engine Maintenance Manual") {
		t.Error("Title not found in frontmatter")
	}

	// Verify data module references are listed
	if !strings.Contains(contentStr, "data_module_references:") {
		t.Error("Data module references section not found")
	}

	if !strings.Contains(contentStr, "MYAIRCRAFT-A-00-00-00-00A-040A-D") {
		t.Error("Expected DM reference not found in frontmatter")
	}

	if !strings.Contains(contentStr, "MYAIRCRAFT-A-72-10-00-00A-520A-C") {
		t.Error("Expected DM reference not found in frontmatter")
	}

	// Verify table of contents
	if !strings.Contains(contentStr, "## Table of Contents") {
		t.Error("Table of contents heading not found")
	}

	// Verify hierarchical structure
	if !strings.Contains(contentStr, "**Chapter 1 - General Information**") {
		t.Error("Chapter 1 entry not found")
	}

	if !strings.Contains(contentStr, "**Chapter 2 - Maintenance Procedures**") {
		t.Error("Chapter 2 entry not found")
	}

	if !strings.Contains(contentStr, "**Section 2.1 - Engine Oil System**") {
		t.Error("Section 2.1 entry not found")
	}

	// Verify links to data modules
	if !strings.Contains(contentStr, "[MYAIRCRAFT-A-00-00-00-00A-040A-D](MYAIRCRAFT_A_00_00_00_00A_040A_D.md)") {
		t.Error("Expected link to first data module not found")
	}

	if !strings.Contains(contentStr, "[MYAIRCRAFT-A-72-10-00-00A-520A-C](MYAIRCRAFT_A_72_10_00_00A_520A_C.md)") {
		t.Error("Expected link to second data module not found")
	}
}

func TestDMCToFilename(t *testing.T) {
	renderer := NewRenderer(nil)

	tests := []struct {
		dmc      string
		expected string
	}{
		{"MYAIRCRAFT-A-00-0-0-00A-040A-D", "MYAIRCRAFT_A_00_0_0_00A_040A_D.md"},
		{"TEST-B-12-3-4-56C-789D-E", "TEST_B_12_3_4_56C_789D_E.md"},
	}

	for _, tt := range tests {
		result := renderer.dmcToFilename(tt.dmc)
		if result != tt.expected {
			t.Errorf("dmcToFilename(%s) = %s, expected %s", tt.dmc, result, tt.expected)
		}
	}
}

func TestPMCToFilename(t *testing.T) {
	renderer := NewRenderer(nil)

	tests := []struct {
		pmc      string
		expected string
	}{
		{"MYAIRCRAFT-12345-00001-01", "MYAIRCRAFT_12345_00001_01.md"},
		{"TEST-67890-12345-02", "TEST_67890_12345_02.md"},
	}

	for _, tt := range tests {
		result := renderer.pmcToFilename(tt.pmc)
		if result != tt.expected {
			t.Errorf("pmcToFilename(%s) = %s, expected %s", tt.pmc, result, tt.expected)
		}
	}
}

func TestRendererWithDefaultOptions(t *testing.T) {
	renderer := NewRenderer(nil)

	if renderer.options == nil {
		t.Fatal("Renderer options should not be nil")
	}

	if renderer.options.OutputDir != "output" {
		t.Errorf("Default OutputDir = %s, expected 'output'", renderer.options.OutputDir)
	}

	if renderer.options.LinkExtension != ".md" {
		t.Errorf("Default LinkExtension = %s, expected '.md'", renderer.options.LinkExtension)
	}
}

func TestFileCreation(t *testing.T) {
	// Parse a data module
	dm, err := datamodule.ParseFile("../../examples/example_datamodule.xml")
	if err != nil {
		t.Fatalf("Failed to parse data module: %v", err)
	}

	// Create a temporary directory
	tmpDir := t.TempDir()

	// Create renderer
	opts := &RendererOptions{
		OutputDir: tmpDir,
	}
	renderer := NewRenderer(opts)

	// Render the data module
	filepath, err := renderer.RenderDataModule(dm)
	if err != nil {
		t.Fatalf("Failed to render data module: %v", err)
	}

	// Verify the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("File was not created: %s", filepath)
	}

	// Verify the file is in the correct directory
	if !strings.HasPrefix(filepath, tmpDir) {
		t.Errorf("File was not created in the correct directory. Expected prefix: %s, got: %s", tmpDir, filepath)
	}

	// Verify the filename is correct
	expectedFilename := "MYAIRCRAFT_A_00_00_00_00A_040A_D.md"
	if !strings.HasSuffix(filepath, expectedFilename) {
		t.Errorf("Filename is incorrect. Expected to end with: %s, got: %s", expectedFilename, filepath)
	}
}

func TestNestedDirectoryCreation(t *testing.T) {
	// Parse a data module
	dm, err := datamodule.ParseFile("../../examples/example_datamodule.xml")
	if err != nil {
		t.Fatalf("Failed to parse data module: %v", err)
	}

	// Create a temporary directory with a nested structure
	tmpDir := t.TempDir()
	nestedDir := filepath.Join(tmpDir, "level1", "level2", "level3")

	// Create renderer
	opts := &RendererOptions{
		OutputDir: nestedDir,
	}
	renderer := NewRenderer(opts)

	// Render the data module
	filepath, err := renderer.RenderDataModule(dm)
	if err != nil {
		t.Fatalf("Failed to render data module: %v", err)
	}

	// Verify the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("File was not created: %s", filepath)
	}

	// Verify the directory was created
	if _, err := os.Stat(nestedDir); os.IsNotExist(err) {
		t.Errorf("Directory was not created: %s", nestedDir)
	}
}
