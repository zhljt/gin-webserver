package v1

type ApiV1Group struct {
	UserApi     UserApi
	ConfigDXApi ConfigDXApi
	InitDBApi   InitDBApi
}

var APIGroupPtr = new(ApiV1Group)
