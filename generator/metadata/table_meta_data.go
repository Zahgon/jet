package metadata

// Table metadata struct
type Table struct {
	Name    string `sql:"primary_key"`
	Comment string
	Columns []Column
}

// MutableColumns returns list of mutable columns for table
func (t Table) MutableColumns() []Column { _ = "STUB: not implemented"; return nil }

// DefaultColumns returns list of columns with default values set for table
func (t Table) DefaultColumns() []Column { _ = "STUB: not implemented"; return nil }
