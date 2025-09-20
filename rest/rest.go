package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zivivle/go/blockchain"
	"github.com/zivivle/go/utils"
)

var port string

type url string

func (u url) MarshalText() (text []byte, err error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	// omitempty 데이터가 없을 경우 필드를 숨김
	Payload string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(w http.ResponseWriter, r *http.Request) {
	// 유저에게 JSON을 보내는 것부터 시작
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	/**
	1. http.ResponseWriter를 사용할 수 없는 상황에서의 JSON 파싱 방법이 될 수 있음
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(w, "%s", b)
	*/

	// 2. http.ResponseWriter를 사용한 쉬운 방법
	json.NewEncoder(w).Encode(data)
}

func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		w.WriteHeader(http.StatusCreated)
	}
}

func block(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// strconv 라이브러리를 사용해서 string -> int로 변환
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(w)

	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Start(aPort int) {
	// ListenAndServe가 생성하는 MUX를 여러개 생성했을 경우 에러가 발생함
	// 해결하기 위해서는 나만에 커스텀 MUX를 만들면됨
	port = fmt.Sprintf(":%d", aPort)
	// handler := http.NewServeMux()
	router := mux.NewRouter()
	//ServerMux는 url과 url 함수를 연결해주는 역할
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	// ListenAndServe에 커스텀 MUX를 사용하고 있다는걸 알려줘야해서
	// 생성한 MUX handler를 두번째 인자로 전달
	log.Fatal(http.ListenAndServe(port, router))
}
