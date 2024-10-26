package errors

import "fmt"

// Code fields: [App][Component][SubComponent][ErrType]
// Tree-like structure
// App: backend
//    Component: handler
//       SubComponents:
//         - unknown
//         - users
//         - records
//         - analytics
//       ErrType: unknown
//       ErrType: validation error
//       ErrType: authorization error
//       ...
//    Component: job
//       SubComponents:
//         - unknown
//         - sync
//         - analytics
//       ErrType: unknown
//       ErrType: database query
//       ErrType: external API request
//       ErrType: timeout
// App: frontend
//    Class

type AppCode uint8
type ComponentCode uint8
type SubComponentCode uint8
type ErrorCode uint8

type AppComponentErrorCode struct {
	App          AppCode
	Component    ComponentCode
	SubComponent SubComponentCode
	ErrType      ErrorCode
}

type ErrorInfo struct {
	Value       ErrorCode
	Name        string
	Description string
}

type SubComponentInfo struct {
	Value       SubComponentCode
	Name        string
	Description string
	ErrorTypes  []ErrorInfo
}

type ComponentInfo struct {
	Value         ComponentCode
	Name          string
	Description   string
	SubComponents []SubComponentInfo
}

type AppInfo struct {
	Value       AppCode
	Name        string
	Description string
	Components  []ComponentInfo
}

var CodeTree = []AppInfo{
	{
		Value:       1,
		Name:        "backend",
		Description: "Backend application services",
		Components: []ComponentInfo{
			{
				Value:       1,
				Name:        "handler",
				Description: "Request handler services",
				SubComponents: []SubComponentInfo{
					{
						Value:       0,
						Name:        "unknown",
						Description: "Unknown handler component",
						ErrorTypes: []ErrorInfo{
							{
								Value:       0,
								Name:        "unknown",
								Description: "Unknown handler error",
							},
						},
					},
					{
						Value:       1,
						Name:        "users",
						Description: "User management handler",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "validation_error",
								Description: "Input validation failed for user operation",
							},
							{
								Value:       2,
								Name:        "authorization_error",
								Description: "User lacks required permissions for operation",
							},
						},
					},
					{
						Value:       2,
						Name:        "records",
						Description: "Record management handler",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "validation_error",
								Description: "Input validation failed for record operation",
							},
							{
								Value:       2,
								Name:        "authorization_error",
								Description: "User lacks required permissions for record operation",
							},
						},
					},
					{
						Value:       3,
						Name:        "analytics",
						Description: "Analytics data handler",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "validation_error",
								Description: "Input validation failed for analytics operation",
							},
							{
								Value:       2,
								Name:        "authorization_error",
								Description: "User lacks required permissions for analytics operation",
							},
						},
					},
				},
			},
			{
				Value:       2,
				Name:        "job",
				Description: "Background job processor",
				SubComponents: []SubComponentInfo{
					{
						Value:       0,
						Name:        "unknown",
						Description: "Unknown job component",
						ErrorTypes: []ErrorInfo{
							{
								Value:       0,
								Name:        "unknown",
								Description: "Unknown job error",
							},
						},
					},
					{
						Value:       1,
						Name:        "sync",
						Description: "Data synchronization job",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "database_error",
								Description: "Database operation failed during sync",
							},
							{
								Value:       2,
								Name:        "external_api_error",
								Description: "External API call failed during sync",
							},
							{
								Value:       3,
								Name:        "timeout",
								Description: "Operation timed out during sync",
							},
						},
					},
					{
						Value:       2,
						Name:        "analytics",
						Description: "Analytics processing job",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "database_error",
								Description: "Database operation failed during analytics processing",
							},
							{
								Value:       2,
								Name:        "external_api_error",
								Description: "External API call failed during analytics processing",
							},
							{
								Value:       3,
								Name:        "timeout",
								Description: "Operation timed out during analytics processing",
							},
						},
					},
				},
			},
		},
	},
	{
		Value:       2,
		Name:        "frontend",
		Description: "Frontend application",
		Components: []ComponentInfo{
			{
				Value:       1,
				Name:        "ui",
				Description: "User interface components",
				SubComponents: []SubComponentInfo{
					{
						Value:       0,
						Name:        "unknown",
						Description: "Unknown UI component",
						ErrorTypes: []ErrorInfo{
							{
								Value:       0,
								Name:        "unknown",
								Description: "Unknown UI error",
							},
						},
					},
					{
						Value:       1,
						Name:        "forms",
						Description: "Form handling and validation",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "validation_error",
								Description: "Form validation failed",
							},
							{
								Value:       2,
								Name:        "submission_error",
								Description: "Form submission failed",
							},
						},
					},
					{
						Value:       2,
						Name:        "routing",
						Description: "Client-side routing",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "not_found",
								Description: "Route not found",
							},
							{
								Value:       2,
								Name:        "unauthorized",
								Description: "Route access unauthorized",
							},
						},
					},
				},
			},
			{
				Value:       2,
				Name:        "state",
				Description: "State management",
				SubComponents: []SubComponentInfo{
					{
						Value:       0,
						Name:        "unknown",
						Description: "Unknown state management error",
						ErrorTypes: []ErrorInfo{
							{
								Value:       0,
								Name:        "unknown",
								Description: "Unknown state error",
							},
						},
					},
					{
						Value:       1,
						Name:        "store",
						Description: "State store operations",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "update_failed",
								Description: "State update operation failed",
							},
							{
								Value:       2,
								Name:        "invalid_action",
								Description: "Invalid state action dispatched",
							},
						},
					},
					{
						Value:       2,
						Name:        "persistence",
						Description: "State persistence",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "storage_error",
								Description: "Local storage operation failed",
							},
							{
								Value:       2,
								Name:        "sync_error",
								Description: "State synchronization failed",
							},
						},
					},
				},
			},
			{
				Value:       3,
				Name:        "api",
				Description: "API client",
				SubComponents: []SubComponentInfo{
					{
						Value:       0,
						Name:        "unknown",
						Description: "Unknown API client error",
						ErrorTypes: []ErrorInfo{
							{
								Value:       0,
								Name:        "unknown",
								Description: "Unknown API error",
							},
						},
					},
					{
						Value:       1,
						Name:        "request",
						Description: "API request handling",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "network_error",
								Description: "Network request failed",
							},
							{
								Value:       2,
								Name:        "timeout",
								Description: "Request timed out",
							},
							{
								Value:       3,
								Name:        "invalid_response",
								Description: "Invalid response received",
							},
						},
					},
					{
						Value:       2,
						Name:        "cache",
						Description: "API response caching",
						ErrorTypes: []ErrorInfo{
							{
								Value:       1,
								Name:        "cache_miss",
								Description: "Cache miss error",
							},
							{
								Value:       2,
								Name:        "cache_invalid",
								Description: "Cache invalidation error",
							},
						},
					},
				},
			},
		},
	},
}

func (AppComponentErrorCode) GetType() CodeType {
	return CodeTypeAppComponent
}

// Encode returns the byte representation of the error code
func (e AppComponentErrorCode) Encode() []byte {
	result := make([]byte, 4)

	result[0] = byte(e.GetType())

	// Byte 1: [AAAACCCC]
	result[1] = (byte(e.App) & 0x0F) << 4        // 4 bits for App
	result[1] |= (byte(e.Component) >> 2) & 0x0F // Upper 4 bits of Component

	// Byte 2: [CCSSSSSS]
	result[2] = (byte(e.Component) & 0b00000011) << 6 // Lower 2 bits of Component
	result[2] |= byte(e.SubComponent) & 0b00111111    // 6 bits for SubComponent

	// Byte 3: [EEEEEEEE]
	result[3] = byte(e.ErrType) // 8 bits for ErrorType

	return result
}

// Decode creates an AppComponentErrorCode from a byte slice
func DecodeAppComponentErrorCode(data []byte) (AppComponentErrorCode, error) {
	if len(data) != 4 {
		return AppComponentErrorCode{}, fmt.Errorf("expected 4 bytes, got %d", len(data))
	}

	codeType := CodeType(data[0])
	if codeType != CodeTypeAppComponent {
		return AppComponentErrorCode{}, fmt.Errorf("invalid code type: %d", codeType)
	}

	return AppComponentErrorCode{
		App:          AppCode((data[1] >> 4) & 0x0F),
		Component:    ComponentCode(((data[1] & 0x0F) << 2) | ((data[2] >> 6) & 0b00000011)),
		SubComponent: SubComponentCode(data[2] & 0b00111111),
		ErrType:      ErrorCode(data[3]),
	}, nil
}

// String returns a human-readable representation of the error code
func (e AppComponentErrorCode) String() string {
	for _, app := range CodeTree {
		if app.Value != e.App {
			continue
		}
		for _, comp := range app.Components {
			if comp.Value != e.Component {
				continue
			}
			for _, subComp := range comp.SubComponents {
				if subComp.Value != e.SubComponent {
					continue
				}
				for _, errType := range subComp.ErrorTypes {
					if errType.Value != e.ErrType {
						continue
					}
					return fmt.Sprintf("%s.%s.%s.%s",
						app.Name, comp.Name, subComp.Name, errType.Name)
				}
			}
		}
	}
	return "invalid"
}

func (AppComponentErrorCode) GetPrefix() string {
	return "E"
}

func (AppComponentErrorCode) GetDocSection() DocSection {
	return DocSection{
		Title: "App Component Format",
		Description: `Each error code is composed of three bytes encoded as follows:
- App (4 bits): Identifies the application (allows up to 16 apps)
- Component (6 bits): Identifies the major component (allows up to 64 components per app)
- SubComponent (6 bits): Identifies the specific sub-component (allows up to 64 sub-components per component)
- ErrorType (8 bits): Identifies the specific error (allows up to 256 error types per sub-component)

The format provides:
- Up to 16 different applications
- Up to 64 different components per application
- Up to 64 different sub-components per component
- Up to 256 different error types per sub-component
- Total of 16,777,216 possible unique error codes (16 * 64 * 64 * 256)

Byte layout:
` + "```" + `
[AAAACCCC][CCSSSSSS][EEEEEEEE]
A: App bits
C: Component bits
S: SubComponent bits
E: ErrorType bits
` + "```" + ``,
		Headers: []string{"Code", "App.Component.SubComponent.Type", "Description"},
	}
}

func (AppComponentErrorCode) GetFieldInfo() []FieldInfo {
	return []FieldInfo{
		{
			Name:        "App",
			Bits:        4,
			Description: "Identifies the application (0-15)",
			Values:      "backend(1), frontend(2)",
		},
		{
			Name:        "Component",
			Bits:        6,
			Description: "Identifies the major component (0-63)",
			Values:      "handler(1), job(2), ui(1), state(2), api(3)",
		},
		{
			Name:        "SubComponent",
			Bits:        6,
			Description: "Identifies the specific sub-component (0-63)",
			Values:      "unknown(0), users(1), records(2), analytics(3), forms(1), routing(2), store(1), persistence(2), request(1), cache(2)",
		},
		{
			Name:        "ErrorType",
			Bits:        8,
			Description: "Identifies the specific error (0-255)",
			Values:      "unknown(0), validation_error(1), authorization_error(2), database_error(1), external_api_error(2), timeout(3), etc",
		},
	}
}

// GetPermutations returns all possible error code combinations
func (AppComponentErrorCode) GetPermutations() []Permutation {
	var perms []Permutation

	for _, app := range CodeTree {
		for _, comp := range app.Components {
			for _, subComp := range comp.SubComponents {
				for _, errType := range subComp.ErrorTypes {
					ace := AppComponentErrorCode{
						App:          app.Value,
						Component:    comp.Value,
						SubComponent: subComp.Value,
						ErrType:      errType.Value,
					}
					encoded := ace.Encode()

					// Format code with all bytes needing 2 hex chars
					code := fmt.Sprintf("%s%02X%02X%02X%02X",
						ace.GetPrefix(),
						encoded[0], // type byte
						encoded[1], // first data byte  [AAAACCCC]
						encoded[2], // second data byte [CCSSSSSS]
						encoded[3]) // third data byte  [EEEEEEEE]

					// Validate by decoding and comparing
					decoded, err := DecodeAppComponentErrorCode(encoded)
					if err != nil {
						panic(fmt.Sprintf("failed to decode app component code: %v", err))
					}
					if decoded.App != ace.App ||
						decoded.Component != ace.Component ||
						decoded.SubComponent != ace.SubComponent ||
						decoded.ErrType != ace.ErrType {
						panic(fmt.Sprintf("decode mismatch: got %v, want %v", decoded, ace))
					}

					path := fmt.Sprintf("%s.%s.%s.%s",
						app.Name, comp.Name, subComp.Name, errType.Name)

					perms = append(perms, Permutation{
						Type: CodeTypeAppComponent,
						Code: code,
						Fields: map[string]string{
							"App":          app.Name,
							"Component":    comp.Name,
							"SubComponent": subComp.Name,
							"ErrorType":    errType.Name,
							"Description":  errType.Description,
						},
						TableFields: []string{
							code,
							path,
							errType.Description,
						},
					})
				}
			}
		}
	}
	return perms
}
