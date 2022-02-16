package model

type InputYaml struct {
	ProjectName string
	Tests       struct {
		Http struct {
			Cases []*HttpCase
		}
	}
}

type HttpCase struct {
	Name string
	Jobs []*HttpJob
}

type HttpJob struct {
	Name    string
	Request struct {
		Url    string
		Method string
	}
	Response struct {
		StatusCode string
	}
}
