package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	models "github.com/3Thiago/Multithreading-Go/pkg/model"
	"github.com/spf13/viper"
)

type ServicePkg struct {
}

func (s *ServicePkg) GetApiCep(cep string) (*models.ApiCep, error) {
	url := viper.GetString("API_CEP_URL")
	res, err := http.Get(fmt.Sprintf("%s%s.json", url, cep))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp models.ApiCep
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil

}

func (s *ServicePkg) GetViaCep(cep string) (*models.ViaCep, error) {
	url := viper.GetString("VIA_CEP_URL")
	res, err := http.Get(fmt.Sprintf("%s%s/json/", url, cep))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp models.ViaCep
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
