// Code generated from /Users/qinggniq/IdeaProjects/Cacl/src/zip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // zip

import "github.com/antlr/antlr4/runtime/Go/antlr"

// zipListener is a complete listener for a parse tree produced by zipParser.
type zipListener interface {
	antlr.ParseTreeListener

	// EnterSingal is called when entering the singal production.
	EnterSingal(c *SingalContext)

	// EnterTimesSignal is called when entering the timesSignal production.
	EnterTimesSignal(c *TimesSignalContext)

	// EnterSingalfComb is called when entering the singalfComb production.
	EnterSingalfComb(c *SingalfCombContext)

	// EnterTimesComb is called when entering the timesComb production.
	EnterTimesComb(c *TimesCombContext)

	// EnterTimes is called when entering the times production.
	EnterTimes(c *TimesContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// ExitSingal is called when exiting the singal production.
	ExitSingal(c *SingalContext)

	// ExitTimesSignal is called when exiting the timesSignal production.
	ExitTimesSignal(c *TimesSignalContext)

	// ExitSingalfComb is called when exiting the singalfComb production.
	ExitSingalfComb(c *SingalfCombContext)

	// ExitTimesComb is called when exiting the timesComb production.
	ExitTimesComb(c *TimesCombContext)

	// ExitTimes is called when exiting the times production.
	ExitTimes(c *TimesContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)
}
