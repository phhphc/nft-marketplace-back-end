package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch"
	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch/nft"
	"log"
	"net/http"
)

func main() {
	elastic, err := elasticsearch.NewElasticsearch([]string{"http://localhost:9200"})
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := elasticsearch.NewNFTStorage(elastic, false)
	if err != nil {
		log.Fatalln(err)
	}

	nftAPI := nft.NewHandler(storage)

	router := httprouter.New()
	router.HandlerFunc("GET", "/nfts/:token/:identifier", nftAPI.FindOneNFT)
	router.HandlerFunc("POST", "/nfts", nftAPI.CreateNFT)
	router.HandlerFunc("DELETE", "/nfts/:token/:identifier", nftAPI.DeleteNFT)

	log.Fatalln(http.ListenAndServe(":8080", router))
}
