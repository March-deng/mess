package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/satori/go.uuid"

	"github.com/Jeffail/gabs"
	nats "github.com/nats-io/go-nats"

	"github.com/go-redis/redis"
)

func TestTransport(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func TestTime(t *testing.T) {
	meal := "18:03"
	abs := time.Now().Add(time.Hour).Unix()
	nowaday := time.Unix(abs, 0).Format("2006-01-02")
	//取得年月日
	log.Println("nowaday:", nowaday)
	mealabs := nowaday + "T" + meal
	log.Println("meal abs:", mealabs)
	ti, err := time.Parse("2006-01-02T15:04", mealabs)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(ti)
}

func TestJsonParseToMap(t *testing.T) {
	container, err := gabs.ParseJSON([]byte(`[
		{
			"id": "1111",
			"code": "2222"
		},
		{
			"id": "3333",
			"code": "4444"
		}
	]`))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(container.Data())
	co, err := container.Children()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range co {
		log.Println(c.String())
	}
}

func TestNatsClient(t *testing.T) {
	nc, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	nc.Publish("log.example", []byte("-----"))
}

func TestMapInit(t *testing.T) {
	type Item struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
	items := []Item{
		Item{
			Name: "dengcong",
			ID:   "111111",
		},
		Item{
			Name: "yangzhiqiang",
			ID:   "222222",
		},
		Item{
			Name: "xuepengcheng",
			ID:   "333333",
		},
	}
	itemMap := make(map[string]Item)
	var item Item
	for _, i := range items {
		// itemMap[i.Name] = i
		item = i
		itemMap[item.Name] = item
	}
	log.Println(itemMap)
}

type Reserve struct {
	ID   string `json:"id"`
	Date string `json:"date"`
	Time string `json:""`
}

// func TestStormRange(t *testing.T) {
// 	s, err := storm.Open("./test.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer s.Close()
// 	rs := []Reserve{
// 		Reserve{
// 			ID: "111111",
// 			Date: "2019-"
// 		}
// 	}
// }

func TestFormatDate(t *testing.T) {
	dayMinute := time.Unix(1559307730000/1000, 0).Format("2006-01-02 15:04")
	log.Println(dayMinute)
}

func TestTimeStamp(t *testing.T) {
	uid, _ := uuid.NewV4()
	log.Println(uid.String())
	ti, _ := time.ParseInLocation("2006-01-02 15:04", "2019-06-03 15:02", time.Local)

	log.Println(ti)
	log.Println(ti.Unix() * 1000)
}

func TestSocketConnection(t *testing.T) {
	conn, err := net.DialTimeout("tcp", "192.168.1.108:9100", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	log.Println(conn.Write([]byte("ping")))
}

func TestArrayAppend(t *testing.T) {
	type Item struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
	var item Item
	var items []Item
	for i := 0; i < 5; i++ {
		uid, _ := uuid.NewV1()
		item = Item{
			Name: "fault",
			ID:   uid.String(),
		}
		items = append(items, item)
	}
	log.Println(items)
}

func TestBase64(t *testing.T) {
	input := `G0AbAAAbVQAdIREbcgAbRQEgINbY06G1pSAgG0UAHSEAHSEAG3IAICAgICAgICAgICAgICAgICAgICAoKSAgICAgICAgICAgICAgICAgICAgHSEAG3IAICAgICAgICAgICAgICAgICAgtdjWtyAgICAgICAgICAgICAgICAgIB0hABtyAD09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PR0hABtyAL270tfI1cbaOiAbcgAyMDE5LTA2LTAzG3IAvbvS18qxvOQ6ICAgG3IAIDExOjU2OjAzIB0hABtyANOq0rXI1cbaOiAbcgAyMDE5LTA1LTI3G3IAvbvS17WlusUgIBtyACAgICAxNCAgICAdIQAbcgDXwLrFOiAgICAgG3IAICAgIEMyICAgIBtyAL7Nss3Iy8r9OiAgIBtyACAgICAgMyAgICAdIQAbcgDK1dL41LE6ICAgG3IAICAgc2RmMSAgIBtyAL270tfA4NDNOiAgIBtyACAgz/rK27WlICAdIQAbcgAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0dIQAbcgDJzMa3ICAgICAgICAgICAgICAbcgAgICAgICC1pbzbG3IAICDK/cG/G3IAICAgIL3wtu4dIQAbcgAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0dIQAbcgC1sLjiICAgICAgICAgICAgICAgICAgG3IAICAgMTAwG3IAICAgIHgxG3IAICAgICAxMDAdIQAbcgC/ydGhzNeyzSAgICAgICAgICAgICAgG3IAICAgMjAwG3IAICAgIHgxG3IAICAgICAgIDIdIQAbcgAgIMWjxcUgICAgICAgICAgICAbcgAgICAgICAgICAgICAgICAxHSEAG3IALS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tHSEAG3IAt/7O8bfROiAgICAgICAgICAgICAgIBtyACAgICAgICAgICAgICAgICAgICAwHSEAG3IALS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tHSEAG3IA1fu1pTE1JSAgICAgICAgICAgICAgIBtyACAgICAgICAgICAgICAgIC0xNS4zHSEAG3IA1du/27rPvMY6ICAgICAgICAgICAgIBtyACAgICAgICAgICAgICAgIC0xNS4zHSEAG3IALS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tHSEAG3IAxKjB4zogICAgICAgICAgICAgICAgIBtyACAgICAgICAgICAgICAgICAgICAwHSEAG3IALS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tHSEAHSERG3IA19y8xlJNQjobcgAgICAgIDg2LjcdIQAbcgAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0dIQAbcgDP1r3wICAgICAgICAgICAgICAbcgAgICAgICAgICAgG3IAICAgICAgG3IAICAgIDg2LjcdIQAbcgAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0dIQAdIQAbcgC9u9LXtqm1pbrFOiAgG3IAICAgICAgICAgICAgICAgICAgICAgICAgICAgIB0hAB0hABtyALe/vOQv1cu6xVJPT00vQUNDLik6ICAgG3IAICAgICAgICAgICAgICAdIQAdIQAdIQAbcgDQ1cP7L7mry75HVUVTVC9DTy4pOiAgICAgICAgICAgICAgICAgICAgHSEAHSEAHSEAG3IAX19fX19fX19fX19fX19fX19fX19fX19fX19fX19fX19fX18gICAgICAgHSEAG3IAPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09HSEAG3IAICAgtMvK1b7dysfE+r3h1cu1xM6o0ru21NXLxr7WpMfrsaPB9CAgIB0hABtyACAgICAgIFBsZWFzZSByZXRhaW4gY2hhcmdlIGRvY2tldCBhcyAgICAgIBtyACAgICAgICAgIHlvdXIgc3VwcG9ydGluZyBkb2N1bWVudC4gICAgICAgIB0hABtyACAgICAgICAgICAgILbg0Lu73bnLICDH69TZueLB2SAgICAgICAgICAgIB0hABtyACAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIB0hABtyACAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIB0hABtyACAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIB0hABtyACAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIB1WAA==`
	body, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s", string(body))
}
