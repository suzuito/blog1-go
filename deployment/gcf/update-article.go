package gcf

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"golang.org/x/xerrors"
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

func BlogUpdateArticle(ctx context.Context, ev GCSEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return xerrors.Errorf("metadata.FromContext: %w", err)
	}
	fmt.Printf("%+v\n", meta)
	fmt.Printf("%+v\n", ev)
	if ev.Bucket != env.GCPBucketArticle {
		return xerrors.Errorf("Invalid backet name exp:%s != real:%s", env.GCPBucketArticle, ev.Bucket)
	}
	u := usecase.NewImpl(gdeps.DB, gdeps.Storage, gdeps.MDConverter)
	log.Info().Str(
		"file", fmt.Sprintf("%s/%s", ev.Bucket, ev.Name),
	).Send()
	if filepath.Ext(ev.Name) != ".md" {
		return nil
	}
	if err := u.UpdateArticle(ctx, ev.Name); err != nil {
		return xerrors.Errorf("Cannot u.UpdateArticle : %w", err)
	}
	return nil
}
