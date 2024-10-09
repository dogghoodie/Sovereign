package visuals

// Struct for our terminal colors
type TerminalColors struct {
	ANSI_BG      string
	ANSI_FG      string
	ANSI_BOLD    string
	ANSI_RESET   string
	BACKGROUND   string
	CURRENT_LINE string
	FOREGROUND   string
	COMMENT      string
	CYAN         string
	GREEN        string
	MATRIX1      string
	MATRIX2      string
	MATRIX3      string
	ORANGE       string
	PINK         string
	PURPLE       string
	RED          string
	YELLOW       string
	BLACK        string
	WHITE        string
	WHITE2       string
	WHITE3       string
	WHITE4       string
}

//TODO: Solve gocui color limit issue or simplify theme

// Color list from original project.
// Color codes are all from :
// https://en.wikipedia.org/wiki/Dracula_(color_scheme)

var Colors = TerminalColors{
	ANSI_BOLD:    "\u001B[1m",
	ANSI_BG:      "\033[48;2;",
	ANSI_FG:      "\033[38;2;",
	ANSI_RESET:   "\033[0m",
	BACKGROUND:   "\033[48;2;40;42;54m",
	CURRENT_LINE: "\033[48;2;68;71;90m",
	FOREGROUND:   "\033[38;2;248;248;242m",
	COMMENT:      "\033[38;2;98;114;164m",
	CYAN:         "\033[38;2;139;233;253m",
	GREEN:        "\033[38;2;80;250;123m",
	MATRIX1:      "\033[38;2;57;255;20m",
	MATRIX2:      "\033[38;2;80;200;120m",
	MATRIX3:      "\033[38;2;0;128;0m",
	ORANGE:       "\033[38;2;255;184;108m",
	PINK:         "\033[38;2;255;121;198m",
	PURPLE:       "\033[38;2;189;147;249m",
	RED:          "\033[38;2;255;85;85m",
	YELLOW:       "\033[38;2;241;250;140m",
	BLACK:        "\033[38;2;28;28;28m",
	WHITE:        "\033[38;2;255;255;255m",
	WHITE2:       "\033[38;2;230;230;230m",
	WHITE3:       "\033[38;2;195;195;195m",
	WHITE4:       "\033[38;2;175;175;175m",
}
