package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sisu-network/lib/log"
)

type VaultHandler struct {
	ServerUrl string
	Path      string
}

func NewGatewayHandler(server, path string) *VaultHandler {
	return &VaultHandler{
		ServerUrl: server,
		Path:      path,
	}
}

func (v *VaultHandler) GetGatewayAddress(w http.ResponseWriter, r *http.Request) {
	queryStr := r.URL.Query()
	chain := queryStr.Get("chain")

	addr, err := v.getVaultAddr(chain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bz, err := json.Marshal(map[string]string{
		"address": addr,
	})
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(bz); err != nil {
		log.Error(err)
		return
	}
}

func (v *VaultHandler) getVaultAddr(chain string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, v.ServerUrl+v.Path, nil)
	if err != nil {
		log.Error("error when new request ", err)
		return "", err
	}

	q := req.URL.Query()
	q.Add("chain", chain)
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer resp.Body.Close()

	type response struct {
		Address string `json:"address"`
	}

	res := &response{}
	if err := json.Unmarshal(body, res); err != nil {
		log.Error(err)
		return "", err
	}

	return res.Address, nil
}
