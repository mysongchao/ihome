package controllers

import (
	"github.com/astaxie/beego"
	//"ihome/utils"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"github.com/astaxie/beego/cache"
	"time"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"ihome/utils"
	"encoding/json"
	"reflect"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"strconv"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"ihome/models"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) Retdata(resp interface{}){
	c.Data["json"] = resp
	c.ServeJSON()
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *UserController) GetImage() {

	// 打印调用的函数
	beego.Info("---------------- GET  /api/v1.0/imagecode/* GetImage() ------------------")
	// 创建返回空间
	//resp := make(map[string]interface{})


	//resp["errno"] = utils.RECODE_SESSIONERR
	//resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//延迟调用发送给前端json数据
	//defer c.Retdata(resp)

	/* 获取 前端uu id */
	uuid := c.Ctx.Input.Param(":splat")
	beego.Info(uuid)

	/*生成随机数杨郑码图片*/
	cap := captcha.New()

					// 需将comic.ttf 源文件考到本地项目
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	//设置图片的大小
	cap.SetSize(91, 41)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	//SetFrontColor(colors ...color.Color)  这两个颜色设置的函数属于不定参函数
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//生成图片 返回图片和 字符串(图片内容的文本形式)
	img, str := cap.Create(4, captcha.NUM)
	beego.Info("验证码："+str)


	/*将uuid与 随机数验证码对应的存储在redis缓存中*/
	//初始化缓存全局变量的对象																			  // 6039 redis 默认端口
	bm, err := cache.NewCache("redis", `{"key":"ihome","conn":"127.0.0.1:6379","dbNum":"0"}`)
	if err != nil  {
		beego.Info("GetImage() cache.NewCache err",err)
	}

	// redis 存储操作  插入数据库中
	bm.Put(uuid,str,600*time.Second) // 进行一小时缓存-

	/* 向前端返回杨怎么图片 */
	//将图片发送给前端的 直接发送图片
	png.Encode(c.Ctx.ResponseWriter,img)

	return
}



// 短信验证马
func (c *UserController) Getsmscode() {

	// 打印调用的函数
	beego.Info("---------------- GET  /api/v1.0/smscode/* Getsmscode() ------------------")
	//创建返回空间
	resp := make(map[string]interface{})


	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//延迟调用发送给前端json数据
	defer c.Retdata(resp)

	// 获取前端发送过来的id  正则获取
	uuid := c.Ctx.Input.Param(":id")
	beego.Info(uuid)

	// 获取url发送过来的参数
	var text string
	c.Ctx.Input.Bind(&text,"text")

	var id string
	c.Ctx.Input.Bind(&id,"id")

	beego.Info(text,id)



	redis_config_map := map[string]string{
		"key":"ihome",
		"conn":utils.G_mysql_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}


	redis_config  ,_ := json.Marshal(redis_config_map)
	beego.Info("jsonsss:=",string(redis_config))

	pbm, err := cache.NewCache("redis",string(redis_config))
	if err != nil  {
		beego.Info(" () cache.NewCache err1",err)
	}

	value := pbm.Get(id)
	if value !=nil{
		beego.Info("value := pbm.Get(id) err",value)
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	}

	beego.Info(value,reflect.TypeOf(value))
	value_str ,_:= redis.String(value,nil)
	beego.Info("后："+value_str,reflect.TypeOf(value_str))
	// 数据对比
	if text != value_str {
		beego.Info("验证马错误")
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	}

	//手机号验证
	//myreg := regexp.MustCompile(`0?()`)
	//bo := myreg.MatchString(id) // 判断字符窗是否与正则匹配
	//if bo ==false {
	//	beego.Info("手机号码验证错误")
	//	resp["errno"] = utils.RECODE_DATAERR
	//	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//}

	// 连接 缓存数据库 获取缓存信息，验证是否正确

	// 通过已有的验证马捷克 模拟短信发送

	// 将短信验证马 缓存
	/*通过已有的短信验证码的接口 模拟发送短信*/
	//type Values map[string][]string
	v := url.Values{}  //创建1个url.values map
	//格式化当前的时间
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	beego.Info(_now)
	_account := "C10921244"  //账户名 需要花钱购买的
	_password := "da3018614650fa96137bb61dd71e85d8" //查看密码请登录用户中心->验证码、通知短信->帐户及签名设置->APIKEY

	_mobile := string(id)  //手机号  通过id进行赋值
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //生成随机数

	sms_code := r.Intn(9999)
	beego.Info("短信验证码是",sms_code)
	//拼接短信发送的内容
	_content := "您的验证码是：" + strconv.Itoa(sms_code) + "。请不要把验证码泄露给其他人。"
	//添加url的内容
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)

	//body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	//client := &http.Client{}
	//req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)
	//
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	////fmt.Printf("%+v\n", req) //看下发送的结构
	//beego.Info(req)
	//resp1, err := client.Do(req) //发送
	//defer resp1.Body.Close()     //一定要关闭resp.Body
	//data, _ := ioutil.ReadAll(resp1.Body)
	//fmt.Println(string(data), err)

	/*将短信验证码存入缓存数据库 */
	pbm.Put("smscode",strconv.Itoa(sms_code) , 600 *time.Second)

	/*成功返回ok */
	return

}

// 用户注册
func (c *UserController) Postuserret() {

	// 打印调用的函数
	beego.Info("---------------- GET  /api/v1.0/users  Postuserret() ------------------")
	//创建返回空间
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//延迟调用发送给前端json数据
	defer c.Retdata(resp)

	// 验证 所自段

	//  获取前端发送的数据
	//var  Requestmap = make (map[string]interface{})
	//json.Unmarshal(c.Ctx.Input.RequestBody.)

}

// 用户注册
func (c *UserController) Postlogin() {

	// 打印调用的函数
	beego.Info("---------------- GET  /api/v1.0/users  Postuserret() ------------------")
	//创建返回空间
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//延迟调用发送给前端json数据
	defer c.Retdata(resp)
	/* 获得用户注册信息*/
	var Requestmap = make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &Requestmap)

	for key, value := range Requestmap {
		beego.Info(key, value)
	}
	beego.Info( Requestmap["mobile"])
	beego.Info( Requestmap["password"])
	beego.Info( Requestmap["sms_code"])

	/*校验信息准确信*/
	if Requestmap["mobile"] == ""|| Requestmap["password"] == "" || Requestmap["sms_code"] == ""{
		beego.Info("注册数据为空")
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	}

	/*验证短信验证码 */
	//构建连接缓存的数据
	redis_config_map := map[string]string{
		"key":"ihome",
		//"conn":"127.0.0.1:6379",
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	beego.Info(redis_config_map)
	redis_config ,_:=json.Marshal(redis_config_map)
	beego.Info( string(redis_config) )
	//连接redis数据库 创建句柄
	bm, err := cache.NewCache("redis", string(redis_config) )
	if err !=nil{
		beego.Info("GetImage()   cache.NewCache err ",err)
	}
	//获取我们存在缓存数据库中的短信验证码
	value :=bm.Get("smscode")
	if  value == nil{

		beego.Info("Postuserret()  bm.Get(uuid) err  ",value)
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
		return
	}
	beego.Info(value,reflect.TypeOf(value))

	value_str ,_ :=redis.String(value,nil)
	/*
	第一个参数填写通过bm.get 获取到的返回值
	第二个一般情况下写nil
	*/
	beego.Info(value_str,reflect.TypeOf(value_str))
	//短信验证码对比
	if  Requestmap["sms_code"].(string) != value_str{
		beego.Info("短信验证码错误 错误 ")
		resp["errno"] = utils.RECODE_PWDERR
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
		return
	}

	/*将用户信息存储在 mysql 中 */

	//beego.Info( Requestmap["mobile"])
	//beego.Info( Requestmap["password"])
	//beego.Info( Requestmap["sms_code"])
	//创建1个mysql对象
	user:= models.User{}
	user.Name = Requestmap["mobile"].(string)
	user.Mobile = Requestmap["mobile"].(string)
	//正常情况下我们需要吧 password 转成md5 或sha256 等等密文格式   为了调试所以进行直接赋值
	//正常情况下密码的相关对比都要转成密文进行操作
	user.Password_hash =Requestmap["password"].(string)
	//操作数据库
	o :=orm.NewOrm()
	id ,err := o.Insert(&user)
	if err != nil{
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
		return
	}
	beego.Info(id)

	/*添加session字段 */
	//SetSession(name string, value interface{})
	c.SetSession("name", Requestmap["mobile"].(string))
	c.SetSession("user_id",id)
	c.SetSession("mobile", Requestmap["mobile"].(string))

	/*进行返回*/

	return
}

// 用户fdfs 上传文件
func (c *UserController) Postxxx() {

	// 打印调用的函数
	beego.Info("---------------- GET  /api/v1.0/users  Postuserret() ------------------")
	//创建返回空间
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//延迟调用发送给前端json数据
	defer c.Retdata(resp)

	// 获取前段发过来的文件数据
	file, hander,err := c.GetFile("avatar")
	if err != nil {
		beego.Info("c.GetFile(avatar) err",err)
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
		return
	}
	beego.Info(file,hander)
	beego.Info("文件大小：",hander.Size)
	beego.Info("文件名：",hander.Filename)



	//filebuffer := make([]byte ,hander.Size)
	//if err != nil {
	//	beego.Info("c.GetFile(avatar) err",err)
	//	resp["errno"] = utils.RECODE_OK
	//	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//	return
	//}
	//
	//
	//// 获取文件后槜名
	//beego.Info("文件后缀名：",path.Ext(hander.Filename))
	//
	//// 存储到fastdfs中并获取URL
	//models.UploadByFilename("",path.Ext(hander.Filename))
	//if err != nil {
	//	beego.Info("c.GetFile(avatar) err",err)
	//	resp["errno"] = utils.RECODE_OK
	//	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//	return
	//}



}