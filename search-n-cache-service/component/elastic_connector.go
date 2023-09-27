package component

import (
	"crypto/tls"
	"net/http"
	"os"
	"strconv"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
)

var ElasticTypedClient *elasticsearch.TypedClient

func InitializeElasticSearchClient() *elasticsearch.TypedClient {
	log.Info("Initializing ElasticSearch Client")
	skipVerifySSL, _ := strconv.ParseBool(os.Getenv("ELASTIC_SKIP_VERIFY_SSL"))
	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTIC_URL")},
		Username:  os.Getenv("ELASTIC_USERNAME"),
		Password:  os.Getenv("ELASTIC_PASSWORD"),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipVerifySSL,
			},
		},
	})

	if err != nil {
		log.Error(err)
		return nil
	}
	log.Info("ElasticSearch Client Initialized Successfully")
	return typedClient
}
