<template>
  <div>
    <warning-bar
      title="基本信息，关联华泰数据"
    />
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="标识">
          <el-input v-model="searchInfo.symbol" placeholder="标识" />
        </el-form-item>
        <el-form-item label="标识" disabled="false">
          <el-input v-model="searchInfo.ts_code" placeholder="标识" />
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="名称" />
        </el-form-item>
        <el-form-item label="财务">
          <el-select v-model="searchInfo.caiwufenxi" clearable filterable placeholder="请选择">
            <el-option
              v-for="item in cwfxOptions"
              :key="item.label"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
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

        <el-form-item label="分类">
          <el-select v-model="searchInfo.fenlei" clearable filterable placeholder="请选择">
            <el-option
              v-for="item in fenleiOptions"
              :key="item.label"
              :label="`${item.value}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button icon="refresh" @click="stockRank">行业排行</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!--    <warning-bar title="在资源权限中将此角色的资源权限清空" />-->
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @sort-change="sortChange"
      >
        <el-table-column align="left" label="标识" prop="symbol" width="80" />
        <el-table-column align="left" label="名称" prop="name" width="80" />
        <el-table-column align="left" label="行业" prop="industry" width="80" sortable="custom"/>
<!--        <el-table-column align="left" label="城市" prop="area" width="100" />-->
        <el-table-column align="left" label="盈利" prop="profit_percent" width="70" sortable="custom"/>
        <el-table-column align="left" label="跌幅" prop="lat_low" width="70" sortable="custom"/>
        <el-table-column align="left" label="最小" prop="lowerest_price" width="60" />
        <el-table-column align="left" label="昨收" prop="close" width="60" />
        <el-table-column align="left" label="概念" prop="gainian" width="180" show-overflow-tooltip/>
        <el-table-column align="left" label="财务" prop="caiwufenxi" width="80" />
        <el-table-column align="left" label="分类" prop="fenlei" width="60" />
        <el-table-column align="left" label="市盈静" prop="shiyingjing" width="60"  sortable="custom"/>
        <el-table-column align="left" label="市盈动" prop="shiyingdong" width="60"  sortable="custom"/>
        <el-table-column align="left" label="主营" prop="zhuying" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="申万" prop="shenwan" width="60" show-overflow-tooltip/>
        <el-table-column align="left" label="亮点" prop="liangdian" width="170" />
        <el-table-column align="center" label="操作" min-width="240">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="huoli(scope.row)">获利比</el-button>
            <el-button type="primary" link icon="edit" @click="hangyepaiming(scope.row)">股票排名</el-button>
            <el-button type="primary" link icon="edit" @click="zhongdianhangye(scope.row)">重点行业</el-button>
            <el-button type="primary" link icon="edit" @click="openDialog(scope.row)">行业</el-button>
            <el-button type="primary" link icon="edit" @click="zhongdianshenwan(scope.row)">重点申万</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100, 500]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          height="calc(100vh - 180px)"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">

    </el-dialog>
  </div>
</template>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { cwfxOptions, getParams } from '@/core/config'
import { getStockBasic, getIndustry } from '@/api/stockBasic'
import { toSQLLine } from '@/utils/stringFun'

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

const huoli = (val) => {
  window.open('/#/layout/stockMessage/zl_daily?ts_code=' + val.ts_code)
}

const zhongdianhangye = (val) => {
  window.open('/#/layout/stockMessage/dailyGood?industry=' + val.industry)
}

const zhongdianshenwan = (val) => {
  window.open('/#/layout/stockMessage/dailyGood?shenwan=' + val.shenwan)
}

const hangyepaiming = (val) => {
  window.open('/#/layout/stockMessage/stock_rank?industry=' + escape(val.industry))
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 100
  getTableData()
}

// const stockRank = () => {
//   console.log(searchInfo.value.industry)
//   window.location.href = '/#/layout/stockMessage/stock_rank?industry=' + searchInfo.value.industry
// }

const getParamsSearch = () => {
  const params = ref({})
  params.value = getParams()
  searchInfo.value.symbol = params.value.symbol
  searchInfo.value.ts_code = params.value.ts_code
  if (params.value.industry !== undefined) {
    searchInfo.value.industry = unescape(decodeURI(params.value.industry))
  }
  console.log(params.value.industry)
}
getParamsSearch()
// 查询
const getTableData = async() => {
  console.log(searchInfo.value.symbol)
  const table = await getStockBasic({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    customerName: '',
    customerPhoneData: ''
  }
}

const openDialog = (data) => {
  dialogFormVisible.value = true
  getIndustryLatest(data)
}

const getIndustryLatest = async(data) => {
  const table = await getStockBasic({ industry: data.industry })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
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
  name: 'StockBasic'
}
</script>

<style></style>
