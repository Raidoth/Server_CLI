package handlefunc

import (
	"encoding/json"
	"log"
	"net/http"
	"server/service/model"
	"server/service/opfunc"
	"strconv"
)

func SubString(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		if r.Header.Get("Content-Type") != "application/json" {

			t := model.RespError{}
			t.Success = false
			t.Error = "incorrect content type"
			res, _ := json.Marshal(t)
			w.Header().Set("Content-Length", strconv.Itoa(len(res)))
			w.Write(res)
			return
		}
		req := model.ReqLongestSubstring{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := model.RespLongestSubstring{}
		if req.SubString == "" {
			resp.Success = false
			resp.Error = "no data in substring"
		} else {
			res := opfunc.LongestSubstring(req.SubString)
			resp.Success = true
			resp.SubString = res
		}
		answer, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)

		w.Header().Set("Content-Length", strconv.Itoa(len(answer)))
		w.Header()
		w.Write(answer)
	default:
		http.Error(w, "incorrect method", http.StatusMethodNotAllowed)
		t := model.RespError{}
		t.Success = false
		t.Error = "incorrect method"
		res, _ := json.Marshal(t)
		w.Header().Set("Content-Length", strconv.Itoa(len(res)))
		w.Write(res)
	}

}
