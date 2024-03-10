package openrtb3_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

var _ = DescribeTable(
	"Marshaling",
	func(file string) {
		golden, err := ioutil.ReadFile(filepath.Join("testdata", file))
		Expect(err).NotTo(HaveOccurred())

		v := new(Body)
		Expect(json.Unmarshal(golden, v)).To(Succeed())

		b, err := json.Marshal(v)
		Expect(err).NotTo(HaveOccurred())

		Expect(b).To(MatchJSON(golden))
	},

	Entry("request", "request.json"),
	Entry("response", "response.json"),
)
