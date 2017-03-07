package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brettallred/go-logit"
	"github.com/brettallred/go-rabbit"
	"github.com/nuvi/go-airbrake"
	"github.com/nuvi/go-social-activity-parser/social_activity"
	"github.com/streadway/amqp"
)

var publisherConnection *rabbit.Connection
var activitiesPublisher *rabbit.AssuredPublisher

// StartRabbit is the intializtion method for all of the RabbitMQ Subscribers.
func startRabbit() {
	logit.Info("INFO: Starting RabbitMQ")

	publisherURL := os.Getenv("RABBITMQ_PUBLISHER_URL")
	if publisherURL == "" {
		panic("RABBITMQ_PUBLISHER_URL is empty")
	}
	publisherConnection = rabbit.NewConnectionWithURL(publisherURL)
	activitiesPublisher = rabbit.NewAssuredPublisherWithConnection(publisherConnection, exitChannel)
	activitiesPublisher.DisableRepublishing()
	activitiesPublisher.SetConfirmationHandler(publisherConfirmationHandler)
	activitiesPublisher.SetExplicitWaiting()

	subscribers := map[string]string{
		"DISQUS_FIREHOSE":    "streamer.disqus_firehose_activity.created",
		"FACEBOOK":           "data_collector.facebook_activity.created",
		"GOOGLE_PLUS":        "data_collector.google_plus_activity.created",
		"INSTAGRAM":          "data_collector.instagram_activity.created",
		"LEXIS_NEXIS":        "data_collector.lexis_nexis_activity.created",
		"PINTEREST":          "data_collector.pinterest_activity.created",
		"REDDIT":             "data_collector.reddit_activity.created",
		"SOCIALGIST":         "streamer.socialgist_activity.created",
		"STACK_OVERFLOW":     "data_collector.stack_overflow_activity.created",
		"TUMBLR":             "streamer.tumblr_activity.created",
		"TWITTER":            "streamer.twitter_activity.created",
		"TWITTER_SEARCH":     "twitter_search.twitter_activity.created",
		"VK":                 "data_collector.vk_activity.created",
		"WEBHOSE":            "gazette.webhose_activity.created",
		"WORDPRESS_FIREHOSE": "streamer.wordpress_firehose_activity.created",
		"YOUTUBE":            "data_collector.youtube_activity.created",
	}

	handlers := map[string]func(amqp.Delivery) bool{
		"DISQUS_FIREHOSE":    disqusFirehoseActivityCreatedHandler,
		"FACEBOOK":           facebookActivityCreatedHandler,
		"GOOGLE_PLUS":        googleplusActivityCreatedHandler,
		"INSTAGRAM":          instagramActivityCreatedHandler,
		"LEXIS_NEXIS":        lexisNexisActivityCreatedHandler,
		"PINTEREST":          pinterestActivityCreatedHandler,
		"REDDIT":             redditActivityCreatedHandler,
		"SOCIALGIST":         socialgistActivityCreatedHandler,
		"STACK_OVERFLOW":     stackOverflowActivityCreatedHandler,
		"TUMBLR":             tumblrActivityCreatedHandler,
		"TWITTER":            twitterActivityCreatedHandler,
		"TWITTER_SEARCH":     twitterSearchActivityCreatedHandler,
		"VK":                 vkActivityCreatedHandler,
		"WEBHOSE":            webhoseActivityCreatedHandler,
		"WORDPRESS_FIREHOSE": wordpressFirehoseActivityCreatedHandler,
		"YOUTUBE":            youtubeActivityCreatedHandler,
	}

	logit.Info("Starting processing for %s", ProcessingNetwork)
	subscriber := buildSubscriber(subscribers[ProcessingNetwork], ProcessingNetwork)
	isProd := os.Getenv("PLATFORM_ENV") == "prod"
	handler := func(delivery amqp.Delivery) bool {
		// Catch anything bad that happens
		defer func() {
			if r := recover(); r != nil {
				delivery.Nack(false, true)
				logit.Error("Error: %#v", r)
				go airbrake.GetNotifier().Notify(r, nil)
				if !isProd {
					time.Sleep(3 * time.Second) // needed for AirBrake
					panic(r)
				}
			}
		}()
		return handlers[ProcessingNetwork](delivery)
	}
	rabbit.Register(subscriber, handler)
	rabbit.StartSubscribers()
}

func disqusFirehoseActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateDisqusFirehoseSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "disqus_firehose", delivery)
}

func facebookActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateFacebookSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "facebook", delivery)
}

func googleplusActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateGoogleplusSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "google_plus", delivery)
}

// InstagramActivityCreatedHandler is the RabbitMQ subscriber handler for
func instagramActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateInstagramSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "instagram", delivery)
}

func lexisNexisActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateLexisNexisSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "lexis_nexis", delivery)
}

func pinterestActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GeneratePinterestSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "pinterest", delivery)
}

func redditActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateRedditSocialActivity(delivery.Body)
	if len(activity.ID) == 0 {
		return true
	}
	return publishCreatedActivity(activity, "reddit", delivery)
}

func socialgistActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateSocialgistSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "socialgist", delivery)
}

func stackOverflowActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateStackOverflowSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "stack_overflow", delivery)
}

func tumblrActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateTumblrSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "tumblr", delivery)
}

func twitterActivityCreatedHandler(delivery amqp.Delivery) bool {
	logit.Debug(fmt.Sprintf("GNIP PAYLOAD %s", string(delivery.Body)))
	activity := socialactivity.GenerateTwitterSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "twitter", delivery)
}

func twitterSearchActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateTwitterSearchSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "twitter_search", delivery)
}

func vkActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateVkSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "vk", delivery)
}

func webhoseActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateWebhoseSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "webhose", delivery)
}

func wordpressFirehoseActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateWordpressFirehoseSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "wordpress_firehose", delivery)
}

func youtubeActivityCreatedHandler(delivery amqp.Delivery) bool {
	activity := socialactivity.GenerateYoutubeSocialActivity(delivery.Body)
	return publishCreatedActivity(activity, "youtube", delivery)
}

func publishCreatedActivity(socialActivity *socialactivity.SocialActivity, network string, delivery amqp.Delivery) bool {
	socialActivity.NormalizeAllURLS()
	socialActivity.ResolveLocation(db)

	activitiesSubscriber := rabbit.Subscriber{
		Durable:    true,
		Exchange:   "events",
		RoutingKey: fmt.Sprintf("social_activity_parser.%s_activity.created", network),
	}

	publishItem, _ := json.Marshal(socialActivity)

	if !activitiesPublisher.PublishBytesWithArg(publishItem, &activitiesSubscriber, delivery) {
		delivery.Nack(false, true)
	}

	return true
}

func buildSubscriber(routingKey string, network string) rabbit.Subscriber {
	concurrency, _ := strconv.Atoi(os.Getenv(strings.ToUpper(network) + "_SUBSCRIBER_CONCURRENCY"))
	if concurrency <= 0 {
		concurrency, _ = strconv.Atoi(os.Getenv("SUBSCRIBER_CONCURRENCY"))
		if concurrency <= 0 {
			concurrency = 1
		}
	}

	var subscriber = rabbit.Subscriber{
		Concurrency:   concurrency,
		Durable:       true,
		Exchange:      "events",
		Queue:         "social_activity_parser." + routingKey,
		RoutingKey:    routingKey,
		PrefetchCount: 100,
		ManualAck:     true,
	}
	subscriber.PrefixQueueInDev()
	subscriber.AutoDeleteInDev()

	return subscriber
}

func publisherConfirmationHandler(confirmation amqp.Confirmation, arg interface{}) {
	delivery := arg.(amqp.Delivery)
	if confirmation.Ack {
		if err := delivery.Ack(false); err != nil {
			logit.Error("Error on ACK: %v", err)
		}
	} else {
		if err := delivery.Nack(false, true); err != nil {
			logit.Error("Error on NACK: %v", err)
		}
	}
}
