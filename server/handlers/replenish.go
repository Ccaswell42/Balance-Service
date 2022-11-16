package handlers

import (
	"avito/storage/user_balance"
	"log"
	"net/http"
)

func (d *Data) Replenish(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodPost {
	//	JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
	//	return
	//}
	//var bal user_balance.UserBalance
	//err := json.NewDecoder(r.Body).Decode(&bal)
	//
	//if err != nil {
	//	JsonResponse(false, w, "no request body", http.StatusBadRequest)
	//	log.Println(err)
	//	return
	//}
	//contentType := r.Header.Get("Content-Type")
	//if contentType != "application/json" {
	//	JsonResponse(false, w, "content-type is not application/json", http.StatusUnsupportedMediaType)
	//	log.Println(contentType)
	//	return
	//}

	errStr := ValidateRequest(r, w, http.MethodPost)
	if errStr != OK {
		return
	}

	bal, errStr := ValidateBodyUserBalance(r.Body)
	if errStr != OK {
		JsonResponse2(ResponseError, w, errStr, http.StatusOK)
		return
	}

	err := user_balance.ReplenishBalance(d.DB, bal)
	if err != nil {
		JsonResponse2(ResponseError, w, "can't replenish the balance", http.StatusInternalServerError)
		//JsonResponse(false, w, "can't replenish the balance", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	JsonResponse2(OK, w, "balance successfully replenished", http.StatusOK)
	//JsonResponse(true, w, "balance successfully replenished", http.StatusOK)
}
