package keystore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestN(t *testing.T) {
	println(LightScryptN, StandardScryptN)
}

func TestKeyStore(t *testing.T) {
	key, err := NewKey()

	require.NoError(t, err)

	_, err = WriteScryptKeyStore(key, "test")

	require.NoError(t, err)
}
