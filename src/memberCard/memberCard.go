// Package: openapi.app.models.memberCard
// File: openapi.app.models.memberCard.go
// Created by mint
// Useage: 会员卡
/*
主要产业标识符
0	ISO/TC 68 和其他行业使用
1	航空
2	航空和其他未来行业使用
3	运输、娱乐和金融财务
4	金融财务
5	金融财务
6	商业和金融财务
7	石油和其他未来行业使用
8	医疗、电信和其他未来行业使用
9	由本国标准机构分配
*/
// DATE: 14-7-4 15:30
package memberCard

import (
	"fmt"
	"utils/algorith"
)

type CardNumber string

type MemberCard struct {
	MII byte   // 1     主要产业标识符（Major Industry Identifier (MII)）
	CPI byte   // 2-3   公司标识符，默认: 32
	CDI uint16 // 4-6   国家域标识符（Country Domain Identifier）
	PII uint64 // 7-18  个人信息标识（Personal identifying information）
	IVC byte   // 19    验证码标识（Identity verification code）
}

func (c *MemberCard) GetMII() byte {
	return c.MII
}

func (c *MemberCard) GetCPI() byte {
	return c.CPI
}

func (c *MemberCard) GetCDI() uint16 {
	return c.CDI
}

func (c *MemberCard) GetPII() uint64 {
	return c.PII
}

func (c *MemberCard) GetIVC() byte {
	return c.IVC
}

func (c *MemberCard) SetIVC(code byte) {
	c.IVC = code
}

// 新建会员卡
func NewMemberCard(mii, cpi byte, cdi uint16, pii uint64) *MemberCard {
	return &MemberCard{
		MII: mii,
		CPI: cpi,
		CDI: cdi,
		PII: pii,
		IVC: algorith.GenLuhnCheckDigit([]byte(fmt.Sprintf("%v%v%.*d%.*d", mii, cpi, 3, cdi, 12, pii))),
	}
}

// 格式化输出
func (c *MemberCard) String() CardNumber {
	return CardNumber(fmt.Sprintf("%v%v%.*d%.*d%v", c.MII, c.CPI, 3, c.CDI, 12, c.PII, c.IVC))
}

// 验证会员卡号的正确性
func (c *MemberCard) ValidateCard() bool {
	return algorith.ValidateLuhn(string(c.String()))
}
