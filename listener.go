package autotest

// java -Dwebdriver.chrome.driver=vendor/chromedriver-mac64-2.42 -jar vendor/selenium-server-standalone-3.14.0.jar -port 53188

// antlr -Dlanguage=Go -o parser tl.g4
import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/scaumiao/autotest/utils"

	"github.com/golang/protobuf/ptypes"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/tebeka/selenium"
	"github.com/scaumiao/autotest/parser"
)

const (
	// DefaultWaitInterval is the default polling interval for selenium.Wait
	// function.
	DefaultWaitInterval = 100 * time.Millisecond

	// DefaultWaitTimeout is the default timeout for selenium.Wait function.
	DefaultWaitTimeout = 60 * time.Second
)

type ltListener struct {
	*parser.BasetlListener
	wd selenium.WebDriver

	stack        selenium.WebElement
	stacks       []selenium.WebElement
	currentStage int
	stages       Stages
	mux          *sync.RWMutex
	errBlock     bool
}

func NewListener(wd selenium.WebDriver) *ltListener {
	return &ltListener{wd: wd, mux: &sync.RWMutex{}}
}

// EnterErrorNode is called when a terminal node is visited.
func (l *ltListener) EnterErrorNode(node antlr.TerminalNode) {
	// fmt.Println("VisitTerminal", node)
}

// VisitTerminal is called when a terminal node is visited.
func (l *ltListener) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Println("VisitTerminal", node)
}

// VisitErrorNode is called when an error node is visited.
func (l *ltListener) VisitErrorNode(node antlr.ErrorNode) {
	// fmt.Println("VisitErrorNode", node)

}

func (l *ltListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	command := ctx.GetStart().GetText()
	line := ctx.GetStart().GetLine()
	tp := ctx.GetRuleIndex()
	switch tp {
	case 5:
		return
	}
	switch command {
	case "##": //如果是satges，不处理
		break
	default:
		a := NewAction()
		a.Command = string(command)
		a.Line = int32(line)
		a.StartAt = ptypes.TimestampNow()
		AddAction(l.stages[l.currentStage], a)
		fmt.Println("EnterEveryRule-", command)
	}

}

func (l *ltListener) ExitRule(ctx antlr.ParserRuleContext) {
	//转到最后一个浏览器tab
	getLastWindow(l.wd)
}

func (l *ltListener) Report(err interface{}) {
	last := len(l.stages[l.currentStage].Actions)
	if err != nil {
		l.stages[l.currentStage].State = 1
		l.stages[l.currentStage].Actions[last-1].Msg = fmt.Sprintf("%s", err)
		l.stages[l.currentStage].Actions[last-1].State = 1
	} else {
		l.stages[l.currentStage].Actions[last-1].Msg = "ok"
	}
	t, err := ptypes.Timestamp(l.stages[l.currentStage].Actions[last-1].StartAt)
	if err != nil {
		return
	}

	l.stages[l.currentStage].Actions[last-1].TimeElapsed = time.Since(t).Seconds()
	l.stages[l.currentStage].TimeElapsed += l.stages[l.currentStage].Actions[last-1].TimeElapsed
}

func (l *ltListener) stripQuotes(s string) string {
	if s == "" || !strings.Contains(s, "\"") {
		return s
	}
	s = strings.Replace(s, "\\\"", "\"", -1)
	return s[1 : len(s)-1]
}

const (
	base = 300
)

func (l *ltListener) ExitRunStages(c *parser.RunStagesContext) {
	fmt.Printf("ExitRunStages%v----\n", c.ID())
	//转到最后一个浏览器tab
	getLastWindow(l.wd)

	b, err := l.wd.Screenshot() //截图
	if err == nil {

		nb, err := utils.MakeThumbnail(b) //压缩大小比例
		if err != nil {
			return
		}
		l.stages[l.currentStage].Image = base64.StdEncoding.EncodeToString(nb)
	}
}

func (l *ltListener) EnterRunStages(c *parser.RunStagesContext) {
	id := fmt.Sprintf("Stages %s\n", c.ID())
	line := c.ID().GetSymbol().GetLine()
	l.stages = append(l.stages, NewStage(id, "", int32(line)))
	l.currentStage = len(l.stages) - 1
	fmt.Printf("EnterRunStages%v----\n", c.AllGexpression())
}

// ExitOPEN(ctx *OPENContext)
func (l *ltListener) ExitOpen(c *parser.OpenContext) {

	command := c.OPEN().GetText()
	url := l.stripQuotes(c.QUOTED_STRING().GetText())
	fmt.Printf("%v----%v\n", command, url)

	defer func() {
		l.Report(recover())
	}()
	if err := l.wd.Get(url); err != nil {
		panic(err)
		// panic(c.GetParser().GetCurrentToken().GetLine()) //, err)
	}
}

func (l *ltListener) ExitType(c *parser.TypeContext) {
	command := c.TYPE().GetText()

	str := l.stripQuotes(c.QUOTED_STRING().GetText())
	target := l.stripQuotes(c.Target().(*parser.TargetContext).QUOTED_STRING().GetText())

	fmt.Printf("%s----%s set %s\n", command, target, str)
	defer func() {
		l.Report(recover())
	}()

	findAndSendKey(l.wd, target, str)
}

func (l *ltListener) isStackClick(c *parser.ClickContext) bool {
	if c.QUOTED_STRING() == nil {
		return true
	}
	return false
}
func (l *ltListener) ExitClick(c *parser.ClickContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.CLICK()

	if l.isStackClick(c) {
		fmt.Printf("%v----\n", command)
		l.mux.Lock()
		if l.stack != nil {
			l.stack.Click()
			l.stack = nil
		}
		l.mux.Unlock()
		return
	}
	target := l.stripQuotes(c.QUOTED_STRING().GetText())
	fmt.Printf("%v----%v\n", command, target)

	findAndClick(l.wd, target)
}

func (l *ltListener) ExitFind(c *parser.FindContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.FIND()

	target := l.stripQuotes(c.QUOTED_STRING().GetText())
	fmt.Printf("%v----%v\n", command, target)
	elem := find(l.wd, target)
	l.mux.Lock()
	l.stack = elem
	l.mux.Unlock()
}

func (l *ltListener) ExitFindAll(c *parser.FindAllContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.FINDALL()
	target := l.stripQuotes(c.QUOTED_STRING().GetText())
	fmt.Printf("%v----%v\n", command, target)
	elems := findAll(l.wd, target)
	l.mux.Lock()
	l.stacks = elems
	l.mux.Unlock()
}

func (l *ltListener) ExitActiveElement(c *parser.ActiveElementContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.ACTIVEELEMENT()
	fmt.Printf("%v----\n", command)
	elem := activeElement(l.wd)
	l.mux.Lock()
	l.stack = elem
	l.mux.Unlock()
}

func (l *ltListener) ExitWait(c *parser.WaitContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.WAIT()
	fmt.Printf("%v----\n", command)
	i := l.hasNumber(c)
	interval := DefaultWaitInterval
	if i > 0 {
		interval = time.Duration(i * int(time.Second))
	}
	time.Sleep(interval)

}

func (l *ltListener) ExitRandom(c *parser.RandomContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.RANDOM()
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.stacks == nil || len(l.stacks) == 0 {
		return
	}
	target := random(0, len(l.stacks))
	l.stack = l.stacks[target]

	fmt.Printf("%v----%v\n", command, target)

}

func (l *ltListener) hasID(c interface{}) bool {
	switch v := c.(type) {
	case *parser.IframeContext:
		if v.ID() != nil {
			return true
		}
	}

	return false
}
func (l *ltListener) hasNumber(c interface{}) int {
	switch v := c.(type) {
	case *parser.WaitContext:
		if v.NUMBER() != nil {
			i, _ := strconv.Atoi(v.NUMBER().GetText())
			return i
		}
	}
	return -1
}
func (l *ltListener) ExitIframe(c *parser.IframeContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.IFRAME()
	target := ""
	if l.hasID(c) {
		target = l.stripQuotes(c.ID().GetText())
	} else {
		target = l.stripQuotes(c.QUOTED_STRING().GetText())
	}

	fmt.Printf("%v----%v\n", command, target)
	switchFrame(l.wd, target)

}

func (l *ltListener) ExitText(c *parser.TextContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.TEXT()
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.stack == nil {
		return
	}
	str := text(l.stack)
	l.stack = nil

	fmt.Printf("%v----%v\n", command, str)

}

func (l *ltListener) ExitSendKey(c *parser.SendKeyContext) {
	defer func() {
		l.Report(recover())
	}()
	command := c.SENDKEY()
	target := l.stripQuotes(c.ID().GetText())
	fmt.Printf("%v----%v\n", command, target)

	switch target {
	case "enter":
		target = selenium.EnterKey
	}
	var elem selenium.WebElement
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.stack == nil {
		elem = activeElement(l.wd)
		l.stack = elem
	} else {
		elem = l.stack
	}
	err := elem.SendKeys(target)
	if err != nil {
		panic(err)
	}

}

func findAndSendKey(wd selenium.WebDriver, where, set string) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, where)
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(set)
	if err != nil {
		panic(err)
	}
}

func findAndClick(wd selenium.WebDriver, where string) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, where)
	if err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.Click()
	if err != nil {
		panic(err)
	}
}

func find(wd selenium.WebDriver, where string) selenium.WebElement {
	err := wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		elem, _ := wd.FindElement(selenium.ByCSSSelector, where)
		if elem != nil {
			return true, nil
		}
		return false, nil
	})

	if err != nil {
		panic(err)
	}

	elem, err := wd.FindElement(selenium.ByCSSSelector, where)
	if err != nil {
		panic(err)
	}
	return elem
}

func findAll(wd selenium.WebDriver, where string) []selenium.WebElement {
	elems, err := wd.FindElements(selenium.ByCSSSelector, where)
	if err != nil {
		panic(err)
	}
	return elems
}

func activeElement(wd selenium.WebDriver) selenium.WebElement {
	elem, err := wd.ActiveElement()
	if err != nil {
		panic(err)
	}
	return elem
}

func switchFrame(wd selenium.WebDriver, id string) {
	err := wd.SwitchFrame(id)
	if err != nil {
		panic(err)
	}
}

func text(elem selenium.WebElement) (output string) {
	var err error
	for {
		output, err = elem.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	return output
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//getLastWindow 访问最后一个浏览器窗口
func getLastWindow(wd selenium.WebDriver) {
	//转到最后一个浏览器tab
	windows, err := wd.WindowHandles()
	if err != nil {
		return
	}
	current, err := wd.CurrentWindowHandle()
	if err != nil {
		return
	}

	last := windows[len(windows)-1]

	if last != current {
		wd.SwitchWindow(last)
	}
}
