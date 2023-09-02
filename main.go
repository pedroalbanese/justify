package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const marker = "\u0008"

func justifyText(text string, width, margin int) string {
	lineBreak := detectLineBreak(text)
	paragraphs := strings.Split(text, lineBreak+lineBreak)
	justifiedText := ""

	for _, paragraph := range paragraphs {
		// Replace initial spaces with underlines (or any other character you prefer)
		for i := 0; i < len(paragraph) && paragraph[i] == ' '; i++ {
			paragraph = paragraph[:i] + marker + paragraph[i+1:]
		}

		lines := strings.Split(paragraph, "\n")
		justifiedLines := make([]string, 0)

		for _, line := range lines {
			words := strings.Fields(line)
			currentLine := strings.Repeat(" ", margin)
			for _, word := range words {
				if len(currentLine)+len(word)+1 <= margin+width {
					currentLine += word + " "
				} else {
					justifiedLines = append(justifiedLines, currentLine)
					currentLine = strings.Repeat(" ", margin) + word + " "
				}
			}

			justifiedLines = append(justifiedLines, currentLine)
		}

		for i, line := range justifiedLines {
			if i < len(justifiedLines)-1 {
				words := strings.Fields(line)
				numWords := len(words)
				if numWords > 1 {
					extraSpaces := width - len(strings.Join(words, ""))
					spaces := extraSpaces / (numWords - 1)
					extraSpaces %= numWords - 1

					justifiedLine := strings.Repeat(" ", margin)
					for i, word := range words {
						justifiedLine += word
						if i < numWords-1 {
							justifiedLine += strings.Repeat(" ", spaces)
							if extraSpaces > 0 {
								justifiedLine += " "
								extraSpaces--
							}
						}
					}

					justifiedLines[i] = justifiedLine
				}
			}
		}

		justifiedText += strings.Replace(strings.Join(justifiedLines, "\n"), marker, " ", -1) + "\n\n"
	}

	return justifiedText
}

func detectLineBreak(text string) string {
	if strings.Contains(text, "\r\n") {
		return "\r\n" // Windows CRLF
	} else if strings.Contains(text, "\n") {
		return "\n" // Unix LF
	} else {
		return ""
	}
}

func main() {
	width := flag.Int("width", 60, "Desired line width")
	margin := flag.Int("margin", 0, "Left margin width")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	textBytes, _ := ioutil.ReadAll(reader)

	text := string(textBytes)

	// Justify the text
	justifiedText := justifyText(text, *width, *margin)

	// Display the justified text
	fmt.Print(justifiedText)
}
