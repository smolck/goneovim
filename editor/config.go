package editor

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
	homedir "github.com/mitchellh/go-homedir"
)

// gonvimConfig is the following toml file
// # Goneovim config toml
// [editor]
// ui = "trans"
// width = 1000  # >= 400
// height = 800  # >= 300
// fontFamily = "FuraCode Nerd Font Mono"
// fontsize = 18
// linespace = 10
// clipboard = true
// cursorBlink = true
// indentGuide = true
// cachedDrawing = false
// disableIMEinNormal = true
// startFullScreen = true
// transparent = 0.5
// desktopNotifications = true
// // -- diffpattern enum --
// // SolidPattern             1
// // Dense1Pattern            2
// // Dense2Pattern            3
// // Dense3Pattern            4
// // Dense4Pattern            5
// // Dense5Pattern            6
// // Dense6Pattern            7
// // Dense7Pattern            8
// // HorPattern               9
// // VerPattern               10
// // CrossPattern             11
// // BDiagPattern             12
// // FDiagPattern             13
// // DiagCrossPattern         14
// // LinearGradientPattern    15
// // RadialGradientPattern    16
// // ConicalGradientPattern   17
// // TexturePattern           24
// diffdeletepattern = 12
// diffchangepattern = 12
// diffaddpattern = 1
// SkipGlobalId = true
//
// [palette]
// AreaRatio = 0.8
// MaxNumberOfResultItems = 40
//
// [statusLine]
// visible = true
// # textLabel / icon / background / none
// modeIndicatorType = "icon"
// normalModeColor = "#123456"
// commandModeColor = "#123456"
// insertModeColor = "#123456"
// replaceModeColor = "#123456"
// visualModeColor = "#123456"
// termnalModeColor = "#123456"
// left = [ "mode", "filepath", "filename" ]
// right = [ "message", "git", "filetype", "fileformat", "fileencoding", "curpos", "lint" ]
//
// [tabline]
// visible = true
//
//
// [lint]
// visible = true
//
// [scrollBar]
// visible = true
//
// [activityBar]
// visible = true
// dropshadow = true
//
// [miniMap]
// visible = true
//
// [sideBar]
// visible = false
// dropshadow = true
// width = 360
// accentColor = "#5596ea"
//
// [workspace]
// # Path style
// #   full: fullpath,
// #   name: directory name only,
// #   minimum: only the last directory is full name, middle directory is short form
// pathStyle = minimum
// FileExplorerOpenCmd = ":tabew"
//
// # restore the previous sessions if there are exists.
// restoreSession = false
//
// [dein]
// tomlFile
type gonvimConfig struct {
	Editor       editorConfig
	Palette      paletteConfig
	Statusline   statusLineConfig
	Tabline      tabLineConfig
	Lint         lintConfig
	ScrollBar    scrollBarConfig
	ActivityBar  activityBarConfig
	MiniMap      miniMapConfig
	SideBar      sideBarConfig
	Workspace    workspaceConfig
	FileExplorer fileExplorerConfig
	Dein         deinConfig
}

type paletteConfig struct {
	AreaRatio              float64
	MaxNumberOfResultItems int
}
type editorConfig struct {
	Ui         string
	Width      int
	Height     int
	FontFamily string
	FontSize   int
	Linespace  int
	ExtCmdline bool
	// ExtWildmenu          bool
	ExtPopupmenu bool
	ExtTabline   bool
	// ExtMultigrid         bool
	ExtMessages bool
	Clipboard   bool
	// CursorBlink          bool
	CachedDrawing        bool
	DisableImeInNormal   bool
	GinitVim             string
	StartFullscreen      bool
	Transparent          float64
	DrawBorder           bool
	SkipGlobalId         bool
	IndentGuide          bool
	DesktopNotifications bool
	DiffAddPattern       int
	DiffDeletePattern    int
	DiffChangePattern    int
}

type statusLineConfig struct {
	Visible           bool
	ModeIndicatorType string
	NormalModeColor   string
	CommandModeColor  string
	InsertModeColor   string
	ReplaceModeColor  string
	VisualModeColor   string
	TerminalModeColor string
	Left              []string
	Right             []string
}

type tabLineConfig struct {
	Visible bool
}

type lintConfig struct {
	Visible bool
}

type miniMapConfig struct {
	Visible bool
}

type scrollBarConfig struct {
	Visible bool
}

type activityBarConfig struct {
	Visible    bool
	DropShadow bool
}

type sideBarConfig struct {
	Visible     bool
	DropShadow  bool
	Width       int
	AccentColor string
}

type workspaceConfig struct {
	RestoreSession bool
	PathStyle      string
}

type fileExplorerConfig struct {
	OpenCmd  string
	MaxItems int
}

type deinConfig struct {
	TomlFile string
}

func newGonvimConfig(home string) gonvimConfig {
	var config gonvimConfig

	config.init()

	// Read toml
	_, _ = toml.DecodeFile(filepath.Join(home, ".goneovim", "setting.toml"), &config)

	if config.Editor.Transparent < 1.0 {
		config.Editor.DrawBorder = true
	}

	if config.Editor.DiffAddPattern < 1 || config.Editor.DiffAddPattern > 24 {
		config.Editor.DiffAddPattern = 1
	}
	if config.Editor.DiffDeletePattern < 1 || config.Editor.DiffDeletePattern > 24 {
		config.Editor.DiffDeletePattern = 1
	}
	if config.Editor.DiffChangePattern < 1 || config.Editor.DiffChangePattern > 24 {
		config.Editor.DiffChangePattern = 1
	}

	if config.Editor.Width <= 400 {
		config.Editor.Width = 400
	}
	if config.Editor.Height <= 300 {
		config.Editor.Height = 300
	}
	if config.Editor.Transparent <= 0.1 {
		config.Editor.Transparent = 1.0
	}
	if config.Statusline.ModeIndicatorType == "" {
		config.Statusline.ModeIndicatorType = "textLabel"
	}

	if config.Editor.Linespace < 0 {
		config.Editor.Linespace = 6
	}

	if config.Statusline.NormalModeColor == "" {
		config.Statusline.NormalModeColor = newRGBA(60, 171, 235, 1).Hex()
	}
	if config.Statusline.CommandModeColor == "" {
		config.Statusline.CommandModeColor = newRGBA(82, 133, 184, 1).Hex()
	}
	if config.Statusline.InsertModeColor == "" {
		config.Statusline.InsertModeColor = newRGBA(42, 188, 180, 1).Hex()
	}
	if config.Statusline.VisualModeColor == "" {
		config.Statusline.VisualModeColor = newRGBA(153, 50, 204, 1).Hex()
	}
	if config.Statusline.ReplaceModeColor == "" {
		config.Statusline.ReplaceModeColor = newRGBA(255, 140, 10, 1).Hex()
	}
	if config.Statusline.TerminalModeColor == "" {
		config.Statusline.TerminalModeColor = newRGBA(119, 136, 153, 1).Hex()
	}

	if config.SideBar.Width == 0 {
		config.SideBar.Width = 200
	}
	if config.SideBar.AccentColor == "" {
		config.SideBar.AccentColor = "#5596ea"
	}

	if config.Workspace.PathStyle == "" {
		config.Workspace.PathStyle = "minimum"
	}

	return config
}

func (c *gonvimConfig) init() {
	// Set default value
	c.Editor.Width = 800
	c.Editor.Height = 600
	c.Editor.Transparent = 1.0

	// UI options:
	c.Editor.SkipGlobalId = false
	c.Editor.CachedDrawing = true

	// basic UI
	c.Editor.ExtCmdline = true
	// c.Editor.ExtWildmenu = false
	c.Editor.ExtPopupmenu = false
	c.Editor.ExtTabline = true
	c.Editor.ExtMessages = false
	c.Editor.DrawBorder = false

	// Indent guide
	c.Editor.IndentGuide = true

	// replace diff color drawing pattern
	c.Editor.DiffAddPattern = 12
	c.Editor.DiffDeletePattern = 12

	// palette size
	c.Palette.AreaRatio = 0.5
	c.Palette.MaxNumberOfResultItems = 30

	c.Statusline.Visible = false
	c.Statusline.ModeIndicatorType = "textLabel"
	c.Statusline.Left = []string{"mode", "filename"}
	c.Statusline.Right = []string{"git", "filetype", "fileformat", "fileencoding", "curpos", "lint"}

	c.Tabline.Visible = true

	c.Lint.Visible = true

	// c.ActivityBar.Visible = true

	c.ScrollBar.Visible = true

	c.SideBar.Width = 200
	c.SideBar.AccentColor = "#5596ea"

	c.Workspace.PathStyle = "minimum"
}

func outputGonvimConfig() {
	home, err := homedir.Dir()
	if err != nil {
		home = "~"
	}
	filepath := filepath.Join(home, ".goneovim", "setting.toml")
	if isFileExist(filepath) {
		return
	}
	buf := new(bytes.Buffer)
	toml.NewEncoder(buf).Encode(editor.config)
	err = ioutil.WriteFile(filepath, buf.Bytes(), 664)
	if err != nil {
		editor.pushNotification(NotifyWarn, -1, "[Goneovim] "+err.Error())
	}
}
