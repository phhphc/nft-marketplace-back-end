package nft

//type DummyHttpClient struct {
//	responseMock string
//}
//
//func DummyElasticSearchClient(endpoint string, responseMock string) (*elasticsearch.Client, error) {
//	client, err := elasticsearch.NewClient(elasticsearch.Config{
//		Addresses: []string{endpoint},
//		Transport: MockHttpClient(responseMock),
//	})
//	if err != nil {
//		return nil, err
//	}
//}
//
//func (c *DummyHttpClient) Do(r *http.Request) (*http.Response, error) {
//	recoder := httptest.NewRecorder()
//	recoder.Write([]byte(c.responseMock))
//	recoder.Header().Set("Content-Type", "application/json")
//
//	return recoder.Result(), nil
//}
//
//func MockHttpClient(responseMock string) *DummyHttpClient {
//	return &DummyHttpClient{
//		responseMock: responseMock,
//	}
//}
//
//func TestNFTStorage(t *testing.T) {
//	assert.Equal(t, 1, 1, "1 should equal 1")
//}
