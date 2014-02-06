package ColorPrint

import (
	"fmt"
)

const CL_LIGHT_GREEN string = "\x1b[1;32;1m"
const CL_GREEN string = "\x1b[0;32;2m"
const CL_BLACK string = "\x1b[0;30;1m"
const CL_DARK_GRAY string = "\x1b[1;30;1m"
const CL_BLUE string = "\x1b[0;34;2m"
const CL_LIGHT_BLUE string = "\x1b[1;34;1m"
const CL_CYAN string = "\x1b[0;36;2m"
const CL_LIGHT_CYAN string = "\x1b[1;36;1m"
const CL_RED string = "\x1b[1;31;1m"
const CL_LIGHT_RED string = "\x1b[0;31;2m"
const CL_YELLOW string = "\x1b[1;33;1m"
const CL_WHITE string = "\x1b[1;37;1m"
const CL_LIGHT_GRAY string = "\x1b[1;37;1m"
const CL_GRAY string = "\x1b[0;37;2m"

func ToColor(msg string, col string) string {
	return col + msg + "\x1b[0m"
}

func ColWrite(msg string, col string) {

	fmt.Printf("%s", ToColor(msg, col))
}
