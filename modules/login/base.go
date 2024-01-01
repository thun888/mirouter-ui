package login

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	password   string
	key        string
	ip         string
	token      string
	routername string
	hardware   string
)

func createNonce() string {
	typeVar := 0
	deviceID := "" //无效参数
	timeVar := int(time.Now().Unix())
	randomVar := rand.Intn(10000)
	return fmt.Sprintf("%d_%s_%d_%d", typeVar, deviceID, timeVar, randomVar)
}

func hashPassword(pwd string, nonce string, key string) string {
	pwdKey := pwd + key
	pwdKeyHash := sha1.New()
	pwdKeyHash.Write([]byte(pwdKey))
	pwdKeyHashStr := fmt.Sprintf("%x", pwdKeyHash.Sum(nil))

	noncePwdKey := nonce + pwdKeyHashStr
	noncePwdKeyHash := sha1.New()
	noncePwdKeyHash.Write([]byte(noncePwdKey))
	noncePwdKeyHashStr := fmt.Sprintf("%x", noncePwdKeyHash.Sum(nil))

	return noncePwdKeyHashStr
}
func newhashPassword(pwd string, nonce string, key string) string {
	pwdKey := pwd + key
	pwdKeyHash := sha256.Sum256([]byte(pwdKey))
	pwdKeyHashStr := hex.EncodeToString(pwdKeyHash[:])

	noncePwdKey := nonce + pwdKeyHashStr
	noncePwdKeyHash := sha256.Sum256([]byte(noncePwdKey))
	noncePwdKeyHashStr := hex.EncodeToString(noncePwdKeyHash[:])

	return noncePwdKeyHashStr
}
func getrouterinfo(ip string) (bool, string, string) {

	// 发送 GET 请求
	ourl := fmt.Sprintf("http://%s/cgi-bin/luci/api/xqsystem/init_info", ip)
	response, err := http.Get(ourl)
	if err != nil {
		return false, "", ""
	}
	defer response.Body.Close()
	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, "", ""
	}

	// 解析 JSON
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, "", ""
	}
	//提取routername
	routername = data["routername"].(string)
	hardware = data["hardware"].(string)
	logrus.Debug("路由器型号为:" + hardware)
	logrus.Debug("路由器名称为:" + routername)
	// 检查 newEncryptMode
	newEncryptMode, ok := data["newEncryptMode"].(float64)
	if !ok {
		logrus.Debug("使用旧加密模式")
		return false, routername, hardware
	}

	if newEncryptMode != 0 {
		logrus.Debug("使用新加密模式")
		logrus.Info("当前路由器可能无法正常获取某些数据！")
		return true, routername, hardware
	}
	return false, routername, hardware
}

func CheckRouterAvailability(ip string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	_, err := client.Get("http://" + ip)
	if err != nil {
		logrus.Info("路由器" + ip + "不可用，请检查配置或路由器状态")
		return false
	}

	return true
}
func GetToken(password string, key string, ip string) (string, string, string) {
	logrus.Debug("检查路由器可用性...")
	if !CheckRouterAvailability(ip) {
		return "", "路由器不可用", ""
	}
	logrus.Debug("获取路由器信息...")
	newEncryptMode, routername, hardware := getrouterinfo(ip)
	logrus.Info("更新token...")
	nonce := createNonce()
	var hashedPassword string

	if newEncryptMode {
		hashedPassword = newhashPassword(password, nonce, key)
	} else {
		hashedPassword = hashPassword(password, nonce, key)
	}

	ourl := fmt.Sprintf("http://%s/cgi-bin/luci/api/xqsystem/login", ip)
	params := url.Values{}
	params.Set("username", "admin")
	params.Set("password", hashedPassword)
	params.Set("logtype", "2")
	params.Set("nonce", nonce)

	resp, err := http.PostForm(ourl, params)
	if err != nil {
		logrus.Info("登录失败，请检查配置或路由器状态")
		logrus.Info(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	var code int
	if result["code"] != nil {
		code = int(result["code"].(float64))
	} else {
		logrus.Info("路由器登录请求返回值为空！请检查配置")
	}

	if code == 0 {
		logrus.Debug("当前token为:" + fmt.Sprint(result["token"]))
		token = result["token"].(string)
	} else {
		logrus.Info("登录失败，请检查配置，以下为返回输出:")
		logrus.Info(string(body))
		logrus.Info("5秒后退出程序")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	return token, routername, hardware
}
