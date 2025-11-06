package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/aldehir/S1000D/pkg/datamodule"
	"github.com/aldehir/S1000D/pkg/pubmodule"
)

func main() {
	fmt.Println("=== S1000D Parser Example ===\n")

	// Parse descriptive data module
	fmt.Println("1. Parsing Descriptive Data Module")
	fmt.Println("-----------------------------------")
	dmDesc, err := datamodule.ParseFile(filepath.Join("example_datamodule.xml"))
	if err != nil {
		log.Printf("Warning: Could not parse descriptive DM: %v\n", err)
	} else {
		printDataModuleInfo(dmDesc)
	}

	// Parse procedural data module
	fmt.Println("\n2. Parsing Procedural Data Module")
	fmt.Println("----------------------------------")
	dmProc, err := datamodule.ParseFile(filepath.Join("example_procedure.xml"))
	if err != nil {
		log.Printf("Warning: Could not parse procedural DM: %v\n", err)
	} else {
		printDataModuleInfo(dmProc)
		if dmProc.Content.Procedure != nil {
			steps := dmProc.Content.Procedure.MainProcedure.ProcedureStep
			fmt.Printf("  Procedure steps: %d\n", len(steps))
			if len(steps) > 0 {
				fmt.Println("  First step:", steps[0].Para[0].Content)
			}
		}
	}

	// Parse publication module
	fmt.Println("\n3. Parsing Publication Module")
	fmt.Println("------------------------------")
	pm, err := pubmodule.ParseFile(filepath.Join("example_pubmodule.xml"))
	if err != nil {
		log.Printf("Warning: Could not parse PM: %v\n", err)
	} else {
		printPublicationModuleInfo(pm)
	}
}

func printDataModuleInfo(dm *datamodule.DataModule) {
	fmt.Printf("  DMC: %s\n", dm.GetDMC())
	fmt.Printf("  Title: %s\n", dm.GetTitle())
	fmt.Printf("  Issue: %s\n", dm.GetIssueInfo())
	fmt.Printf("  Language: %s\n", dm.GetLanguage())

	// Print security classification
	secClass := dm.IdentAndStatusSection.DMStatus.Security.SecurityClassification
	fmt.Printf("  Security: %s\n", secClass)

	// Print content type
	if dm.Content.Description != nil {
		fmt.Println("  Content Type: Description")
		fmt.Printf("  Paragraphs: %d\n", len(dm.Content.Description.LevelledPara))
	} else if dm.Content.Procedure != nil {
		fmt.Println("  Content Type: Procedure")
	}
}

func printPublicationModuleInfo(pm *pubmodule.PublicationModule) {
	fmt.Printf("  PMC: %s\n", pm.GetPMC())
	fmt.Printf("  Title: %s\n", pm.GetTitle())
	fmt.Printf("  Issue: %s\n", pm.GetIssueInfo())
	fmt.Printf("  Language: %s\n", pm.GetLanguage())

	// Print referenced data modules
	dmRefs := pm.GetDataModuleRefs()
	fmt.Printf("  Referenced DMs: %d\n", len(dmRefs))
	for i, ref := range dmRefs {
		fmt.Printf("    %d. %s\n", i+1, ref)
	}
}
