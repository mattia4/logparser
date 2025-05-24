package utils

import (
	"fmt"
	"io"
	android_logcat_template "logparser/template/android-logcat"
	op_log_template "logparser/template/op-log"
	"os"
	"time"
)

func CopyFile(src, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func HandleTemplatePath(FormatTag string) (string, string, string) {
	log_html_template_path := ""
	log_css_template_path := ""
	log_js_template_path := ""

	switch FormatTag {
	case "AndroidLogcat":
		log_html_template_path = android_logcat_template.GetHtmlTemplatePath()
		log_css_template_path = android_logcat_template.GetCssTemplatePath()
		log_js_template_path = android_logcat_template.GetJsTemplatePath()
	case "OpLog":
		log_html_template_path = op_log_template.GetHtmlTemplatePath()
		log_css_template_path = op_log_template.GetCssTemplatePath()
		log_js_template_path = op_log_template.GetJsTemplatePath()
	default:
		log_html_template_path = android_logcat_template.GetHtmlTemplatePath()
		log_css_template_path = android_logcat_template.GetCssTemplatePath()
		log_js_template_path = android_logcat_template.GetJsTemplatePath()
	}

	return log_html_template_path, log_css_template_path, log_js_template_path
}

func ParseApacheDate(s string) string {
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", s)
	if err != nil {
		fmt.Println("Errore parsing:", err)
		return ""
	}

	date := t.Format("02-01-2006")

	return date
}

func ParseApacheTime(s string) string {
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", s)
	if err != nil {
		fmt.Println("Errore parsing:", err)
		return ""
	}

	hour := fmt.Sprintf("%02d", t.Hour())
	minute := fmt.Sprintf("%02d", t.Minute())
	second := fmt.Sprintf("%02d", t.Second())

	time := hour + ":" + minute + ":" + second

	return time
}
