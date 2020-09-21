/*
Package checkconfigurations defines the configuration of each check:
- metadata
- output template
- under which conditions it's enabled
- the level of a failure
- which function implements it
*/
package checkconfigurations

import (
	"github.com/arduino/arduino-check/check/checkfunctions"
	"github.com/arduino/arduino-check/configuration/checkmode"
	"github.com/arduino/arduino-check/project/projecttype"
)

// Type is the type for check configurations.
type Type struct {
	ProjectType projecttype.Type // The project type the check applies to.
	// The following fields provide arbitrary text for the tool output associated with each check:
	Category        string
	Subcategory     string
	ID              string // Unique check identifier: <project type identifier (L|S|P|I)><category identifier><number>
	Name            string // Short description of the check.
	Description     string // Supplemental information about the check.
	MessageTemplate string // The warning/error message template displayed when the check fails. Will be filled by check function output.
	// The following fields define under which tool configuration modes the check will run:
	DisableModes []checkmode.Type // Check is disabled when tool is in any of these modes.
	EnableModes  []checkmode.Type // Check is only enabled when tool is in one of these modes.
	// The following fields define the check level in each configuration mode:
	InfoModes     []checkmode.Type    // Failure of the check only results in an informational message.
	WarningModes  []checkmode.Type    // Failure of the check is considered a warning.
	ErrorModes    []checkmode.Type    // Failure of the check is considered an error.
	CheckFunction checkfunctions.Type // The function that implements the check.
}

// Checks is an array of structs that define the configuration of each check.
var Configurations = []Type{
	{
		ProjectType:     projecttype.Library,
		Category:        "library.properties",
		Subcategory:     "general",
		ID:              "LP001",
		Name:            "invalid format",
		Description:     "",
		MessageTemplate: "library.properties has an invalid format: {{.}}",
		DisableModes:    nil,
		EnableModes:     []checkmode.Type{checkmode.All},
		InfoModes:       nil,
		WarningModes:    nil,
		ErrorModes:      []checkmode.Type{checkmode.All},
		CheckFunction:   checkfunctions.LibraryPropertiesFormat,
	},
	{
		ProjectType:     projecttype.Library,
		Category:        "library.properties",
		Subcategory:     "name field",
		ID:              "LP002",
		Name:            "missing name field",
		Description:     "",
		MessageTemplate: "missing name field in library.properties",
		DisableModes:    nil,
		EnableModes:     []checkmode.Type{checkmode.All},
		InfoModes:       nil,
		WarningModes:    nil,
		ErrorModes:      []checkmode.Type{checkmode.All},
		CheckFunction:   checkfunctions.LibraryPropertiesNameFieldMissing,
	},
	{
		ProjectType:     projecttype.Library,
		Category:        "library.properties",
		Subcategory:     "name field",
		ID:              "LP003",
		Name:            "disallowed characters",
		Description:     "",
		MessageTemplate: "disallowed characters in library.properties name field. See: https://arduino.github.io/arduino-cli/latest/library-specification/#libraryproperties-file-format",
		DisableModes:    nil,
		EnableModes:     []checkmode.Type{checkmode.All},
		InfoModes:       nil,
		WarningModes:    nil,
		ErrorModes:      []checkmode.Type{checkmode.All},
		CheckFunction:   checkfunctions.LibraryPropertiesNameFieldDisallowedCharacters,
	},
	{
		ProjectType:     projecttype.Library,
		Category:        "library.properties",
		Subcategory:     "version field",
		ID:              "LP004",
		Name:            "missing version field",
		Description:     "",
		MessageTemplate: "missing version field in library.properties",
		DisableModes:    nil,
		EnableModes:     []checkmode.Type{checkmode.All},
		InfoModes:       nil,
		WarningModes:    nil,
		ErrorModes:      []checkmode.Type{checkmode.All},
		CheckFunction:   checkfunctions.LibraryPropertiesVersionFieldMissing,
	},
	{
		ProjectType:     projecttype.Sketch,
		Category:        "structure",
		Subcategory:     "",
		ID:              "SS001",
		Name:            ".pde extension",
		Description:     "The .pde extension is used by both Processing sketches and Arduino sketches. Processing sketches should either be in the \"data\" subfolder of the sketch or in the \"extras\" folder of the library. Arduino sketches should use the modern .ino extension",
		MessageTemplate: "{{.}} uses deprecated .pde file extension. Use .ino for Arduino sketches",
		DisableModes:    nil,
		EnableModes:     []checkmode.Type{checkmode.All},
		InfoModes:       nil,
		WarningModes:    []checkmode.Type{checkmode.Permissive},
		ErrorModes:      []checkmode.Type{checkmode.Default},
		CheckFunction:   checkfunctions.PdeSketchExtension,
	},
}
