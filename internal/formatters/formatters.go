package formatters

func FormatMessage(tree map[string]map[string]any, format string) string {
	var message string
	switch format {
	case "":
		message = FormatterStylish(tree)
	case "stylish":
		message = FormatterStylish(tree)
	case "plain":
		message = FormmaterPlain(tree)
	case "json":
		message = FormmaterJson(tree)
	default:
		message = "unknown format"
	}
	return message
}
