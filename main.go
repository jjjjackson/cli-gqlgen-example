package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/jjjjackson/cli-gqlgen-example/graph"
	"github.com/jjjjackson/cli-gqlgen-example/graph/generated"
)

func fixturePath(fixture string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	pwd := filepath.Dir(currentFile)
	filename := fmt.Sprintf("%v.graphql", fixture)

	return filepath.Join(pwd, "fixture", filename)
}

func loadFixture(fixture string) string {
	readText, _ := ioutil.ReadFile(fixturePath(fixture))
	jsonText, _ := json.Marshal(string(readText))
	graphqlQuery := fmt.Sprintf("{\"query\": %v}", string(jsonText))

	return graphqlQuery
}

func fakeRequest(fixtureName string) *http.Request {
	reqBody := bytes.NewBufferString(loadFixture(fixtureName))
	req, _ := http.NewRequest(http.MethodPost, "foo.com.tw", reqBody)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func main() {
	// 可以在這裡切換你想嘗試的 GraphQL Query file
	req := fakeRequest("query_example")
	// req := fakeRequest("mutation_example")

	record := httptest.NewRecorder()

	// 使用 gqlgen 的 handler，並將結果寫入 record 之中
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	// 執行 gqlgen 的 handler
	srv.ServeHTTP(record, req)

	resp := record.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
