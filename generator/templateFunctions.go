package generator

import "strings"

func FistCapital(input string) string {
	output := ""
	if len(input) == 0 {
		return ""
	}
	output += strings.ToUpper(string(input[0]))
	output += strings.ToLower(input[1:])
	return output
}

func LowerCase(input string) string {
	return strings.ToLower(input)
}

func GetDTO(inputType string) string {
	inputTypeUc := strings.ToUpper(inputType)
	if _, isValidEntity := entities[inputTypeUc]; isValidEntity {
		return "*" + LowerCase(inputType) + ".DTO"
	}
	return TypeConvert(inputType)
}

func basicTypeConversion(inputType string) string {
	switch strings.ToUpper(inputType) {
	case "INT":
		return "int"
	case "BIGINT":
		return "int64"
	case "STRING":
		return "string"
	default:
		panic("type " + inputType + " can not be converted in golang datatye")
	}
}

func CustomPackage(inputType string) string {
	inputTypeUc := strings.ToUpper(inputType)

	if _, isValidEntity := entities[inputTypeUc]; isValidEntity {
		return "*" + LowerCase(inputType)
	}

	return basicTypeConversion(inputType)
}

func TypeConvert(inputType string) string {
	inputTypeUc := strings.ToUpper(inputType)

	if _, isValidEntity := entities[inputTypeUc]; isValidEntity {
		return "*" + LowerCase(inputType) + "." + FistCapital(inputType)
	}

	return basicTypeConversion(inputType)
}
