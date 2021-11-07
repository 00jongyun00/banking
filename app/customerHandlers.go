package app

import (
	"bangking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	status := request.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)
	// fprint 와 비슷한 함수
	// writer 를 전달하고 encode 함수를 호출하고 모든 고객을 json 형식으로 인코딩 하도록 요청
	// 그러나 헤더로 이동하여 콘텐츠 유형을 보면 일반 텍스트로 표시 됩니다.
	// 따라서 이 응답은 JSON 처럼 보이지만 실제로는 JSON 이 아니므로 올바른 내용을 설정해야 합니다.
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(err.Code)
		json.NewEncoder(writer).Encode(err.AsMessage())
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		// 먼저 헤더를 설정한다음에 헤더에 status code 를 작성해야 합니다.
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(err.Code)
		json.NewEncoder(writer).Encode(err.AsMessage())
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
