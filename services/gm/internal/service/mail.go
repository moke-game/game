package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/moke-game/game/services/gm/internal/common"

	pb "github.com/moke-game/game/api/gen/gm"
)

func (s *Service) SendMail(ctx context.Context, request *pb.SendMailRequest) (*pb.SendMailResponse, error) {
	s.logger.Info("send mail request", zap.Any("request", request))
	mailMsg := &pb.MailSendData{}
	var err error
	if data, e := common.CBCDecrypt([]byte(s.aesKey), request.Data); e != nil {
		err = fmt.Errorf("decrypt data failed: %w", e)
	} else if e := protojson.Unmarshal([]byte(data), mailMsg); e != nil {
		err = fmt.Errorf("unmarshal data failed: %w,%v", e, data)
	} else if req, e := common.TransMail2MailReqMsg(mailMsg); e != nil {
		err = fmt.Errorf("make mail from request failed: %w", e)
	} else if _, e := s.mailCli.SendMail(ctx, req); e != nil {
		err = fmt.Errorf("send mail failed: %w", e)
	}
	status, code, info := "ok", "1", "success"
	if err != nil {
		status, code, info = "error", "-1", err.Error()
		s.logger.Error("send mail failed", zap.Error(err))
	}
	return &pb.SendMailResponse{
		Status: status,
		Code:   code,
		Info:   info,
		Data: &pb.MailSendResult{
			UserId: mailMsg.RoleId,
		},
	}, nil

}
