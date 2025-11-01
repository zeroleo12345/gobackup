package compressor

import (
	"path"
	"strings"
	"testing"
	"time"

	"github.com/gobackup/gobackup/config"
	"github.com/longbridgeapp/assert"
	"github.com/spf13/viper"
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
	viper.Set("format", "backup-2006.01.02.15.04.05")
	base := newBase(config.ModelConfig{
		CompressWith: config.SubConfig{
			Type:  "compress_with",
			Name:  "tar",
			Viper: viper,
		},
	})
	prefixPath := path.Join(base.model.TempPath, time.Now().Format("backup-2006.01.02.15.04"))
	archivePath := base.archiveFilePath(".tar")
	assert.True(t, strings.HasPrefix(archivePath, prefixPath))
	assert.True(t, strings.HasSuffix(archivePath, ".tar"))
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
