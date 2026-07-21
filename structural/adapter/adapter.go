package main

import "fmt"

// 适配器模式：让接口不兼容的类能够协作。
// 场景：我们的系统只认识 PaymentProcessor 接口，
// 但要接入一个第三方支付 SDK，它的方法签名完全不同。

// ---------- 目标接口（客户端期望的接口）----------

type PaymentProcessor interface {
	Pay(amountCents int) string
}

// ---------- 被适配者（第三方 SDK，我们不能改它）----------

type ThirdPartyPay struct{}

// 它的接口用「元」为单位、方法名也不同。
func (ThirdPartyPay) MakePayment(yuan float64) string {
	return fmt.Sprintf("[第三方SDK] 已支付 %.2f 元", yuan)
}

// ---------- 适配器：把 ThirdPartyPay 适配成 PaymentProcessor ----------

type ThirdPartyAdapter struct {
	sdk ThirdPartyPay
}

func NewThirdPartyAdapter() *ThirdPartyAdapter { return &ThirdPartyAdapter{} }

// Pay 做接口转换：分 -> 元，并转调 SDK 的方法。
func (a *ThirdPartyAdapter) Pay(amountCents int) string {
	yuan := float64(amountCents) / 100.0
	return a.sdk.MakePayment(yuan)
}

// 客户端只依赖目标接口，不知道背后是谁。
func checkout(p PaymentProcessor, cents int) {
	fmt.Println("结算 ->", p.Pay(cents))
}

func main() {
	// 通过适配器，第三方 SDK 就能被当作 PaymentProcessor 使用
	adapter := NewThirdPartyAdapter()
	checkout(adapter, 12999) // 129.99 元
}
