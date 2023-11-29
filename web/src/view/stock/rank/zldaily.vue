<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="股票名称">
          <el-input v-model="searchInfo.ts_code" placeholder="名称" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.trade_date" placeholder="开始日期" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-card-box display_n" >
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-dot">{{ basic.name }} {{ basic.industry }}</div>

          <div class="gva-top-card-left-rows">
            <el-row>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center" style="font-size: large">
                  <el-icon class="dashboard-icon">
                    <sort />
                  </el-icon>
                  max_180_90_60d: {{basic.max_180d}} / {{basic.max_90d}} / {{basic.max_60d}}
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center" style="font-size: large">
                  <el-icon class="dashboard-icon">
                    <avatar />
                  </el-icon>
                    min_180_90_60d:   {{basic.min_180d}} / {{basic.min_90d}} / {{basic.min_60d}}
                </div>
              </el-col>
            </el-row>
          </div>

          <div class="gva-top-card-left-rows">
            <el-row>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center" style="font-size: large">
                  <el-icon class="dashboard-icon">
                    <sort />
                  </el-icon>
                  max_23_22_21: {{basic.close_max_2023}} / {{basic.close_max_2022}} / {{basic.close_max_2021}}
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center" style="font-size: large">
                  <el-icon class="dashboard-icon">
                    <avatar />
                  </el-icon>
                  max_23_22_21: {{basic.close_min_2023}} / {{basic.close_min_2022}}/ {{basic.close_min_2021}}
                </div>
              </el-col>
            </el-row>
          </div>

        </div>
      </div>
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
        <el-table-column align="left" label="年" prop="year" width="60" />
        <el-table-column align="left" label="周" prop="week" width="60" />
        <el-table-column align="left" label="获利占比" prop="profit_percent" width="120" />
        <el-table-column align="left" label="平均值" prop="avg_close" width="120" />
        <el-table-column align="left" label="买入跌1%" prop="buy1" width="120" />
        <el-table-column align="left" label="卖出" prop="sale1" width="60" />
        <el-table-column align="left" label="买入跌2%" prop="buy2" width="120" />
        <el-table-column align="left" label="卖出" prop="sale2" width="60" />
        <el-table-column align="left" label="一周合并" prop="concat_close" width="360" />
      </el-table>
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
import { getBasicInfo, getIndustry } from '@/api/stockBasic'
import { getZlDaily } from '@/api/stockRank'

const form = ref({
  customerName: '',
  customerPhoneData: ''
})
const paramsObj = ref({});
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const basic = ref({})
const methodOptions = ref({})

const onReset = () => {
  searchInfo.value = {}
}
const getParams = () => {
  var url = window.location.href;
  const params = url.split('?')[1];
  if (params) {
    let paramsArr = params.split('&');
    for (let i = 0; i < paramsArr.length; i++) {
      let param = paramsArr[i].split('=');
      paramsObj[param[0]] = param[1];
    }
  }
}

getParams()
console.log(paramsObj['ts_code'])

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()

}

// 查询
const getTableData = async() => {
  const table = await getZlDaily({ page: page.value, pageSize: pageSize.value, ts_code: paramsObj['ts_code'] })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.list
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 查询
const getStockBasic = async() => {
  const table = await getBasicInfo({ ts_code: paramsObj['ts_code'] })
  if (table.code === 0) {
    basic.value = table.data.list
    console.log(basic.value)
    console.log(basic['symbol'])
    console.log(basic.value.symbol)
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
getStockBasic()

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
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createExaCustomer(form.value)
      break
    case 'update':
      res = await updateExaCustomer(form.value)
      break
    default:
      res = await createExaCustomer(form.value)
      break
  }

  if (res.code === 0) {
    closeDialog()
    getTableData()
  }
}
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
.display_n {
  /*display: none;*/
}
</style>
