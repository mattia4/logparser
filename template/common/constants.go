package common

const outputDir = "logs_output"

const html_template_path = "template/common/template.html"
const css_template_path = "template/common/template.css"
const js_template_path = "template/common/template.js"
const js_pdf_utils_path = "template/common/pdf-utils.js"

func GetHtmlTemplatePath() string { return html_template_path }
func GetCssTemplatePath() string  { return css_template_path }
func GetJsTemplatePath() string   { return js_template_path }
func GetJsPdfUtilsPath() string   { return js_pdf_utils_path }

func GetOutputDir() string { return outputDir }
