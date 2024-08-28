package app

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Component struct {
	Html string
}

var dialogOptions = runtime.OpenDialogOptions{
	ShowHiddenFiles: true,
	Filters: []runtime.FileFilter{
		{
			DisplayName: "Excel or Csv files",
			Pattern:     "*.xlsx;*.xls;*.csv",
		},
	},
}
