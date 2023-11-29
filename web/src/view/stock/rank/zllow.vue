<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="股票名称">
          <el-input v-model="searchInfo.ts_code" placeholder="名称" />
        </el-form-item>
        <el-form-item label="日期">
          <el-input v-model="searchInfo.trade_date" placeholder="开始日期" />
        </el-form-item>

        <el-form-item label="获利占比">
          <el-input v-model="searchInfo.profit_percent" placeholder="获利占比" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          height="calc(100vh - 180px )"
          tooltip-effect="dark"
          row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="代码" prop="symbol" width="120" />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="行业" prop="industry" width="120" />
        <el-table-column align="left" label="23MAX/23min" prop="close_max_2023" :formatter="close2023" width="150" />
        <el-table-column align="left" label="22MAX/22min" prop="close_max_2022" :formatter="close2022" width="150" />
        <el-table-column align="left" label="21MAX/21min" prop="close_max_2021" :formatter="close2021" width="150" />
        <el-table-column align="left" label="收盘" prop="close" width="60" />
        <el-table-column align="left" label="昨收盘" prop="pre_close" width="60" />
        <el-table-column align="left" label="涨跌" prop="pct_chg" width="100" />
        <el-table-column align="left" fixed="right" label="明细" width="100">
          <template #default="scope">
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDialog(scope.row)"
            >获利占比</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>
<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import { getIndustry } from '@/api/stockBasic'
import { getZlLow } from '@/api/stockRank'
import { close2023, close2021, close2022 } from '@/core/config'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const methodOptions = ref({})

const onReset = () => {
  searchInfo.value = {}
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// console.log(this.)
// 查询
const getTableData = async() => {
  const table = await getZlLow({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const getIndustryData = async() => {
  // 查询
  const options = await getIndustry()
  if (options.code === 0) {
    methodOptions.value = options.data.list
  }
}

getTableData()
getIndustryData()

const dialogFormVisible = ref(false)
const type = ref('')

const openDialog = (data) => {
  console.log(data)
  type.value = 'create'
  dialogFormVisible.value = true
  window.location.href = '/#/layout/stockMessage/zl_daily?ts_code=' + data.symbol
}

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style>
:deep(.el-table)  {
  overflow: visible;
}

:deep(.el-table__header-wrapper)  {
  position: sticky;
  top: 100px;
z-index: 10;
}

</style>
