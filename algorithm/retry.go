package algorithm

import (
    "errors"
    "log"
    "testing"
    "time"
)

// 自定义Stop错误
type Stop struct {
    error
}
func NewStop(msg string) Stop {
    return Stop{errors.New(msg)}
}

// 短信
type Sms struct {
    Title /*标题*/, From /*发送者*/, To /*接收者*/, Content /*消息内容*/ string
}

// 模拟发短信
func (p *Sms) Send() error {
    if (time.Now().Second())%7 == 0 {
        log.Printf("【%s】向【%s】发送了一条短信:「%s」%s\n", p.From, p.To, p.Title, p.Content)
        return nil
    }
    //
    if (time.Now().Second())%9 == 0 {
        return NewStop("已取消重试")
    }
    return errors.New("( ¯▽¯；) 网络异常( ¯▽¯；)  ")
}

// 发送短信(失败重试)
//  count: 尝试发送次数
//  interval: 第一次失败重试间隔时间(后续间隔翻倍)
func (p *Sms) SendWithRetry(count uint8, interval time.Duration) {
    if err := p.Send(); err != nil {
        if stop, ok := err.(Stop); ok {
            log.Println(stop.Error())
            return
        }

        log.Println(err.Error())
        if count--; count > 0 {
            log.Printf("短信发送失败，%s后再次重试 #%d. \n", interval, count)
            <-time.After(interval)
            p.SendWithRetry(count, interval)
        }
    }
}

func TestRetry(t *testing.T) {
    sms := &Sms{Title: "今晚吃鸡", From: "Lucy", To: "Lily", Content: "八点半开黑"}
    sms.SendWithRetry(5, time.Second*2)
}
