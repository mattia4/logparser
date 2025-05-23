package template

const html_template_path = "template/template.html"
const css_template_path = "template/template.css"
const js_template_path = "template/template.js"
const outputDir = "logs_output"

func GetHtmlTemplatePath() string { return html_template_path }
func GetCssTemplatePath() string  { return css_template_path }
func GetJsTemplatePath() string   { return js_template_path }
func GetOutputDir() string        { return outputDir }
