package results

import (
	"bufio"
	"fmt"
	"io"
	"os"

	errors "github.com/apenella/go-common-utils/error"
)

const (
	// PrefixTokenSeparator is and string printed between prefix and ansible output
	PrefixTokenSeparator = "\u2500\u2500"
)

// DefaultStdoutCallbackResults is the default method to print ansible-playbook results
func DefaultStdoutCallbackResults(prefix string, r io.Reader, w io.Writer) error {

	if r == nil {
		return errors.New("(results::DefaultStdoutCallbackResults)", "Reader is not defined")
	}

	if w == nil {
		w = os.Stdout
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintf(w, "%s %s %s\n", prefix, PrefixTokenSeparator, scanner.Text())
	}

	return nil
}
