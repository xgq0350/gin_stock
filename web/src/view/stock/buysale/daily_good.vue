<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">涨幅前300</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
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
          height="calc(100vh - 180px)"
          @sort-change="sortChange"
          row-key="ID"
       >
        <el-table-column align="left" label="代码" prop="ts_code" width="80" />
        <el-table-column align="left" label="名称" prop="name" width="80" />
        <el-table-column align="left" label="概念" prop="gainian" width="280" />
        <el-table-column align="left" label="行业" prop="industry" width="60" />
        <el-table-column align="left" label="分类" prop="fenlei" width="60" />
        <el-table-column align="left" label="亮点" prop="liangdian" width="220" />
        <el-table-column align="left" label="涨幅" prop="pct_chg" width="60" />
        <el-table-column align="left" label="昨日" prop="close" width="60" />
        <el-table-column align="left" label="23MAX/min" prop="close_max_2023" :formatter="close2023" width="100" />
        <el-table-column align="left" label="22MAX/min" prop="close_max_2022" :formatter="close2022" width="100" />
        <el-table-column align="left" label="21MAX/min" prop="close_max_2021" :formatter="close2021" width="100" />
        <el-table-column align="left" label="申万" prop="shenwan" width="60" />
        <el-table-column align="left" label="市盈动" prop="shiyingdong" width="70" />
        <el-table-column align="left" label="财务" prop="caiwufenxi" width="60" />
        <el-table-column align="left" label="明细" min-width="160">
          <template #default="scope" >
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDialog(scope.row)"
            >明细</el-button>
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openWencai(scope.row)"
            >问财</el-button>
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openHangye(scope.row)"
            >行业</el-button>
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openShenwan(scope.row)"
            >申万</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="detailFormVisible" :before-close="closeDialogDetail" :title="codeName">
      <el-table :data="detailData"  width="1800">
        <el-table-column align="left" label="年" width="80" prop="year"  />
        <el-table-column align="left" label="月" width="60" prop="month" />
        <el-table-column align="left" label="幅度" min-width="300" prop="pct_chg" />
        <el-table-column align="left" label="涨跌4" min-width="60" prop="pct_popular4" />
        <el-table-column align="left" label="涨4" min-width="60" prop="pct_pos4" />
        <el-table-column align="left" label="涨跌2" min-width="60" prop="pct_popular2" />
        <el-table-column align="left" label="涨2" min-width="60" prop="pct_pos2" />
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialogDetail">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import { ref } from 'vue'
// import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import { getDetailPctChg,getDailyLimit } from '@/api/stockBuy'
import { close2023, close2021, close2022 } from '@/core/config'
import { toSQLLine } from '@/utils/stringFun'
import { getParams } from '@/core/config'

const form = ref({
  customerName: '',
  customerPhoneData: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const detailData = ref([])
const searchInfo = ref({})
const detailFormVisible = ref(false)
const code = ref('')
const industry = ref('')
const shenwan = ref('')
const params = ref({})

const getParamsSearch = () => {
  params.value = getParams()
  industry.value = unescape(decodeURI(params.value.industry))
  shenwan.value = unescape(decodeURI(params.value.shenwan))
}
getParamsSearch()

const onReset = () => {
  searchInfo.value = {}
  industry.value = {}
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  searchInfo.value = ref({})
  getTableData()
}

const openHangye = async(data) => {
  const table = await getDailyLimit({ industry: data.industry, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

const openShenwan = async(data) => {
  const table = await getDailyLimit({ shenwan: data.shenwan, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}
// 查询
const getTableData = async() => {
  if (industry.value !== '' && industry.value !== 'undefined') {
    searchInfo.value.industry = industry.value
  }
  if (shenwan.value !== '' && shenwan.value !== 'undefined') {
    searchInfo.value.shenwan = shenwan.value
  }
  const table = await getDailyLimit({ ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

// 详情
const getDetailPct = async() => {
  const table = await getDetailPctChg({ page: page.value, pageSize: pageSize.value, ts_code: code.value})
  if (table.code === 0) {
    detailData.value = table.data.list
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const openWencai = (data) => {
  window.open('http://iwencai.com/unifiedwap/result?w=' + data.ts_code + '&querytype=stock')
}

const openDialog = (data) => {
  // dialogFormVisible.value = true
  code.value = data.ts_code
  // const isHan = data.symbol[0].indexOf('0')
  // if (isHan === 0) {
  //   code.value = data.symbol + '.SZ'
  // } else {
  //   code.value = data.symbol + '.SH'
  // }
  detailFormVisible.value = true
  console.log(code.value)
  getDetailPct()
}

const closeDialogDetail =() => {
  detailFormVisible.value = false
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
</script>

<script>

export default {
  name: 'BuySale'
}
</script>

<style></style>
