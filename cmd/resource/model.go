// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	TPSCode            *string  `json:",omitempty"`
	EKSClusterName     *string  `json:",omitempty"`
	CoverSheetIncluded *bool    `json:",omitempty"`
	DueDate            *string  `json:",omitempty"`
	ApprovalDate       *string  `json:",omitempty"`
	Memo               *Memo    `json:",omitempty"`
	SecondCopyOfMemo   *Memo    `json:",omitempty"`
	Authors            []string `json:",omitempty"`
	Tags               []Tag    `json:",omitempty"`
}

// Memo is autogenerated from the json schema
type Memo struct {
	Heading *string `json:",omitempty"`
	Body    *string `json:",omitempty"`
}

// Tag is autogenerated from the json schema
type Tag struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}
