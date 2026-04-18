package processor

import "jaytaylor.com/html2text"

func SafeHtmlToText(input string) (output string) {
	output, err := html2text.FromString(input, html2text.Options{})
	if err != nil {
		return input
	}
	return output
}
