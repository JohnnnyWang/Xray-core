package dispatcher

import (
	"sync"

	"github.com/juju/ratelimit"
	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/buf"
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/protocol"
)

type Bucket struct {
	User   *protocol.MemoryUser
	Bucket *ratelimit.Bucket
}
type BucketManage struct {
	Users map[string]*Bucket
	sync.RWMutex
}

type Writer struct {
	writer  buf.Writer
	limiter *ratelimit.Bucket
	logger  errors.ExportOption
}

func RateWriter(writer buf.Writer, limiter *ratelimit.Bucket, log errors.ExportOption) buf.Writer {
	return &Writer{
		writer:  writer,
		limiter: limiter,
		logger:  log,
	}
}

func (w *Writer) Close() error {
	return common.Close(w.writer)
}

func (w *Writer) WriteMultiBuffer(mb buf.MultiBuffer) error {
	// newError("limiter Wait:", mb.Len(), " ---- ", len(mb)).AtWarning().WriteToLog(w.logger)
	w.limiter.Wait(int64(mb.Len()))
	return w.writer.WriteMultiBuffer(mb)
}

func (b *BucketManage) GetUserBucket(u *protocol.MemoryUser, rate int64) *ratelimit.Bucket {
	if len(u.Email) > 0 && b.Users[u.Email] != nil {
		return b.Users[u.Email].Bucket
	} else {
		bucket := ratelimit.NewBucketWithRate(float64(rate), rate)
		bu := &Bucket{
			User:   u,
			Bucket: bucket,
		}
		b.Lock()
		defer b.Unlock()
		b.Users[u.Email] = bu
		return bucket
	}
}
func BucketMange() *BucketManage {
	return newBucketMange
}

var newBucketMange *BucketManage

func init() {
	newBucketMange = new(BucketManage)
	newBucketMange.Users = make(map[string]*Bucket)
}
