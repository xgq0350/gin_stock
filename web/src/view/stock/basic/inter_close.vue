<template>
  <div class="zhishu">
    <warning-bar title="间隔30天，180天，年度" />
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
        <el-table-column align="left" label="标识" prop="ts_code" width="120" sortable="custom"/>
        <el-table-column align="left" label="名称" prop="name" width="100" sortable="custom"/>
        <el-table-column align="left" label="行业" prop="industry" width="80" sortable="custom"/>
        <el-table-column align="left" label="今日" prop="close" width="80" sortable="custom"/>
        <el-table-column align="left" label="180最大" prop="max_180d" width="100" sortable="custom"/>
        <el-table-column align="left" label="180最小" prop="min_180d" width="60" />
        <el-table-column align="left" label="90最大" prop="max_90d" width="60" />
        <el-table-column align="left" label="90最小" prop="min_90d" width="60" />
        <el-table-column align="left" label="60最大" prop="min_60d" width="60" />
        <el-table-column align="left" label="60最小" prop="min_60d" width="60" />
        <el-table-column align="left" label="year" prop="year_close" width="60" />
        <el-table-column align="left" label="y_chg" prop="year_chg" width="60" />
        <el-table-column align="left" label="y_chg" prop="month_close" width="60" />
        <el-table-column align="left" label="m_chg" prop="month_chg" width="60" />

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
import {getStockYearClose} from '@/api/stockRank';

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
  const table = await getStockYearClose({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  window.open('/#/layout/stockMessage/zl_daily?ts_code=' + val.ts_code)
}

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style>
</style>
