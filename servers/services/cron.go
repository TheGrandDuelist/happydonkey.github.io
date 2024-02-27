package scheduler

import (
	"bbs-go/pkg/sitemap"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"bbs-go/services"
)

func Start() {
	c := cron.New()

	// Generate sitemap
	addCronFunc(c, "0 0 4 ? * *", func() {
		sitemap.Generate()
	})
	// Generate RSS
	addCronFunc(c, "@every 30m", func() {
		services.ArticleService.GenerateRss()
		services.TopicService.GenerateRss()
	})

	c.Start()
}

func StartTest() {
	c := cron.New()

	// Generate sitemap
	addCronFunc(c, "0 0 4 ? * *", func() {
		sitemap.Generate()
	})
	// Generate RSS
	addCronFunc(c, "@every 30m", func() {
		services.ArticleService.GenerateRss()
		services.TopicService.GenerateRss()
	})

	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
