package metadata

// Schema struct
type Schema struct {
	Name           string
	TablesMetaData []Table
	ViewsMetaData  []Table
	EnumsMetaData  []Enum
}

// IsEmpty returns true if schema info does not contain any table, views or enums metadata
func (s Schema) IsEmpty() bool { _ = "STUB: not implemented"; return false }
