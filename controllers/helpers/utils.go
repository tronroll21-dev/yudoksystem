package helpers

import (
	"fmt"
	"html/template"
	"io"
	"os/exec"

	"github.com/dustin/go-humanize"
)

// ParseTemplateWithFunc parses the template at the given path and sets the formatNumber function as 'format'.
func ParseTemplateWithFunc(path string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"format": formatNumber,
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// formatNumber formats an integer with thousands separators
// func formatNumber(n int) string {
// 	// You can use fmt or a more advanced library if needed
// 	return template.HTMLEscapeString(fmt.Sprintf("%d", n))
// }

func formatNumber(n interface{}) string {
	switch v := n.(type) {
	case int:
		return humanize.Comma(int64(v))
	case int64:
		return humanize.Comma(v)
	case float64:
		return humanize.Comma(int64(v))
	default:
		return fmt.Sprintf("%v", n)
	}
}

func GeneratePDFfromHTML(html []byte) ([]byte, error) {
	// Dummy implementation for compilation; replace with real logic.

	// Use wkhtmltopdf to convert HTML to PDF
	cmd := exec.Command("wkhtmltopdf", "-q", "-", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to create stdin pipe")
		return nil, err
	}
	defer stdin.Close()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to create stdout pipe")
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		//c.String(http.StatusInternalServerError, "Failed to start wkhtmltopdf")
		return nil, err
	}
	_, err = stdin.Write(html)
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to write HTML to wkhtmltopdf")
		return nil, err
	}
	stdin.Close()
	pdfBytes, err := io.ReadAll(stdout)
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to read PDF output")
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		//c.String(http.StatusInternalServerError, "wkhtmltopdf failed: "+err.Error())
		return nil, err
	}
	return pdfBytes, nil
}
