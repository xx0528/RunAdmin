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
            v-model="searchInfo.pageName"
            placeholder="请选择落地页名"
            filterable
            allow-create
            clearable
            style="width: 150px"
          >
            <el-option
              v-for="(name, index) in searchOptions.pageName"
              :key="index"
              :label="name"
              :value="name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="国家：">
          <el-select
            v-model="searchInfo.country"
            placeholder="请选择国家"
            filterable
            allow-create
            clearable
            style="width: 150px"
          >
            <el-option
              v-for="(name, index) in searchOptions.countryName"
              :key="index"
              :label="name"
              :value="name"
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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: auto;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
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
        <el-table-column align="left" label="日期" width="100">
          <template #default="scope">{{ formatDateShort(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="落地页名" prop="pageName" width="200" />
        <el-table-column align="left" label="国家" prop="country" width="150" />
        <el-table-column align="left" label="用户数" prop="userNum" width="70" />
        <!-- <el-table-column align="left" label="用户id" prop="userId" width="120" /> -->
        <el-table-column align="left" label="状态" prop="state" width="60">
          <template #default="scope">
            {{ filterDict(scope.row.state,stateOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="链接" prop="url" width="400" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <!-- <el-button type="primary" link icon="edit" @click="openAddNums(scope.row)">添加号码</el-button> -->
            <el-button type="primary" link icon="edit" class="table-button" @click="updateRunPageFunc(scope.row)">变更</el-button>
            <!-- <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名字:" prop="pageName">
          <el-input v-model="formData.pageName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="国家:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-show="type === 'update'" label="链接:" prop="url">
          <el-input v-model.trim="formData.url" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-show="type === 'update'" label="备注:" prop="remark">
          <el-input v-model.trim="formData.remark" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item v-show="false" label="用户数:" prop="userNum">
          <el-input v-model.number="formData.userNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-show="false" label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-show="false" label="状态:" prop="state">
          <el-select v-model="formData.state" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'RunPage'
}
</script>

<script setup>
import {
  createRunPage,
  deleteRunPage,
  deleteRunPageByIds,
  updateRunPage,
  findRunPage,
  getRunPageList
} from '@/api/runPage'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatDateShort, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { computed, watch } from 'vue'
import { init } from 'events'
import { async } from 'q'
import { ref, reactive } from 'vue'

const stateOptions = ref([])
const phoneTypeOptions = ref([])
const formData = ref({
  name: '',
  url: '',
  remark: '',
  country: '',
  userNum: 0,
  userId: 0,
  state: undefined,
})

// 验证规则
const rule = reactive({
})

const elFormRef = ref()
// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchOptions = ref({})
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
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
  const table = await getRunPageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteRunPageFunc(row)
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
  const res = await deleteRunPageByIds({ ids })
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
const updateRunPageFunc = async(row) => {
  const res = await findRunPage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rerunPage
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteRunPageFunc = async(row) => {
  const res = await deleteRunPage({ ID: row.ID })
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
    name: '',
    url: '',
    remark: '',
    country: '',
    userNum: 0,
    userId: 0,
    state: undefined,
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createRunPage(formData.value)
           break
         case 'update':
           res = await updateRunPage(formData.value)
           break
         default:
           res = await createRunPage(formData.value)
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

const playAlertAndSpeak = (alertSound, text) => {
  // 播放警报声音
  alertSound.play()

  // 在警报声音播放完之后，播放语音
  alertSound.onended = () => {
    const utterance = new SpeechSynthesisUtterance(text)
    utterance.lang = 'zh-CN'
    window.speechSynthesis.speak(utterance)
  }
}

// playAlertAndSpeak(new Audio('src/assets/alert.mp3'), "001号管理员的印度工单进粉数已达到上限")

</script>

<style>
</style>
