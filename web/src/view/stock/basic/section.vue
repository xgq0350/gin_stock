<template>
  <div>
    <warning-bar
        title="3月 / 6月 / 1年 / 2年 / 3年，均分10份"
    />
    <div class="">
      <div style="display:flex;flex-wrap:wrap">
        <div style="flex-basis: 18%">
          <warning-bar
              title="3 月"
          />
        <el-table
            :data="tableM3Data"
        >
          <el-table-column align="left" label="金额" prop="money" width="120" />
          <el-table-column align="left" label="数量" prop="num" width="120" />
        </el-table>
          <warning-bar
              :title="fudu3"
          />
      </div>

      <div style="flex-basis: 20%">
        <warning-bar
            title="6 月"
        />
        <el-table
            :data="tableM6Data"
            tooltip-effect="dark"
        >
          <el-table-column align="left" label="金额" prop="money" width="120" />
          <el-table-column align="left" label="数量" prop="num" width="120" />
        </el-table>
        <warning-bar
            :title="fudu6"
        />
      </div>

      <div style="flex-basis: 18%">
        <warning-bar
            title="1 年"
        />
        <el-table
            :data="tableM12Data"
            tooltip-effect="dark"
        >
          <el-table-column align="left" label="金额" prop="money" width="120" />
          <el-table-column align="left" label="数量" prop="num" width="120" />
        </el-table>
        <warning-bar
            :title="fudu12"
        />
      </div>

      <div style="flex-basis: 20%">
        <warning-bar
            title="2 年"
        />
        <el-table
            :data="tableM24Data"
            tooltip-effect="dark"
        >
          <el-table-column align="left" label="金额" prop="money" width="120" />
          <el-table-column align="left" label="数量" prop="num" width="120" />
        </el-table>
        <warning-bar
            :title="fudu24"
        />
      </div>

        <div style="flex-basis: 18%">
          <warning-bar
              title="3 年"
          />
          <el-table
              :data="tableM36Data"
              tooltip-effect="dark"
          >
            <el-table-column align="left" label="金额" prop="money" width="120" />
            <el-table-column align="left" label="数量" prop="num" width="120" />
          </el-table>
          <warning-bar
              :title="fudu36"
          />
        </div>

      </div>
    </div>
  </div>
</template>
<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import { getIndustry } from '@/api/stockBasic'
import { getZlLow, getSection } from '@/api/stockRank'
import WarningBar from '@/components/warningBar/warningBar.vue'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableM3Data = ref([])
const tableM6Data = ref([])
const tableM12Data = ref([])
const tableM24Data = ref([])
const tableM36Data = ref([])
const methodOptions = ref({})

const fudu36 = ref()
const fudu24 = ref()
const fudu12 = ref()
const fudu6 = ref()
const fudu3 = ref()

const paramsObj = ref({})
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
  const table = await getSection({ page: page.value, pageSize: pageSize.value, ts_code: paramsObj['ts_code'], type: 'm3' })
  if (table.code === 0) {
    tableM3Data.value = table.data.list
    if (table.data.list.length >= 1) {
      var list = table.data.list
      var start = list[0].money
      console.log(start)
      var end = list[list.length - 1].money
      console.log(end)
      fudu3.value = Math.round((end - start) * 100 / start) + '%'
    }
  }
}

const getTableDataM6 = async() => {
  const table = await getSection({ page: page.value, pageSize: pageSize.value, ts_code: paramsObj['ts_code'], type: 'm6' })
  if (table.code === 0) {
    tableM6Data.value = table.data.list
    if (table.data.list.length >= 1) {
      var list = table.data.list
      var start = list[0].money
      console.log(start)
      var end = list[list.length - 1].money
      console.log(end)
      fudu6.value = Math.round((end - start) * 100 / start) + '%'
    }
  }
}
getTableDataM6()

const getTableDataM12 = async() => {
  const table = await getSection({ page: page.value, pageSize: pageSize.value, ts_code: paramsObj['ts_code'], type: 'm12' })
  if (table.code === 0) {
    tableM12Data.value = table.data.list
    if (table.data.list.length >= 1) {
      var list = table.data.list
      var start = list[0].money
      console.log(start)
      var end = list[list.length - 1].money
      console.log(end)
      fudu12.value = Math.round((end - start) * 100 / start) + '%'
    }
  }
}
getTableDataM12()

const getTableDataM24 = async() => {
  const table = await getSection({ page: page.value, pageSize: pageSize.value, ts_code: paramsObj['ts_code'], type: 'm24' })
  if (table.code === 0) {
    tableM24Data.value = table.data.list
    if (table.data.list.length >= 1) {
      var list = table.data.list
      var start = list[0].money
      console.log(start)
      var end = list[list.length - 1].money
      console.log(end)
      fudu24.value = Math.round((end - start) * 100 / start) + '%'
    }
  }
}
getTableDataM24()

const getTableDataM36 = async() => {
  const table = await getSection({ page: page.value, pageSize: pageSize.value, ts_code: paramsObj['ts_code'], type: 'm36' })
  if (table.code === 0) {
    tableM36Data.value = table.data.list
    if (table.data.list.length >= 1) {
      var list = table.data.list
      var start = list[0].money
      console.log(start)
      var end = list[list.length - 1].money
      console.log(end)
      fudu36.value = Math.round((end - start) * 100 / start) + '%'
    }
  }
}
getTableDataM36()

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
