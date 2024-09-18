package v1

type ApiV1Group struct {
	UserApi     UserApi
	ConfigDXApi ConfigDXApi
}

var APIGroupPtr = new(ApiV1Group)
