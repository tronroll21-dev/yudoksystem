package helpers

import (
	"bytes"
	"fmt"
	html_template "html/template"
	"io"
	"os"
	"os/exec"
	text_template "text/template"

	"github.com/dustin/go-humanize"
)

// ParseTemplateWithFunc parses the template at the given path and sets the formatNumber function as 'format'.
func ParseTemplateWithFunc(path string) (*html_template.Template, error) {
	funcMap := html_template.FuncMap{
		"format": formatNumber,
	}
	tmpl, err := html_template.New("").Funcs(funcMap).ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func ParseSVGTemplateWithFunc(path string) (*text_template.Template, error) {
	funcMap := text_template.FuncMap{
		"format": formatNumber,
		"yOffset": func(i int, base, step float64) float64 {
			return base + float64(i)*step
		},
	}
	tmpl, err := text_template.New("").Funcs(funcMap).ParseFiles(path)
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

func GeneratePDFfromSVG(svg []byte) ([]byte, error) {
	// Write SVG to a temp file
	svgFile, err := os.CreateTemp("", "input-*.svg")
	if err != nil {
		return nil, err
	}
	defer os.Remove(svgFile.Name())

	if _, err := svgFile.Write(svg); err != nil {
		svgFile.Close()
		return nil, err
	}
	svgFile.Close()

	// Prepare a temp path for the output PDF
	pdfFile, err := os.CreateTemp("", "output-*.pdf")
	if err != nil {
		return nil, err
	}
	pdfPath := pdfFile.Name()
	pdfFile.Close()
	defer os.Remove(pdfPath)

	// Inkscape 1.2+ exports all pages by default with --export-type=pdf
	cmd := exec.Command("inkscape",
		"--export-type=pdf",
		"--export-filename="+pdfPath,
		"--export-area-page",
		svgFile.Name(),
	)

	// Capture stderr for useful error messages
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("inkscape failed: %w: %s", err, stderr.String())
	}

	pdfBytes, err := os.ReadFile(pdfPath)
	if err != nil {
		return nil, err
	}

	return pdfBytes, nil
}
