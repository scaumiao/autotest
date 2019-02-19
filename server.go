package autotest

import (
	"fmt"

	"github.com/tebeka/selenium/log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/scaumiao/autotest/parser"
	"github.com/scaumiao/autotest/proto/stages"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	// These paths will be different on your system.
	// seleniumPath = "vendor/selenium-server-standalone-3.14.0.jar"
	seleniumPath     = ""
	chrom2DriverPath = "vendor/chromedriver-mac64-2.42"
	chromebinary     = "vendor/chrome-mac/Chromium.app/Contents/MacOS/Chromium"
	port             = 53188
)

//Server ...
type Server struct {
	wd    selenium.WebDriver
	close func() error
}

func NewServer() *Server {
	s := Server{}
	s.Start()
	return &s
}

func (s *Server) GetWd() selenium.WebDriver {
	return s.wd
}

func (s *Server) GetReport() string {
	return "report"
}

func (s *Server) Start() {
	s.wd, s.close = newRemote()

}

func (s *Server) Stop() error {
	return s.close()
}

func (s *Server) RunTest(rawtl string) []*stages.Stage {

	// Setup the input
	is := antlr.NewInputStream(rawtl)

	// Create the Lexer
	lexer := parser.NewtlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewtlParser(stream)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(p.GetParserRuleContext())
		}
	}()

	ls := NewListener(s.wd) //Create listener

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(ls, p.Start())

	return ls.stages
}

func newRemote() (selenium.WebDriver, func() error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  chromebinary,
		Args: []string{
			"--whitelisted-ips",
			// "--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	//设置日志
	caps.SetLogLevel(log.Browser, log.All)
	caps.SetLogLevel(log.Performance, log.All)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	return wd, wd.Quit
}
