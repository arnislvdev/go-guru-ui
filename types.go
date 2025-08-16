package guruui

// TextStyle represents different text styling options.
type TextStyle int

const (
	StylePrimary TextStyle = iota
	StyleMuted
	StyleAccent
)

// BorderStyle represents different border styling options.
type BorderStyle int

const (
	BorderNone BorderStyle = iota
	BorderThin
	BorderThick
	BorderDouble
	BorderRounded
	BorderGeometric
	BorderLayered
	BorderFractal
)

// Alignment represents text alignment options.
type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
	AlignJustify
)

// Icons for different message types and statuses
const (
	IconInfo     = "ℹ" // Information
	IconSuccess  = "✓" // Success/Checkmark
	IconWarning  = "⚠" // Warning
	IconError    = "✗" // Error/Cross
	IconArrow    = "→" // Arrow/Flow
	IconStar     = "★" // Star/Favorite
	IconHeart    = "♥" // Heart/Like
	IconCheck    = "☑" // Checkbox checked
	IconUncheck  = "☐" // Checkbox unchecked
	IconDot      = "●" // Dot/Bullet
	IconSquare   = "■" // Square
	IconCircle   = "●" // Circle
	IconTriangle = "▲" // Triangle
	IconDiamond  = "◆" // Diamond
	IconCross    = "✚" // Cross/Plus
	IconInfinity = "∞" // Infinity
	IconPi       = "π" // Pi
	IconSigma    = "Σ" // Sigma
	IconOmega    = "Ω" // Omega
	IconDelta    = "Δ" // Delta
)

// Neo-Brutalist Design Patterns
const (
	// Geometric Patterns
	PatternHexagon  = "⬡⬢⬣⬤⬥⬦⬧⬨⬩⬪⬫⬬⬭⬮⬯"
	PatternDiamond  = "◆◇◈◊◈◇◆"
	PatternTriangle = "▲△▲△▲△▲"
	PatternSquare   = "■□▣▤▥▦▧▨▩"
	PatternCircle   = "●○◐◑◒◓◔◕"
	PatternCross    = "✚✛✜✝✞✟✠✡"
	PatternStar     = "★☆✮✯✰✱✲✳✴✵"
	PatternArrow    = "→←↑↓↔↕↖↗↘↙"
	PatternWave     = "~≈∽∿∼≁≂≃≄≅≆≇≈"
	PatternFractal  = "◢◣◤◥◢◣◤◥◢◣◤◥"

	// Brutalist Borders
	BorderBrutalTop    = "╔══╗"
	BorderBrutalMid    = "║  ║"
	BorderBrutalBottom = "╚══╝"

	// Layered Depth Patterns
	LayerDepth1 = "▁▂▃▄▅▆▇█"
	LayerDepth2 = "░▒▓█"
	LayerDepth3 = "▏▎▍▌▋▊▉█"

	// Dynamic Progress Patterns
	ProgressRing    = "◐◑◒◓◐◑◒◓"
	ProgressWave    = "▁▂▃▄▅▆▇█▇▆▅▄▃▂▁"
	ProgressPulse   = "●○●○●○●○"
	ProgressMatrix  = "█░█░█░█░"
	ProgressFractal = "◢◣◤◥◢◣◤◥"

	// Unique Layout Elements
	CornerTopLeft     = "╭"
	CornerTopRight    = "╮"
	CornerBottomLeft  = "╰"
	CornerBottomRight = "╯"
	DividerThick      = "━━━"
	DividerThin       = "───"
	DividerDotted     = "┄┄┄"
	DividerWave       = "∿∿∿"

	// Brutalist Accents
	AccentBrutal    = "█▓▒░"
	AccentGeometric = "◆◇◈◊"
	AccentOrganic   = "●○◐◑"
	AccentTechnical = "⚡⚙️🔧⚗️"
)

// Color codes for enhanced visual appeal
const (
	ColorReset     = "\033[0m"
	ColorBold      = "\033[1m"
	ColorDim       = "\033[2m"
	ColorItalic    = "\033[3m"
	ColorUnderline = "\033[4m"
	ColorBlink     = "\033[5m"
	ColorReverse   = "\033[7m"
	ColorHidden    = "\033[8m"

	// Standard colors
	ColorBlack   = "\033[30m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"

	// Bright colors
	ColorBrightBlack   = "\033[90m"
	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"

	// Background colors
	ColorBgBlack   = "\033[40m"
	ColorBgRed     = "\033[41m"
	ColorBgGreen   = "\033[42m"
	ColorBgYellow  = "\033[43m"
	ColorBgBlue    = "\033[44m"
	ColorBgMagenta = "\033[45m"
	ColorBgCyan    = "\033[46m"
	ColorBgWhite   = "\033[47m"
)

// Animation frames for dynamic elements
var (
	SpinnerFrames = []string{
		"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
	}

	PulseFrames = []string{
		"●", "○", "●", "○", "●", "○", "●", "○",
	}

	WaveFrames = []string{
		"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█",
	}

	MatrixFrames = []string{
		"█░", "░█", "█░", "░█", "█░", "░█",
	}

	FractalFrames = []string{
		"◢◣", "◤◥", "◢◣", "◤◥", "◢◣", "◤◥",
	}
)
