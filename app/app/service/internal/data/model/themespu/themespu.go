// Code generated by entc, DO NOT EDIT.

package themespu

const (
	// Label holds the string label denoting the themespu type in the database.
	Label = "theme_spu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldThemeID holds the string denoting the theme_id field in the database.
	FieldThemeID = "theme_id"
	// FieldSpuID holds the string denoting the spu_id field in the database.
	FieldSpuID = "spu_id"
	// EdgeTheme holds the string denoting the theme edge name in mutations.
	EdgeTheme = "theme"
	// Table holds the table name of the themespu in the database.
	Table = "theme_spu"
	// ThemeTable is the table that holds the theme relation/edge.
	ThemeTable = "theme_spu"
	// ThemeInverseTable is the table name for the Theme entity.
	// It exists in this package in order to avoid lower111 circular dependency with the "theme" package.
	ThemeInverseTable = "theme"
	// ThemeColumn is the table column denoting the theme relation/edge.
	ThemeColumn = "theme_id"
)

// Columns holds all SQL columns for themespu fields.
var Columns = []string{
	FieldID,
	FieldThemeID,
	FieldSpuID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
