package rclone

import (
	"context"

	"github.com/loqutus/rdl/pkg/types"
	"github.com/rclone/rclone/fs/sync"
)

func Sync() {
	srcDir := types.RunConfig.TempDir
	sync.Sync(context.Context{}, srcDir, types.RunConfig.RClonePath)
}
