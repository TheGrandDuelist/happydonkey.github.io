package services

import (
	"bbs-go/model"
	"bbs-go/pkg/common"
	"bbs-go/repositories"
	"net/http"

	"github.com/mlogclub/simple/common/dates"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"github.com/sirupsen/logrus"
)

var OperateLogService = newOperateLogService()

func newOperateLogService() *operateLogService {
	return &operateLogService{}
}

type operateLogService struct {
}


func (s *operateLogService) Find(cnd *sqls.Cnd) []model.OperateLog {
	return repositories.OperateLogRepository.Find(sqls.DB(), cnd)
}

func (s *operateLogService) FindOne(cnd *sqls.Cnd) *model.OperateLog {
	return repositories.OperateLogRepository.FindOne(sqls.DB(), cnd)
}

func (s *operateLogService) Get(id int64) *model.OperateLog {
	return repositories.OperateLogRepository.Get(sqls.DB(), id)
}

func (s *operateLogService) FindPageByCnd(cnd *sqls.Cnd) (list []model.OperateLog, paging *sqls.Paging) {
	return repositories.OperateLogRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *operateLogService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.OperateLogRepository.Updates(sqls.DB(), id, columns)
}

func (s *operateLogService) Take(where ...interface{}) *model.OperateLog {
	return repositories.OperateLogRepository.Take(sqls.DB(), where...)
}

func (s *operateLogService) FindPageByParams(params *params.QueryParams) (list []model.OperateLog, paging *sqls.Paging) {
	return repositories.OperateLogRepository.FindPageByParams(sqls.DB(), params)
}

func (s *operateLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.OperateLogRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *operateLogService) Delete(id int64) {
	repositories.OperateLogRepository.Delete(sqls.DB(), id)
}

// SendEmailNotice 发送邮件通知
func (s *messageService) SendEmailNotice(t *model.Message) {
	msgType := msg.Type(t.Type)

	// 话题被删除不发送邮件提醒
	if msgType == msg.TypeTopicDelete {
		return
	}
	user := cache.UserCache.Get(t.UserId)
	if user == nil || len(user.Email.String) == 0 {
		return
	}
	var (
		siteTitle  = cache.SysConfigCache.GetValue(constants.SysConfigSiteTitle)
		emailTitle = siteTitle + " - 新消息提醒"
	)

	if msgType == msg.TypeTopicComment {
		emailTitle = siteTitle + " - 收到话题评论"
	} else if msgType == msg.TypeCommentReply {
		emailTitle = siteTitle + " - 收到他人回复"
	} else if msgType == msg.TypeTopicFavorite {
		emailTitle = siteTitle + " - 话题被收藏"
	} else if msgType == msg.TypeTopicRecommend {
		emailTitle = siteTitle + " - 话题被设为推荐"
	} else if msgType == msg.TypeTopicDelete {
		emailTitle = siteTitle + " - 话题被删除"
	} else if msgType == msg.TypeTopicLike {
		emailTitle = siteTitle + " - 收到点赞"
	} else if msgType == msg.TypeArticleComment {
		emailTitle = siteTitle + " - 收到文章评论"
	}

	var from *model.User
	if t.FromId > 0 {
		from = cache.UserCache.Get(t.FromId)
	}
	err := email.SendTemplateEmail(from, user.Email.String, emailTitle, emailTitle, t.Content,
		t.QuoteContent, &model.ActionLink{
			Title: "点击查看详情",
			Url:   bbsurls.AbsUrl("/user/messages"),
		})
	if err != nil {
		logrus.Error(err)
	}
}

func (s *operateLogService) AddOperateLog(userId int64, opType, dataType string, dataId int64,
	description string, r *http.Request) {

	operateLog := &model.OperateLog{
		UserId:      userId,
		OpType:      opType,
		DataType:    dataType,
		DataId:      dataId,
		Description: description,
		CreateTime:  dates.NowTimestamp(),
	}
	if r != nil {
		operateLog.Ip = common.GetRequestIP(r)
		operateLog.UserAgent = common.GetUserAgent(r)
		operateLog.Referer = r.Header.Get("Referer")
	}
	if err := repositories.OperateLogRepository.Create(sqls.DB(), operateLog); err != nil {
		logrus.Error(err)
	}
}

func (s *messageService) SendMsg(from, to int64, msgType msg.Type,
	title, content, quoteContent string, extraData interface{}) {

	t := &model.Message{
		FromId:       from,
		UserId:       to,
		Title:        title,
		Content:      content,
		QuoteContent: quoteContent,
		Type:         int(msgType),
		ExtraData:    jsons.ToJsonStr(extraData),
		Status:       msg.StatusUnread,
		CreateTime:   dates.NowTimestamp(),
	}
	if err := s.Create(t); err != nil {
		logrus.Error(err)
	} else {
		s.SendEmailNotice(t)
	}
}

func (s *operateLogService) Count(cnd *sqls.Cnd) int64 {
	return repositories.OperateLogRepository.Count(sqls.DB(), cnd)
}

func (s *operateLogService) UpdateLog(t *model.OperateLog) error {
	return repositories.OperateLogRepository.Update(sqls.DB(), t)
}

func (s *operateLogService) FindOneLog(cnd *sqls.Cnd) *model.OperateLog {
	return repositories.OperateLogRepository.FindOne(sqls.DB(), cnd)
}

func (s *operateLogService) GetLogById(id int64) *model.OperateLog {
	return repositories.OperateLogRepository.Get(sqls.DB(), id)
}

func (s *operateLogService) FindPageByCnd(cnd *sqls.Cnd) (list []model.OperateLog, paging *sqls.Paging) {
	return repositories.OperateLogRepository.FindPageByCnd(sqls.DB(), cnd)
}


func (s *operateLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.OperateLogRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *operateLogService) Create(t *model.OperateLog) error {
	return repositories.OperateLogRepository.Create(sqls.DB(), t)
}

func (s *operateLogService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.OperateLogRepository.Updates(sqls.DB(), id, columns)
}

func (s *operateLogService) Take(where ...interface{}) *model.OperateLog {
	return repositories.OperateLogRepository.Take(sqls.DB(), where...)
}

func (s *operateLogService) FindPageByParams(params *params.QueryParams) (list []model.OperateLog, paging *sqls.Paging) {
	return repositories.OperateLogRepository.FindPageByParams(sqls.DB(), params)
}

func (s *operateLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.OperateLogRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *operateLogService) Delete(id int64) {
	repositories.OperateLogRepository.Delete(sqls.DB(), id)
}

// SendEmailNotice 发送邮件通知
func (s *messageService) SendEmailNotice(t *model.Message) {
	msgType := msg.Type(t.Type)

	// 话题被删除不发送邮件提醒
	if msgType == msg.TypeTopicDelete {
		return
	}
	user := cache.UserCache.Get(t.UserId)
	if user == nil || len(user.Email.String) == 0 {
		return
	}
	var (
		siteTitle  = cache.SysConfigCache.GetValue(constants.SysConfigSiteTitle)
		emailTitle = siteTitle + " - 新消息提醒"
	)

	if msgType == msg.TypeTopicComment {
		emailTitle = siteTitle + " - 收到话题评论"
	} else if msgType == msg.TypeCommentReply {
		emailTitle = siteTitle + " - 收到他人回复"
	} else if msgType == msg.TypeTopicFavorite {
		emailTitle = siteTitle + " - 话题被收藏"
	} else if msgType == msg.TypeTopicRecommend {
		emailTitle = siteTitle + " - 话题被设为推荐"
	} else if msgType == msg.TypeTopicDelete {
		emailTitle = siteTitle + " - 话题被删除"
	} else if msgType == msg.TypeTopicLike {
		emailTitle = siteTitle + " - 收到点赞"
	} else if msgType == msg.TypeArticleComment {
		emailTitle = siteTitle + " - 收到文章评论"
	}

	var from *model.User
	if t.FromId > 0 {
		from = cache.UserCache.Get(t.FromId)
	}
	err := email.SendTemplateEmail(from, user.Email.String, emailTitle, emailTitle, t.Content,
		t.QuoteContent, &model.ActionLink{
			Title: "点击查看详情",
			Url:   bbsurls.AbsUrl("/user/messages"),
		})
	if err != nil {
		logrus.Error(err)
	}
}
