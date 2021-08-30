package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// To Do:
	// Use static webpage to get data
	for {
		var originInput string
		var originJson map[string]interface{}
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("请输入原生JSON：")
		originInput, err := inputReader.ReadString('\n')
		if originInput=="exit"{
			break
		}
		if err != nil {
			fmt.Println("请导入单行JSON")
			continue
		}
		err = json.Unmarshal([]byte(originInput), &originJson)
		if err != nil {
			fmt.Println("\nJSON Error：" + err.Error())
			continue
		}
		fmt.Println("结果如下：")
		fmt.Println(ParseObj(originJson))
	}
	return
}

func ParseObj(origin map[string]interface{}) string {
	result:="object{"
	i:=0
	for k,v:=range origin{
		if i!=0{
			result+=","
		}
		if typeres:=TypeCheck(v);typeres=="object"{
			result+=k+"="+ParseObj(v.(map[string]interface{}))
		}else if typeres=="array"{
			result+=k+"="+ParseArray(v.([]interface{}))
		}else if typeres==""{
			fmt.Print(v)
		}else{
			result+=k+"="+typeres
		}
		i++
	}
	result+="}"
	return result
}

func ParseArray(v []interface{}) string {
	result:=""
	if len(v)==0 {
		return "[]object{}"
	}
	if typeres:=TypeCheck(v[0]);typeres=="object"{
		result="[]"+ParseObj(v[0].(map[string]interface{}))
	}else if typeres=="array"{
		result="[]"+ParseArray(v[0].([]interface{}))
	}else if typeres==""{
		fmt.Print(v)
	}else{
		result="[]"+typeres
	}
	return result
}

func TypeCheck(origin interface{}) string{
	switch origin.(type){
		case string:
			return "string"
		case int:
			return "int"
		case float32:
			return "number"
	    case float64:
			return "number"
	case bool:
			return "boolean"
	case []interface{}:
		return "array"
	case interface{}:
		return "object"
	default:
		return ""
	}
}
