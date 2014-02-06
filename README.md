ColorPrint
==========

a Color-Printf library for GO

Examples
==========

    package main

    import (
	    "fmt"
	    "github.com/rawoke083/ColorPrint"
    )

    func main() {
	    fmt.Println("Hello in default")
	    ColorPrint.ColWrite("Hello in GREEN", ColorPrint.CL_GREEN)
    }
