package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func upperFirst(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func replaceJsonType(str string) string {
	if str == "String" {
		return "string"
	}
	return str
}

func GenMessageElement(jsonStr, dir string) error {
	data := []byte(jsonStr)
	if len(data) == 0 || data[0] == 'n' { // copied from the Q, can be improved
		return nil
	}
	messageType := gjson.GetBytes(data, "type").Str
	messageData := gjson.GetBytes(data, "data")
	structName := "MessageElement" + upperFirst(messageType)
	structString := fmt.Sprintf("type %v struct{", structName)
	messageData.ForEach(func(key, value gjson.Result) bool {
		propType := replaceJsonType(value.Type.String())
		var propKey string
		if key.Str == "type" {
			propKey = fmt.Sprintf(messageType + "Type")
		} else {
			propKey = key.Str
		}
		propStr := fmt.Sprintf("\n\t%v %v `json:\"%v\"`", upperFirst(propKey), propType, key)
		structString += propStr
		return true
	})
	structString += "\n}"
	structTypeMethod := fmt.Sprintf(
		"func (msg %v) Type() string {\n\treturn \"%v\"\n}",
		structName,
		messageType,
	)
	elementContent := fmt.Sprintf(
		"package model\n\nimport \"encoding/json\"\n\n%v\n\n%v",
		structString,
		structTypeMethod,
	) + fmt.Sprintf(
		"\n\nfunc init(){\n\tunmarshalJSONMap[\"%v\"] = func(data []byte) (MessageElement, error) {"+
			"\n\tvar result %v"+
			"\n\terr := json.Unmarshal(data, &result)"+
			"\n\treturn result, err"+
			"\n\t}"+
			"\n}",
		messageType,
		structName,
	)
	elementFileName := dir + "/message_element_" + messageType + ".go"
	err := writeToFile(elementContent, elementFileName)
	if err != nil {
		return err
	}
	//test文件
	testFileName := dir + "/message_element_" + messageType + "_test.go"
	testContent := "package model\n\n" +
		"import (\n" +
		"\t\"encoding/json\"\n" +
		"\t\"errors\"\n" +
		"\t\"testing\"\n" +
		")\n\n" +
		"func Test" + upperFirst(messageType) + "(t *testing.T) {\n" +
		"\tdata := []byte(\"[" + strings.ReplaceAll(jsonStr, "\"", "\\\"") + "]\")\n" +
		"\tvar msg []MessageSegment\n" +
		"\terr := json.Unmarshal(data, &msg)\n" +
		"\tif err != nil {\n" +
		"\t\tpanic(err)\n" +
		"\t}\n" +
		"\tif msg[0].Data.Type() != \"" + messageType + "\" {\n" +
		"\t\tpanic(errors.New(\"类型错误\"))\n" +
		"\t}\n" +
		"\t_, ok := msg[0].Data.(" + structName + ")\n" +
		"\tif !ok {\n" +
		"\t\tpanic(errors.New(\"类型错误\"))\n" +
		"\t}\n" +
		"}"
	err = writeToFile(testContent, testFileName)
	if err != nil {
		return err
	}
	return nil
}

func writeToFile(content, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		return err
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		return err
	}
	return nil
}
