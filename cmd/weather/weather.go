package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const BaseURL = "https://api.openweathermap.org"

type Conditions struct {
	Summary string
	Temp    Temperature
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp Temperature
	}
}

type Client struct {
	ApiKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string) Client {
	return Client{
		ApiKey:  apiKey,
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.ApiKey)
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)

	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status %q", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}
	return conditions, nil
}

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf(
			"invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one Weather element", data)
	}
	conditions := Conditions{
		Summary: resp.Weather[0].Main,
		Temp:    resp.Main.Temp,
	}
	return conditions, nil
}

func RunCLI() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the environment variable OPENWEATHERMAP_API_KEY.")
	}

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s LOCATION\n\nExample: %[1]s London,UK", os.Args[0])
	}
	location := os.Args[1]

	conditions, err := Get(location, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %.1fºC\n", conditions.Summary,
		conditions.Temp.Celsius())
}

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}
