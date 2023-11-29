<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="股票名称">
          <el-input v-model="searchInfo.ts_code" placeholder="名称" />
        </el-form-item>

        <el-form-item label="日期">
          <el-input v-model="searchInfo.trade_date" placeholder="日期" />
        </el-form-item>

        <el-form-item label="行业">
          <el-select v-model="searchInfo.industry" clearable filterable placeholder="请选择">
            <el-option
                v-for="item in methodOptions"
                :key="item.industry"
                :label="`${item.industry}`"
                :value="item.industry"
            />
          </el-select>
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
          tooltip-effect="dark"
          row-key="ID"
          height="calc(100vh - 180px)"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="行业" prop="industry" width="120" />
        <el-table-column align="left" label="排名" prop="rank1" width="500" />
        <el-table-column align="left" label="财务" prop="caiwufenxi" width="500" />
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
import { getIndustry } from '@/api/stockBasic'
import { getParams } from '@/core/config'
import { getStockRank } from '@/api/stockRank'

const form = ref({
  customerName: '',
  customerPhoneData: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(100)
const tableData = ref([])
const searchInfo = ref({})
const methodOptions = ref({})
const params = ref({})
const onReset = () => {
  searchInfo.value = {}
}

const getParamsSearch = () => {
  params.value = getParams()
  searchInfo.value.industry = unescape(decodeURI(params.value.industry))
}
getParamsSearch()

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

// console.log(this.)
// 查询
const getTableData = async() => {
  const table = await getStockRank({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.list.length
    page.value = 1
    pageSize.value = 100
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
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const fenleiOptions = ref([
  {
    value: '超大盘股',
    label: '超大盘股'
  },
  {
    value: '大盘股',
    label: '大盘股'
  },
  {
    value: '中盘股',
    label: '中盘股'
  },
  {
    value: '小盘股',
    label: '小盘股'
  }
])

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style>

</style>
