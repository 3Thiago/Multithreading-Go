package service

import (
	"time"

	"github.com/3Thiago/Multithreading-Go/configs"
	models "github.com/3Thiago/Multithreading-Go/pkg/model"
	"github.com/3Thiago/Multithreading-Go/pkg/service"
	"github.com/sirupsen/logrus"
)

type ServiceInfra struct {
	servicePkg *service.ServicePkg
}

func NewServiceInfra() *ServiceInfra {
	return &ServiceInfra{}
}

func (s *ServiceInfra) GetCep(cep string) (any, error) {
	var ch1 = make(chan models.ViaCep)
	var ch2 = make(chan models.ApiCep)
	configs, err := configs.LoadEnvironmentVariables(".")
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		resp, err := s.servicePkg.GetViaCep(cep)
		if err != nil {
			logrus.Print(err)
		}
		ch1 <- *resp
	}()
	go func() {
		resp, err := s.servicePkg.GetApiCep(cep)
		if err != nil {
			logrus.Print(err)
		}
		ch2 <- *resp
	}()
	select {
	case viaCepResponse := <-ch1:
		logrus.Printf(" Resquest ViaCep: %s", viaCepResponse)
		return viaCepResponse, nil
	case apiCepResponse := <-ch2:
		logrus.Printf(" Resquest ApiCep: %s", apiCepResponse)
		return apiCepResponse, nil
	case <-time.After(1 * time.Second):
		logrus.Print("Timeout")
		res := map[string]interface{}{
			"msg":    configs.TIMEOUTMESSAGE,
			"status": configs.TIMEOUTCODE,
		}
		return res, nil
	}
}
