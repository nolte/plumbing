package pkg

import (
	"encoding/json"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParsingVersionFromOutput(output string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(output), &data)
	if err == nil {
		return data["terraform_version"].(string), nil
	} else {
		return strings.Replace(strings.Split(output, "\n")[0], "Terraform v", "", -1), nil
	}

}
