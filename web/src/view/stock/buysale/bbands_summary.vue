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
          <el-button type="primary" icon="search" @click="onLowerSubmit">1月低lower</el-button>
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
        <el-table-column align="left" label="符号" prop="ts_code" width="120" />
        <el-table-column align="left" label="名称" prop="name" width="100" />
        <el-table-column align="left" label="行业" prop="industry" width="100" />
        <el-table-column align="left" label="lower" prop="lower" width="100" />
        <el-table-column align="left" label="close" prop="close" width="100" />
        <el-table-column align="left" label="今日" prop="value" width="100" />
        <el-table-column align="left" label="日期" prop="trade_date" width="120" />
        <el-table-column align="left" label="昨涨" prop="pct_chg" width="100" />
        <el-table-column align="left" label="上证" prop="szzs_c" width="100" />
        <el-table-column align="left" fixed="right" label="明细" width="240">
          <template #default="scope">
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDialog(scope.row)"
            >查询</el-button>
            <el-button
                type="primary"
                link
                @click="openBasic(scope.row)"
            >基本</el-button>
            <el-button
                type="primary"
                link
                @click="openFenduan(scope.row)"
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

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="codeName">
      <el-table :data="listData" width="1200" >
        <el-table-column align="left" label="日期" min-width="120" prop="trade_date" />
        <el-table-column align="left" label="close" min-width="80" prop="close" />
        <el-table-column align="left" label="lower" min-width="80" prop="lower" />
        <el-table-column align="left" label="pct_chg" min-width="100" prop="pct_chg" />
        <el-table-column align="left" label="上证" min-width="100" prop="szzs_c"  />
        <el-table-column align="left" fixed="right" label="明细" width="240">
          <template #default="scope">
            <el-button
                type="primary"
                link
                @click="openList(scope.row)"
            >细节</el-button>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="listVisible" :before-close="closeList" :title="codeName">
      <el-table :data="bandListData" width="800" >
        <el-table-column align="left" label="日期" min-width="120" prop="trade_date" />
        <el-table-column align="left" label="close" min-width="80" prop="close" />
        <el-table-column align="left" label="lower" min-width="80" prop="lower" />
        <el-table-column align="left" label="pct_chg" min-width="80" prop="pct_chg" />
        <el-table-column align="left" label="上证" min-width="80" prop="szzs_c"  />
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeList">取 消</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="mainVisible" :before-close="closeMainVisible" title="9月1号开始累计低于布林带低点值">
      <el-table :data="mainListData" width="800" >
        <el-table-column align="left" label="ts_code" min-width="120" prop="ts_code" />
        <el-table-column align="left" label="name" min-width="120" prop="name" />
        <el-table-column align="left" label="industry" min-width="120" prop="industry" />
        <el-table-column align="left" label="num" min-width="80" prop="num" />
        <el-table-column align="left" fixed="right" label="明细" width="240">
          <template #default="scope">
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDialog(scope.row)"
            >查询</el-button>
            <el-button
                type="primary"
                link
                @click="openBasic(scope.row)"
            >基本</el-button>
            <el-button
                type="primary"
                link
                @click="openFenduan(scope.row)"
            >分段</el-button>

          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeList">取 消</el-button>
        </div>
      </template>
    </el-dialog>



  </div>
</template>

<script setup>

import { ref } from 'vue'
// import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import { getBandSummary, getBandDetail, getBandList } from '@/api/stockBuy'

const page = ref(1)
const mainVisible = ref(false)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const listData = ref([])
const bandListData = ref([])
const detailData = ref([])
const mainListData = ref([])
const searchInfo = ref({})
const codeName = ref('')
const close = ref(0.0)

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

const isHolder = (data) => {
  if (data.is_holder === true) {
    return '持有'
  } else {
    return '清仓'
  }
}

// 查询
const getTableData = async() => {
  const table = await getBandSummary({ ...searchInfo.value })
  if (table.code === 0) {
    console.log(table.data.list)
    tableData.value = table.data.list
    console.log(table.data.list)
    total.value = table.data.list.length
    page.value = table.data.page
    pageSize.value = table.data.list.length
  }
}

// 查询
const getStockBuyList = async(tsCode) => {
  const table = await getBandDetail({ ts_code: tsCode })
  if (table.code === 0) {
    listData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 详情
const getBandListData = async(data) => {
  console.log(data)
  const table = await getBandList({ ts_code: data })
  if (table.code === 0) {
    bandListData.value = table.data.list
    total.value = table.data.list.length
  }
}

getTableData()

const dialogFormVisible = ref(false)
const detailFormVisible = ref(false)
const listVisible = ref(false)

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
}

const closeMainVisible = () => {
  mainVisible.value = false
}

const closeList = () => {
  listVisible.value = false
}
const onLowerSubmit = async() => {
  mainVisible.value = true
  const table = await getBandSummary({ type: 2 })
  if (table.code === 0) {
    mainListData.value = table.data.list
  }
}

const openDialog = (key) => {
  dialogFormVisible.value = true
  codeName.value = key.name + '-' + key.industry + '-' + key.close
  close.value = key.close
  getStockBuyList(key.ts_code)
}
const openList = (key) => {
  listVisible.value = true
  getBandListData(key.ts_code)
}

const openBasic = (key) => {
  window.open('/#/layout/stockMessage/stockBasic?ts_code=' + key.ts_code)
}

const openFenduan = (key) => {
  window.open('#/layout/stockHide/stockSection?ts_code=' + key.ts_code)
}

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style></style>
