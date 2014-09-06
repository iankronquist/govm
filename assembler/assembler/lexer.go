package assembler

import "strings"

/*
func fileChunker(fileName string, tokens chan string) {
	out := make(chan string)
	file, err := os.Open(fileName)
	if err != nil {
		close(out)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		numRead, err := file.Read(buf)
		if err != nil {
			close(out)
		}
		go SplitOnWhitespace(string(buf), out)
		tokens <- out
	}
}
*/

func Tokenize(chunk string, out chan string) {
	chunk = strings.Replace(chunk, "\t", " ", -1)
	lines := strings.Split(chunk, "\n")
	for line := range lines {
		tokens := strings.Split(lines[line], " ")
		for token := range tokens {
			if tokens[token] == "" {
				continue
			}
			out <- tokens[token]
		}
	}
	close(out)
}

