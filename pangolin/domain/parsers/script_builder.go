package parsers

import "errors"

type scriptBuilder struct {
	values []ScriptValue
}

func createScriptBuilder() ScriptBuilder {
	out := scriptBuilder{
		values: []ScriptValue{},
	}

	return &out
}

// Create initializes the builder
func (app *scriptBuilder) Create() ScriptBuilder {
	return createScriptBuilder()
}

// WithValues add values to the builder
func (app *scriptBuilder) WithValues(values []ScriptValue) ScriptBuilder {
	app.values = values
	return app
}

// Now builds a new Script instance
func (app *scriptBuilder) Now() (Script, error) {
	if app.values == nil {
		app.values = []ScriptValue{}
	}

	name := ""
	version := ""
	var script RelativePath
	var language RelativePath
	for _, oneValue := range app.values {
		if oneValue.IsName() {
			name = oneValue.Name()
			continue
		}

		if oneValue.IsVersion() {
			version = oneValue.Version()
			continue
		}

		if oneValue.IsScript() {
			script = oneValue.Script()
			continue
		}

		if oneValue.IsLanguage() {
			language = oneValue.Language()
			continue
		}
	}

	if name == "" {
		return nil, errors.New("the name is mandatory in order to build a Script instance")
	}

	if version == "" {
		return nil, errors.New("the version is mandatory in order to build a Script instance")
	}

	if script == nil {
		return nil, errors.New("the script is mandatory in order to build a Script instance")
	}

	if language == nil {
		return nil, errors.New("the language is mandatory in order to build a Script instance")
	}

	return createScript(name, version, script, language), nil
}
