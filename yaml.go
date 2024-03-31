package yamlfile

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// ReadYAML updates the ioutil.ReadFile call to os.ReadFile
func ReadYAML(filePath string, out interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, out)
	return err
}

// WriteYAML updates the ioutil.WriteFile call to os.WriteFile
func WriteYAML(filePath string, data interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, yamlData, 0644)
}

func YamlDemo() (err error) {
	// 读取 YAML 文件到 Config 结构体
	var cfg ConfigDemo
	err = ReadYAML("./yamlfile/cfgdemo.yaml", &cfg)
	if err != nil {
		return
	}
	fmt.Println(cfg)

	// 修改 cfg 或添加新的数据...

	// 将修改后的 Config 结构体写回到新的 YAML 文件
	err = WriteYAML("./yamlfile/updated_config.yaml", cfg)
	if err != nil {
		return
	}

	// 将 map 写入 YAML 文件
	dataMap := map[string]interface{}{
		"key":    "value",
		"number": 5,
	}
	err = WriteYAML("./yamlfile/map_data.yaml", dataMap)

	return
}
