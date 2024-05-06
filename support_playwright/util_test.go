package support_playwright

import (
	"bcd-util/util"
	"bytes"
	"encoding/base64"
	"github.com/playwright-community/playwright-go"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	//err := playwright.Install()
	//if err != nil {
	//	util.Log.Errorf("%+v", err)
	//	return
	//}
	pw, err := playwright.Run()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	page, err := browser.NewPage()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}

	//进入网站
	_, err = page.Goto("http://oa.szsoling.com:8088/wui/index.html#/?logintype=1&_key=6lvlz5")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//输入用户名
	locator := page.Locator("#loginid")
	err = locator.Fill("changdabai")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//输入密码
	locator = page.Locator("#userpassword")
	err = locator.Fill("5221527")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}

	count := 0
	for {
		count++
		//主动刷新验证码、避免验证码请求已经完成
		locator = page.Locator("#validatecode ~ div")
		err = locator.Click()
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		//等待获取验证码
		response, err := page.ExpectResponse("http://oa.szsoling.com:8088/weaver/weaver.file.MakeValidateCode*", func() error {
			return nil
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		codeImageData, err := response.Body()
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		//验证码识别
		encodeToString := base64.StdEncoding.EncodeToString(codeImageData)
		resp, err := http.Post("http://127.0.0.1:5000/code", "text/plain", bytes.NewBuffer([]byte(encodeToString)))
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		codeRes, err := io.ReadAll(resp.Body)
		code := string(codeRes)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		util.Log.Infof("%s", code)
		//填入验证码
		locator = page.Locator("#validatecode")
		err = locator.Fill(code)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		//登陆
		locator = page.Locator(".loginBtn")
		err = locator.Click()
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		//等待登陆
		response, err = page.ExpectResponse("http://oa.szsoling.com:8088/api/hrm/login/checkLogin*", func() error {
			return nil
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		loginData, err := response.Body()
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		loginSucceed := false
		util.Log.Infof("%s", string(loginData))
		parseBytes := gjson.ParseBytes(loginData)
		if parseBytes.Get("msgcode").Int() == 0 {
			loginSucceed = true
		}

		//判断是否登陆成功
		if loginSucceed {
			break
		} else {
			if count == 3 {
				util.Log.Info("尝试3次、登陆失败、退出操作")
				return
			}
		}
	}

	//点击门户
	locator = page.GetByText("门户")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//点击流程
	locator = page.GetByText("流程")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//等待页面加载完成
	err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{State: playwright.LoadStateDomcontentloaded})
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//展开左侧菜单
	locator = page.Locator("#portal-intro2 > i")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//点击新建流程
	locator = page.GetByText("新建流程")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//等待界面加载
	err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{State: playwright.LoadStateLoad})
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//点击请假申请单
	locator = page.GetByText("8.1.14 请假申请单")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//等待新页面打开
	evt, err := page.WaitForEvent("popup")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	newPage := evt.(playwright.Page)
	//等待所有网络请求结束
	err = newPage.WaitForLoadState(playwright.PageWaitForLoadStateOptions{State: playwright.LoadStateNetworkidle})
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//填写请假原因
	locator = newPage.Locator("#field34182")
	err = locator.Fill("调休(调试、不要批准)")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//选择请假类型
	locator = newPage.Locator("#weaSelect_1")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	locator = newPage.GetByTitle("调休")
	err = locator.Focus()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}

	//选择开始日期
	locator = newPage.Locator(".field26683_swapDiv > div > span")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	all, err := newPage.GetByPlaceholder("请选择日期").All()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	locator = all[0]
	err = locator.Fill("2024-04-26")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	err = locator.Press("Enter")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//选择结束日期
	locator = newPage.Locator(".field26684_swapDiv > div > span")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	all, err = newPage.GetByPlaceholder("请选择日期").All()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	locator = all[1]
	err = locator.Fill("2024-04-26")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	err = locator.Press("Enter")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//选择开始时间
	locator = newPage.Locator(".field26685_swapDiv > span > span")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	all, err = newPage.GetByPlaceholder("请选择时间").All()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	locator = all[0]
	err = locator.Fill("13:00")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	err = locator.Press("Enter")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//选择结束时间
	locator = newPage.Locator(".field26686_swapDiv > span > span")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	all, err = newPage.GetByPlaceholder("请选择时间").All()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	locator = all[1]
	err = locator.Fill("18:00")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	err = locator.Press("Enter")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//提交
	locator = newPage.Locator("button[title='提交']")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//返回流程页面查看
	locator = page.GetByTitle("我的请求")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//点击进入详情页面
	locator = page.Locator("div[class='ant-table-body'] > table > tbody > tr:first-child > td:nth-child(2) > span:nth-child(2) > a")
	err = locator.Click()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//等待新页面打开
	evt, err = page.WaitForEvent("popup")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	newPage = evt.(playwright.Page)
	//等待所有网络请求结束
	err = newPage.WaitForLoadState(playwright.PageWaitForLoadStateOptions{State: playwright.LoadStateNetworkidle})
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	//截图
	screenshot, err := newPage.Screenshot(playwright.PageScreenshotOptions{FullPage: playwright.Bool(true)})
	err = os.WriteFile("res.png", screenshot, os.ModePerm)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}

	time.Sleep(60 * time.Minute)
	err = browser.Close()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	err = pw.Stop()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
}
