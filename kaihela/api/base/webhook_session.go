package base

import (
	"errors"
	"github.com/bytedance/sonic"
	log "github.com/sirupsen/logrus"
)

type WebhookSession struct {
	*StateSession
	EncryptKey  string
	VerifyToken string
}

func NewWebhookSession(encryptKey, verityToken string, compress int) *WebhookSession {
	session := &WebhookSession{}
	if encryptKey != "" {
		session.EncryptKey = encryptKey
	}
	if verityToken != "" {
		session.VerifyToken = verityToken
	}
	session.Compressed = compress
	session.StateSession.ProcessDataHandler = session.ProcessData
	return session
}

func (s *WebhookSession) ProcessData(data []byte) (err error, data2 []byte) {
	//如果有加密，则对数据进行解密
	if s.EncryptKey == "" {
		data2 = data
		return
	}
	jdata, err := sonic.Get(data)
	if err != nil {
		return err, nil
	}
	if jdata.Get("encrypt").Exists() == false {
		log.Error("Encrypt_Data Not Exist", string(data))
		err = errors.New("Encrypt_Data Not Exist")
		return
	}
	
}
