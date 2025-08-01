package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

const marker = "\u0008"

func justifyText(text string, width, margin int) string {
	lineBreak := detectLineBreak(text)
	paragraphs := strings.Split(text, lineBreak+lineBreak)
	justifiedText := ""

	for i, paragraph := range paragraphs {
		// Marcar espaços iniciais com um marcador invisível
		for j := 0; j < len(paragraph) && paragraph[j] == ' '; j++ {
			paragraph = paragraph[:j] + marker + paragraph[j+1:]
		}

		lines := strings.Split(paragraph, "\n")
		justifiedLines := make([]string, 0)

		for _, line := range lines {
			words := strings.Fields(line)
			currentLine := strings.Repeat(" ", margin)
			currentLen := utf8.RuneCountInString(currentLine)

			for _, word := range words {
				wordLen := utf8.RuneCountInString(word)
				extraSpace := 0
				if currentLen > margin {
					extraSpace = 1
				}
				if currentLen+wordLen+extraSpace <= margin+width {
					if currentLen > margin {
						currentLine += " "
						currentLen++
					}
					currentLine += word
					currentLen += wordLen
				} else {
					justifiedLines = append(justifiedLines, currentLine)
					currentLine = strings.Repeat(" ", margin) + word
					currentLen = margin + wordLen
				}
			}

			justifiedLines = append(justifiedLines, currentLine)
		}

		// Justifica todas as linhas exceto a última do parágrafo
		for j, line := range justifiedLines {
			if j < len(justifiedLines)-1 {
				words := strings.Fields(line)
				numWords := len(words)
				if numWords > 1 {
					totalChars := 0
					for _, word := range words {
						totalChars += utf8.RuneCountInString(word)
					}
					extraSpaces := width - totalChars
					spaceBetween := extraSpaces / (numWords - 1)
					remaining := extraSpaces % (numWords - 1)

					justifiedLine := strings.Repeat(" ", margin)
					for k, word := range words {
						justifiedLine += word
						if k < numWords-1 {
							spaces := spaceBetween
							if remaining > 0 {
								spaces++
								remaining--
							}
							justifiedLine += strings.Repeat(" ", spaces)
						}
					}

					justifiedLines[j] = justifiedLine
				}
			}
		}

		// Monta o parágrafo justificado
		justifiedParagraph := strings.Replace(strings.Join(justifiedLines, "\n"), marker, " ", -1)
		justifiedText += justifiedParagraph

		// Adiciona quebra dupla apenas entre parágrafos
		if i < len(paragraphs)-1 {
			justifiedText += "\n\n"
		}
	}

	return justifiedText
}

func detectLineBreak(text string) string {
	if strings.Contains(text, "\r\n") {
		return "\r\n"
	}
	return "\n"
}

func main() {
	width := flag.Int("width", 60, "Desired line width")
	margin := flag.Int("margin", 0, "Left margin width")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	textBytes, _ := ioutil.ReadAll(reader)
	text := string(textBytes)

	justified := justifyText(text, *width, *margin)
	fmt.Print(justified)
}
