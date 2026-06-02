package prmaven

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReportSchemaTracksJSONContractFields(t *testing.T) {
	schema := loadReportSchema(t)

	assertSchemaProperties(t, schema, jsonFields(reflect.TypeOf(Report{})), "properties")
	assertSchemaRequired(t, schema, requiredJSONFields(reflect.TypeOf(Report{})), "required")
	defs := schemaObject(t, schema, "$defs")
	assertSchemaProperties(t, defs, jsonFields(reflect.TypeOf(Summary{})), "summary.properties")
	assertSchemaRequired(t, defs, requiredJSONFields(reflect.TypeOf(Summary{})), "summary.required")
	assertSchemaProperties(t, defs, jsonFields(reflect.TypeOf(Module{})), "module.properties")
	assertSchemaRequired(t, defs, requiredJSONFields(reflect.TypeOf(Module{})), "module.required")
	assertSchemaProperties(t, defs, jsonFields(reflect.TypeOf(Finding{})), "finding.properties")
	assertSchemaRequired(t, defs, requiredJSONFields(reflect.TypeOf(Finding{})), "finding.required")
}

func TestGeneratedJSONReportsValidateAgainstSchema(t *testing.T) {
	schema := loadReportSchema(t)

	tests := []struct {
		name       string
		projectDir string
	}{
		{name: "multi-module failure demo", projectDir: "../../demo/multi-module-failure"},
		{name: "no-failure demo", projectDir: "../../demo/no-failure"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report, err := Analyze(Options{ProjectDir: tt.projectDir})
			if err != nil {
				t.Fatal(err)
			}

			var output bytes.Buffer
			if err := WriteJSON(&output, report); err != nil {
				t.Fatal(err)
			}

			var generated any
			if err := json.Unmarshal(output.Bytes(), &generated); err != nil {
				t.Fatalf("generated JSON is not parseable: %v", err)
			}

			validateSchemaValue(t, schema, schema, generated, "$")
		})
	}
}

func loadReportSchema(t *testing.T) map[string]any {
	t.Helper()

	data, err := os.ReadFile("../../schema/prmaven-report.schema.json")
	if err != nil {
		t.Fatal(err)
	}

	var schema map[string]any
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatal(err)
	}
	return schema
}

func assertSchemaProperties(t *testing.T, schema map[string]any, expected []string, path string) {
	t.Helper()

	properties := schemaMapAtPath(t, schema, path)
	for _, field := range expected {
		if _, ok := properties[field]; !ok {
			t.Fatalf("%s missing JSON field %q", path, field)
		}
	}
}

func assertSchemaRequired(t *testing.T, schema map[string]any, expected []string, path string) {
	t.Helper()

	requiredValue := schemaObjectAtPath(t, schema, path)
	required, ok := requiredValue.([]any)
	if !ok {
		t.Fatalf("%s is %T, want []any", path, requiredValue)
	}

	got := map[string]bool{}
	for _, value := range required {
		name, ok := value.(string)
		if !ok {
			t.Fatalf("%s contains %T, want string", path, value)
		}
		got[name] = true
	}
	for _, field := range expected {
		if !got[field] {
			t.Fatalf("%s missing required JSON field %q", path, field)
		}
	}
}

func schemaObjectAtPath(t *testing.T, schema map[string]any, path string) any {
	t.Helper()

	parts := strings.Split(path, ".")
	var current any = schema
	for _, part := range parts {
		object, ok := current.(map[string]any)
		if !ok {
			t.Fatalf("%s parent is %T, want object", path, current)
		}
		current = object[part]
	}
	return current
}

func schemaMapAtPath(t *testing.T, schema map[string]any, path string) map[string]any {
	t.Helper()

	value := schemaObjectAtPath(t, schema, path)
	object, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("%s is %T, want object", path, value)
	}
	return object
}

func schemaObject(t *testing.T, schema map[string]any, key string) map[string]any {
	t.Helper()

	value, ok := schema[key]
	if !ok {
		t.Fatalf("schema missing key %q", key)
	}
	object, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("schema key %q is %T, want object", key, value)
	}
	return object
}

func jsonFields(typ reflect.Type) []string {
	fields := make([]string, 0, typ.NumField())
	for i := range typ.NumField() {
		if name := jsonFieldName(typ.Field(i)); name != "" {
			fields = append(fields, name)
		}
	}
	return fields
}

func requiredJSONFields(typ reflect.Type) []string {
	fields := make([]string, 0, typ.NumField())
	for i := range typ.NumField() {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if strings.Contains(tag, "omitempty") {
			continue
		}
		if name := jsonFieldName(field); name != "" {
			fields = append(fields, name)
		}
	}
	return fields
}

func jsonFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" || tag == "-" {
		return ""
	}
	return strings.Split(tag, ",")[0]
}

func validateSchemaValue(t *testing.T, root map[string]any, schema map[string]any, value any, path string) {
	t.Helper()

	if ref, ok := schema["$ref"].(string); ok {
		validateSchemaValue(t, root, resolveSchemaRef(t, root, ref), value, path)
		return
	}

	if typ, ok := schema["type"].(string); ok {
		validateSchemaType(t, typ, value, path)
	}

	if enumValues, ok := schema["enum"].([]any); ok {
		validateSchemaEnum(t, enumValues, value, path)
	}

	switch schema["type"] {
	case "object":
		validateSchemaObject(t, root, schema, value, path)
	case "array":
		validateSchemaArray(t, root, schema, value, path)
	case "integer":
		validateSchemaMinimum(t, schema, value, path)
	}
}

func validateSchemaType(t *testing.T, typ string, value any, path string) {
	t.Helper()

	ok := false
	switch typ {
	case "object":
		_, ok = value.(map[string]any)
	case "array":
		_, ok = value.([]any)
	case "string":
		_, ok = value.(string)
	case "integer":
		number, isNumber := value.(float64)
		ok = isNumber && number == math.Trunc(number)
	default:
		t.Fatalf("%s uses unsupported schema type %q in test validator", path, typ)
	}
	if !ok {
		t.Fatalf("%s is %T, want schema type %q", path, value, typ)
	}
}

func validateSchemaEnum(t *testing.T, enumValues []any, value any, path string) {
	t.Helper()

	for _, enumValue := range enumValues {
		if enumValue == value {
			return
		}
	}
	t.Fatalf("%s value %q is not allowed by enum %v", path, value, enumValues)
}

func validateSchemaObject(t *testing.T, root map[string]any, schema map[string]any, value any, path string) {
	t.Helper()

	object, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("%s is %T, want object", path, value)
	}

	for _, name := range schemaStringArray(t, schema, "required", path) {
		if _, ok := object[name]; !ok {
			t.Fatalf("%s missing required property %q", path, name)
		}
	}

	propertiesValue, ok := schema["properties"]
	if !ok {
		return
	}
	properties, ok := propertiesValue.(map[string]any)
	if !ok {
		t.Fatalf("%s.properties is %T, want object", path, propertiesValue)
	}

	for name, childSchemaValue := range properties {
		childValue, ok := object[name]
		if !ok {
			continue
		}
		childSchema, ok := childSchemaValue.(map[string]any)
		if !ok {
			t.Fatalf("%s.properties.%s is %T, want object", path, name, childSchemaValue)
		}
		validateSchemaValue(t, root, childSchema, childValue, path+"."+name)
	}
}

func validateSchemaArray(t *testing.T, root map[string]any, schema map[string]any, value any, path string) {
	t.Helper()

	values, ok := value.([]any)
	if !ok {
		t.Fatalf("%s is %T, want array", path, value)
	}

	itemsValue, ok := schema["items"]
	if !ok {
		return
	}
	itemsSchema, ok := itemsValue.(map[string]any)
	if !ok {
		t.Fatalf("%s.items is %T, want object", path, itemsValue)
	}

	for i, childValue := range values {
		validateSchemaValue(t, root, itemsSchema, childValue, fmt.Sprintf("%s[%d]", path, i))
	}
}

func validateSchemaMinimum(t *testing.T, schema map[string]any, value any, path string) {
	t.Helper()

	minimum, ok := schema["minimum"].(float64)
	if !ok {
		return
	}
	number, ok := value.(float64)
	if !ok {
		t.Fatalf("%s is %T, want number", path, value)
	}
	if number < minimum {
		t.Fatalf("%s value %v is below schema minimum %v", path, number, minimum)
	}
}

func resolveSchemaRef(t *testing.T, root map[string]any, ref string) map[string]any {
	t.Helper()

	if !strings.HasPrefix(ref, "#/") {
		t.Fatalf("unsupported schema ref %q", ref)
	}
	current := any(root)
	for _, part := range strings.Split(strings.TrimPrefix(ref, "#/"), "/") {
		object, ok := current.(map[string]any)
		if !ok {
			t.Fatalf("schema ref %q parent is %T, want object", ref, current)
		}
		current, ok = object[part]
		if !ok {
			t.Fatalf("schema ref %q missing part %q", ref, part)
		}
	}
	object, ok := current.(map[string]any)
	if !ok {
		t.Fatalf("schema ref %q resolves to %T, want object", ref, current)
	}
	return object
}

func schemaStringArray(t *testing.T, schema map[string]any, key, path string) []string {
	t.Helper()

	value, ok := schema[key]
	if !ok {
		return nil
	}
	values, ok := value.([]any)
	if !ok {
		t.Fatalf("%s.%s is %T, want array", path, key, value)
	}
	result := make([]string, 0, len(values))
	for _, item := range values {
		name, ok := item.(string)
		if !ok {
			t.Fatalf("%s.%s contains %T, want string", path, key, item)
		}
		result = append(result, name)
	}
	return result
}
