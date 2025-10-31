package compressor

import (
    "fmt"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/gobackup/gobackup/config"
	"github.com/longbridgeapp/assert"
)

type Monkey struct {
	Base
}

func (c Monkey) perform() (archivePath string, err error) {
	result := "aaa"
	return result, nil
}

func TestBase_archiveFilePath(t *testing.T) {
	viper := viper.New()
	viper.Set("format", "2006.01.02.15.04.05")
	base := newBase(config.ModelConfig{
		CompressWith: config.SubConfig{
			Type:  "compress_with",
			Name:  "tar",
			Viper: viper,
		},
	},
	prefixPath := path.Join(base.model.TempPath, time.Now().Format("2006.01.02.15.04"))
// 	filepath := base.archiveFilePath(".tar")
// 	fmt.Println("GoBackup starting...")
	assert.True(t, strings.HasPrefix(base.archiveFilePath(".tar"), prefixPath))
	assert.True(t, strings.HasSuffix(base.archiveFilePath(".tar"), ".tar"))
}

func TestBaseInterface(t *testing.T) {
	model := config.ModelConfig{
		Name: "TestMoneky",
	}
	base := newBase(model)
	assert.Equal(t, base.name, model.Name)
	assert.Equal(t, base.model, model)

	c := Monkey{Base: base}
	result, err := c.perform()
	assert.Equal(t, result, "aaa")
	assert.Nil(t, err)
}
