package handler

import (
    `context`
    `fmt`
    `github.com/gin-gonic/gin`
    `github.com/micro/go-micro/client`
    user "micro-go/user/proto/user"
    "github.com/micro/go-micro/util/log"
    "encoding/json"
)

type Users struct {
}

var (
    cli1 user.UserService
)


// setup Greeter Server Client Init
func Init() {
    cli1 = user.NewUserService("lu.micro.srv.user", client.DefaultClient)
}

func (us *Users) Login(c *gin.Context) {
    // GET PostForm
    // POST json body // FUNC
    LoginMap := make(map[string]interface{})
    buf := make([]byte, 1024)
    // n => json.Unmarshal(buf[0: `>n<` ])
    n, _ := c.Request.Body.Read(buf)
    // ==> json_string > map
    err := json.Unmarshal(buf[0:n], &LoginMap)
    if err != nil {
        log.Fatal(err)
    }

    // names := LoginMap["name"].(string)
    phone := LoginMap["phone"].(string)
    password := LoginMap["password"].(string)

    resp, err := cli1.Login(context.TODO(), &user.LoginRequest{
        Phone: phone,
        Password: password,
    })

    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

func (us *Users) Register(c *gin.Context) {
    registerMap := make(map[string]map[string]interface{})
    buf := make([]byte, 1024)
    n,_ := c.Request.Body.Read(buf)
    err := json.Unmarshal(buf[0:n], &registerMap)
    if err != nil {
        log.Fatal(err)
    }

    username := registerMap["user"]["name"].(string)
    phone := registerMap["user"]["phone"].(string)
    password := registerMap["user"]["password"].(string)

    fmt.Println(username, phone, password)


    resp, err := cli1.Register(context.TODO(), &user.RegisterRequest{
        // type => * && &   ====  指针地址值
        User: &user.Userinfo{
            Name: username,
            Phone: phone,
            Password: password,
        },
    })
    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

func (us *Users) UpdatePassword(c *gin.Context) {
    // 1、get post_json_body
    // 2、get string []byte[0:n] => map ==> json.unmarshal
    updateMap := make(map[string]interface{})
    buf := make([]byte, 1024)
    n, _ := c.Request.Body.Read(buf)
    err := json.Unmarshal(buf[0:n], &updateMap)
    if err != nil {
        log.Fatal(err)
    }

    uids := updateMap["uid"].(uint32)
    oldpassword := updateMap["oldpassword"].(string)
    newpassword := updateMap["newpassword"].(string)

    resp, err := cli1.UpdatePassword(context.TODO(), &user.UpdatePasswordRequest{
        OldPassword: oldpassword,
        NewPassword: newpassword,
        Uid: uids,
    })
    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

func (us *Users) List(c *gin.Context) {
    log.Debug("Received Say.Anything API request")
    c.JSON(200, map[string]string{
        "message": "Hi",
    })
}

func (us *Users) DeleUser(c *gin.Context) {

}
