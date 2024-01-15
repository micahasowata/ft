package ft_test

import (
	"testing"

	"github.com/spobly/ft/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCategory(t *testing.T) {
	extension := ".jpg"
	want := "Images"

	got, err := ft.GetFileCategory(extension)
	require.NoError(t, err)
	assert.Equal(t, got, want)
}

func TestMissingExtension(t *testing.T) {
	extension := ".img"

	_, err := ft.GetFileCategory(extension)

	require.Error(t, err)
	require.EqualError(t, err, ft.ErrNotFound.Error())
}

func TestInvalidExtension(t *testing.T) {
	extension := "png"
	_, err := ft.GetFileCategory(extension)

	require.Error(t, err)
	require.EqualError(t, err, ft.ErrInvalidFormat.Error())
}
