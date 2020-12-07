package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"syreclabs.com/go/faker"
)

func main() {
	e := echo.New()
	e.GET("/d", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339))
	})
	e.GET("/events", handleEvents)
	e.GET("/participans", handleParticipans)
	e.GET("/test", hand)
	e.Logger.Fatal(e.Start(":1323"))

}

func handleEvents(c echo.Context) error {
	limit, err := strconv.ParseInt(c.QueryParams().Get("limit"), 10, 32)
	if err != nil {
		limit = 1
	}
	events := ""
	for i := 0; i < int(limit); i++ {
		randomID, _ := uuid.NewRandom()
		duration1, _ := time.ParseDuration("150h")
		duration2, _ := time.ParseDuration("50h")
		duration3, _ := time.ParseDuration("10h")

		event := fmt.Sprintf(`
	{
		"eventId": "%s",
		"type": "raceroster",
		"name": "%s",
		"city": "San Francisco",
		"region": {
				"code": "%s",
				"name": "%s"
		},
		"country": {
				"code": "US",
				"name": "United States"
		},
		"startDate": "%s",
		"timeZone": "America/Vancouver",
		"registrationOpenDate": "%s",
		"registrationCloseDate": "%s",
		"lastModifiedDate": "2015-02-11T15:19:12+00:00",
		"url": "https://raceroster.com/events/2015/761/2015-alcatraz-xxxv-escape-from-the-rock",
		"organizerName": "%s",
		"organizerEmail": "%s",
		"organizerPhone": "%s",
		"organizerDashboardUrl": "https://raceroster.com/dashboard/event-organizers/legacy/Home?eid=1",
		"address": "%s",
		"latitude": 43.7820344,
		"longitude": -88.4437064,
		"locale": "en_US",
		"description": "Lorem Ipsum description",
		"facebook": "https://www.facebook.com/RaceRoster",
		"twitter": "RaceRoster",          
		"branding": {
				"logo": "https://cdn.raceroster.com/event_logo/2013-06-17_10-48-06_logo_uid20354.png",
				"backgroundImage": "https://cdn.raceroster.com/event_bg/2013-08-08_13-02-02_eventImage.jpg",
				"backgroundPattern": "https://raceroster.com/patterns/pattern.png",
				"backgroundColor": "#F0F0F0",
				"primaryColor": "#2A6979",
				"primaryFontColor": "#FFFFFF",
				"secondaryColor": "#0099FF",
				"secondaryFontColor": "#FFFFFF",
				"tertiaryColor": "#0099FF",
				"tertiaryFontColor": "#FFFFFF"
			  },
		"resultsUrl": "https://results.raceroster.com/results/am52payyhk52ma9r",
		"subEvents": {
			"data": [
				{
					"subEventId": 12,
					"name": "Individual",
					"maxParticipants": "-1",
					"participantMinAge": "",
					"participantMaxAge": "",
					"distance": "",
					"distanceType": ""
				},
				{
					"subEventId": 14,
					"name": "Relay Team",
					"maxParticipants": "-1",
					"participantMinAge": "",
					"participantMaxAge": "",
					"distance": "",
					"distanceType": ""
				},
				{
					"subEventId": 15,
					"name": "Swim-Only",
					"maxParticipants": "-1",
					"participantMinAge": "",
					"participantMaxAge": "",
					"distance": "",
					"distanceType": ""
				}
			]
		}
	}`,
			randomID,
			faker.Name().String(),
			faker.Address().StateAbbr(),
			faker.Address().State(),
			faker.Time().Backward(duration1).Format(time.RFC3339),
			faker.Time().Backward(duration2).Format(time.RFC3339),
			faker.Time().Backward(duration3).Format(time.RFC3339),
			faker.Company().Name(),
			faker.Internet().Email(),
			faker.PhoneNumber(),
			faker.Address().String(),
		)

		events = fmt.Sprintf(events+"%s,", event)
	}
	events = strings.TrimSuffix(events, ",")

	empArray := fmt.Sprintf(`
{
	"data": [
		%s
	]
}`, events)
	// Declared an empty interface of type Array
	var results struct {
		Data []map[string]interface{} `json:"data"`
	}

	json.Unmarshal([]byte(empArray), &results)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(results)
}

func handleParticipans(c echo.Context) error {
	limit, err := strconv.ParseInt(c.QueryParams().Get("limit"), 10, 32)
	if err != nil {
		limit = 1
	}

	offset, err := strconv.ParseInt(c.QueryParams().Get("offset"), 10, 32)
	if err != nil {
		offset = 0
	}
	total, err := strconv.ParseInt(c.QueryParams().Get("total"), 10, 32)
	if err != nil {
		total = limit
	}

	events := ""
	max := limit
	if limit+offset > total {
		max = total - offset
	}
	for i := 0; i < int(max); i++ {
		randomID, _ := uuid.NewRandom()
		duration1, _ := time.ParseDuration("150h")
		duration2, _ := time.ParseDuration("50h")
		// duration3, _ := time.ParseDuration("10h")
		id := i + faker.Number().NumberInt(3)
		event := fmt.Sprintf(`
		{
            "participantId": "%s",
            "uniqueParticipantId": "%s",
            "transactionId": 8281947,
            "subEventId": 12,
            "registrationDate": "%s",
            "lastModifiedDate": "%s",
            "name": "%s",
            "dob": "%s",
            "age": %d,
            "gender": 1,
            "city": "%s",
            "region": {
                "code": "%s",
                "name": "%s"
            },
            "country": {
                "code": "%s",
                "name": "%s"
            },
            "firstName": "%s",
            "lastName": "%s",
            "teamId": null,
            "bibNumber": 2992,
            "usatNumber": "1A2B3C4D5E",                         
            "profile": {
                "email": "%s",
                "address": "%s",
                "phone": "%s"
            },
            "bundleRegistration": false,
            "swag": [],
            "registrationQuestions": {
                "data": [
                    {
                        "questionId": 24098,
                        "questionOptionId": 76598,
                        "questionTitle": "Question Two",
                        "questionAnswer": "Yes"
                    },
                    {
                        "questionId": 24097,
                        "questionOptionId": null,
                        "questionTitle": "Question One",
                        "questionAnswer": "test"
                    },
                    {
                        "questionId": 24000,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact first name",
                        "questionAnswer": "John"
                    },
                    {
                        "questionId": 24001,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact last name",
                        "questionAnswer": "Deer"
                    },
                    {
                        "questionId": 24002,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact phone number",
                        "questionAnswer": "555-555-5555"
                    },
                    {
                        "questionId": 24003,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact email address",
                        "questionAnswer": "john_deer@example.com"
                    }
                ]
            },
            "emergencyContact": {
                "data": [
                    {
                        "questionId": 24000,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact first name",
                        "questionAnswer": "John",
                        "type": "contact_first_name"
                    },
                    {
                        "questionId": 24001,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact last name",
                        "questionAnswer": "Deer",
                        "type": "contact_last_name"
                    },
                    {
                        "questionId": 24002,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact phone number",
                        "questionAnswer": "555-555-5555",
                        "type": "contact_phone"
                    },
                    {
                        "questionId": 24003,
                        "questionOptionId": null,
                        "questionTitle": "Emergency contact email address",
                        "questionAnswer": "john_deer@example.com",
                        "type": "contact_email"
                    }
                ]
            },
            "registrationTotals": {
                "netRegistrationTotal": "100.00",
                "netSwagTotal": "0.00"
            }
        }`,
			string(rune(id)),
			randomID,
			faker.Time().Backward(duration1).Format(time.RFC3339),
			faker.Time().Backward(duration2).Format(time.RFC3339),
			faker.Name().Name(),
			faker.Date().Birthday(1, 60).Format("1939-11-20"),
			faker.Number().NumberInt(2),
			faker.Address().City(),
			faker.Address().StateAbbr(),
			faker.Address().State(),
			faker.Address().StateAbbr(),
			faker.Address().State(),
			faker.Name().FirstName(),
			faker.Name().LastName(),
			faker.Internet().Email(),
			faker.Address().String(),
			faker.PhoneNumber(),
		)

		events = fmt.Sprintf(events+"%s,", event)
	}
	events = strings.TrimSuffix(events, ",")

	empArray := fmt.Sprintf(`
{
	"data": [
		%s
	],
	"metadata":{
        "totalResults": %d
    }
}`, events, total)
	// Declared an empty interface of type Array
	var results struct {
		Data     []map[string]interface{} `json:"data"`
		Metadata interface{}              `json:"metadata"`
	}

	json.Unmarshal([]byte(empArray), &results)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(results)
}

func hand(c echo.Context) error {
	empArray := fmt.Sprintf(`[
		{
			"id": 1,
			"name": "%s",
			"department": "",
			"designation": "Director"
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager"
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead"
		}
	]`, faker.App().Name())
	// Declared an empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empArray), &results)
	fmt.Println(results)
	for key, result := range results {

		fmt.Println("Reading Value for Key :", key)
		//Reading each value by its key
		fmt.Println("Id :", result["id"],
			"- Name :", result["name"],
			"- Department :", result["department"],
			"- Designation :", result["designation"])
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(results)
}
