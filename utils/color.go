package utils

import "github.com/fatih/color"

var (
	ColorTitle   = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	ColorPrompt  = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	ColorSuccess = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	ColorError   = color.New(color.FgHiRed, color.Bold).SprintFunc()
	ColorInfo    = color.New(color.FgHiBlue, color.Bold).SprintFunc()
)
