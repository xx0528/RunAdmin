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
      <el-form :model="searchInfo" :inline="true" label-width="68px" @keyup.enter="onSubmit">
        <el-form-item label="落地页：">
          <el-select
            v-model="searchInfo.pageName"
            placeholder="请选择落地页名"
            filterable
            allow-create
            clearable
            style="width: 140px"
          >
            <el-option
              v-for="(pageId, index) in searchOptions.pageName"
              :key="index"
              :label="pageId"
              :value="pageId"
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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openAddNums">新增</el-button>
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
        <el-table-column align="left" label="工单名" prop="orderName" width="220" />
        <el-table-column align="left" label="落地页名" prop="pageName" width="220" />
        <el-table-column align="left" label="进粉限制" prop="maxEnterNum" width="120" />
        <el-table-column align="left" label="平均进粉" prop="eachEnterNum" width="120" />
        <el-table-column align="left" label="绑定号码数" prop="userNum" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateRunOrderFunc(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作" @keyup.enter="enterDialog">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="100px">
        <!-- <el-form-item label="工单名:" prop="orderName">
          <el-input v-model="formData.orderName" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="落地页ID:" prop="pageId">
          <el-input v-model="formData.pageId" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="进粉限制:" prop="maxEnterNum">
          <el-input v-model.number="formData.maxEnterNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="号码状态:" prop="state">
          <el-select v-model="formData.state" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="平均进粉:" prop="eachEnterNum">
          <el-input v-model.number="formData.eachEnterNum" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="addNumsFormVisible" :before-close="closeAddNums" title="弹窗操作" @keyup.enter="enterAddNums">
      <el-form ref="elAddNumsFormRef" :model="formAddNumsData" :rules="addNumsRule" label-width="120px">
        <el-row>
          <el-form-item label="落地页名" prop="pageId">
            <el-select
              v-model="formAddNumsData.pageId"
              placeholder="请选择落地页"
              filterable
              clearable
              style="width: auto"
            >
              <el-option
                v-for="(info, index) in pageNameOptions"
                :key="index"
                :label="info.pageName"
                :value="info.pageId"
              />
            </el-select>
          </el-form-item>
        </el-row>
        <!-- 第一行：网页地址和密码输入 -->
        <el-row :gutter="10">
          <el-col :span="12">
            <el-form-item label="网页地址" prop="orderUrl">
              <el-input v-model.trim="formAddNumsData.orderUrl" :clearable="true" placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="提取密码" prop="orderPsw">
              <el-input v-model.trim="formAddNumsData.orderPsw" :clearable="true" placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-button type="primary" :loading="isLoading" @click="getNumsData">提取</el-button>
          </el-col>
        </el-row>

        <!-- 第二行：手机号码展示 -->
        <el-row>
          <el-col :span="24">
            <el-checkbox v-if="formAddNumsData.numbers.length" v-model="formAddNumsData.isSelectAll" style="text-align: left;" @change="selectAll(0)">{{ formAddNumsData.isSelectAll ? '取消所有' : '全选所有' }}</el-checkbox>
            <el-row v-if="formAddNumsData.numbers.length">
              <el-col :span="8">
                <el-checkbox v-model="formAddNumsData.isSelectAllColumn1" style="text-align: left;" @change="selectColumn(2)">{{ formAddNumsData.isSelectAllColumn1 ? '取消整列' : '选择整列' }}</el-checkbox>
                <div v-for="(number, index) in filterNums1" :key="index">
                  <el-checkbox v-model="number.isSelected">{{ number.numId }}</el-checkbox>
                </div>
              </el-col>
              <el-col :span="8">
                <el-checkbox v-model="formAddNumsData.isSelectAllColumn2" style="text-align: left;" @change="selectColumn(3)">{{ formAddNumsData.isSelectAllColumn2 ? '取消整列' : '选择整列' }}</el-checkbox>
                <div v-for="(number, index) in filterNums2" :key="index">
                  <el-checkbox v-model="number.isSelected">{{ number.numId }}</el-checkbox>
                </div>
              </el-col>
              <el-col :span="8">
                <el-checkbox v-model="formAddNumsData.isSelectAllColumn3" style="text-align: left;" @change="selectColumn(1)">{{ formAddNumsData.isSelectAllColumn3 ? '取消整列' : '选择整列' }}</el-checkbox>
                <div v-for="(number, index) in filterNums3" :key="index">
                  <el-checkbox v-model="number.isSelected">{{ number.numId }}</el-checkbox>
                </div>
              </el-col>
            </el-row>
          </el-col>
        </el-row>

        <!-- 第三行：手动输入号码 -->
        <el-row :gutter="20">
          <el-col :span="20">
            <el-form-item label="手动输入粘贴匹配" prop="numsText" label-width="80px">
              <el-input ref="elNumsInput" v-model="formAddNumsData.numsText" type="textarea" :rows="3" :clearable="true" placeholder="请输入号码 可同时用回车或，或,分割&#13;如果是要监控的工单，将号码复制过来点击匹配会自动勾选" />
            </el-form-item>
            <el-form-item v-if="arrInputNums.length" :label="arrInputNums.length+'个号码'" prop="arrInputNums">
              <div v-for="(number, index) in arrInputNums" :key="index">
                {{ number }} &nbsp;
              </div>
            </el-form-item>
          </el-col>
          <el-button v-if="arrInputNums.length && formAddNumsData.numbers.length" type="primary" style="height:70px; width:55px;" @click="matchNums">匹配</el-button>
        </el-row>

        <!-- 第四行：限制进粉数 -->
        <el-row :gutter="5">
          <el-col :span="10">
            <el-form-item label="进粉限制" prop="maxEnterNum">
              <el-input v-model.number="formAddNumsData.maxEnterNum" type="number" :clearable="true" placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="平均进粉" prop="eachEnterNum">
              <el-input v-model.number="formAddNumsData.eachEnterNum" type="number" :clearable="true" placeholder="请输入" />
            </el-form-item>
          </el-col>
        </el-row>
        <!-- 第五行：工单名称和号码类型 -->
        <el-row :gutter="5">
          <el-col :span="10">
            <el-form-item label="工单名称" prop="orderName">
              <el-input v-model.trim="formAddNumsData.orderName" :clearable="true" placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="号码类型" prop="numType">
              <el-select v-model="formAddNumsData.numType" placeholder="请选择类型" style="width:100%" :clearable="true">
                <el-option v-for="(item,key) in phoneTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <!-- 第六行：打招呼内容 -->
        <el-row v-show="formAddNumsData.numType===1" :gutter="20">
          <el-col :span="22">
            <el-form-item label="打招呼内容" prop="sayHi">
              <el-input v-model.lazy="formAddNumsData.sayHi" type="textarea" :rows="1" :clearable="true" placeholder="请输入打招呼内容" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-form-item v-if="allNums.length" :label="'共'+allNums.length+'个号'" prop="allNums">
            <div v-for="(number, index) in allNums" :key="index">
              {{ number }} &nbsp;
            </div>
          </el-form-item>
        </el-row>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddNums">取 消</el-button>
          <el-button type="primary" @click="enterAddNums">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'RunOrder'
}
</script>

<script setup>
import {
  createRunOrder,
  deleteRunOrder,
  deleteRunOrderByIds,
  updateRunOrder,
  findRunOrder,
  getRunOrderList,
  getRunPages,
  getOrderNums,
  createRunNums
} from '@/api/runOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatDateShort, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { computed, ref, reactive } from 'vue'
import clipboard from 'clipboardy'
// 自动化生成的字典（可能为空）以及字段

const stateOptions = ref([])
const phoneTypeOptions = ref([])

const formData = ref({
  orderName: '',
  pageId: '',
  pageName: '',
  maxEnterNum: 999,
  eachEnterNum: 99,
  sayHi: '',
})

const formAddNumsData = ref({
  id: 0,
  orderUrl: '',
  orderPsw: '',
  orderUrlReal: '',
  orderUrlType: '',
  isSelectAll: false,
  isSelectAllColumn1: false,
  isSelectAllColumn2: false,
  isSelectAllColumn3: false,
  numbers: [],
  orderName: '',
  numType: 1,
  numsText: '',
  maxEnterNum: 999,
  eachEnterNum: 99,
  sayHi: '',
  pageId: '',
  pageName: '',
})

// 验证规则
const rule = reactive({
  orderName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  pageId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  pageName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  maxEnterNum: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  eachEnterNum: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 添加号码规则
const addNumsRule = reactive({
  orderName: [
    { required: true, message: '请输入工单名', trigger: 'blur' }
  ],
  numType: [
    { required: true, message: '请选择类型', trigger: ['input', 'blur'] }
  ],
  pageName: [
    { required: true, message: '请选择落地页', trigger: ['input', 'blur'] }
  ],
  eachEnterNum: [{
    validator: (rule, value, callback) => {
      // 验证maxEnterNum是否大于0
      if (value <= 0) {
        callback(new Error('平均进粉必须大于0'))
      } else {
        callback()
      }
    },
    trigger: 'blur',
  }],
  maxEnterNum: [{
    validator: (rule, value, callback) => {
      // 验证maxEnterNum是否大于0
      if (value <= 0) {
        callback(new Error('进粉数量必须大于0'))
      } else {
        callback()
      }
    },
    trigger: 'blur',
  }]
})

const elAddNumsFormRef = ref()
const elNumsInput = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchOptions = ref({})
const searchInfo = ref({})
const pageNameOptions = ref([])

const filterNums1Value = ref([])
const filterNums2Value = ref([])
const filterNums3Value = ref([])
const arrInputNumsValue = ref([])
const allNumsValue = ref([])
const isLoading = ref(false)
// 重置
const onReset = () => {
  searchInfo.value = {}
  pageNameOptions.value = []
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
  const table = await getRunOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    searchOptions.value = table.data.searchOptions
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 点击“提取”按钮，获取数据
const getNumsData = async() => {
  console.log(formAddNumsData)
  if (!formAddNumsData.value.orderUrl) {
    ElMessage('请输入网页地址！')
    return
  }
  isLoading.value = true
  formAddNumsData.value.numbers = []
  const res = await getOrderNums({ Url: formAddNumsData.value.orderUrl, Psw: formAddNumsData.value.orderPsw })
  // console.log('res -- ' + JSON.stringify(res))
  isLoading.value = false
  if (res.code === 0) {
    // console.log('data --- ' + JSON.stringify(res.data))
    const content = res.data
    if (content.code !== 200) {
      ElMessage(content.msg)
      return
    }
    console.log('进粉总数--' + content.data.intoAllFuns)
    formAddNumsData.value.orderUrlReal = content.orderUrl
    formAddNumsData.value.orderUrlType = content.orderType
    formAddNumsData.value.numbers = content.data.list
  }
}

const matchNums = async() => {
  for (var idx = 0; idx < formAddNumsData.value.numbers.length; idx++) {
    formAddNumsData.value.numbers[idx].isSelected = false
    for (var strIdx = 0; strIdx < arrInputNums.value.length; strIdx++) {
      if (arrInputNums.value[strIdx].trim() === formAddNumsData.value.numbers[idx].numId.trim()) {
        formAddNumsData.value.numsText = formAddNumsData.value.numsText.replace(arrInputNums.value[strIdx], '')
        formAddNumsData.value.numbers[idx].isSelected = true
      }
    }
  }
  // console.log('formAddNumsData.value.numsText--' + formAddNumsData.value.numsText)
  // elNumsInput.value.modelValue = formAddNumsData.value.numsText
}

// 点击“全选”或“取消全选”按钮
const selectAll = (columnIndex) => {
  formAddNumsData.value.numbers.forEach(number => {
    number.isSelected = !number.isSelected
  })
}

// 点击每列“全选”或“取消全选”按钮
const selectColumn = (columnIndex) => {
  formAddNumsData.value.numbers.forEach((number, index) => {
    if ((index + 1) % 3 === columnIndex - 1) {
      number.isSelected = !number.isSelected
    }
  })
}

const filterNums1 = computed(() => {
  return formAddNumsData.value.numbers.filter((number, index) => (index + 1) % 3 === 1)
})

const filterNums2 = computed(() => {
  return formAddNumsData.value.numbers.filter((number, index) => (index + 1) % 3 === 2)
})

const filterNums3 = computed(() => {
  return formAddNumsData.value.numbers.filter((number, index) => (index + 1) % 3 === 0)
})

const arrInputNums = computed(() => {
  const numberArr = getInputNums()
  return numberArr
})

const allNums = computed(() => {
  const arr = getAllNums()
  return arr
})
// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  stateOptions.value = await getDictFunc('state')
  phoneTypeOptions.value = await getDictFunc('phone_type')
}

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
    deleteRunOrderFunc(row)
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
  const res = await deleteRunOrderByIds({ ids })
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
const updateRunOrderFunc = async(row) => {
  const res = await findRunOrder({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rerunOrder
    dialogFormVisible.value = true
  }
}

// 弹窗控制标记
const addNumsFormVisible = ref(false)

const openAddNums = async() => {
  // addNumsFormVisible.value = true
  // 这里就不去查询信息了
  const res = await getRunPages()
  if (res.code === 0) {
    addNumsFormVisible.value = true
    pageNameOptions.value = res.data
    // formAddNumsData.value.userId = res.data.rerunPage.userId
    // formAddNumsData.value.pageId = res.data.rerunPage.pageId
    // formAddNumsData.value.pageName = res.data.rerunPage.pageName
    // formAddNumsData.value.id = row.ID
  }
}

const closeAddNums = () => {
  addNumsFormVisible.value = false
  // 保存计算属性的值
  filterNums1.value = filterNums1Value.value
  filterNums2.value = filterNums2Value.value
  filterNums3.value = filterNums3Value.value
  arrInputNums.value = arrInputNumsValue.value
  allNums.value = allNumsValue.value

  formAddNumsData.value = {
    id: 0,
    orderUrl: '',
    orderPsw: '',
    orderUrlReal: '',
    orderUrlType: '',
    isSelectAll: false,
    isSelectAllColumn1: false,
    isSelectAllColumn2: false,
    isSelectAllColumn3: false,
    numbers: [],
    orderName: '',
    numType: 1,
    numsText: '',
    maxEnterNum: 999,
    eachEnterNum: 99,
    pageId: '',
    pageName: '',
  }
}
// 弹窗确定
const enterAddNums = async(row) => {
    elAddNumsFormRef.value?.validate(async(valid) => {
      if (!valid) return

      const arrNums = getAllNums()
      var pageName = ''
      pageNameOptions.value.forEach(pInfo => {
        if (pInfo.pageId === formAddNumsData.value.pageId) {
          pageName = pInfo.pageName
        }
      })
      const sendData = {
        ID: formAddNumsData.value.id,
        NumType: Number(formAddNumsData.value.numType),
        OrderName: formAddNumsData.value.orderName,
        OrderUrl: formAddNumsData.value.orderUrlReal,
        OrderUrlType: formAddNumsData.value.orderUrlType,
        PageId: formAddNumsData.value.pageId,
        PageName: pageName,
        MaxEnterNum: formAddNumsData.value.maxEnterNum,
        EachEnterNum: formAddNumsData.value.eachEnterNum,
        SayHi: formAddNumsData.value.sayHi,
        Nums: arrNums
      }
      console.log('sendData -- ' + JSON.stringify(sendData))
      const res = await createRunNums(sendData)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建成功'
        })
        closeAddNums()
        getTableData()
        // 重置计算属性
        filterNums1.value = filterNums1Value.value
        filterNums2.value = filterNums2Value.value
        filterNums3.value = filterNums3Value.value
        arrInputNums.value = arrInputNumsValue.value
        allNums.value = allNumsValue.value
      }
    })
}

const getInputNums = () => {
  if (!formAddNumsData.value.numsText || formAddNumsData.value.numsText === '') {
    return []
  }
  // 将中文逗号、英文逗号、回车符等多个分隔符替换成英文逗号
  const processedStr = formAddNumsData.value.numsText.replace(/，|\n/g, ',')
  // 使用 split() 方法将处理后的字符串按照英文逗号分隔成数组
  const numberArr = processedStr.split(',')
  // 过滤掉空的字符串
  return numberArr.filter(str => Boolean(str.trim()))
}

const getAllNums = () => {
  const arrNums = []
  if (formAddNumsData.value.numbers && formAddNumsData.value.numbers.length > 0) {
    formAddNumsData.value.numbers.forEach(number => {
      if (number.isSelected) {
        arrNums.push(number.numId)
      }
    })
  }

  const arrInput = getInputNums()
  arrNums.push.apply(arrNums, arrInput)
  return arrNums
}

// 删除行
const deleteRunOrderFunc = async(row) => {
  const res = await deleteRunOrder({
    ID: row.ID,
    OrderName: row.orderName,
    PageId: row.pageId,
  })
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
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    orderName: '',
    pageId: '',
    pageName: '',
    maxEnterNum: 999,
    eachEnterNum: 99,
    sayHi: '',
  }
}

// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createRunOrder(formData.value)
           break
         case 'update':
           res = await updateRunOrder(formData.value)
           break
         default:
           res = await createRunOrder(formData.value)
           break
       }
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: '创建/更改成功'
         })
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>
