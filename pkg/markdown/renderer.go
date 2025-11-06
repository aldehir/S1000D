package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aldehir/S1000D/pkg/common"
	"github.com/aldehir/S1000D/pkg/datamodule"
	"github.com/aldehir/S1000D/pkg/pubmodule"
)

// RendererOptions configures the markdown renderer behavior
type RendererOptions struct {
	// OutputDir is the directory where markdown files will be written
	OutputDir string
	// LinkExtension is the extension to use for links (default: .md)
	LinkExtension string
	// LinkFormat controls how links are formatted:
	//   - "relative" or empty: use relative file paths (default)
	//   - URI scheme (e.g., "s1000d://dmc/"): use absolute URI references
	LinkFormat string
}

// DefaultOptions returns the default renderer options
func DefaultOptions() *RendererOptions {
	return &RendererOptions{
		OutputDir:     "output",
		LinkExtension: ".md",
		LinkFormat:    "relative",
	}
}

// Renderer handles conversion of S1000D documents to Markdown
type Renderer struct {
	options *RendererOptions
}

// NewRenderer creates a new Markdown renderer with the given options
func NewRenderer(opts *RendererOptions) *Renderer {
	if opts == nil {
		opts = DefaultOptions()
	}
	return &Renderer{options: opts}
}

// RenderDataModule converts a DataModule to a Markdown file
// Returns the path to the created file
func (r *Renderer) RenderDataModule(dm *datamodule.DataModule) (string, error) {
	// Generate filename from DMC
	filename := r.dmcToFilename(dm.GetDMC())
	filePath := filepath.Join(r.options.OutputDir, filename)

	// Create the markdown content
	content := r.dataModuleToMarkdown(dm)

	// Write to file
	if err := os.MkdirAll(r.options.OutputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return filePath, nil
}

// RenderPublicationModule converts a PublicationModule to a hierarchy of Markdown files
// Returns a map of filenames to their paths
func (r *Renderer) RenderPublicationModule(pm *pubmodule.PublicationModule) (map[string]string, error) {
	files := make(map[string]string)

	// Generate filename for the publication module itself
	filename := r.pmcToFilename(pm.GetPMC())
	filePath := filepath.Join(r.options.OutputDir, filename)

	// Create the markdown content
	content := r.publicationModuleToMarkdown(pm)

	// Write to file
	if err := os.MkdirAll(r.options.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	files[filename] = filePath
	return files, nil
}

// dataModuleToMarkdown converts a DataModule to Markdown format
func (r *Renderer) dataModuleToMarkdown(dm *datamodule.DataModule) string {
	var sb strings.Builder

	// Write YAML frontmatter
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("dmc: %s\n", dm.GetDMC()))
	sb.WriteString(fmt.Sprintf("title: %s\n", dm.GetTitle()))
	sb.WriteString(fmt.Sprintf("issue: %s\n", dm.GetIssueInfo()))
	sb.WriteString(fmt.Sprintf("language: %s\n", dm.GetLanguage()))

	// Add issue date
	issueDate := dm.IdentAndStatusSection.DMAddress.DMAddressItems.IssueDate
	sb.WriteString(fmt.Sprintf("issue_date: %s-%s-%s\n", issueDate.Year, issueDate.Month, issueDate.Day))

	// Add security classification
	sb.WriteString(fmt.Sprintf("security_classification: %s\n",
		dm.IdentAndStatusSection.DMStatus.Security.SecurityClassification))

	// Add responsible party if available
	if dm.IdentAndStatusSection.DMStatus.ResponsiblePartnerCompany != nil {
		rpc := dm.IdentAndStatusSection.DMStatus.ResponsiblePartnerCompany
		sb.WriteString(fmt.Sprintf("responsible_party: \"%s (%s)\"\n",
			strings.TrimSpace(rpc.EnterpriseName), rpc.EnterpriseCode))
	}

	// Add originator if available
	if dm.IdentAndStatusSection.DMStatus.Originator != nil {
		orig := dm.IdentAndStatusSection.DMStatus.Originator
		sb.WriteString(fmt.Sprintf("originator: \"%s (%s)\"\n",
			strings.TrimSpace(orig.EnterpriseName), orig.EnterpriseCode))
	}

	// Add applicability if available
	if dm.IdentAndStatusSection.DMStatus.AppLicability != nil {
		sb.WriteString(fmt.Sprintf("applicability: %s\n",
			dm.IdentAndStatusSection.DMStatus.AppLicability.DisplayText))
	}

	// Add business rule exchange references if available
	if len(dm.IdentAndStatusSection.DMStatus.BrExDmRef) > 0 {
		sb.WriteString("brex_references:\n")
		for _, ref := range dm.IdentAndStatusSection.DMStatus.BrExDmRef {
			dmc := r.dmRefToString(&ref)
			sb.WriteString(fmt.Sprintf("  - %s\n", dmc))
		}
	}

	sb.WriteString("---\n\n")

	// Write title
	sb.WriteString(fmt.Sprintf("# %s\n\n", dm.GetTitle()))

	// Write content based on type
	if dm.Content.Description != nil {
		r.writeDescription(&sb, dm.Content.Description)
	} else if dm.Content.Procedure != nil {
		r.writeProcedure(&sb, dm.Content.Procedure)
	}

	return sb.String()
}

// publicationModuleToMarkdown converts a PublicationModule to Markdown format
func (r *Renderer) publicationModuleToMarkdown(pm *pubmodule.PublicationModule) string {
	var sb strings.Builder

	// Write YAML frontmatter
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("pmc: %s\n", pm.GetPMC()))
	sb.WriteString(fmt.Sprintf("title: %s\n", pm.GetTitle()))
	sb.WriteString(fmt.Sprintf("issue: %s\n", pm.GetIssueInfo()))
	sb.WriteString(fmt.Sprintf("language: %s\n", pm.GetLanguage()))

	// Add issue date
	issueDate := pm.IdentAndStatusSection.PMAddress.PMAddressItems.IssueDate
	sb.WriteString(fmt.Sprintf("issue_date: %s-%s-%s\n", issueDate.Year, issueDate.Month, issueDate.Day))

	// Add security classification
	sb.WriteString(fmt.Sprintf("security_classification: %s\n",
		pm.IdentAndStatusSection.PMStatus.Security.SecurityClassification))

	// Add data module references
	dmRefs := pm.GetDataModuleRefs()
	if len(dmRefs) > 0 {
		sb.WriteString("data_module_references:\n")
		for _, ref := range dmRefs {
			sb.WriteString(fmt.Sprintf("  - %s\n", ref))
		}
	}

	sb.WriteString("---\n\n")

	// Write title
	sb.WriteString(fmt.Sprintf("# %s\n\n", pm.GetTitle()))

	// Write table of contents
	sb.WriteString("## Table of Contents\n\n")
	r.writePMEntries(&sb, pm.Content.PMEntry, 1)

	return sb.String()
}

// writeDescription writes descriptive content to the string builder
func (r *Renderer) writeDescription(sb *strings.Builder, desc *datamodule.Description) {
	for _, lp := range desc.LevelledPara {
		r.writeLevelledPara(sb, &lp, 2)
	}
}

// writeLevelledPara writes a levelled paragraph and its children recursively
func (r *Renderer) writeLevelledPara(sb *strings.Builder, lp *datamodule.LevelledPara, level int) {
	// Write the heading
	if lp.Title != "" {
		sb.WriteString(strings.Repeat("#", level))
		sb.WriteString(fmt.Sprintf(" %s\n\n", lp.Title))
	}

	// Write paragraphs
	for _, para := range lp.Para {
		sb.WriteString(r.cleanParaContent(para.Content))
		sb.WriteString("\n\n")
	}

	// Write nested levelled paragraphs
	for _, nestedLp := range lp.LevelledPara {
		r.writeLevelledPara(sb, &nestedLp, level+1)
	}
}

// writeProcedure writes procedural content to the string builder
func (r *Renderer) writeProcedure(sb *strings.Builder, proc *datamodule.Procedure) {
	sb.WriteString("## Procedure\n\n")

	for i, step := range proc.MainProcedure.ProcedureStep {
		sb.WriteString(fmt.Sprintf("%d. ", i+1))

		for j, para := range step.Para {
			if j > 0 {
				sb.WriteString("\n   ")
			}
			sb.WriteString(r.cleanParaContent(para.Content))
		}
		sb.WriteString("\n\n")
	}
}

// writePMEntries writes publication module entries recursively
func (r *Renderer) writePMEntries(sb *strings.Builder, entries []pubmodule.PMEntry, level int) {
	for _, entry := range entries {
		indent := strings.Repeat("  ", level-1)

		if entry.PMEntryTitle != "" {
			sb.WriteString(fmt.Sprintf("%s- **%s**\n", indent, entry.PMEntryTitle))
		}

		if entry.DMRef != nil {
			dmc := r.dmRefToString(entry.DMRef)
			link := r.dmcToLink(dmc)
			sb.WriteString(fmt.Sprintf("%s  - [%s](%s)\n", indent, dmc, link))
		}

		if len(entry.PMEntry) > 0 {
			r.writePMEntries(sb, entry.PMEntry, level+1)
		}
	}
}

// cleanParaContent cleans paragraph content for markdown
func (r *Renderer) cleanParaContent(content string) string {
	// Remove leading/trailing whitespace
	content = strings.TrimSpace(content)

	// Basic XML tag handling - this could be expanded for more sophisticated conversion
	// For now, we'll strip most tags but preserve some common ones
	content = strings.ReplaceAll(content, "<emphasis>", "*")
	content = strings.ReplaceAll(content, "</emphasis>", "*")
	content = strings.ReplaceAll(content, "<emphasis emphasisType=\"em01\">", "*")
	content = strings.ReplaceAll(content, "<emphasis emphasisType=\"em02\">", "**")

	// Remove other XML tags (simple approach)
	// A more sophisticated implementation would use proper XML parsing

	return content
}

// dmcToFilename converts a DMC string to a valid filename
func (r *Renderer) dmcToFilename(dmc string) string {
	// Replace special characters with underscores
	filename := strings.ReplaceAll(dmc, "-", "_")
	return filename + ".md"
}

// pmcToFilename converts a PMC string to a valid filename
func (r *Renderer) pmcToFilename(pmc string) string {
	// Replace special characters with underscores
	filename := strings.ReplaceAll(pmc, "-", "_")
	return filename + ".md"
}

// dmcToLink converts a DMC string to a markdown link
func (r *Renderer) dmcToLink(dmc string) string {
	// If LinkFormat is empty or "relative", use relative file path
	if r.options.LinkFormat == "" || r.options.LinkFormat == "relative" {
		filename := r.dmcToFilename(dmc)
		return filename
	}

	// Otherwise, use the LinkFormat as a URI scheme prefix
	return r.options.LinkFormat + dmc
}

// dmRefToString converts a DMRef to a string representation
func (r *Renderer) dmRefToString(dmRef *common.DMRef) string {
	return dmRef.DMRefIdent.DMCode.String()
}
