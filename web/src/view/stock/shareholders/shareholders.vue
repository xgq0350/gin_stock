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
        <el-form-item label="API组">
          <el-input v-model="searchInfo.apiGroup" placeholder="api组" />
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
          row-key="ID"
          height="calc(100vh - 180px)"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="代码" prop="symbol" width="60" />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="2023股东数" prop="holder_2023" width="300" />
        <el-table-column align="left" label="2022股东数" prop="holder_2022" width="300" />
        <el-table-column align="left" label="2021股东数" prop="holder_2021" width="200" />
        <el-table-column align="left" label="行业" prop="industry" width="120" />
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="客户名">
          <el-input v-model="form.customerName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="客户电话">
          <el-input v-model="form.customerPhoneData" autocomplete="off" />
        </el-form-item>
      </el-form>
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
import { getStockBasic, getBolders } from '@/api/stockBasic'

const form = ref({
  customerName: '',
  customerPhoneData: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
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

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getBolders({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

<style></style>
