package input

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/vaiktorg/portfolio/src/ui/models"

	"gopkg.in/yaml.v2"
)

type YamlInput struct {
	once  sync.Once
	Posts map[string]models.PostCard `yaml:"posts"`
}

func (y *YamlInput) Init() {
	file, err := ioutil.ReadFile("web/posts.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(file, y)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(y.Posts)
}
