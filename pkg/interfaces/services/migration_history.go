// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import "github.com/itcloudy/base-framework/pkg/models"

type IMigrationHistory interface {
	GetCurrentVersion() (version string, err error)
	FirstMigration() (err error)
	UpdateToOneVersion() (err error)
	GetAllListMigration() (migrates []models.MigrationHistory, err error)
}
