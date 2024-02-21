package imgprint

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilDestinationImage = erorr.Error("imgprint: nil destination image")
	errNilFontDrawer       = erorr.Error("imgprint: nil font drawer")
)
