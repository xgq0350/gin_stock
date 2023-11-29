<template>
  <div class="gainian">
    <warning-bar title="概念展示，查询可查看相关的股票" />
    <div class="gva-table-box">
      <!--      <div class="gva-btn-list">-->
      <!--        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
      <!--      </div>-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          height="calc(100vh - 180px )"
          row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="名词" prop="mingci" width="120" sortable="custom"/>
        <el-table-column align="left" label="名词定义" prop="mcdingyi" width="600" sortable="custom"/>
        <el-table-column align="left" fixed="right" label="关联股票" width="200">
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

    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="">

      <el-table :data="basicData"  width="800" @sort-change="sortChange" >
        <el-table-column align="left" label="名称" min-width="150" prop="name"  />
        <el-table-column align="left" label="行业" min-width="150" prop="industry" sortable="custom" />
        <el-table-column align="left" label="财务" min-width="150" prop="caiwufenxi" sortable="custom" />
        <el-table-column align="left" label="标识" min-width="150" prop="symbol"  />
        <el-table-column align="left" fixed="right" label="关联股票" width="100">
          <template #default="scope">
            <el-button
                type="primary"
                link
                @click="openPvgWeek(scope.row)"
            >获利比</el-button>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
import { getMingci, getStockBasic } from '@/api/stockBasic'
import {toSQLLine} from "@/utils/stringFun";

const form = ref({
  customerName: '',
  customerPhoneData: ''
})

const page = ref(1)
const gainian = ref("")
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const basicData = ref([])
const searchInfo = ref({})
const searchBasicInfo = ref({})

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchBasicInfo.value.orderKey = toSQLLine(prop)
    searchBasicInfo.value.desc = order === 'descending'
  }
  getStockBasicList()
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
  const table = await getMingci({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const getStockBasicList = async() => {
  // 查询
  const basic = await getStockBasic({ page: page.value, pageSize: pageSize.value, ...searchBasicInfo.value })
  if (basic.code === 0) {
    basicData.value = basic.data.list
    console.log(basic.data)
  }
}

getTableData()

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

const openPvgWeek = (key) => {
  window.open('/#/layout/stockMessage/zl_daily?ts_code=' + key.ts_code)
}

const openDialog = (key) => {
  dialogFormVisible.value = true
  page.value = 1
  pageSize.value = 100
  console.log(key)
  console.log(page.value)
  searchBasicInfo.value.gainian = key.mingci
  getStockBasicList()
}

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style>
.gainian {
.el-input-number {
  margin-left: 15px;
}
}
.tree-content{
  margin-top: 10px;
  height: calc(100vh - 148px);
  overflow: auto;
}

</style>
