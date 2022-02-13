package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
	"sort"

	"github.com/jaslife1/shippingcost-server/model"
)

func CalculateJnTShippingCost(sender string, receiver string, weight float64) float64 {

	// Calculate cost using J&T calculation and shipping utility
	// curl -i -X POST \
	// -H "Content-Type:application/x-www-form-urlencoded; charset=UTF-8" \
	// --data "method=app.findRate&data%5BsenderAddr%5D=TANAUAN&data%5BreceiverAddr%5D=DAGAMI&data%5BserviceType%5D=&data%5Bweight%5D=1&pId=testtesttest" https://www.jtexpress.ph/index/router/index.html

	requestMethod := "POST"
	routerURL := "https://www.jtexpress.ph/index/router/index.html"

	//Format for J&T data
	// method=app.findRate&data%5BsenderAddr%5D=TANAUAN&data%5BreceiverAddr%5D=DAGAMI&data%5BserviceType%5D=&data%5Bweight%5D=1&pId=testtesttest
	// Use only the minimum requirements needed to query the shipping cost
	data := url.Values{}
	data.Add("method", "app.findRate")
	data.Add("data[senderAddr]", sender)
	data.Add("data[receiverAddr]", receiver)
	//Convert weight to string. Use the smallest possible digit (-1)
	data.Add("data[weight]", strconv.FormatFloat(weight, 'f', -1, 64))

	encodedData := data.Encode()

	body := strings.NewReader(encodedData)
	req, err := http.NewRequest(requestMethod, routerURL, body)
	if err != nil {
		fmt.Printf("Error making new request. Details: %+v", err)
		return 0
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Printf("Error sending request to J&T Server. Details: %+v", err)
		return 0
	}

	var respData model.JNTRate
	err = json.NewDecoder(resp.Body).Decode(&respData)

	if err != nil {
		fmt.Printf("Error Decoding Response Body: %+v\n", err)
		return 0
	}

	//Decode the Data
	reader := strings.NewReader(respData.Data)
	err = json.NewDecoder(reader).Decode(&respData.NewData)

	if err != nil {
		fmt.Printf("Error decoding data: %+v\n", err)
		return 0
	}

	// fmt.Printf("Decode: %+v\n", respData)
	// fmt.Printf("The shipping cost is: %+v\n", respData.NewData[0].TotalFee) //Since there is only one content always

	defer resp.Body.Close()

	ret, _ := strconv.ParseFloat(respData.NewData[0].TotalFee, 8)

	return ret
}

func GetAllProvinces() ([]*string, error) {

	//TODO: Query all of this from the database

	// Read the JSON file
	var filename = "files/provinces/ALL.json"
	var ret []*string;

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("GetAllProvinces: Failed to open file. Error: ", err)
		return ret, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("GetAllProvinces: Reading file. Error: ", err)
		return ret, err
	}

	var data []string
	json.Unmarshal([]byte(byteValue),&data)
	sort.Strings(data)

	for i:=0; i < len(data); i++ {
		ret = append(ret, &data[i])
	}

	//fmt.Println(ret)

	return ret, nil

}


func GetAllCities(province string) ([]*string, error) {
	//fmt.Println("Get Cities for "+province)
	//TODO: Query from the database
	
	// Read the JSON file
	var filename = "files/provinces/" + province + ".json"
	var ret []*string;

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("GetAllCities: Failed to open file. Error: ", err)
		return ret, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("GetAllCities: Reading file. Error: ", err)
		return ret, err
	}

	var data []model.AreaAddress
	json.Unmarshal([]byte(byteValue),&data)

	cityMap := make(map[string]string, 0)
	for _, val := range data {
		//Check if key already exists, if not, add the key and value
		_, ok := cityMap[val.City]
		if !ok {
			cityMap[val.City] = val.City
		}
	}

	var temp []string
	for _, val := range cityMap {
		temp = append(temp, val)
	}
	sort.Strings(temp)

	for i:=0; i < len(temp); i++ {
		ret = append(ret, &temp[i])
	}

	return ret, nil
}


//TO BE REMOVED
func GetAllTowns() []*string {
	towns := make([]*string, 0) //initially with 0 length
	temp1 := new(string)
	*temp1 = "ABUYOG";

	temp2 := new(string)
	*temp2 = "BAYBAY-CITY";

	temp3 := new(string)
	*temp3 = "CEBU-CITY";

	towns = append(towns, temp1)
	towns = append(towns, temp2)
	towns = append(towns, temp3)

	return towns
}