package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://www.emsifa.com/api-wilayah-indonesia/api"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type ProvCityUsecase interface {
	GetAllProvinces() ([]Province, error)
	GetCitiesByProvinceID(provinceID string) ([]City, error)
	GetProvinceByID(provinceID string) (Province, error)
	GetCityByID(cityID string) (City, error)
}

type provCityUsecase struct{}

func NewProvCityUsecase() ProvCityUsecase {
	return &provCityUsecase{}
}

func (uc *provCityUsecase) GetAllProvinces() ([]Province, error) {
	var provinces []Province
	resp, err := http.Get(fmt.Sprintf("%s/provinces.json", baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (uc *provCityUsecase) GetCitiesByProvinceID(provinceID string) ([]City, error) {
	var cities []City
	url := fmt.Sprintf("%s/regencies/%s.json", baseURL, provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func (uc *provCityUsecase) GetProvinceByID(provinceID string) (Province, error) {
	var province Province
	url := fmt.Sprintf("%s/province/%s.json", baseURL, provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return Province{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&province); err != nil {
		return Province{}, err
	}
	return province, nil
}

func (uc *provCityUsecase) GetCityByID(cityID string) (City, error) {
	var city City
	url := fmt.Sprintf("%s/regency/%s.json", baseURL, cityID)
	resp, err := http.Get(url)
	if err != nil {
		return City{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&city); err != nil {
		return City{}, err
	}
	return city, nil
}