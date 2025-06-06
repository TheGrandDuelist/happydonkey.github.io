package api

import (
	"bbs-go/model/constants"
	"bbs-go/pkg/errs"
	"bbs-go/services"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

type LikeController struct {
	Ctx iris.Context
}

func (c *LikeController) PostLike() *web.JsonResult {
	var (
		entityType = params.FormValue(c.Ctx, "entityType")
		entityId   = params.FormValueInt64Default(c.Ctx, "entityId", 0)
		user       = services.UserTokenService.GetCurrent(c.Ctx)
		err        error
	)
	if user == nil {
		return web.JsonError(errs.NotLogin)
	}
	if entityType == constants.EntityTopic {
		err = services.UserLikeService.TopicLike(user.Id, entityId)
	} else if entityType == constants.EntityArticle {
		err = services.UserLikeService.ArticleLike(user.Id, entityId)
	} else if entityType == constants.EntityComment {
		err = services.UserLikeService.CommentLike(user.Id, entityId)
	}
	if err != nil {
		return web.JsonError(err)
	}
	return web.JsonSuccess()
}

func (c *LikeController) PostUnlike() *web.JsonResult {
	var (
		entityType = params.FormValue(c.Ctx, "entityType")
		entityId   = params.FormValueInt64Default(c.Ctx, "entityId", 0)
		user       = services.UserTokenService.GetCurrent(c.Ctx)
		err        error
	)
	if user == nil {
		return web.JsonError(errs.NotLogin)
	}
	if entityType == constants.EntityTopic {
		err = services.UserLikeService.TopicUnLike(user.Id, entityId)
	} else if entityType == constants.EntityArticle {
		err = services.UserLikeService.ArticleUnLike(user.Id, entityId)
	} else if entityType == constants.EntityComment {
		err = services.UserLikeService.CommentUnLike(user.Id, entityId)
	}
	if err != nil {
		return web.JsonError(err)
	}
	return web.JsonSuccess()
}

func (c *LikeController) GetIsLiked() *web.JsonResult {
	var (
		user           = services.UserTokenService.GetCurrent(c.Ctx)
		entityType     = params.FormValue(c.Ctx, "entityType")
		entityIds      = params.FormValueInt64Array(c.Ctx, "entityIds")
		likedEntityIds []int64
	)
	if user == nil || strs.IsBlank(entityType) || entityId <= 0 {
		return web.NewEmptyRspBuilder().Put("liked", false).JsonResult()
	} else {
		liked := services.UserLikeService.Exists(user.Id, entityType, entityId)
		return web.NewEmptyRspBuilder().Put("liked", liked).JsonResult()
	}
	if user != nil {
		likedEntityIds = services.UserLikeService.IsLiked(user.Id, entityType, entityIds)
	}
	return web.NewEmptyRspBuilder().Put("liked", likedEntityIds).JsonResult()
}

func (c *LikeController) GetLiked() *web.JsonResult {
	var (
		user       = services.UserTokenService.GetCurrent(c.Ctx)
		entityType = params.FormValue(c.Ctx, "entityType")
		entityId   = params.FormValueInt64Default(c.Ctx, "entityId", 0)
	)
	if user == nil || strs.IsBlank(entityType) || entityId <= 0 {
		return web.NewEmptyRspBuilder().Put("liked", false).JsonResult()
	} else {
		liked := services.UserLikeService.Exists(user.Id, entityType, entityId)
		return web.NewEmptyRspBuilder().Put("liked", liked).JsonResult()
	}
	
	if user == nil {
		return web.JsonError(errs.NotLogin)
	}
	if entityType == constants.EntityTopic {
		err = services.UserLikeService.TopicUnLike(user.Id, entityId)
	} else if entityType == constants.EntityArticle {
		err = services.UserLikeService.ArticleUnLike(user.Id, entityId)
	} else if entityType == constants.EntityComment {
		err = services.UserLikeService.CommentUnLike(user.Id, entityId)
	}
}

func (c *LikeController) GetIsLiked() *web.JsonResult {
	var (
		user           = services.UserTokenService.GetCurrent(c.Ctx)
		entityType     = params.FormValue(c.Ctx, "entityType")
		entityIds      = params.FormValueInt64Array(c.Ctx, "entityIds")
		likedEntityIds []int64
	)
	if user == nil || strs.IsBlank(entityType) || entityId <= 0 {
		return web.NewEmptyRspBuilder().Put("liked", false).JsonResult()
	} else {
		liked := services.UserLikeService.Exists(user.Id, entityType, entityId)
		return web.NewEmptyRspBuilder().Put("liked", liked).JsonResult()
	}
	if user != nil {
		likedEntityIds = services.UserLikeService.IsLiked(user.Id, entityType, entityIds)
	}
	return web.NewEmptyRspBuilder().Put("liked", likedEntityIds).JsonResult()
}

func (c *LikeController) GetLiked() *web.JsonResult {
	var (
		user       = services.UserTokenService.GetCurrent(c.Ctx)
		entityType = params.FormValue(c.Ctx, "entityType")
		entityId   = params.FormValueInt64Default(c.Ctx, "entityId", 0)
	)
	if user == nil || strs.IsBlank(entityType) || entityId <= 0 {
		return web.NewEmptyRspBuilder().Put("liked", false).JsonResult()
	} else {
		liked := services.UserLikeService.Exists(user.Id, entityType, entityId)
		return web.NewEmptyRspBuilder().Put("liked", liked).JsonResult()
	}
	
	if user == nil {
		return web.JsonError(errs.NotLogin)
	}
	if entityType == constants.EntityTopic {
		err = services.UserLikeService.TopicUnLike(user.Id, entityId)
	} else if entityType == constants.EntityArticle {
		err = services.UserLikeService.ArticleUnLike(user.Id, entityId)
	} else if entityType == constants.EntityComment {
		err = services.UserLikeService.CommentUnLike(user.Id, entityId)
	}
}
