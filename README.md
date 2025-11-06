# S1000D Parser

A Go library for parsing S1000D technical publications. S1000D is an international specification for technical publications using a common source database for complex systems, particularly in aerospace, defense, and heavy industry.

## Features

- Parse S1000D Data Modules (DM) - Issue 5.0 compatible
- Parse S1000D Publication Modules (PM) - Issue 5.0 compatible
- Support for both descriptive and procedural content
- Extract data module codes (DMC), publication module codes (PMC), and metadata
- Type-safe Go structures for S1000D XML elements
- **Markdown Renderer**: Convert S1000D documents to Markdown format with YAML frontmatter

## Installation

```bash
go get github.com/aldehir/S1000D
```

## Quick Start

### Parsing a Data Module

```go
package main

import (
    "fmt"
    "log"

    "github.com/aldehir/S1000D/pkg/datamodule"
)

func main() {
    // Parse from file
    dm, err := datamodule.ParseFile("path/to/datamodule.xml")
    if err != nil {
        log.Fatal(err)
    }

    // Access data module information
    fmt.Println("DMC:", dm.GetDMC())
    fmt.Println("Title:", dm.GetTitle())
    fmt.Println("Issue:", dm.GetIssueInfo())
    fmt.Println("Language:", dm.GetLanguage())

    // Access content
    if dm.Content.Description != nil {
        fmt.Println("Content type: Description")
        for _, para := range dm.Content.Description.LevelledPara {
            fmt.Println("  Paragraph:", para.Title)
        }
    }

    if dm.Content.Procedure != nil {
        fmt.Println("Content type: Procedure")
        steps := dm.Content.Procedure.MainProcedure.ProcedureStep
        fmt.Printf("  Number of steps: %d\n", len(steps))
    }
}
```

### Parsing a Publication Module

```go
package main

import (
    "fmt"
    "log"

    "github.com/aldehir/S1000D/pkg/pubmodule"
)

func main() {
    // Parse from file
    pm, err := pubmodule.ParseFile("path/to/pubmodule.xml")
    if err != nil {
        log.Fatal(err)
    }

    // Access publication module information
    fmt.Println("PMC:", pm.GetPMC())
    fmt.Println("Title:", pm.GetTitle())
    fmt.Println("Issue:", pm.GetIssueInfo())

    // Get all referenced data modules
    dmRefs := pm.GetDataModuleRefs()
    fmt.Println("\nReferenced Data Modules:")
    for _, ref := range dmRefs {
        fmt.Println("  -", ref)
    }
}
```

### Parsing from Different Sources

```go
// From file
dm, err := datamodule.ParseFile("datamodule.xml")

// From io.Reader
file, _ := os.Open("datamodule.xml")
dm, err := datamodule.Parse(file)

// From byte slice
xmlData := []byte(`<dmodule>...</dmodule>`)
dm, err := datamodule.ParseBytes(xmlData)
```

## Markdown Renderer

The markdown renderer converts parsed S1000D documents into Markdown files with YAML frontmatter for metadata.

### Rendering Data Modules

```go
package main

import (
    "log"

    "github.com/aldehir/S1000D/pkg/datamodule"
    "github.com/aldehir/S1000D/pkg/markdown"
)

func main() {
    // Parse a data module
    dm, err := datamodule.ParseFile("datamodule.xml")
    if err != nil {
        log.Fatal(err)
    }

    // Create markdown renderer
    opts := &markdown.RendererOptions{
        OutputDir:     "markdown_output",
        LinkExtension: ".md",
    }
    renderer := markdown.NewRenderer(opts)

    // Render to markdown
    mdPath, err := renderer.RenderDataModule(dm)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Rendered to:", mdPath)
}
```

### Rendering Publication Modules

```go
// Parse a publication module
pm, err := pubmodule.ParseFile("pubmodule.xml")
if err != nil {
    log.Fatal(err)
}

// Render to markdown
renderer := markdown.NewRenderer(nil) // Uses default options
files, err := renderer.RenderPublicationModule(pm)
if err != nil {
    log.Fatal(err)
}

// Print all generated files
for filename, path := range files {
    fmt.Printf("%s -> %s\n", filename, path)
}
```

### Markdown Output Format

The renderer creates markdown files with:

**YAML Frontmatter** containing:
- DMC/PMC identifier
- Title
- Issue information
- Language
- Issue date
- Security classification
- Responsible party and originator
- Data module references (for PMs)
- Business rule exchange references (if applicable)

**Content** formatted as:
- Hierarchical headings from levelled paragraphs
- Numbered lists for procedural steps
- Links to referenced data modules using relative paths

Example output:

```markdown
---
dmc: MYAIRCRAFT-A-00-00-00-00A-040A-D
title: Engine - Description
issue: 001-00
language: en-US
issue_date: 2024-01-15
security_classification: 01
responsible_party: "Aerospace Manufacturing Inc. (12345)"
---

# Engine - Description

## General

This data module provides a description of the aircraft engine system.

### Engine Components

The engine consists of the following major components:

1. Compressor section
2. Combustion chamber
...
```

## Project Structure

```
S1000D/
├── pkg/
│   ├── common/          # Common structures (DMC, PMC, identifiers)
│   ├── datamodule/      # Data Module parser and structures
│   ├── pubmodule/       # Publication Module parser and structures
│   └── markdown/        # Markdown renderer for S1000D documents
├── examples/            # Example S1000D XML files and usage
│   ├── main.go
│   ├── example_datamodule.xml
│   ├── example_procedure.xml
│   └── example_pubmodule.xml
└── internal/            # Internal utilities
```

## S1000D Components

### Data Module (DM)
Data Modules are the smallest self-contained units of technical information in S1000D. Each DM has:
- **Data Module Code (DMC)**: Unique identifier
- **Identification and Status Section**: Metadata about the module
- **Content**: Either descriptive or procedural content

Example DMC: `MYAIRCRAFT-A-00-00-00-00A-040A-D`

### Publication Module (PM)
Publication Modules define how Data Modules are organized into publications. A PM contains:
- **Publication Module Code (PMC)**: Unique identifier
- **Identification and Status Section**: Metadata about the publication
- **Content**: Hierarchical structure of PM entries and DM references

Example PMC: `MYAIRCRAFT-12345-00001-01`

## Supported Content Types

### Descriptive Content
- Levelled paragraphs with hierarchical structure
- Technical names and information names
- Nested paragraph support

### Procedural Content
- Main procedures with sequential steps
- Procedural steps with detailed instructions
- Support for complex multi-step procedures

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Examples

See the `examples/` directory for complete S1000D XML files that demonstrate:
- Descriptive data modules
- Procedural data modules
- Publication modules with nested structure

## API Documentation

For detailed API documentation, see the [GoDoc](https://pkg.go.dev/github.com/aldehir/S1000D).

### Key Types

#### Data Module
```go
type DataModule struct {
    IdentAndStatusSection IdentAndStatusSection
    Content               Content
}
```

#### Publication Module
```go
type PublicationModule struct {
    IdentAndStatusSection IdentAndStatusSection
    Content               Content
}
```

#### Common Structures
- `DMCode`: Data Module Code
- `PMCode`: Publication Module Code
- `IssueInfo`: Issue and in-work information
- `Language`: Language and country codes
- `Security`: Security classification

## Limitations

This parser currently supports:
- S1000D Issue 5.0
- Basic descriptive and procedural content
- Standard data module and publication module structures

Not yet implemented:
- Common Information Repository (CIR) parsing
- Information Management (IM) parsing
- All content types (only description and procedure)
- Full applicability cross-reference tables (ACTs)
- Illustration graphics (ICN/CGM)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source. Please check the LICENSE file for details.

## References

- [S1000D Official Website](http://www.s1000d.org/)
- [S1000D Specification](http://www.s1000d.org/Pages/SpecificationDownload.aspx)
- [ASD S1000D](https://www.asd-europe.org/s-series)

## Support

For issues, questions, or contributions, please open an issue on GitHub.
