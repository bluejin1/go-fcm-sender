package helper

import (
	"github.com/common-nighthawk/go-figure"
	"runtime"
	"strings"
)

func GetServicePackageName() string {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	pkage := ""
	if parts[pl-2][0] == '(' {
		pkage = strings.Join(parts[0:pl-2], ".")
	} else {
		pkage = strings.Join(parts[0:pl-1], ".")
	}
	packagenames := strings.Split(pkage, "/")
	return packagenames[0]
}

func ShowServicelogoPrint() {
	serviceLogo := strings.ToUpper("SG Feature. " + GetServicePackageName())
	myFigure := figure.NewColorFigure(serviceLogo, "eftitalic", "blue", true)
	myFigure.Print()
}
