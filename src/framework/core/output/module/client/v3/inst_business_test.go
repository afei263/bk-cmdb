package v3_test

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/config"
	"configcenter/src/framework/core/output/module/client/v3"
	"configcenter/src/framework/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchBusiness(t *testing.T) {
	cli := v3.New(config.Config{"supplierAccount": "0", "user": "build_user", "ccaddress": "http://test.apiserver:8080"}, nil)

	cond := common.CreateCondition().Field("bk_biz_name").Eq("蓝鲸")
	rets, err := cli.Business().SearchBusiness(cond)
	t.Logf("search business result: %v", rets)
	assert.NoError(t, err)
	assert.NotEmpty(t, rets)
}

func TestBusiness(t *testing.T) {
	cli := v3.New(config.Config{"supplierAccount": "0", "user": "build_user", "ccaddress": "http://test.apiserver:8080"}, nil)

	data := types.MapStr{
		"bk_biz_name":       "testBiz",
		"bk_biz_maintainer": "build_user",
	}
	id, err := cli.Business().CreateBusiness(data)
	t.Logf("search business result: %v", id)
	require.NoError(t, err)
	require.True(t, id > 0)

	data.Set("bk_biz_maintainer", "test_user")
	err = cli.Business().UpdateBusiness(data, id)
	require.NoError(t, err)

	cond := common.CreateCondition().Field("bk_biz_name").Eq("testBiz")
	rets, err := cli.Business().SearchBusiness(cond)
	require.NoError(t, err)
	require.NotEmpty(t, rets)
	require.Equal(t, "test_user", rets[0]["bk_biz_maintainer"])

	err = cli.Business().DeleteBusiness(id)
	require.NoError(t, err)

}