package gcf

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

type GCSEvent struct {
	Kind                    string                 `json:"kind"`
	ID                      string                 `json:"id"`
	SelfLink                string                 `json:"selfLink"`
	Name                    string                 `json:"name"`
	Bucket                  string                 `json:"bucket"`
	Generation              string                 `json:"generation"`
	Metageneration          string                 `json:"metageneration"`
	ContentType             string                 `json:"contentType"`
	TimeCreated             time.Time              `json:"timeCreated"`
	Updated                 time.Time              `json:"updated"`
	TemporaryHold           bool                   `json:"temporaryHold"`
	EventBasedHold          bool                   `json:"eventBasedHold"`
	RetentionExpirationTime time.Time              `json:"retentionExpirationTime"`
	StorageClass            string                 `json:"storageClass"`
	TimeStorageClassUpdated time.Time              `json:"timeStorageClassUpdated"`
	Size                    string                 `json:"size"`
	MD5Hash                 string                 `json:"md5Hash"`
	MediaLink               string                 `json:"mediaLink"`
	ContentEncoding         string                 `json:"contentEncoding"`
	ContentDisposition      string                 `json:"contentDisposition"`
	CacheControl            string                 `json:"cacheControl"`
	Metadata                map[string]interface{} `json:"metadata"`
	CRC32C                  string                 `json:"crc32c"`
	ComponentCount          int                    `json:"componentCount"`
	Etag                    string                 `json:"etag"`
	CustomerEncryption      struct {
		EncryptionAlgorithm string `json:"encryptionAlgorithm"`
		KeySha256           string `json:"keySha256"`
	}
	KMSKeyName    string `json:"kmsKeyName"`
	ResourceState string `json:"resourceState"`
}

func (ev *GCSEvent) ArticleID() entity.ArticleID {
	return entity.ArticleID(strings.Replace(filepath.Base(ev.Name), filepath.Ext(ev.Name), "", -1))
}

func BlogUpdateArticle(ctx context.Context, meta *metadata.Metadata, ev GCSEvent) error {
	if meta.EventType != "google.storage.object.finalize" {
		return nil
	}
	if ev.Bucket != setting.E.GCPBucketArticle {
		return errors.Errorf("Invalid backet name exp:%s != real:%s", setting.E.GCPBucketArticle, ev.Bucket)
	}
	if filepath.Ext(ev.Name) != ".md" {
		return nil
	}
	u := usecase.NewImpl(gdeps.DB, gdeps.Storage, gdeps.MDConverter)
	log.Info().Str(
		"file", fmt.Sprintf("%s/%s", ev.Bucket, ev.Name),
	).Send()
	if err := u.UpdateArticle(ctx, ev.Name); err != nil {
		return errors.Wrapf(err, "Cannot u.UpdateArticle : %s", ev.Name)
	}
	return nil
}

func BlogDeleteArticle(ctx context.Context, meta *metadata.Metadata, ev GCSEvent) error {
	if meta.EventType != "google.storage.object.delete" {
		return nil
	}
	if ev.Bucket != setting.E.GCPBucketArticle {
		return errors.Errorf("Invalid backet name exp:%s != real:%s", setting.E.GCPBucketArticle, ev.Bucket)
	}
	if filepath.Ext(ev.Name) != ".md" {
		return nil
	}
	articleID := ev.ArticleID()
	log.Info().Str(
		"file", fmt.Sprintf("%s/%s", ev.Bucket, ev.Name),
	).Send()
	u := usecase.NewImpl(gdeps.DB, gdeps.Storage, gdeps.MDConverter)
	if err := u.DeleteArticle(ctx, articleID); err != nil {
		return errors.Wrapf(err, "cannot delete article '%s'", articleID)
	}
	return nil
}
