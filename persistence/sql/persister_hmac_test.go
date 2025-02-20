package sql

import (
	"context"
	"testing"

	"github.com/ory/x/configx"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/schema"

	"github.com/gobuffalo/pop/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/logrusx"

	"github.com/ory/kratos/driver/config"
)

type logRegistryOnly struct {
	l *logrusx.Logger
	c *config.Provider
}

func (l *logRegistryOnly) IdentityTraitsSchemas(ctx context.Context) schema.Schemas {
	panic("implement me")
}

func (l *logRegistryOnly) IdentityValidator() *identity.Validator {
	panic("implement me")
}

func (l *logRegistryOnly) Logger() *logrusx.Logger {
	if l.l == nil {
		l.l = logrusx.New("kratos", "testing")
	}
	return l.l
}

func (l *logRegistryOnly) Configuration(_ context.Context) *config.Provider {
	return l.c
}

func (l *logRegistryOnly) Audit() *logrusx.Logger {
	panic("implement me")
}

var _ persisterDependencies = &logRegistryOnly{}

func TestPersisterHMAC(t *testing.T) {
	conf := config.MustNew(logrusx.New("", ""), configx.SkipValidation())
	conf.MustSet(config.ViperKeySecretsDefault, []string{"foobarbaz"})
	c, err := pop.NewConnection(&pop.ConnectionDetails{URL: "sqlite://foo?mode=memory"})
	require.NoError(t, err)
	p, err := NewPersister(&logRegistryOnly{c: conf}, c)
	require.NoError(t, err)

	assert.True(t, p.hmacConstantCompare(context.Background(), "hashme", p.hmacValue(context.Background(), "hashme")))
	assert.False(t, p.hmacConstantCompare(context.Background(), "notme", p.hmacValue(context.Background(), "hashme")))
	assert.False(t, p.hmacConstantCompare(context.Background(), "hashme", p.hmacValue(context.Background(), "notme")))

	hash := p.hmacValue(context.Background(), "hashme")
	conf.MustSet(config.ViperKeySecretsDefault, []string{"notfoobarbaz"})
	assert.False(t, p.hmacConstantCompare(context.Background(), "hashme", hash))
	assert.True(t, p.hmacConstantCompare(context.Background(), "hashme", p.hmacValue(context.Background(), "hashme")))

	conf.MustSet(config.ViperKeySecretsDefault, []string{"notfoobarbaz", "foobarbaz"})
	assert.True(t, p.hmacConstantCompare(context.Background(), "hashme", hash))
	assert.True(t, p.hmacConstantCompare(context.Background(), "hashme", p.hmacValue(context.Background(), "hashme")))
	assert.NotEqual(t, hash, p.hmacValue(context.Background(), "hashme"))
}
