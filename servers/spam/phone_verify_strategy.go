package spam

import (
	"bbs-go/model"
	"bbs-go/pkg/errs"
	"bbs-go/services"
)

type PhoneVerifyStrategy struct{}

func (PhoneVerifyStrategy) Name() string {
	return "PhoneVerifyStrategy"
}

func (PhoneVerifyStrategy) CheckTopic(user *model.User, form model.CreateTopicForm) error {
	if services.SysConfigService.IsCreateTopicphoneVerified() && !user.phoneVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (PhoneVerifyStrategy) CheckArticle(user *model.User, form model.CreateArticleForm) error {
	if services.SysConfigService.IsCreateArticlephoneVerified() && !user.phoneVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (PhoneVerifyStrategy) CheckComment(user *model.User, form model.CreateCommentForm) error {
	if services.SysConfigService.IsCreateCommentphoneVerified() && !user.phoneVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (PhoneVerifyStrategy) CheckTopicBak(user *model.User, form model.CreateTopicForm) error {
	if services.SysConfigService.IsCreateTopicphoneVerified() && !user.phoneVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (PhoneVerifyStrategy) CheckArticleBak(user *model.User, form model.CreateArticleForm) error {
	if services.SysConfigService.IsCreateArticlephoneVerified() && !user.phoneVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (PhoneVerifyStrategy) CheckCommentBak(user *model.User, form model.CreateCommentForm) error {
	if services.SysConfigService.IsCreateCommentphoneVerified() && !user.phoneVerified {
		return errs.EmailNotVerified
	}
	return nil
}

