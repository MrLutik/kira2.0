package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
)

type TOMLEditor struct {
	log *logging.Logger
}

func NewTOMLEditor(logger *logging.Logger) *TOMLEditor {
	return &TOMLEditor{
		log: logger,
	}
}

// SetTomlVar updates a specific configuration value in a TOML file represented by the 'config' string.
// The function takes the 'tag', 'name', and 'value' of the configuration to update and
// returns the updated 'config' string. It ensures that the provided 'value' is correctly
// formatted in quotes if necessary and handles the update of configurations within a specific tag or section.
// The 'tag' parameter allows specifying the configuration section where the 'name' should be updated.
// If the 'tag' is empty ("") or not found, the function updates configurations in the [base] section.
func (t *TOMLEditor) SetTomlVar(config *config.TomlValue, configStr string) (string, error) {
	tag := strings.TrimSpace(config.Tag)
	name := strings.TrimSpace(config.Name)
	value := strings.TrimSpace(config.Value)

	t.log.Infof("Trying to update the ([%s] %s = %s)", tag, name, value)

	if tag != "" {
		tag = "[" + tag + "]"
	}

	lines := strings.Split(configStr, "\n")

	tagLine, nameLine, nextTagLine := -1, -1, -1

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if tag == "" && StrStartsWith(trimmedLine, name+" =") {
			t.log.Debugf("Found base config '%s' on line: %d", name, i)
			nameLine = i
			break
		}
		if tagLine == -1 && IsSubStr(line, tag) {
			t.log.Debugf("Found tag config '%s' on line: %d", tag, i)
			tagLine = i
			continue
		}

		if tagLine != -1 && nameLine == -1 && IsSubStr(line, name+" =") {
			t.log.Debugf("Found config '%s' from section '%s' on line: %d", tag, name, i)
			nameLine = i
			continue
		}

		if tagLine != -1 && nameLine != -1 && nextTagLine == -1 && IsSubStr(line, "[") && !IsSubStr(line, tag) {
			t.log.Debugf("Found next section after '%s' on line: %d", tag, i)
			nextTagLine = i
			break
		}
	}

	if nameLine == -1 || (nextTagLine != -1 && nameLine > nextTagLine) {
		return "", &ConfigurationVariableNotFoundError{
			VariableName: name,
			Tag:          tag,
		}
	}

	if IsNullOrWhitespace(value) {
		t.log.Warnf("Quotes will be added, value '%s' is empty or a seq. of whitespaces", value)
		value = fmt.Sprintf("\"%s\"", value)
	} else if StrStartsWith(value, "\"") && StrEndsWith(value, "\"") {
		t.log.Warnf("Nothing to do, quotes already present in '%q'", value)
	} else if (!StrStartsWith(value, "[")) || (!StrEndsWith(value, "]")) {
		if IsSubStr(value, " ") {
			t.log.Warnf("Quotes will be added, value '%s' contains whitespaces", value)
			value = fmt.Sprintf("\"%s\"", value)
		} else if (!IsBoolean(value)) && (!IsNumber(value)) {
			t.log.Warnf("Quotes will be added, value '%s' is neither a number nor boolean", value)
			value = fmt.Sprintf("\"%s\"", value)
		}
	}

	lines[nameLine] = name + " = " + value
	t.log.Debugf("New line is: %q", lines[nameLine])

	return strings.Join(lines, "\n"), nil
}

// IsNullOrWhitespace checks if the given string is either empty or consists of only whitespace characters.
func IsNullOrWhitespace(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

// IsBoolean checks if the given string represents a valid boolean value ("true" or "false").
func IsBoolean(input string) bool {
	_, err := strconv.ParseBool(input)
	return err == nil
}

// IsNumber checks if the given string represents a valid integer number.
func IsNumber(input string) bool {
	_, err := strconv.ParseInt(input, 0, 64)
	return err == nil
}

// StrStartsWith checks if the given string 's' starts with the specified prefix.
func StrStartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// StrEndsWith checks if the given string 's' ends with the specified suffix.
func StrEndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// IsSubStr checks if the specified substring 'substring' exists in the given string 's'.
func IsSubStr(s, substring string) bool {
	return strings.Contains(s, substring)
}
