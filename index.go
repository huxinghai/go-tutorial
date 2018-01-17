import "github.com/elastic/go-elasticsearch/client"

c, _ = client.New(client.WithHosts([]string{"https://elasticsearch:9200"}))
body := map[string]interface{}{
  "query": map[string]interface{}{
    "term": map[string]interface{}{
      "user": "kimchy",
    },
  },
}
resp, err := c.Search(body)
