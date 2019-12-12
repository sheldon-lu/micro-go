package handler

import (
    `context`
    `github.com/gin-gonic/gin`
    `github.com/micro/go-micro/client`
    user "micro-go/user/proto/user"
    "github.com/micro/go-micro/util/log"
    "encoding/json"
)

type User struct {
}

var (
    cli1 user.UserService
)


// setup Greeter Server Client Init
func Init() {
    cli1 = user.NewUserService("lu.micro.srv.user", client.DefaultClient)
}

func (us *User) Login(c *gin.Context) {
    // GET PostForm
    // POST json body // FUNC
    mapResult := make(map[string]interface{})
    buf := make([]byte, 1024)
    // n => json.Unmarshal(buf[0: `>n<` ])
    n, _ := c.Request.Body.Read(buf)
    // ==> json_string > map
    err := json.Unmarshal(buf[0:n], &mapResult)
    if err != nil {
        log.Fatal(err)
    }

    phone := mapResult["phone"].(string)
    password := mapResult["password"].(string)

    resp, err := cli1.Login(context.TODO(), &user.LoginRequest{
        Phone: phone,
        Password: password,
    })

    if err != nil {
        log.Fatal(err)
        c.JSON(505, err)
    }

    c.JSON(200, resp)
}

func (us *User) Register(c *gin.Context) {

}

func (us *User) UpdatePassword(c *gin.Context) {

}

func (us *User) List(c *gin.Context) {
    log.Debug("Received Say.Anything API request")
    c.JSON(200, map[string]string{
        "message": "Hi",
    })
}
