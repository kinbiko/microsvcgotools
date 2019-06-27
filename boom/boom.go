package boom

// DoSomethingDangerous does something dangerous
func DoSomethingDangerous()        { shipNaughtyCodeToProduction() }
func shipNaughtyCodeToProduction() { skipCodeReviewAndShip() }
func skipCodeReviewAndShip()       { dontWriteTestsAndDeploy() }
func dontWriteTestsAndDeploy()     { shipToProd() }

func shipToProd() {
	panic("Boom!")
}
