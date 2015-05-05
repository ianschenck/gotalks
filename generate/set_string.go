//go:generate ./gen.sh StringSet string set.go
package generate

// StringSet is a set of strings.
type StringSet map[string]struct{}
