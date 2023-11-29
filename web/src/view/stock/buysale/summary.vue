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
          height="calc(100vh - 180px)"
          row-key="ID">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="符号" prop="ts_code" width="120" />
        <el-table-column align="left" label="名称" prop="name" width="100" />
        <el-table-column align="left" label="行业" prop="industry" width="100" />
        <el-table-column align="left" label="持有" prop="is_holder" width="60" :formatter="isHolder"/>
        <el-table-column align="left" label="清仓" prop="buy_num" width="60" />
        <el-table-column align="left" label="清仓日" prop="last_sale" width="120" />
        <el-table-column align="left" label="买入" prop="buy_his_rate" width="100" />
        <el-table-column align="left" label="卖出" prop="sale_his_rate" width="100" />
        <el-table-column align="left" label="最高" prop="high" width="60" />
        <el-table-column align="left" label="最低" prop="low" width="80" />
        <el-table-column align="left" label="收盘" prop="close" width="60" />
        <el-table-column align="left" label="数量"  width="60" :formatter="cishuliang"/>
        <el-table-column align="left" label="明细" width="200">
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
        <el-table-column align="left" label="收益" min-width="120" prop="shouyi" />
        <el-table-column align="left" label="买入日" min-width="120" prop="buy_date" />
        <el-table-column align="left" label="卖出日" min-width="120" prop="sale_date" />
        <el-table-column align="left" label="买入" min-width="60" prop="buy_rate" />
        <el-table-column align="left" label="卖出" min-width="60" prop="sale_rate"  />
        <el-table-column align="left" label="持有" min-width="120" prop="is_holder"  />
        <el-table-column align="left" fixed="right" label="明细" width="120">
          <template #default="scope">
            <el-button
                icon="edit"
                type="primary"
                link
                @click="openDetail(scope.row)"
            >查询</el-button>
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

    <el-dialog v-model="detailFormVisible" :before-close="closeDialogDetail" :title="codeName" width="1000">
      <el-table :data="detailData"  width="1200">
        <el-table-column align="left" label="日期" min-width="80" prop="trade_date"  />
        <el-table-column align="left" label="类型" min-width="60" prop="trade_date" :formatter="handleType"/>
        <el-table-column align="left" label="金额" min-width="80" prop="money" />
        <el-table-column align="left" label="幅度" min-width="60" prop="zhang"/>
        <el-table-column align="left" label="金额" min-width="100" prop="leiji"/>
        <el-table-column align="left" label="成本" min-width="80" prop="pingjun"/>
        <el-table-column align="left" label="跌3%/6%/10%" min-width="150" :formatter="fall" />
        <el-table-column align="left" label="数量" min-width="60" prop="num"  />
        <el-table-column align="left" label="最低" min-width="100" prop="low"  />
        <el-table-column align="left" label="最高" min-width="60" prop="high"  />
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
import { getBuySummary, getBuyList, getBuyDetail } from '@/api/stockBuy'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const listData = ref([])
const detailData = ref([])
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

const cishuliang = (data) => {
  return Math.round(40000 / data.close / 100) * 100
}

const handleType = (data) => {
  if (data.type === '1') {
    return '买入'
  } else {
    return '卖出'
  }
}

const fall = (data) => {
  return Math.round(data.money * 0.97 * 100) / 100 + '/' + Math.round(data.money * 0.94 * 100) / 100 + '/' + Math.round(data.money * 0.9 * 100) / 100
}

// 查询
const getTableData = async() => {
  const table = await getBuySummary({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    console.log(table.data.list)
    tableData.value = table.data.list
    console.log(table.data.list)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 查询
const getStockBuyList = async(symbol) => {
  const table = await getBuyList({ page: page.value, pageSize: pageSize.value, symbol: symbol })
  if (table.code === 0) {
    listData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 详情
const getStockBuyDetail = async(data) => {
  console.log(data)
  const table = await getBuyDetail({ symbol: data.ts_code, buy_date: data.buy_date, sale_date: data.sale_date })
  if (table.code === 0) {
    var kk = Math.round((close.value - table.data.list[0].money) * 100 / table.data.list[0].money)
    console.log(close.value)
    console.log(table.data.list[0].money)
    console.log(kk)
    codeName.value = codeName.value + '/ ' + kk + '%'
    detailData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const dialogFormVisible = ref(false)
const detailFormVisible = ref(false)

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

const closeDialogDetail = () => {
  detailFormVisible.value = false
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
const openDialog = (key) => {
  dialogFormVisible.value = true
  codeName.value = key.name + '-' + key.industry + '-' + key.close
  close.value = key.close
  getStockBuyList(key.symbol)
}

const openBasic = (key) => {
  console.log(key.symbol)
  window.open('/#/layout/stockMessage/stockBasic?symbol=' + key.symbol)
}

const openDetail = (key) => {
  detailFormVisible.value = true
  getStockBuyDetail(key)
}

const openFenduan = (key) => {
  window.open('#/layout/stockHide/stockSection?ts_code=' + key.ts_code)
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
