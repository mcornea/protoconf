package testdata

type FirstMessage struct {
	AString string
	AnInt int
	AStringWithJSONTag string `json:"a_string_with_json_tag"`
	// AnotherMessage SecondMessage
}

type SecondMessage struct {
	RepeatedString []string
	MappedString map[string]float32
	AnInterface interface{}
};
