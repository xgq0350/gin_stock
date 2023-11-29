<template>
  <div class="zhishu">
    <warning-bar title="指数展示，上证100，沪深300，中证100，深证300" />
    <div class="gva-table-box">
      <!--      <div class="gva-btn-list">-->
      <!--        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
      <!--      </div>-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="ID"
          height="calc(100vh - 180px)"
          @sort-change="sortChange"
      >
<!--        <el-table-column type="selection" width="55" />-->
        <el-table-column align="left" label="标识" prop="code_name" width="120" sortable="custom"/>
        <el-table-column align="left" label="指数" prop="zs_name" width="80" sortable="custom"/>
        <el-table-column align="left" label="名称" prop="name" width="100" sortable="custom"/>
        <el-table-column align="left" label="行业" prop="industry" width="80" sortable="custom"/>
        <el-table-column align="left" label="财务" prop="caiwufenxi" width="100" sortable="custom"/>
        <el-table-column align="left" label="收盘" prop="close" width="60" />
        <el-table-column align="left" label="23MAX/23min" prop="close_max_2023" :formatter="close2023" width="120" />
        <el-table-column align="left" label="22MAX/22min" prop="close_max_2022" :formatter="close2022" width="120" />
        <el-table-column align="left" label="21MAX/21min" prop="close_max_2021" :formatter="close2021" width="120" />
        <el-table-column align="left" label="昨收盘" prop="pre_close" width="80" />
        <el-table-column align="left" label="涨跌" prop="pct_chg" width="100" :formatter="zhangdie"/>
        <el-table-column align="left" fixed="right" label="关联" width="200">
          <template #default="scope">
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDialog(scope.row)"
            >查询</el-button>
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
import {
  createExaCustomer,
  updateExaCustomer,
  // deleteExaCustomer,
  // getExaCustomer,
} from '@/api/customer'
import { ref } from 'vue'
// import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { getZhishu } from '@/api/stockBasic'
import { toSQLLine } from '@/utils/stringFun'
import { close2023, close2021, close2022 } from '@/core/config'

const form = ref({
  customerName: '',
  customerPhoneData: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const zhangdie = (data) => {
  return Math.round(data.pct_chg * 100 ) / 100
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    page.value = 1
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// const onSubmit = () => {
//   page.value = 1
//   pageSize.value = 10
//   getTableData()
// }

// 查询
const getTableData = async() => {
  const table = await getZhishu({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
  }
}

getTableData()

const dialogFormVisible = ref(false)

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

const openDialog = (val) => {
  window.location.href = '/#/layout/stockMessage/zl_daily?ts_code=' + val.ts_code
}

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style>
</style>
