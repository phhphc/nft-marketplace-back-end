package nft

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(storage elasticsearch.NFTStorer) *Handler {
	return &Handler{
		service: NewService(storage),
	}
}

// GET /api/v1/search/nfts/{token}/{identifier}
func (h *Handler) FindOneNFT(w http.ResponseWriter, r *http.Request) {
	var req FindOneRequest
	ps := httprouter.ParamsFromContext(r.Context())
	req.Token = ps.ByName("token")
	req.Identifier = ps.ByName("identifier")

	res, err := h.service.FindOneNFT(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bdy, _ := json.Marshal(res)
	_, _ = w.Write(bdy)
}

// POST /api/v1/search/nfts
func (h *Handler) CreateNFT(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	res, err := h.service.CreateNFT(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	bdy, _ := json.Marshal(res)
	_, _ = w.Write(bdy)
}

// DELETE /api/v1/search/nfts/{token}/{identifier}
func (h *Handler) DeleteNFT(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	bdy, _ := json.Marshal("Delete NFT from elastic successfully")
	_, _ = w.Write(bdy)
}
