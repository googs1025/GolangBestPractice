package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// https://www.cnblogs.com/f-ck-need-u/p/10080793.html

/*
	1. Marshal()：Go数据对象 -> json数据
       func Marshal(v interface{}) ([]byte, error)
	2. UnMarshal()：Json数据 -> Go数据对象
	   func Unmarshal(data []byte, v interface{}) error

 */

/*
	Marshal()和MarshalIndent()函数可以将数据封装成json数据。

	struct、slice、array、map都可以转换成json
	struct转换成json的时候，只有字段首字母大写的才会被转换
	map转换的时候，key必须为string
	封装的时候，如果是指针，会追踪指针指向的对象进行封装
 */


type Post struct {
	ID int
	Content string
	Author  string
}

func main() {

	UseStructToJson()
	UseSliceToJson()
	UseMapToJson()
	UseStructToJson1()
	UseStructToJson2()
	UseStructToJson3()
	JsonStreamStruct()
	UseJsonStream()


}

func UseStructToJson() {
	instance := &Post{
		ID: 1,
		Content: "jianjian",
		Author: "jiang",
	}
	// 两种转换的方式
	//res, err := json.Marshal(instance)
	res, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	// Marshal()返回的是一个[]byte类型，现在变量b就已经存储了一段[]byte类型的json数据，
	fmt.Println(string(res))
}


func UseSliceToJson() {
	s := []string{"a", "b", "c"}
	d, _ := json.MarshalIndent(s, "", "\t")
	fmt.Println(string(d))

}

func UseMapToJson() {
	dict := map[string]string{
		"a": "aaaa",
		"b": "bbbb",
		"C": "cccc",
	}
	d, _ := json.MarshalIndent(dict, "", "\t")
	fmt.Println(string(d))

}

//////////////////////////////////////////////////////////

type Post1 struct {
	ID int	`json:"ID"`
	Content string	`json:"content"`
	Author  string	`json:"author"`
	Label   []string	`json:"label"`
}

// struct能被转换的字段都是首字母大写的字段，但如果想要在json中使用小写字母开头的key，可以使用struct的tag来辅助反射。
func UseStructToJson1() {
	postp := &Post1{
		2,
		"Hello World",
		"userB",
		[]string{"linux", "shell"},
	}

	p, _ := json.MarshalIndent(postp, "", "\t")
	fmt.Println(string(p))
}


///////////////////////////////////////////////////////////////////

type Post3 struct {
	ID        int64         `json:"id"`
	Content   string        `json:"content"`
	Author    Author        `json:"author"`
	Published bool          `json:"published"`
	Label     []string      `json:"label"`
	NextPost  *Post         `json:"nextPost"`
	Comments  []*Comment    `json:"comments"`
}

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func UseStructToJson2() {
	// 打开json文件
	fh, err := os.Open("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/json-handle-learning/a.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}

	var post Post3
	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}

func UseStructToJson3() {
	// 读取json文件
	fh, err := os.Open("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/json-handle-learning/a.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 定义空接口接收解析后的json数据
	var unknown interface{}
	// 或：map[string]interface{} 结果是完全一样的
	err = json.Unmarshal(jsonData, &unknown)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(unknown)

	// 进行断言，并switch匹配
	m := unknown.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "type: string\nvalue: ", vv)
			fmt.Println("------------------")
		case float64:
			fmt.Println(k, "type: float64\nvalue: ", vv)
			fmt.Println("------------------")
		case bool:
			fmt.Println(k, "type: bool\nvalue: ", vv)
			fmt.Println("------------------")
		case map[string]interface{}:
			fmt.Println(k, "type: map[string]interface{}\nvalue: ", vv)
			for i, j := range vv {
				fmt.Println(i,": ",j)
			}
			fmt.Println("------------------")
		case []interface{}:
			fmt.Println(k, "type: []interface{}\nvalue: ", vv)
			for key, value := range vv {
				fmt.Println(key, ": ", value)
			}
			fmt.Println("------------------")
		default:
			fmt.Println(k, "type: nil\nvalue: ", vv)
			fmt.Println("------------------")
		}
	}


}


func JsonStreamStruct() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func UseJsonStream() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}