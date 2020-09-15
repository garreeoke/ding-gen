package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var appName = flag.String("appName", "new application", "Application Name")
var pipelineName = flag.String("pipelineName", "new pipeline", "Pipeline Name")
var inputFile = flag.String("inputFile", "pipeline.json", "File to read in")
var pipelineFileName = flag.String("pipelineFileName", "dinghyFile", "Pipeline file to create")
var moduleFolder = flag.String("moduleFolder", "local_modules", "Folder to place modules in")

func main() {
	flag.Parse()
	appName := *appName
	//pipelineName := *pipelineName
	appFile := *inputFile
	log.Println("Building dinghyfile for appName: ", appName)
	app := Application{
		AppName: appName,
		Globals: map[string]string{},
	}
	// Read the inputFile ... future will pull directly from spin
	err := app.Process(appFile)
	if err != nil {
		log.Fatalln("Cannot read inputFile: ", err)
	}
	// Move out of this function
	err = WriteFile(*pipelineFileName,"pipeline",app)
	if err != nil {
		log.Fatalln("error writing file", err)
	}
}

// Process ... starts processing to build the pipelineName
func (a *Application) Process(input string) error {
	flag.Parse()
	log.Println("Reading inputFile: ", input)
	pipeline := Pipeline{}
	file, err := ioutil.ReadFile(input)
	if err != nil {
		log.Printf("Error reading pipelineName json (%v): %v ", input, err)
		return err
	}
	err = json.Unmarshal(file, &pipeline)
	if err != nil {
		log.Printf("Error creating pipelineName struct (%v): %v ", input, err)
		return err
	}

	// Process all the sections
	err = pipeline.ParameterCfg()
	if err != nil {
		log.Println("Error creating parameters", err)
		return err
	}

	err = pipeline.TriggerCfg()
	if err != nil {
		log.Println("Error creating triggers", err)
		return err
	}

	pipeline.Name = *pipelineName
	pipeline.Application = a.AppName
	a.Pipelines = append(a.Pipelines, pipeline)
	return nil
}

// Trigger
func (p *Pipeline) TriggerCfg() error {

	log.Println("Processing triggers")
	for i,input := range p.Triggers {
		t,err := json.Marshal(input)
		if err != nil {
			log.Println("error marshaling triggers ", err)
			return err
		}
		genericTrigger := TriggerGeneric{}
		err = json.Unmarshal(t, &genericTrigger)
		if err != nil {
			log.Println("error unmarshalling to generic trigger ", err)
			return err
		}
		switch genericTrigger.Type {
		case "docker":
			trigger := TriggerDocker{}
			err = json.Unmarshal(t, &trigger)
			if err != nil {
				log.Println("Error unmarshalling trigger to DockerTrigger", err)
				return err
			}
			triggerString,err := trigger.Build()
			if err != nil {
				log.Println("error building triggers", err)
				return err
			}
			p.Triggers[i] = triggerString
		}
	}
	return nil
}

// Build ... docker trigger builder
func (td *TriggerDocker) Build() (string,error) {
	flag.Parse()
	log.Println("Building docker trigger")
	modTd := *td
	modTd.Account = "{{ var \"dockerAccount\" }}"
	modTd.Enabled = "{{ var \"triggerEnabled\" }}"
	modTd.ExpectedArtifactIds = []string{}
	modTd.Organization = "{{ var \"dockerOrg\" }}"
	modTd.Registry = "{{ var \"dockerRegistry\" }}"
	modTd.Repository = "{{ var \"dockerRepo\" }}"
	modTd.Tag = "{{ var \"dockerTag\" }}"
	fileName := *moduleFolder + "/docker.trigger.module"
	err := WriteFile(fileName,"module",modTd)
	if err != nil {
		log.Println("error writing docker module file", err)
		return "",err
	}

	pmls := []string{"{{ module \"" + fileName + "\"" }
	pmls = append(pmls, "\"dockerAccount\" " + "\"" +  td.Account + "\"")
	pmls = append(pmls, "\"enabled\" " + strconv.FormatBool(td.Enabled.(bool)))
	pmls = append(pmls, "\"dockerOrg\"" + "\"" + td.Organization + "\"")
	pmls = append(pmls, "\"dockerRegistry\"" + "\"" + td.Registry + "\"")
	pmls = append(pmls, "\"dockerRepo\"" + "\"" + td.Repository + "\"")
	pmls = append(pmls, "\"dockerTag\"" + "\"" + td.Tag + "\"")
	pmls = append(pmls, "}}")
	pml := strings.Join(pmls, " ")
	return pml,nil
}

// ParameterCfg
func (p *Pipeline) ParameterCfg() error {
	flag.Parse()
	log.Println("Building ParameterCFG")
	// Create module inputFile first
	module := ParameterConfig{}
	module.Default = "{{ var \"paramDefault\" ?: \"None\" }}"
	module.Description = "{{ var \"paramDescription\" }}"
	module.HasOptions = "{{ var \"paramHasOptions\" }}"
	module.Label = "{{ var \"paramLabel\" }}"
	module.Name = "{{ var \"paramName\" }}"
	module.Options = []string{}
	module.Pinned = "{{ var \"paramPinned\" }}"
	module.Required = "{{ var \"Required\" }}"
	err := WriteFile(*moduleFolder + "/param_config.module","module", module)
	if err != nil {
		log.Println(err)
		return err
	}
	for i,e := range p.ParameterConfigs {
		m, err := json.Marshal(e)
		if err != nil {
			log.Println("Error marshalling parameter config: ", err)
			return err
		}
		x := ParameterConfig{}
		err = json.Unmarshal(m, &x)
		if err != nil {
			log.Println("Error unmarshalling parameter config: ", err)
			return err
		}
		pmls := []string{"{{ module \"local_modules/param_config.module\""}
		pmls = append(pmls, "\"paramName\" " + "\"" + x.Name + "\"")
		pmls = append(pmls, "\"paramDefaultValue\" " + "\"" + x.Default + "\"")
		pmls = append(pmls, "\"paramDescription\" " + "\"" + x.Description + "\"")
		pmls = append(pmls, "\"paramHasOptions\" " + strconv.FormatBool(x.HasOptions.(bool)))
		pmls = append(pmls, "\"paramLabel\" " + "\"" + x.Label + "\"")
		pmls = append(pmls, "\"paramPinned\" " + strconv.FormatBool(x.Pinned.(bool)))
		pmls = append(pmls, "\"paramRequired\" " + strconv.FormatBool(x.Required.(bool)))
		pmls = append(pmls, "}}")
		pml := strings.Join(pmls, " ")
		p.ParameterConfigs[i] = pml
	}

	return nil
}

// FormatModule ...
func FormatModule(mod interface{}) []byte {

	a, err := json.Marshal(mod)
	if err != nil {
		log.Println("error marshalling module ", err)
	}
	b := string(a)
	c := []byte(strings.Replace(b, `\`,"",-1))
	return c
}

// FormatPipeline ... cleanup module lines in the dinghy inputFile
func FormatPipeline(object interface{}) []byte {

	a,err := json.Marshal(object)
	if err != nil {
		log.Println(err)
	}

	b := string(a)
	c := strings.Replace(b, "\"{{", "{{", -1)
	d := strings.Replace(c, "}}\"", "}}", -1)
	e := []byte(strings.Replace(d, `\`,"",-1))
	return e
}

func WriteFile(filename,filetype string, i interface{}) error {

	var err error
	switch filetype {
	case "module":
		err = ioutil.WriteFile(filename, FormatModule(i), 0644)
	case "pipeline":
		err = ioutil.WriteFile(filename, FormatPipeline(i), 0644)
	}
	if err != nil {
		log.Println("error writing file", filename, err)
	}
	return nil
}