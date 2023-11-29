<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="路径">
          <el-input v-model="searchInfo.path" placeholder="路径" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.description" placeholder="描述" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!--    <warning-bar title="在资源权限中将此角色的资源权限清空" />-->
    <div class="gva-table-box">
      <!--      <div class="gva-btn-list">-->
      <!--        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
      <!--      </div>-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          height="calc(100vh - 180px)"
          row-key="ID">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="名称" prop="name" width="100" />
        <el-table-column align="left" label="日期" prop="trade_date" width="100" />
        <el-table-column align="left" label="意向价" prop="buy_rate" width="80" />
        <el-table-column align="left" label="收盘" prop="close" width="60" />
        <el-table-column align="left" label="23MAX/min" prop="close_max_2023" :formatter="close2023" width="110" />
        <el-table-column align="left" label="22MAX/min" prop="close_max_2022" :formatter="close2022" width="110" />
        <el-table-column align="left" label="21MAX/min" prop="close_max_2021" :formatter="close2021" width="110" />
        <el-table-column align="left" label="昨收" prop="pre_close" width="60" />
        <el-table-column align="left" label="涨跌" prop="pct_chg" width="80" />
        <el-table-column align="left" label="购买过" prop="record" width="80" :formatter="isTention"/>
        <el-table-column align="left" fixed="right" label="明细" width="200">
          <template #default="scope">
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDialog(scope.row)"
            >查询</el-button>
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openTen(scope.row)"
            >分段</el-button>
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
// import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import { getBuyIntention } from '@/api/stockBuy'
import { close2023, close2021, close2022 } from '@/core/config'

const form = ref({
  customerName: '',
  customerPhoneData: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const listData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const isTention = (data) => {
  console.log(data.record)
  if (data.record === true) {
    return '买过'
  } else {
    return '意向'
  }
}
// 查询
const getTableData = async() => {
  const table = await getBuyIntention({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 查询
const getStockBuyList = async(symbol) => {
  const table = await getBuyList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    listData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 详情
const getStockBuyDetail = async() => {
  const table = await getBuyDetail({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    detailData.value = table.data.list
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const dialogFormVisible = ref(false)
const detailFormVisible = ref(false)

// const updateCustomer = async(row) => {
//   const res = await getExaCustomer({ ID: row.ID })
//   type.value = 'update'
//   if (res.code === 0) {
//     form.value = res.data.customer
//     dialogFormVisible.value = true
//   }
// }
const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    customerName: '',
    customerPhoneData: ''
  }
}
// const deleteCustomer = async(row) => {
//   row.visible = false
//   const res = await deleteExaCustomer({ ID: row.ID })
//   if (res.code === 0) {
//     ElMessage({
//       type: 'success',
//       message: '删除成功'
//     })
//     if (tableData.value.length === 1 && page.value > 1) {
//       page.value--
//     }
//     getTableData()
//   }
// }
const openDialog = (data) => {
  // dialogFormVisible.value = true
  const code = ref()
  const isHan = data.symbol[0].indexOf('0')
  if (isHan === 0) {
    code.value = data.symbol + '.SZ'
  } else {
    code.value = data.symbol + '.SH'
  }
  console.log(code.value)
  window.open('#/layout/stockMessage/zl_daily?ts_code=' + code.value)
  // getStockBuyList(key.symbol)
}

const openTen = (data) => {
  const code = ref()
  const isHan = data.symbol[0].indexOf('0')
  if (isHan === 0) {
    code.value = data.symbol + '.SZ'
  } else {
    code.value = data.symbol + '.SH'
  }
  console.log(code.value)
  window.open('#/layout/stockHide/stockSection?ts_code=' + code.value)

}

const openDetail = (data) => {
  // detailFormVisible.value = true
  // getStockBuyDetail(key.symbol)
}

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style></style>
