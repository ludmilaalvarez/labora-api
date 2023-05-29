package services

import (
	"Api-Wallet/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
	//	"Api-Wallet/models"
)

var client = &http.Client{}

func Post(Person_id string) {
	var nuevaRespuesta models.Respuesta

	API := "https://api.checks.truora.com/v1/checks/"
	TOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k4YWJkOWE1ZGFmNzM1NGQ1YjVlZjVjYTI4MjJhMjA3OSIsImV4cCI6MzI2MTY4OTIwMiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ4ODkyMDIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX3hUSGxqU1d2RCIsImp0aSI6IjM2YTZiNGJlLTM3NTUtNGQzMC04ZTM0LTNmZDMyOGI3ZDk3NCIsImtleV9uYW1lIjoidHJ1Y29kZSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoidHJ1b3JhdGVhbW5ld3Byb2QtdHJ1Y29kZSJ9.PuE6cS6938PbQz_4qMLySs9dr3fywFqqGdfcF6Suw0U"

	body, _ := json.Marshal(map[string]string{
		"national_id":     Person_id,
		"country":         "BR",
		"type":            "person",
		"user_authorized": "true",
	})

	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, API, payload)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Truora-API-Key", TOKEN)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//nuevaRespuesta=

	err = json.Unmarshal(body, &nuevaRespuesta)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	//err = json.Unmarshal(body, &nuevaRespuesta)

	//log.Println(body.check.check_id)
	//body= string(body)
	//log.Printf("%s", body)
	//body.check_id = models.Respuesta.Check_id
	log.Println(nuevaRespuesta.Check.Check_id)
	log.Println(nuevaRespuesta.Check.Score)
	log.Println(nuevaRespuesta.Check.Summary.NamesFound[0])
	//log.Println(nuevaRespuesta.Check.Summary.NamesFound[0].LastName)
	if err != nil {
		log.Println(err)
	}
}

/* func TemporalPost() {

	url := "https://api.checks.truora.com/v1/checks"
	method := "POST"

	payload := strings.NewReader(`{` + " " + `"national_id": "064.707.957-73",` + " " + ` "country": "BR",` + " " + ` "type": "person",` + " " + ` "user_authorized": "true"` + "" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Truora-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k4YWJkOWE1ZGFmNzM1NGQ1YjVlZjVjYTI4MjJhMjA3OSIsImV4cCI6MzI2MTY4OTIwMiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ4ODkyMDIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX3hUSGxqU1d2RCIsImp0aSI6IjM2YTZiNGJlLTM3NTUtNGQzMC04ZTM0LTNmZDMyOGI3ZDk3NCIsImtleV9uYW1lIjoidHJ1Y29kZSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoidHJ1b3JhdGVhbW5ld3Byb2QtdHJ1Y29kZSJ9.PuE6cS6938PbQz_4qMLySs9dr3fywFqqGdfcF6Suw0U")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}
*/
