package routines

func Init() {
	AutoCalculateScoreBoard()
	go StartProcessingReportsPeriodically()
}
