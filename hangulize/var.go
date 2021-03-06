package hangulize

import (
	"regexp"
	"strings"
)

// <...>
//  └─┴─ (1)
var reVar = re(`<(.+?)>`)

// expandVars replaces <var> to corresponding content regexp such as (a|b|c).
func expandVars(expr string, vars map[string][]string) (string, [][]string) {
	var usedVars [][]string

	expr = reVar.ReplaceAllStringFunc(expr, func(varExpr string) string {
		// Retrieve variable name and values.
		name, vals := getVar(varExpr, vars)

		usedVars = append(usedVars, vals)

		// Build as RegExp like /(a|b|c)/.
		escapedVals := make([]string, len(vals))
		for i, val := range vals {
			escapedVals[i] = regexp.QuoteMeta(val)
		}

		// return `(?P<` + name + `>` + strings.Join(escapedVals, `|`) + `)`
		_ = name
		return `(` + strings.Join(escapedVals, `|`) + `)`
	})

	return expr, usedVars
}

// getVar parses a var expression, which looks like "<var>",
// and returns the var name and values.
func getVar(expr string, vars map[string][]string) (string, []string) {
	name := strings.Trim(expr, `<>`)
	vals := vars[name]

	return name, vals
}
