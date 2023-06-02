<template>
  <div>
    <div class="gva-search-box">
      <!-- <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form> -->
      <el-form :model="searchInfo" :inline="true" label-width="68px">
        <el-form-item label="落地页：">
          <el-select
            v-model="searchInfo.page_name"
            placeholder="请选择落地页名"
            filterable
            allow-create
            clearable
            style="width: 120px"
          >
            <el-option
              v-for="(pageName, index) in searchOptions.pageName"
              :key="index"
              :label="pageName"
              :value="pageName"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="工单名：">
          <el-select
            v-model="searchInfo.order_name"
            placeholder="请选择工单"
            filterable
            allow-create
            clearable
            style="width: 120px"
          >
            <el-option
              v-for="(orderName, index) in searchOptions.orderName"
              :key="index"
              :label="orderName"
              :value="orderName"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态:">
          <el-select
            v-model="searchInfo.state"
            placeholder="请选择状态"
            style="width: 120px"
            clearable
          >
            <el-option
              v-for="(item,key) in stateOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="号码:">
          <el-input v-model="searchInfo.searchNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button type="primary" icon="edit" :disabled="!multipleSelection.length" @click="updateRunNumsFunc">变更</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDateShort(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="号码" prop="num" width="180" />
        <el-table-column align="left" label="状态" prop="state" width="80">
          <template #default="scope">
            <span :class="{'status-0': scope.row.state === 0, 'status-1': scope.row.state === 1, 'status-2': scope.row.state === 2, 'status-3': scope.row.state === 3, 'status-4': scope.row.state === 4}">
              {{ filterDict(scope.row.state, stateOptions) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="号码类型" prop="numType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.numType,phone_typeOptions) }}
          </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="用户id" prop="userId" width="80" /> -->
        <el-table-column align="left" label="用户数" prop="userNum" width="80" />
        <el-table-column align="left" label="限制进粉" prop="eachEnterNum" width="80" />
        <el-table-column align="left" label="工单名" prop="orderName" width="220" />
        <el-table-column align="left" label="落地页名" prop="pageName" width="220" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateRunNumFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="100px">
        <el-form-item v-show="type==='create'" label="工单:" prop="num">
          <el-select
            v-model="formData.orderName"
            placeholder="请选择工单"
            clearable
          >
            <el-option
              v-for="(orderName, index) in orderNameOptions"
              :key="index"
              :label="orderName"
              :value="orderName"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-show="type === 'update' && type !== 'updates'" label="号码:" prop="num">
          <el-input v-model="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-show="type === 'create' && type !== 'updates'" label="号码" prop="numsText">
          <el-input v-model.lazy="formData.numsText" type="textarea" :rows="3" :clearable="true" placeholder="请输入号码 可同时用回车或，或,分割" />
        </el-form-item>
        <el-form-item label="状态:" prop="state">
          <el-select v-model="formData.state" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item v-show="type !== 'updates'" label="号码类型:" prop="numType">
          <el-select v-model="formData.numType" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in phone_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="进粉限制:" prop="eachEnterNum">
          <el-input v-model.number="formData.eachEnterNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="打招呼内容" prop="sayHi">
          <el-input v-model.lazy="formData.sayHi" type="textarea" :rows="1" :clearable="true" placeholder="请输入打招呼内容" />
        </el-form-item>
        <el-form-item v-if="arrInputNums.length" :label="arrInputNums.length+'个号码'" prop="arrInputNums">
          <div v-for="(number, index) in arrInputNums" :key="index">
            {{ number }} &nbsp;
          </div>
        </el-form-item>
        <!-- <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
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

<script>
export default {
  name: 'RunNum'
}
</script>

<script setup>
import {
  createRunNum,
  deleteRunNum,
  deleteRunNumByIds,
  updateRunNum,
  updateRunNumByIds,
  findRunNum,
  getRunOrders,
  getRunNumList
} from '@/api/runNum'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatDateShort, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { computed, ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const stateOptions = ref([])
const phone_typeOptions = ref([])
const formData = ref({
  num: '',
  state: 1,
  numType: 1,
  eachEnterNum: 99,
  sayHi: '',
  numsText: '',
  orderName: undefined,
  userId: 0,
})

// 验证规则
const rule = reactive({
  state: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  orderName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  eachEnterNum: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  numType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  userId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchOptions = ref({})
const searchInfo = ref({})
const orderNameOptions = ref([])
const arrInputNumsValue = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
  orderNameOptions.value = []
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getRunNumList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    searchOptions.value = table.data.searchOptions
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  stateOptions.value = await getDictFunc('state')
  phone_typeOptions.value = await getDictFunc('phone_type')
}

const arrInputNums = computed(() => {
  const numberArr = getInputNums()
  return numberArr
})

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteRunNumFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteRunNumByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateRunNumFunc = async(row) => {
  const res = await findRunNum({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rerunNum
    dialogFormVisible.value = true
  }
}

const updateRunNumsFunc = () => {
  type.value = 'updates'
  dialogFormVisible.value = true
}

// 删除行
const deleteRunNumFunc = async(row) => {
  const res = await deleteRunNum({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = async() => {
  type.value = 'create'
  const res = await getRunOrders()
  if (res.code === 0) {
    dialogFormVisible.value = true
    orderNameOptions.value = res.data
  }
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  arrInputNums.value = arrInputNumsValue.value
  formData.value = {
    num: '',
    state: 1,
    numType: 1,
    eachEnterNum: 99,
    sayHi: '',
    numsText: '',
    orderName: undefined,
    userId: 0,
    ids: undefined,
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return

       if (type.value === 'updates') {
         const ids = []
         if (multipleSelection.value.length === 0) {
           ElMessage({
             type: 'warning',
             message: '请选择要修改的号码'
           })
           return
         }
         multipleSelection.value &&
                multipleSelection.value.map(item => {
                  ids.push(item.ID)
                })
         formData.value.ids = ids
         const res = await updateRunNumByIds(formData.value)
         if (res.code === 0) {
           ElMessage({
             type: 'success',
             message: '修改成功'
           })
           closeDialog()
           getTableData()
         }
       } else {
         let res
         const sendData = {
           num: formData.value.num,
           state: formData.value.state,
           numType: formData.value.numType,
           eachEnterNum: formData.value.eachEnterNum,
           sayHi: formData.value.sayHi,
           nums: getInputNums(),
           orderName: formData.value.orderName,
           userId: formData.value.userId,
         }
         //  console.log('sendData -- ' + JSON.stringify(sendData))
         switch (type.value) {
           case 'create':
             res = await createRunNum(sendData)
             break
           case 'update':
             res = await updateRunNum(formData.value)
             break
           default:
             res = await createRunNum(sendData)
             break
         }
         if (res.code === 0) {
           ElMessage({
             type: 'success',
             message: '创建/更改成功'
           })
           closeDialog()
           getTableData()
           arrInputNums.value = arrInputNumsValue.value
         }
       }
     })
}

const getInputNums = () => {
  if (!formData.value.numsText || formData.value.numsText === '') {
    return []
  }
  // 将中文逗号、英文逗号、回车符等多个分隔符替换成英文逗号
  const processedStr = formData.value.numsText.replace(/，|\n/g, ',')
  // 使用 split() 方法将处理后的字符串按照英文逗号分隔成数组
  const numberArr = processedStr.split(',')
  // 过滤掉空的字符串
  return numberArr.filter(str => Boolean(str.trim()))
}

</script>

<style>
  .status-0 {
    position: relative;
    display: inline-block;
    padding: 0 6px;
    font-size: 14px;
    text-align: center;
    border-radius: 2px;
    color: rgb(255, 255, 255);
    background-color: #FF5722;
  }
  .status-1 {
    position: relative;
    display: inline-block;
    padding: 0 6px;
    font-size: 14px;
    text-align: center;
    border-radius: 2px;
    color: rgb(255, 255, 255);
    background-color: rgb(32, 126, 102);
  }

  .status-2 {
    position: relative;
    display: inline-block;
    padding: 0 6px;
    font-size: 14px;
    text-align: center;
    border-radius: 2px;
    color: rgb(255, 255, 255);
    background-color: rgb(197, 28, 37);
  }

  .status-3 {
    position: relative;
    display: inline-block;
    padding: 0 6px;
    font-size: 14px;
    text-align: center;
    border-radius: 2px;
    color: rgb(255, 255, 255);
    background-color: rgb(85, 9, 226);
  }

  .status-4 {
    position: relative;
    display: inline-block;
    padding: 0 6px;
    font-size: 14px;
    text-align: center;
    border-radius: 2px;
    color: rgb(255, 255, 255);
    background-color: rgb(117, 5, 182);
  }
</style>
