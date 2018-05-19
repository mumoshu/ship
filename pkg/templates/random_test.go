package templates

import (
	"testing"

	"github.com/replicatedcom/ship/pkg/test-mocks/logger"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	ctx := &StaticCtx{Logger: &logger.TestLogger{T: t}}
	str := ctx.RandomString(100)
	assert.Len(t, str, 100)
	assert.Regexp(t, DefaultCharset, str)

}
