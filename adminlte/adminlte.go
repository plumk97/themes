package adminlte

import (
	"strings"

	"github.com/gobuffalo/packr/v2"
	adminTemplate "github.com/plumk97/go-admin/template"
	"github.com/plumk97/go-admin/template/components"
	"github.com/plumk97/go-admin/template/types"
	"github.com/plumk97/themes/adminlte/resource"
	"github.com/plumk97/themes/common"
)

const (
	ColorschemeSkinBlack       = "skin-black"
	ColorschemeSkinBlackLight  = "skin-black-light"
	ColorschemeSkinBlue        = "skin-blue"
	ColorschemeSkinBlueLight   = "skin-blue-light"
	ColorschemeSkinGreen       = "skin-green"
	ColorschemeSkinGreenLight  = "skin-green-light"
	ColorschemeSkinPurple      = "skin-purple"
	ColorschemeSkinPurpleLight = "skin-purple-light"
	ColorschemeSkinRed         = "skin-red"
	ColorschemeSkinRedLight    = "skin-red-light"
	ColorschemeSkinYellow      = "skin-yellow"
	ColorschemeSkinYellowLight = "skin-yellow-light"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Adminlte = Theme{
	ThemeName: "adminlte",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: TemplateList,
	},
}

func init() {
	adminTemplate.Add("adminlte", &Adminlte)
}

func Get() *Theme {
	return &Adminlte
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return TemplateList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	path = strings.Replace(path, "/assets/dist", "", -1)
	box := packr.New("adminlte", "./resource/assets/dist")
	return box.Find(path)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}
