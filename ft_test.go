package ft_test

import (
	"testing"

	"github.com/spobly/ft"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCategory(t *testing.T) {
	extension := ".jpg"
	want := "Images"

	got, err := ft.GetCategory(extension)
	require.NoError(t, err)
	assert.Equal(t, got, want)
}

func TestInvalidExtension(t *testing.T) {
	extension := ".img"

	_, err := ft.GetCategory(extension)

	require.Error(t, err)
	require.EqualError(t, err, ft.ErrNotFound.Error())
}
