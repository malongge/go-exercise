package json

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Job struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo Info `json:"info"`
	JobInfo   Job  `json:"job"`
}
