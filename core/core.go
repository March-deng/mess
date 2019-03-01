package core

import (
	"encoding/json"
	"github.com/xenolf/lego/log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

type Core struct {
	CreateTime time.Time `json:"createTime"`
}

func (c *Core) Start() {

}

func (c *Core) APIHandler() http.Handler {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Group("/")
	r.GET("/rest/v1", c.GetStatus)
	return r
}

type status struct {
	Status     string `json:"status" bson:"status"`
	StatusCode int    `json:"status_code bson:"status_code"`
	Message    string `json:"message" bson:"message"`
}

func (c *Core) GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, status{
		Status:     "verical",
		StatusCode: 1,
		Message:    "all service good",
	})
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Desc string `json:"description"`
}

func GetFileData() Person {
	var p Person
	cmd := exec.Command("cat", "person.json")

	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(out)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	cmd.Wait()

	return p
}

func get() {
	log.Println("dengcong")
}
