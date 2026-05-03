package formatters

func FormatMessage(tree map[string]map[string]any, format string) string {
	var message string
	if format == "" || format == "stylish" {
		message = FormatterStylish(tree)
	}
	if format == "plain" {
		message = FormmaterPlain(tree)
	}
	return message
}
