package models

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/denis-tingajkin/go-header/messages"
)

type Source struct {
	Author string
	Path   string
	header string
	read   bool
}

func (s *Source) Header() string {
	if s.read {
		return s.header
	}
	file, err := os.Open(s.Path)
	if err != nil {
		log.Printf("Source: can't open file: %v. %v", s.Path, messages.ErrorMsg(err))
		return ""
	}
	defer func() {
		s.read = true
		_ = file.Close()
	}()
	s.header = readHeader(file)
	return s.header
}

func (s *Source) Read() bool {
	return s.read
}

func readHeader(reader io.Reader) string {
	result := strings.Builder{}
	r := bufio.NewReader(reader)
	line, err := r.ReadString('\n')
	if cantIgnore(err) {
		return ""
	}
	if strings.HasPrefix(line, "//") {
		_, _ = result.WriteString(line)
		readSingleLineHeader(r, &result)
	}
	if strings.HasPrefix(line, "/*") {
		_, _ = result.WriteString(line)
		readMultiLineHeader(r, &result)
	}
	return strings.TrimSpace(result.String())
}

func readMultiLineHeader(r *bufio.Reader, writer io.StringWriter) {
	for {
		line, err := r.ReadString('\n')
		if cantIgnore(err) {
			return
		}
		_, _ = writer.WriteString(line)
		if strings.HasSuffix(line, "*/") {
			return
		}
		if err == io.EOF {
			return
		}
	}
}

func readSingleLineHeader(r *bufio.Reader, writer io.StringWriter) {
	for {
		line, err := r.ReadString('\n')
		if cantIgnore(err) {
			return
		}
		if !strings.HasPrefix(line, "//") {
			return
		}
		_, _ = writer.WriteString(line)
		if err == io.EOF {
			return
		}
	}
}

func cantIgnore(err error) bool {
	if err != nil && err != io.EOF {
		log.Printf("Source: can not ignore error: %v", err.Error())
		return true
	}
	return false
}
