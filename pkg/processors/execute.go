package processors

import (
	"binvault/pkg/env"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"text/template"

	"github.com/joho/godotenv"
)

var templates = map[string]*template.Template{}

func initializeTemplates() {
	var cmds map[string]string
	var err error
	cmds, err = godotenv.Read(env.GetVars().PROCESSOR_CONFIG_PATH)
	if err != nil {
		panic(fmt.Errorf("error reading compression config file: %s", err))
	}
	var extensions []string
	for ext, cmd := range cmds {
		templates[ext] = generateTemplate(ext, cmd)
		if ext == "default" {
			continue
		}
		extensions = append(extensions, ext)
	}
	log.Println("Templates initialized for extensions:", extensions)
	if templates["default"] == nil {
		panic(fmt.Errorf("default command not defined, 'default' command is required"))
	}
}

func execute(extension string, source string, target string) error {

	template := getTemplate(extension)

	var result bytes.Buffer
	err := template.Execute(&result, map[string]string{
		"Source": source,
		"Target": target,
	})

	if err != nil {
		return err
	}

	command := result.String()
	output, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		return fmt.Errorf("command execution failed: %s\n Output: %s", err, output)
	}
	return nil
}

func generateTemplate(extension string, command string) *template.Template {
	tmpl, err := template.New(extension).Parse(command)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func getTemplate(extension string) *template.Template {
	template := templates[extension]
	if template == nil {
		return templates["default"]
	}
	return template
}
