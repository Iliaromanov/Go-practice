package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)


const key = "secret.key"; // enter your iex api key (its free and easy to obtain, just go to their website)

type StockQuoteResponse struct {
	CompanyName   string  `json:"companyName"`
	LatestPrice   float32 `json:"latestPrice"`
	PreviousClose float32 `json:"previousClose"`
	ChangePercent float32 `json:"changePercent"`
}

func main() {
	// testing below funcs
	quote, err := RequestStockQuote("GOOGL");
	if err != nil {
		panic(err);
	}
	fmt.Printf(
		"GOOGL:\nname: %v, latest price: %v, prev close: %v, change percent: %v\n",
		quote.CompanyName, quote.LatestPrice, quote.PreviousClose, quote.ChangePercent,
	)

	quote, err = RequestStockQuote("AMZN");
	fmt.Println("AMZN:", quote)

	quote, err = RequestStockQuote("MSFT");
	fmt.Println("MSFT:", quote)
}


func RequestStockQuote(sym string) (StockQuoteResponse, error) { // public function (uppercased)
	var resp StockQuoteResponse;
	baseUrl := fmt.Sprintf("https://cloud.iexapis.com/stable/stock/%s/quote", sym);
	err := makeGetRequest(
		baseUrl, map[string]string{"token": key}, &resp,
	)
	if err != nil {
		return StockQuoteResponse{}, fmt.Errorf("error from makeGetRequest:\n %v\n", err);
	}

	return resp, nil;
}


// makes http GET request and returns response inside of passed response struct
func makeGetRequest(
	baseUrl string, params map[string]string, response interface{}, // interface dentoes any type *including ptrs*
) error { // private function (lowercased)
	fullUrl, err := url.Parse(baseUrl); // convert input url string into url.URL struct ptr
	if err != nil {
		return fmt.Errorf("error in parsing baseURL: %v\n", err);
	}
	queryParams := url.Values{}; // create an empty 'dict' for query params
	for key, value := range params { // parse the query params into
		// can use .Set() instead of .Add() => 
		//  .Set() overwrites if key already exists, .Add always appends
		queryParams.Add(key, value);
	}

	// Values.Encode() formats the query params to be usable in url
	//  eg. changes space char to '+'
	// url.URL.RawQuery is a field for encoded query params
	fullUrl.RawQuery = queryParams.Encode();

	fmt.Println("\n---------------------------\nrequestURL: ", fullUrl.String());

	resp, err := http.Get(fullUrl.String());
	if err != nil {
		return fmt.Errorf("error in http get request: %v\n", err);
	}

	defer resp.Body.Close(); // resp.Body is an io.Reader object
	
	if resp.Status != "200 OK" {
		return fmt.Errorf("response status: %v", resp.Status);
	}

	result, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		return fmt.Errorf("error in reading resp.Body: %v\n", err);
	}
	err = json.Unmarshal(result, response); // parse json bytes into response struct
	if err != nil {
		return fmt.Errorf("error in parsing json to response struct: %v\n", err);
	}

	return nil
}

