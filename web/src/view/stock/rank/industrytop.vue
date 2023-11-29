<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="股票名称">
          <el-input v-model="searchInfo.ts_code" placeholder="名称" />
        </el-form-item>
        <el-form-item label="日期">
          <el-input v-model="searchInfo.trade_date" placeholder="开始日期" />
        </el-form-item>

        <el-form-item label="获利占比">
          <el-input v-model="searchInfo.profit_percent" placeholder="获利占比" />
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
        <el-table-column align="left" label="日期" prop="trade_date" width="120" />
        <el-table-column align="left" label="Top1" prop="top1" width="120" />
        <el-table-column align="left" label="Top2" prop="top2" width="120" />
        <el-table-column align="left" label="Top3" prop="top3" width="120" />
        <el-table-column align="left" label="Top4" prop="top4" width="120" />
        <el-table-column align="left" label="Top5" prop="top5" width="120" />
        <el-table-column align="left" label="Top6" prop="top6" width="120" />
        <el-table-column align="left" label="Top7" prop="top7" width="120" />
        <el-table-column align="left" label="Top8" prop="top8" width="120" />
        <el-table-column align="left" label="Top9" prop="top9" width="120" />
        <el-table-column align="left" label="Top10" prop="top10" width="120" />
        <el-table-column align="left" label="Top11" prop="top11" width="120" />
        <el-table-column align="left" label="Top12" prop="top12" width="120" />
        <el-table-column align="left" label="Top13" prop="top13" width="120" />
        <el-table-column align="left" label="Top14" prop="top14" width="120" />
        <el-table-column align="left" label="Top15" prop="top15" width="120" />
        <el-table-column align="left" label="Top16" prop="top16" width="120" />
        <el-table-column align="left" label="Top17" prop="top17" width="120" />
        <el-table-column align="left" label="Top18" prop="top18" width="120" />
        <el-table-column align="left" label="Top19" prop="top19" width="120" />
        <el-table-column align="left" label="Top20" prop="top20" width="120" />
        <el-table-column align="left" label="Top21" prop="top21" width="120" />
        <el-table-column align="left" label="Top22" prop="top22" width="120" />
        <el-table-column align="left" label="Top23" prop="top23" width="120" />
        <el-table-column align="left" label="Top24" prop="top24" width="120" />
        <el-table-column align="left" label="Top25" prop="top25" width="120" />
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
import { ref } from 'vue'
// import { ElMessage } from 'element-plus'
// import { formatDate } from '@/utils/format'
import { getIndustry } from '@/api/stockBasic'
import { getIndustryTop } from '@/api/stockRank'

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

// console.log(this.)
// 查询
const getTableData = async() => {
  const table = await getIndustryTop({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
  }
}

const getIndustryData = async() => {
  // 查
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
// const enterDialog = async() => {
//   let res
//   switch (type.value) {
//     case 'create':
//       res = await createExaCustomer(form.value)
//       break
//     case 'update':
//       res = await updateExaCustomer(form.value)
//       break
//     default:
//       res = await createExaCustomer(form.value)
//       break
//   }
//
//   if (res.code === 0) {
//     closeDialog()
//     getTableData()
//   }
// }
// const openDialog = () => {
//   type.value = 'create'
//   dialogFormVisible.value = true
// }
//
// const fenleiOptions = ref([
//   {
//     value: '超大盘股',
//     label: '超大盘股'
//   },
//   {
//     value: '大盘股',
//     label: '大盘股'
//   },
//   {
//     value: '中盘股',
//     label: '中盘股'
//   },
//   {
//     value: '小盘股',
//     label: '小盘股'
//   }
// ])

</script>

<script>

export default {
  name: 'StockBasic'
}
</script>

<style></style>
