package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb3 "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/mail"
	pb "github.com/moke-game/platform/api/gen/profile"

	pb2 "github.com/moke-game/game/api/gen/gm"
)

func TransKnapsack2Goods(knapsack *pb3.Knapsack) []*pb2.GoodsList {
	if knapsack == nil {
		return nil
	}
	res := make([]*pb2.GoodsList, 0)
	for _, v := range knapsack.Items {
		res = append(res, &pb2.GoodsList{
			GoodsId:  v.Id,
			Quantity: v.Num,
			Status:   fmt.Sprintf("%d", v.Expire),
		})
	}
	return res
}

func TransProfile2PlayerBanInfo(profile *pb.Profile) *pb2.PlayerBannedInfo {
	return &pb2.PlayerBannedInfo{
		PlatformId:     profile.PlatformId,
		PlatformName:   profile.PlatformId,
		ServerId:       profile.PlatformId,
		ServerName:     profile.PlatformId,
		ChannelId:      fmt.Sprintf("%d", profile.Channel),
		UserId:         profile.Account,
		RoleId:         profile.Uid,
		RoleName:       profile.Nickname,
		ChannelUserId:  profile.Uid,
		RechargeAmount: fmt.Sprintf("%d", profile.RechargeAmount),
		CreateTime:     profile.RegisterTime,
	}
}

func TransProfile2PlayerInfo(profile *pb.Profile) *pb2.PlayerInfo {
	channelId := fmt.Sprintf("%d", profile.Channel)
	return &pb2.PlayerInfo{
		PlatformId:     profile.PlatformId,
		ChannelId:      channelId,
		UserId:         profile.Account,
		RoleId:         profile.Uid,
		RoleName:       profile.Nickname,
		ChannelUserId:  profile.Openid,
		RegisterIp:     profile.RegisterIp,
		DeviceId:       profile.DeviceId,
		RechargeAmount: profile.RechargeAmount,
	}
}

func TransProfile2ProfileDetail(profile *pb.Profile, knapsack *pb3.Knapsack) *pb2.GetPlayerInfoResponse {
	diamond := int32(0)
	if v, ok := knapsack.Items[2]; ok {
		diamond = v.Num
	}
	gold := int32(0)
	if v, ok := knapsack.Items[1]; ok {
		gold = v.Num
	}
	return &pb2.GetPlayerInfoResponse{
		PlatformId:    profile.PlatformId,
		RoleId:        profile.Uid,
		RoleName:      profile.Nickname,
		UserId:        profile.Account,
		CreateTime:    profile.RegisterTime,
		LastLoginTime: profile.LastLoginTime,
		Copper:        gold,
		TopUpStarCoin: diamond,
		Code:          "1",
	}
}

func makeMailRewards(items []string) ([]*mail.MailReward, error) {
	rewards := make([]*mail.MailReward, 0)
	for _, v := range items {
		strs := strings.Split(v, ":")
		if len(strs) < 2 {
			return nil, fmt.Errorf("invalid item %s", v)
		}
		if id, err := strconv.ParseInt(strs[0], 10, 32); err != nil {
			return nil, fmt.Errorf("invalid item id %s", v)
		} else if num, err := strconv.ParseInt(strs[1], 10, 32); err != nil {
			return nil, fmt.Errorf("invalid item num %s", v)
		} else {
			rewards = append(rewards, &mail.MailReward{
				Id:   id,
				Num:  int32(num),
				Type: int32(id),
			})
		}
	}
	return rewards, nil
}

// makeContent2ReqMsg make content to request message
// content format: zh:1111111:111111111;en:22222222:22222222222222
// return: map[language]content, map[language]title, error
func makeContent2ReqMsg(content string) (map[string]string, map[string]string, error) {
	res := make(map[string]string)
	title := make(map[string]string)
	strs := strings.Split(content, ";")
	for _, v := range strs {
		langContent := strings.Split(v, ":")
		if len(langContent) != 3 {
			return nil, nil, fmt.Errorf("invalid content %s", v)
		}
		res[langContent[0]] = langContent[2]
		title[langContent[0]] = langContent[1]
	}
	return title, res, nil

}

func makeMailData(sendData *pb2.MailSendData) (*mail.Mail, error) {
	timeLayout := "2006-01-02T15:04:05-0700"
	startTime := time.Now()
	endTime := time.Now().AddDate(0, 0, 90)
	var err error
	if sendData.StartTime != "" {
		startTime, err = time.Parse(timeLayout, sendData.StartTime)
		if err != nil {
			return nil, err
		}
	}

	if sendData.EndTime != "" {
		endTime, err = time.Parse(timeLayout, sendData.EndTime)
		if err != nil {
			return nil, err
		}
	}

	filters := &mail.Mail_Filter{
		RegisterTime: sendData.RegisterTime,
	}

	duration := time.Duration(endTime.Unix()-startTime.Unix()) * time.Second
	if duration.Hours() < 1 {
		return nil, fmt.Errorf("expire time is invalid %v", sendData)
	}
	rewards, err := makeMailRewards(sendData.Items)
	if err != nil {
		return nil, err
	}
	titles, contents, err := makeContent2ReqMsg(sendData.Content)
	if err != nil {
		return nil, err
	}
	res := &mail.Mail{
		Id:       time.Now().UnixMilli(),
		Body:     contents,
		Date:     startTime.Unix(),
		ExpireAt: endTime.Unix(),
		From:     sendData.Sender,
		Rewards:  rewards,
		Title:    titles,
		Filters:  filters,
	}
	return res, nil
}

func TransMail2MailReqMsg(sendData *pb2.MailSendData) (*mail.SendMailRequest, error) {
	m, err := makeMailData(sendData)
	if err != nil {
		return nil, err
	}
	sendType := mail.SendMailRequest_NONE
	if sendData.SendType == "1" {
		sendType = mail.SendMailRequest_ALL
	} else if sendData.SendType == "2" {
		sendType = mail.SendMailRequest_ROLE
	}

	res := &mail.SendMailRequest{
		//ChannelId:  sendData.ChannelId,
		SendType: sendType,
		//PlatformId: sendData.PlatformId,
		RoleIds: []string{sendData.RoleId},
		Mail:    m,
	}
	return res, nil
}

// CBCDecrypt AES-CBC 解密
func CBCDecrypt(key []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphercode, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	if len(ciphercode)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphercode length is not a multiple of the block size")
	}

	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphercode, ciphercode)
	ciphercode = bytes.Trim(ciphercode, "\x00")
	return strings.TrimSpace(string(ciphercode)), nil
}

// EncryptMD5 encrypt data with md5.
func EncryptMD5(data, key string) (string, error) {
	str := fmt.Sprintf("%s.%s", data, key)
	_, err := md5.New().Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5.New().Sum(nil)), nil
}
