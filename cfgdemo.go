package yamlfile

type ConfigDemo struct {
	Title string `yaml:"title"`
	Owner struct {
		Name string `yaml:"name"`
	} `yaml:"owner"`
}
