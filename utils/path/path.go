package path

import (
	android_logcat_template "logparser/template/android-logcat"
	op_log_template "logparser/template/op-log"
)

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
